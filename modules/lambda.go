package main

import (
	"context"
	"encoding/json"
	"fmt"
	"modules/accountmanagement"
	"modules/dboperations"
	"modules/email"
	"modules/mlmodel"
	"modules/parser"
	"modules/payment"
	"modules/validator"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

type InputEvent struct {
	Requestid string `json:"Requestid"`
}

type ResponseEvent struct {
	Message string `json:"Message:"`
}

var db dboperations.DBOperations

func init() {
	// Get DB connection
	fmt.Println("Loading DB Schemas")
	dbfilepath := os.Getenv("DB_FILENAME")
	fmt.Println(dbfilepath)
	db = dboperations.GetDbInstance(dbfilepath)
	fmt.Println("DB connection established", db.GetDB())
}

func HandleRequest(ctx context.Context, requestEvent events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	requestid := requestEvent.PathParameters["requestid"]
	fmt.Println("Path Parameter", requestid)

	req := parser.Request{requestid, db}
	validator := validator.Validator{}
	am := accountmanagement.AccountManagement{db}
	email := email.Email{}
	payment := payment.Payment{}
	model := mlmodel.MLModelCall{}

	msg := req.AuthMessageParser()
	if !validator.PinValidation(msg.GetPin(), 7456) {
		fmt.Println("Invalid PIN")
		os.Exit(1)
	} else {
		fmt.Println("PIN validation...SUCCESS")
	}

	if !validator.AavValidation(msg.GetAav(), 234) {
		fmt.Println("Invalid AAV")
		os.Exit(1)

	} else {
		fmt.Println("AAV validation...SUCCESS")
	}

	// Start account management as one task 1
	amc := make(chan string)
	go am.AccountManagement(amc, msg)

	// Start email in another task 2
	emailc := make(chan string)
	go email.EmailAgeNeustar(emailc)

	// Waiting for return value from above two tasks - AccountManagement & EmailAgeNeustar
	resultAcctMngt := <-amc
	resultEmail := <-emailc

	// Print two return values from tasks
	fmt.Println(resultAcctMngt, " returns from Account Management task.")
	fmt.Println(resultEmail, " returns from Email task")

	//Proceed further after completion of two tasks only
	payment.PaymentVariableCalculation()
	am.SEAccums()
	validator.HSMCIDValidation()
	am.LockAccount()

	// Start ASYNC call to get results from RNN model
	modelc := make(chan mlmodel.MLModelCall)
	go model.RNNModelCall(modelc)

	am.CustAccountDetail()

	// Waiting for return value from above ASYNC task - MLModelCall
	resultMLModelCall := make([]mlmodel.MLModelCall, 1)
	resultMLModelCall[0] = <-modelc

	// Print return values from ML model task
	if resultMLModelCall[0].GetStatus() {
		fmt.Println(resultMLModelCall[0].GetResultData(), " returns from RNN model")

		//Proceed further after completion of RNN model call
		model.GBMModelExecuation()
	} else {
		fmt.Println("No Model data hence stopping program...")
	}

	returnResponse := ResponseEvent{
		Message: fmt.Sprintf("Request processed Successfully for requestid %s", requestid),
	}
	returnResponseJson, _ := json.Marshal(returnResponse)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(returnResponseJson),
	}, nil

}
