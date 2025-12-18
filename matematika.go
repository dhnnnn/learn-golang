package main

import "fmt"

func main() {
	var a = 2
	var b = 10

	var c = a * b

	var d = c + a
	fmt.Println(c)
	fmt.Println(d)


	var i = 10
	i += 5
	fmt.Println(i)

	i -= 3
	fmt.Println(i)

	var j = 20
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
}