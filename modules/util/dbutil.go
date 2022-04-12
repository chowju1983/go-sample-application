package util

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"modules/common"
	"os"

	_ "github.com/lib/pq"
)

const (
	//host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type ProductData struct {
	Cardtypeid int
	Cardtype   string
	Limits     int
	Currency   string
}

type CustomerProfile struct {
	Custid  int
	Name    string
	Address string
	Age     string
	Phone   string
	Email   string
}

type Account struct {
	Accountid  int
	Accountno  int
	Cardtypeid int
	Cardno     int
	Pin        int
	Aav        int
	Custid     int
}

type Session struct {
	Sessionid int
	Accountid int
	Token     string
}

type Request struct {
	Requestid           int
	Browserinfo         string
	Guid                string
	Originoftransaction string
	Transactiontype     string
	Sessionid           int
	Channel             string
	Ipaddress           string
	Deviceid            string
	Devicemodel         string
	Deviceversion       string
	Journeytype         string
	Agentid             string
}

type DBOperations struct {
	db *sql.DB
}

func (dboperations *DBOperations) GetDB() *sql.DB {
	return dboperations.db
}

func loadInitialData(dbfilepath string) DBOperations {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	dboperations := DBOperations{db}
	common.CheckError(err)

	query, readerr := ioutil.ReadFile(dbfilepath)
	common.CheckError(readerr)
	if _, err = db.Exec(string(query)); err != nil {
		log.Fatalf("%v", err)
		panic(err)
	}

	return dboperations
}
