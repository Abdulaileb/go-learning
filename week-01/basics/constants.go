package basics

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

const Constants string = "constant"

func ConstantsDemo() {

	fmt.Println(Constants)

	const n = 500000000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	const maximumRetries = 45
	const serviceName = "api-key"
	const lazyLoad = true
	fmt.Println(maximumRetries, serviceName, lazyLoad)

}
