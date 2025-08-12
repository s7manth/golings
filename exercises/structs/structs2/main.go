// structs2
// Make me compile!

package main

import "fmt"

type Phone struct {
	countryCode int
	number int
}

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age int
	phone Phone
}

func main() {
	// contactDetails := ContactDetails{}
	person := Person{name: "John", age: 32, phone: Phone{countryCode: 45, number: 89442247}}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
