package main

import (
	"fmt"
	"math"
	"time"
)

//var delcaration
var c, python, java bool
var i int

func main2() {
	// println with a param
	fmt.Println("Hey this is formatted printing", time.Now())
	// printf with var substitution
	fmt.Printf("The rand number is %g yo \n", math.Sqrt(7))

	//functional stuffs
	fmt.Println(add(1, 2))

	//return multiple values
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	//implicit named returns
	fmt.Println(split(17))
	//
	fmt.Println(i, c, python, java)

	//for loops w/ init
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	//for w/o init
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// if example
	fmt.Println(sqrt(2), sqrt(-4))
}

// declaring functions
func add(x int, y int) int {
	return x + y
}

// decalring multiple vars via shorthand
func add_alt(x, y int) int {
	return x + y
}

//function with returns
func swap(x, y string) (string, string) {
	return y, x
}

//function with implicit named returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//if's
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
