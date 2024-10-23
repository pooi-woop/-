package main

import "fmt"
import "math"

func circlearea(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	S := circlearea(8)
	fmt.Println(S)
}
