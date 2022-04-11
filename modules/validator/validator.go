package validator

import (

	"fmt"
)
type Validator struct {
	
}
func (v Validator) PinValidation(pin int, comparepin int) bool {
	if pin != comparepin {
		return false
	}	
	return true
}
func (v Validator) AavValidation(aav int, compareaav int) bool {
	if aav != compareaav {
		return false
	}	
	return true
}
func (v Validator) HSMCIDValidation() {
	fmt.Println("HSM CID validation...")
}



