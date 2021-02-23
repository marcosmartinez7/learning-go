package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	corey := person{
		firstName: "Corey",
		lastName:  "Taylor",
		contactInfo: contactInfo{
			email:   "corey@gmail.com",
			zipCode: 123,
		},
	}
	me := person{
		firstName: "Marcos",
		lastName:  "Martinez",
		contactInfo: contactInfo{
			email:   "mmartinezciompi@gmail.com",
			zipCode: 123,
		},
	}
	me.lastName = "Ciompi"
	corey.updateName("Juan")
	corey.print()
	me.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", p)

}
