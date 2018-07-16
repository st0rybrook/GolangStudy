package main
import "fmt"
func main() {
	mySlice := make([]int, 5, 10)
	mySlice = append(mySlice, 1, 2, 3)
	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...)
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	fmt.Println("cap(mySlice2):", cap(mySlice2))
}
