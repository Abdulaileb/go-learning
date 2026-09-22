package main

import (
	"fmt"
	"math"
)

type Severity int

const (
	Infor_Security = iota //
	Low
	Meduim
	High
	Critical
)

func (s Severity) String() string {
	return [...]string{"INFO", "LOW", "MEDIUM", "HIGH", "CRITICAL"}[s]
}

const s string = "constant"

func con() {

	fmt.Println(s)

	const n = 500000000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	const maximumRetires = 45
	const serviceName = "api-key"
	const lazyLoad = true

}

func main() {
	con()
	fmt.Println(Severity(High))
	fmt.Println(Severity(Critical))
	fmt.Printf("%d\n", Severity(Meduim))
}
