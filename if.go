package main

import "fmt"

func main(){
	name := "Eko"

	if name == "Dhani Putra" {
		fmt.Println("Hello", name)
	} else if name == "Eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hi, Siapa Kamu?")
	}
}