package main

import "fmt"

func main(){
	names := [...]string{"Dhani", "Putra", "eko", "budi", "ani", "nugraha"}
	slice := names[1:4]
	fmt.Println(slice)

	names[2] = "Changed"
	fmt.Println(slice)

	slice[0] = "Putra slice"
	fmt.Println(names)

	fmt.Println(len(slice))

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] //menagmbil index ke 5 sampai selesai
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur baru")
	//daysBaru := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu", "Libur baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
}