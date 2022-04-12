package datamanipulation

import (
	"fmt"
	"modules/common"
	"modules/data"
	"util/dbutil"
)

type CustomerDataRecordManagement struct {
	Dboperations dbutil.DBOperations
}

func (cdm CustomerDataRecordManagement) SampleDataManipulate(cd chan string, msg data.CustomerRecord) {
	cdm.startProcess()
	cdm.collectSampleData1(msg)
	cdm.collectSampleData2(msg)
	cdm.collectSampleData3(msg)
	cdm.endProcess()
	cd <- "TRUE"
}

func (acct CustomerDataRecordManagement) SampleProcess3() {
	fmt.Println("SampleProcess3...")
}
func (acct CustomerDataRecordManagement) SampleProcess1() {
	fmt.Println("SampleProcess1 ...")
}
func (acct CustomerDataRecordManagement) SampleProcess2() {
	fmt.Println("SampleProcess2...")
}

func (acct CustomerDataRecordManagement) collectSampleData2(msg data.CustomerRecord) {

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
	fmt.Println("Collect Sample Data 2.....")
	fmt.Println("--------------------------")
	fmt.Println("Card type - ", cardtype)
	fmt.Println("Limits - ", limits)
	fmt.Println("Currency - ", currency)
	fmt.Println("--------------------------")
}
func (acct CustomerDataRecordManagement) collectSampleData3(msg data.CustomerRecord) {

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
func (acct CustomerDataRecordManagement) startProcess() {
	fmt.Println("Starting Data Manipulation Process...")
}
func (acct CustomerDataRecordManagement) collectSampleData1(msg data.CustomerRecord) {

	fmt.Println("-------------------------------")
	fmt.Println("Collect Sample  Data 1...")
	fmt.Println("-------------------------------")
	fmt.Println("Account no - ", msg.GetAccountno())
	fmt.Println("Card type - ", msg.GetCardtype())
	fmt.Println("Customer id - ", msg.GetCustid())
	fmt.Println("-------------------------------")

}
func (acct CustomerDataRecordManagement) endProcess() {
	fmt.Println("Ending Data Manipulation Process")
}
