package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := []int{3, 4, 4}
	arr[1] = 5
	slice = append(slice, 7)

	fmt.Println(arr)
	fmt.Println(slice)
}
