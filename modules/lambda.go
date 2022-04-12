package main

import (
	"context"
	"data/customerdata"
	"datamanipulation/datamanipulate"
	"datareconciliation/reconciledata"
	"encoding/json"
	"fmt"
	"notify/notification"
	"os"
	"processdata/processdata"
	"util/dbutil"

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

var db dbutil.DBOperations

func init() {
	// Get DB connection
	fmt.Println("Loading DB Schemas")
	dbfilepath := os.Getenv("DB_FILENAME")
	fmt.Println(dbfilepath)
	db = dbutil.loadInitialData(dbfilepath)
	fmt.Println("DB connection established", db.GetDB())
}

func HandleRequest(ctx context.Context, requestEvent events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	customerid := requestEvent.PathParameters["customerid"]
	fmt.Println("Path Parameter", customerid)

	req := customerdata.InputRequest{customerid, db}
	cdm := datamanipulate.CustomerDataRecordManagement{db}
	notification := notification.Notification{}
	recondata := reconciledata.ReconcileData{}
	res := processdata.ResultDataFetch{}

	msg := req.CustomerData()

	// Spawn customer data management goroutine
	cd := make(chan string)
	go cdm.SampleDataManipulate(cd, msg)

	// Spawn notification goroutine
	emailChannel := make(chan string)
	go notification.Notify(emailChannel)

	// Await Results - CustomerDataRecordManagement & Notification
	rescdm := <-cd
	resultNotification := <-emailChannel

	// Print two return values from tasks
	fmt.Println(rescdm, " returns from Customer Data Record Management.")
	fmt.Println(resultNotification, " returns from Notification task")

	//Proceed further after completion of two tasks only
	recondata.ReconcileCustomerData()
	cdm.SampleProcess1()
	cdm.SampleProcess2()

	// Spawn another sample goroutine
	resultChannel := make(chan processdata.ResultData)
	go res.ResultDataFetch(resultChannel)

	cdm.CustAccountDetail()

	// Waiting for return value from goroutine
	result := make([]processdata.ResultDataFetch, 1)
	result[0] = <-resultChannel

	if result[0].GetStatus() {
		fmt.Println(result[0].GetResultData(), " returns from processdate")

		//Proceed further after completion
		res.ResultDataSampleExecute()
	} else {
		fmt.Println("Data Processing Error")
	}

	returnResponse := ResponseEvent{
		Message: fmt.Sprintf("Request processed successfully for customerid %s", customerid),
	}
	returnResponseJson, _ := json.Marshal(returnResponse)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(returnResponseJson),
	}, nil

}
