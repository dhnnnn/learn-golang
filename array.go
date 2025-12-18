package main

import "fmt"

func main(){
	var names [3]string
	
	names[0] = "Dhani"
	names[1] = "Dwi"
	names[2] = "Putra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	
	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)

	var NoKTP = [...]int{
		100000,
		200000,
		300000,
		90990909,
	}

	fmt.Println(len(NoKTP))
}