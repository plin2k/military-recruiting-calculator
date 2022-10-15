package main

import (
	"fmt"
)

func main() {
	calc, err := New("dataset.csv")
	if err != nil {
		panic(err)
	}

	fmt.Println(calc.Forward(false, false, false, 2, "В", "IT", "SVO", "IST"))

}
