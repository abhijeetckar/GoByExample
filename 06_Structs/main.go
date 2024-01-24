package main

import "fmt"

type employee struct {
	id   int
	name string
	age  int
}

func main() {
	emps := make([]employee, 3)
	emps[0] = employee{1, `first`, 18}
	emps[1] = employee{2, `second`, 19}
	emps[2] = employee{3, `third`, 20}

	for _, e := range emps {
		fmt.Println(e.id)
		fmt.Println(e.name)
		fmt.Println(e.age)
	}

	emp := employee{4, `Sam`, 21}
	fmt.Printf("Emp: %v\n", emp)
	fmt.Printf("Emp: %+v\n", emp)
	fmt.Printf("Emp: %#v\n", emp)
	fmt.Println(emp)

	e := new(employee)
	fmt.Println(*e)
	fmt.Printf("%+v\n", *e)
}
