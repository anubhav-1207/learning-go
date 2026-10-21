package main

import "fmt"

func main() {
	const PI = 3.141592 //single constant declaration
	const (
		GRAVITY = 9.8
		SPEED   = 299792458
	)				   //multiple constant declaration

	fmt.Println(PI)
	fmt.Println(GRAVITY)
	fmt.Println(SPEED)
	
}