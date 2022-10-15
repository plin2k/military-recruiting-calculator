package main

import (
	"fmt"
)

func main() {
	calc, err := New("dataset.csv")
	if err != nil {
		panic(err)
	}

	//calc.AddCase(true, false, false, 2, "Б", "IT", "SVO", "IST", true)
	//calc.AddCase(true, true, false, 2, "Б", "IT", "SVO", "IST", true)

	//calc.Apply()

	fmt.Println(calc.Forward(false, false, false, 2, "В", "IT", "SVO", "IST"))

	//calc.Load("gonn")

}
