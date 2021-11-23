package new_test

import (
	"fmt"

	"github.com/carlmjohnson/new"
)

func ExampleOf() {
	strptr1 := new.Of("meaning of life")
	strptr2 := new.Of("meaning of life")
	fmt.Println(strptr1 != strptr2)
	fmt.Println(*strptr1 == *strptr2)

	intp1 := new.Of(42)
	intp2 := new.Of(42)
	fmt.Println(intp1 != intp2)
	fmt.Println(*intp1 == *intp2)

	type MyFloat float64
	fp := new.Of[MyFloat](42)
	fmt.Println(fp != nil)
	fmt.Println(*fp == 42)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
