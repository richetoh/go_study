package correctexample

import "fmt"

type Sender interface {
	Send(parcel string)
}

type DaehanSender struct {
}

func (d *DaehanSender) Send(parcel string) {
	fmt.Printf("Daehan sends %v parcel", parcel)
}

type RichetSender struct {
}

func (d *RichetSender) Send(parcel string) {
	fmt.Printf("Richet sends %v parcel", parcel)
}

func SendBook(parcel string, sender Sender) {
	sender.Send(parcel)
}

//SendBook's sender type is related to the method receiver.
//If it's a pointer receiver, you should place pointer type as an argument for SendBook function
//else, you should place value type as an argument for SendBook function
func Correctexample() {
	SendBook("5trillion", &DaehanSender{})
	//below is wrong
	SendBook("5trillion", &RichetSender{})

}
