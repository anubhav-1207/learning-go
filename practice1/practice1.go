package main

import "fmt"

func main(){
	fmt.Print("What is your name?: ")
	var name string
	fmt.Scan(&name)

	fmt.Print("What is your age?: ")
	var age int
	fmt.Scan(&age)

	fmt.Printf("Hello %s I know that you are %d years old\n",name, age)

}