package mlmodel

import (
	"fmt"
)

type MLModelCall struct {
	resultdata string
	status     bool
}

func (m MLModelCall) RNNModelCall(modelc chan MLModelCall) {
	fmt.Println("RNN model call...")
	//time.Sleep(4 * time.Second)
	modelc <- MLModelCall{"RNNModelDATA...", true}
}

func (m MLModelCall) GBMModelExecuation() {
	fmt.Println("GBM model execuation...")
}

func (m *MLModelCall) GetStatus() bool {
	return m.status
}

func (m *MLModelCall) GetResultData() string {
	return m.resultdata
}
