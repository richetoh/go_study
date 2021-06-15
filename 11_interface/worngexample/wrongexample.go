package wrongexample

import "fmt"

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
func SendBook(parcel string, sender *DaehanSender) {
	sender.Send(parcel)
}

func Wrongexample() {
	SendBook("5trillion", &DaehanSender{})
	SendBook("5trillion", &RichetSender{})

}
