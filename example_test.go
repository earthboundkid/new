package new_test

import (
	"fmt"

	"github.com/carlmjohnson/new"
)

func ExampleOf() {
	string1 := new.Of("meaning of life")
	string2 := new.Of("meaning of life")
	fmt.Println(string1 != string2)
	fmt.Println(*string1 == *string2)

	int1 := new.Of(42)
	int2 := new.Of(42)
	fmt.Println(int1 != int2)
	fmt.Println(*int1 == *int2)

	type MyFloat float64
	f := new.Of[MyFloat](42)
	fmt.Println(f != nil)
	fmt.Println(*f == 42)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
