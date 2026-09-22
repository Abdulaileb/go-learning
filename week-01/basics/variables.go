package basics

import (
	"fmt"
)

// In golang, strings are concatenated

var (
	name      string
	age       int
	isStudent bool
)

func VariablesDemo() {
	fmt.Println("go " + "lang")

	var myname string
	myname = "lebbie"

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("My name is = ", myname)

	//Booleans
	fmt.Println(true || false)
	fmt.Println(true && false)
	fmt.Println(!true)

	// variablesDemo() // Commented out to prevent infinite recursion
}

func VariablesDemo2() {
	var mygoals string
	mygoals = "to become successful in this world and the next"

	fmt.Println("i love my goals : ", mygoals)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var uniq string
	uniq = "Udine1"

	var uni2 = "udine3"

	uni := "Udine"
	fmt.Println(uni, uni2, uniq)

	// dates

}

func GroupingVariableDemo() {
	name = "John"
	age = 30
	isStudent = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
}
