package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {
	//harish := person{"Harish", "Gaddale", 39}
	harish02 := person{firstName: "Harish", lastName: "Gaddale"}
	//var harish03 person
	//fmt.Println(harish02, harish03)
	harishPointer := &harish02
	harishPointer.updateName("Hrrish")
	harish02.print()
	harish02.updateName("this")
	harish02.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%v", p)
}
