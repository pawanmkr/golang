package main

import (
	f "fmt"
)

func calculator(integers ...int) (add, sub, product, division int) {
	for i, value := range integers {
		if i == 0 {
			add = integers[0]
		} else {
			add += value
		}
	}
	for i, value := range integers {
		if i == 0 {
			sub = integers[0]
		} else {
			sub -= value
		}
	}
	for i, value := range integers {
		if i == 0 {
			product = integers[0]
		} else {
			product *= value
		}
	}
	for i, value := range integers {
		if i == 0 {
			division = integers[0]
		} else {
			division /= value
		}
	}
	return
}

func main() {
	var add, sub, product, division = calculator(1, 2, 3, 4, 5)

	values := map[string]int{
		"add":      add,
		"sub":      sub,
		"product":  product,
		"division": division,
	}

	var operator string

	f.Printf("~ what you wanna do? ( + - * / )\n")
	f.Printf("> ")
	f.Scanf("%s", &operator)

	switch operator {
	case "+":
		f.Printf("added values make the sum of %d\n", values["add"])
	case "-":
		f.Printf("subtracting the values we get %d\n", values["sub"])
	case "*":
		f.Printf("entered values make the product of %d\n", values["product"])
	case "/":
		f.Printf("after diving the values, this is the result %d\n", values["division"])
	}
}
