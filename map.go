package main

import "fmt"

func main(){
	person := map[string]string{
		"name":    "Dhani Putra",
		"address": "Indonesia",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "Programmer"
	fmt.Println(person)

	delete(person, "address")
	fmt.Println(person)

	//membuat map baru
	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Dhani Putra"
	fmt.Println(book)
}