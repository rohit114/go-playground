package fundamentals

import "fmt"

type Speaker interface {
	Speak() string
}

type Animal struct {
	name string
}

// func (a *Animal) Speak() {
// 	fmt.Println("all animal speaks", a.name)
// }

type Dog struct {
	Animal
	breed string
}

func (d *Dog) Speak() {
	fmt.Println("i am breed", d.breed, d.name)
}

type Cat struct {
	Animal
	color string
}

func (c *Cat) Speak() {
	fmt.Println("i am breed", c.color, c.name)
}

func OopsDemo1() {
	// Creating instances of Dog and Cat
	dog := Dog{Animal: Animal{name: "Buddy"}, breed: "Golden Retriever"}
	cat := Cat{Animal: Animal{name: "Whiskers"}, color: "Tabby"}

	// Using the Speak method from both the Dog and Cat structs
	dog.Speak() // Outputs: Woof! I am a Golden Retriever dog named Buddy
	cat.Speak() // Outputs: Meow! I am a Tabby cat named Whiskers

	p1 := NewPerson("Rohit", 28)
	fmt.Println(p1.Name, " age: ", p1.Age)
}

// Define a struct
type Person2 struct {
	Name string
	Age  int
}

// Constructor-like function to create a new Person instance
func NewPerson(name string, age int) *Person2 {
	// Initialize the struct and return a pointer to it
	return &Person2{Name: name, Age: age}
}
