package main

import "fmt"

func main() {

	type NoKTP string

	var ktpDhani NoKTP = "317xxxxxxxxxxx"

	var contoh string = "222222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpDhani)
	fmt.Println(contohKtp)
}