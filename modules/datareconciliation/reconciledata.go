package datareconciliation

import (
	"fmt"
)

type ReconcileData struct {
}

func (p ReconcileData) ReconcileCustomerData() {
	fmt.Println("Reconciling Customer Data")
}
