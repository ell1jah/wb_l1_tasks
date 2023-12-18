package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

func (h Human) FullName() string {
	return fmt.Sprintf("%s %s", h.Name, h.Surname)
}

type Action struct {
	//embedding
	Human
	Place string
}

func main() {
	h := Human{
		Name:    "Ilya",
		Surname: "Ivanov",
	}

	a := Action{h, "Pdsv"}

	fmt.Println(a.FullName())
}
