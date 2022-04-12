package processdata

import (
	"fmt"
)

type ResultData struct {
	resultdata string
	status     bool
}

func (m ResultData) ResultDataFetch(result chan ResultData) {
	fmt.Println("ResultDataFetch call...")
	result <- ResultData{"RESULT", true}
}

func (m ResultData) ResultDataSampleExecute() {
	fmt.Println("ResultDataSampleExecute..")
}

func (m *ResultData) GetStatus() bool {
	return m.status
}

func (m *ResultData) GetResultData() string {
	return m.resultdata
}
