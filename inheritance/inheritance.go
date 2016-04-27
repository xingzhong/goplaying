package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (animal *Animal) Bark() {
	fmt.Printf("Animal is %s\n", animal.MyName())
}

func (animal *Animal) MyName() string {
	return animal.Name
}

type Dog struct {
	*Animal
}

func (dog *Dog) Bark() {
	fmt.Printf("Dog is %s\n", dog.MyName())
}

type Cat struct {
	*Animal
}

func (cat *Cat) MyName() string {
	return "miew"
}

func main() {
	fmt.Println("Anonymous fields can make sort of inheritance")
	john := Animal{"john"}
	dog := Dog{&john}
	cat := Cat{&john}
	john.Bark()
	dog.Bark()
	cat.Bark()
}
