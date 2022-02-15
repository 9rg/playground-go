package main

import "fmt"

func main(){
	s := []int {3, 4, 5, 6, 8}
	printSlice(s)

	// cap stable
	s = s[:3]
	printSlice(s)

	// cap change
	s = s[1:]
	printSlice(s)

	// secure double amount of cap
	s = append(s, 4)
	printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d s=%v\n", len(s), cap(s), s)
}

