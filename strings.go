package main 

import (
	"fmt";
	"strconv";
	"strings"
)

func main(){
	var var1 string = "Hello" //manual declaration using type
	var var2 = "World"		//automatic type inference when the value is given
	var3 := "Hehe"               //short declaration operator
	
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

	//string methods 
	fmt.Println(len(var1)) //length of the string
	fmt.Println(var1 + var2) //concatenation of strings
	fmt.Printf("%c\n",var3[0]) //accessing the first character of the string

	//string comparison
	fmt.Println(var1 == var2) //false
	fmt.Println(var1 != var2) //true

	//string conversion
	var num int = 10
	var str string = strconv.Itoa(num) //convert int to string
	fmt.Println(str)

	var str2 string = "20"
	var num2, _ = strconv.Atoi(str2) //convert string to int
	fmt.Println(num2)

	//string methods
	var variable string = "golang"
	fmt.Println(strings.ToUpper(variable)) //convert to uppercase
	fmt.Println(strings.ToLower(variable)) //convert to lowercase
	fmt.Println(strings.HasPrefix(variable, "go")) //check if string has prefix
	fmt.Println(strings.HasSuffix(variable, "lang")) //check if string has suffix
	fmt.Println(strings.Contains(variable, "lang")) //check if string contains substring
	fmt.Println(strings.ReplaceAll(variable, "go", "python")) //replace substring		
}
