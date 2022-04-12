package notify

import (
	"fmt"
)

type Notification struct {
}

func (e Notification) Notify(emailChannel chan string) {
	fmt.Println("Sending Email Notification")
	emailChannel <- "TRUE"
}
