package main

import "fmt"

func main() {
	fmt.Print("Enter Degrees in Celcius: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input * 9/5) + 32
	
	fmt.Println(output,"Fahrenheit")
}
