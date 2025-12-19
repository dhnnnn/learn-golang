package main

import "fmt"

func returnMultiple() (string, string) {
	return "Dhani", "Putra"
}

func main(){
	// firstName, lastName := returnMultiple()
	// fmt.Println(firstName, lastName)

	firstName, _ := returnMultiple()
	fmt.Println(firstName)
}