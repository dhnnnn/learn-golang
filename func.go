package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello " + name
}

func main(){
	sayHello("Dhani Putra")
	
	result := getHello("Eko")
	fmt.Println(result)
}