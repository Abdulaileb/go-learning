package main

import (
	"fmt"

	"example.com/go-learning/week-01/basics"
	"example.com/go-learning/week-01/project"

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
	fmt.Println(project.TestPasswordVerification())
	project.SecurityPolicyMapping()
	project.DisplayResult9()
}
