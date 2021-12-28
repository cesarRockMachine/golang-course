package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	blackmamba := person{
		firstName: "Kobe",
		lastName:  "Bryant",
		contact: contactInfo{
			email:   "kobe@lakers.com",
			zipCode: 838132,
		},
	}
	//fmt.Println(blackmamba)
	blackmamba.print()
	blackmamba.updateName("The Goat")
	blackmamba.print()

}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
