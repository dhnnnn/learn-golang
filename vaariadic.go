package main

import "fmt"

func samAll (numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main(){
	result := samAll(10, 20, 30, 40, 50)
	fmt.Println(result)
}