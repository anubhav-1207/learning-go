package main 

import "fmt"

func main(){
	var a,b,c int
	
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if (a > b) && (a > c){
		fmt.Println("a is biggest")
	} else if (b > a) && (b > c){
		fmt.Println("b is biggest")
	} else if (c > b) && (c > a){
		fmt.Println("c is biggest")
	} else{
		fmt.Println("All are equal")
	}

}