package main

import (
	"example.com/go-learning/week-01/basics"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	fmt.Println(basics.Constants)
	basics.ConstantsDemo()
	fmt.Println(basics.High)
	basics.VariablesDemo()
	basics.VariablesDemo2()
	basics.GroupingVariableDemo()
}
