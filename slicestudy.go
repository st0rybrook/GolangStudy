package main

import "fmt"

func main() {
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	var mySlice1 []string
	fmt.Println(mySlice1)
	mySlice1 = append(mySlice1, "tang")
	fmt.Println(mySlice1)
	fmt.Println(len(mySlice1))
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
