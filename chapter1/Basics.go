package main

import "fmt"

func main() {
	and := true && false
	fmt.Println(and)

	or := true || false
	fmt.Println(or)

	not := !true
	fmt.Println(not)

	year := 2023
	fmt.Println(year)

	rate := 1.25
	fmt.Println(rate)

	terms := "Terms"
	fmt.Println(terms)

	/** Go really doesn't care about constants
	  So you can a lot of unused constant :)
	*/
	const language string = "Go"

	var a [3]int
	fmt.Println(a)

	a[0] = 2
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b[1])

	fmt.Println(len(b))

	mySlice := make([]int, 3)
	println(cap(mySlice))
	fmt.Println(language)

	var c int = 10
	var p = &c
	fmt.Println(c)
	fmt.Println(p)

	sum, i := 0, 0
	for i < 10 {
		i++
		sum += i
	}
	fmt.Println(sum)
}
