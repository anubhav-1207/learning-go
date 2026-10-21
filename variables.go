package main

import "fmt"

func main() {
	var fruit string = "apple" //manual declaration using type
	var number = 10            //automatic type inference when the value is given
	name := "John"             //short declaration operator
	var integer int            //manual declaration without initialization -> default 0
	var pi float64 = 3.14

	fmt.Println(fruit)
	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(integer)

	var newPi int = int(pi) //type conversion from float64 to int
	fmt.Println(newPi)
}

