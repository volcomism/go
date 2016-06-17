package main

import "fmt"

func main() {
	fmt.Print("Enter Length in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input * 0.3048)
	var metricType string = "meters"	
	fmt.Println(output, metricType)
}
