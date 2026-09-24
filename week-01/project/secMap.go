// this is a security policy enumeration mapping

package project

import "fmt"

type RiskLevel int

const (
	Non RiskLevel = iota
	Low
	Medium
	High
	Critical
)

func (r RiskLevel) String() string {
	return [...]string{"Non", "Low", "Medium", "High", "Critical"}[r]
}

func SecurityPolicyMapping() {
	fmt.Println("Security policy enumeration mapping")
}

func MapRiskLevelToDescription(r RiskLevel) string {
	return fmt.Sprintf("Risk level %s has a corresponding description.", r)
}

func ShouldBlock(r RiskLevel) bool {
	return r >= High
}

func DisplayResult9() {
	fmt.Println("Displaying security policy mapping results:")
	for _, r := range []RiskLevel{Non, Low, Medium, High, Critical} {
		fmt.Printf("Risk level: %s, Description: %s, Should block: %t\n", r, MapRiskLevelToDescription(r), ShouldBlock(r))
	}
}
