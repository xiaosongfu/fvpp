package fvpp

import "fmt"

func ExamplePretty() {
	fmt.Println(Pretty(u))
	// Output:
	// [Name:'bili'(string) Married:'true'(bool) Cars:'2'(int) Deposit:'200000'(float64)]
}
