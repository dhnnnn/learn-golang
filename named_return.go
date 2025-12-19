package main

import "fmt"

func getCompleteName() (firstName string, lastName string) {
	firstName = "Dhani"
	lastName = "Putra"

	return firstName, lastName
}

func main(){
	firstName, lastName := getCompleteName()
	fmt.Println("Hello,", firstName, lastName)
}