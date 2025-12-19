package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Person struct {
	firstName, lastName string
}

func main(){
	addres1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	addres2 := addres1 //ini adalah copy by value addres 1 jadi addres 1 tidak berubah jika di ubah

	addres2.City = "Bandung"

	fmt.Println(addres1) //tidak berubah
	fmt.Println(addres2) //berubah menjadi bandung

	//============= Menggunakan Pointer ====================//
	person1 := Person{"Dhani", "Putra"}
	person2 := &person1 //ini adalah pointer jadi person 1 akan berubah jika person 2 di ubah
	
	person2.lastName = "Eko"
	fmt.Println(person1) //berubah menjadi eko
	fmt.Println(person2) //berubah menjadi eko
}