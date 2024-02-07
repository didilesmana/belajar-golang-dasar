package main

import (
	"fmt"
)

func main() {

	names := [...]string{"didi", "lesmana", "budi", "joko", "eko", "kurniawan"}

	slices1 := names[4:6]
	fmt.Print(slices1)

	slices2 := names[:3]
	fmt.Println(slices2)

	slices3 := names[3:]
	fmt.Println(slices3)

	slices4 := names[:]
	fmt.Println(slices4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "sabtu lama")
	fmt.Println(daySlice1)
	fmt.Print(daySlice2)
	fmt.Print(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "didi"
	newSlice[1] = "didi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "lesmnana")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "budi"
	fmt.Println(newSlice2)
	fmt.Print(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
