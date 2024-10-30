package main

import "fmt"

func aaa(arr []int) map[int]int {
	mapu := make(map[int]int)
	for _, value := range arr {
		mapu[value]++
	}
	return mapu
}
func main() {
	arr := []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	mapu := aaa(arr)
	fmt.Println("元素出现次数", mapu)
}
