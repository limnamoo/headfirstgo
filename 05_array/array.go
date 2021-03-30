package main

import(
	"fmt"
)

func main(){
	var numbers [3]int
	numbers[0] = 12
	numbers[2] = 34
	fmt.Println("numbers : ", numbers)
	fmt.Printf("%#v\n", numbers) // print array literal

	var letters = [3]string{"a","b","c"}
	fmt.Println("letters : ", letters)
	fmt.Printf("%#v\n", letters)

	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line", // last comma
	}
	fmt.Println("text : ", text)

	notes := [7]string{"do","re","mi","fa","so","la","ti"}

	// loop 1
	for i:=0; i<len(notes); i++{
		fmt.Println(i, notes[i])
	}

	// loop 2
	for index, note := range notes{
		fmt.Println(index, note)
	}

	// loop 2 with _
	var sum int
	for _, number := range numbers{
		sum += number
	}
	fmt.Println("sum = ", sum)


}
