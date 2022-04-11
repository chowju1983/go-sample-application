package email

import (
	"fmt"
)
type Email struct {
	
}
func (e Email) EmailAgeNeustar(emailc chan string)  {
	fmt.Println("Email Age Neustart.....")
	// Set value for channel
	emailc <- "TRUE"	
}