package parcel

import "fmt"

type FedexSender struct {
	Name string
}

func (f *FedexSender) SendParcel(s string) {
	fmt.Println("Fedex sending ", s)
}

type PostSender struct {
	Name string
}

func (p *PostSender) SendParcel(s string) {
	fmt.Println("Post sending ", s)
}

type Sender interface {
	SendParcel(s string)
}

func CheckWhoItIs(name string, s *Sender) {
	fmt.Println("this is ", name)
}
func Compatibility() {

	var sender1 Sender
	var sender2 Sender

	sender1 = &FedexSender{"Fedex"}
	sender2 = &PostSender{"Post"}

	sender1.SendParcel("sender1's product")
	sender2.SendParcel("sender2's product")

	CheckWhoItIs("Fedex", &sender1)
	CheckWhoItIs("Post", &sender2)

}
