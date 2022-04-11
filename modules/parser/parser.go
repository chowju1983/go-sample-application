package parser

import (
	"fmt"
	"modules/common"
	"modules/dboperations"
	"strconv"
)

type Request struct {
	Requestid    string
	Dboperations dboperations.DBOperations
}

type Message struct {
	sessionid  string
	accountno  string
	cardtype   string
	cardno     string
	pin        int
	aav        int
	custid     string
	cardtypeid int
}

func (m *Message) GetSessionid() string {
	return m.sessionid
}
func (m *Message) GetAccountno() string {
	return m.accountno
}

func (m *Message) GetCardno() string {
	return m.cardno
}

func (m *Message) GetPin() int {
	return m.pin
}

func (m *Message) GetCardtype() string {
	return m.cardtype
}

func (m *Message) GetAav() int {
	return m.aav
}

func (m *Message) GetCustid() string {
	return m.custid
}

func (m *Message) GetCardtypeid() int {
	return m.cardtypeid
}

func (r Request) AuthMessageParser() Message {

	db := r.Dboperations.GetDB()
	msg := Message{}
	fmt.Println("Auth message parsing...", db)
	reqid, err := strconv.Atoi(r.Requestid)
	common.CheckError(err)
	fmt.Println("RequestId", reqid)
	rows, dberr := db.Query(`select sess.session_id, acc.account_no, prd.card_type, 
		acc.card_no, acc.pin, acc.aav, acc.cust_id, prd.card_type_id from 
		"FraudRiskSchema"."Request" req, "FraudRiskSchema"."Session" sess,
		"FraudRiskSchema"."Account" acc, "FraudRiskSchema"."ProductData" prd where
		req.session_id = sess.session_id and sess.account_id = acc.account_id and
		acc.card_type_id = prd.card_type_id and req.request_id = $1`, reqid)

	fmt.Println("Query results Retrieved...", dberr)
	common.CheckError(dberr)

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&msg.sessionid, &msg.accountno, &msg.cardtype, &msg.cardno, &msg.pin, &msg.aav, &msg.custid, &msg.cardtypeid)
		common.CheckError(err)
	}

	return msg
}
