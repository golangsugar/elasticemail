package elasticemail

type Person struct {
	Name,
	Address string
}

func (p Person) To() map[string]string {

}

func (p Person) From() map[string]string {

}

func (p Person) ReplyTo() map[string]string {

}
