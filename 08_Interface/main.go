package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

type pet1 struct {
	a    animal
	name string
}

type pet2 struct {
	animal
	name string
}

func main() {
	d := dog{age: 5}
	p1 := pet1{name: "Milo", a: d}

	fmt.Println(p1.name)
	// p1.breathe() err
	// p1.walk() err
	p1.a.breathe()
	p1.a.walk()

	p2 := pet2{name: "Oscar", animal: d}
	fmt.Println(p2.name)
	p2.breathe()
	p2.walk()
	p1.a.breathe()
	p1.a.walk()

	x := lion{age: 10}
	print(x)

}

func print(a animal) {

	l, ok := a.(lion)
	if ok {
		fmt.Println(l)
	} else {
		fmt.Println("a is not of type lion")
	}

	d, ok := a.(dog)
	if ok {
		fmt.Println(d)
	} else {
		fmt.Println("a is not of type lion")
	}

	switch v := a.(type) {
	case lion:
		fmt.Println("Type: lion")
	case dog:
		fmt.Println("Type: dog")
	default:
		fmt.Printf("Unknown Type %T", v)
	}
}
