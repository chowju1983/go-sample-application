package accountmanagement

import (
	"fmt"
	"modules/common"
	"modules/dboperations"
	"modules/parser"
)

type AccountManagement struct {
	Dboperations dboperations.DBOperations
}

func (acct AccountManagement) AccountManagement(amc chan string, msg parser.Message) {
	acct.startAccountManager()
	acct.collectAccountProfileData(msg)
	//time.Sleep(2 * time.Second)
	acct.oneExposerCalculation()
	acct.accountPasswordInformation()
	//time.Sleep(2 * time.Second)
	acct.endAccountManager()
	acct.collectProductData(msg)
	acct.collectCustomerProfileData(msg)
	// Set value for channel
	amc <- "TRUE"
}
func (acct AccountManagement) CustAccountDetail() {
	fmt.Println("Customer accumer...")
	fmt.Println("Exposure/risk variables...")
}
func (acct AccountManagement) SEAccums() {
	fmt.Println("SE accums...")
}
func (acct AccountManagement) LockAccount() {
	fmt.Println("Locking account...")
}
func (acct AccountManagement) accountPasswordInformation() {
	fmt.Println("Account password information...")
}
func (acct AccountManagement) oneExposerCalculation() {
	fmt.Println("One exposer calculation.....")
}
func (acct AccountManagement) collectProductData(msg parser.Message) {

	db := acct.Dboperations.GetDB()
	rows, err := db.Query(`select card_type, limits, currency from "FraudRiskSchema"."ProductData" prd where prd.card_type_id = $1`, msg.GetCardtypeid())
	common.CheckError(err)

	defer rows.Close()
	var cardtype string
	var limits int
	var currency string
	if rows.Next() {
		err = rows.Scan(&cardtype, &limits, &currency)
		common.CheckError(err)
	}

	fmt.Println("--------------------------")
	fmt.Println("Collect Product Data.....")
	fmt.Println("--------------------------")
	fmt.Println("Card type - ", cardtype)
	fmt.Println("Limits - ", limits)
	fmt.Println("Currency - ", currency)
	fmt.Println("--------------------------")
}
func (acct AccountManagement) collectCustomerProfileData(msg parser.Message) {

	db := acct.Dboperations.GetDB()
	rows, err := db.Query(`select cust_id, name, address, age, phone, email from "FraudRiskSchema"."CustomerProfile" cust where cust.cust_id = $1`, msg.GetCustid())
	common.CheckError(err)

	defer rows.Close()
	var custid int
	var name string
	var address string
	var age string
	var phone int
	var email string
	if rows.Next() {
		err = rows.Scan(&custid, &name, &address, &age, &phone, &email)
		common.CheckError(err)
	}

	fmt.Println("--------------------------")
	fmt.Println("Collect Customer Profile..")
	fmt.Println("--------------------------")
	fmt.Println("Cust id - ", custid)
	fmt.Println("Name - ", name)
	fmt.Println("Address - ", address)
	fmt.Println("Age - ", age)
	fmt.Println("Phone - ", phone)
	fmt.Println("Email - ", email)
	fmt.Println("--------------------------")
}
func (acct AccountManagement) startAccountManager() {
	fmt.Println("Start account manager...")
}
func (acct AccountManagement) collectAccountProfileData(msg parser.Message) {

	fmt.Println("-------------------------------")
	fmt.Println("Collect Account Profile Data...")
	fmt.Println("-------------------------------")
	fmt.Println("Account no - ", msg.GetAccountno())
	fmt.Println("Card type - ", msg.GetCardtype())
	fmt.Println("Customer id - ", msg.GetCustid())
	fmt.Println("-------------------------------")

}
func (acct AccountManagement) endAccountManager() {
	fmt.Println("End account manager...")
}
