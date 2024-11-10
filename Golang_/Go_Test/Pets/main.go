package main

import "fmt"

type Pet interface {
	Speak() string
	Type() string
}

type Dog struct{}

type Cat struct{}

type Parrot struct{}

func (a Dog) Speak() string {
	return "Wooof!"
}

func (a Dog) Type() string {
	return "Dog"
}

func (b Cat) Speak() string {
	return "Meow"
}

func (b Cat) Type() string {
	return "Cat"
}

func (c Parrot) Speak() string {
	return "Chick chirick"
}

func (c Parrot) Type() string {
	return "Parrot"
}

func PrinInfo(v Pet) {
	fmt.Printf("Type: %s, Speak: %s \n", v.Type(), v.Speak())
}

func main() {
	Dog := Dog{}
	Cat := Cat{}
	Parrot := Parrot{}

	PrinInfo(Dog)
	PrinInfo(Cat)
	PrinInfo(Parrot)
}
