# new
A helper function to create a pointer to a new object in Go 1.18+.

```
string1 := new.Of("meaning of life")
string2 := new.Of("meaning of life")
string1 != string2 // true
*string1 == *string2 // true

int1 := new.Of(42)
int2 := new.Of(42)
int1 != int2 // true
*int1 == *int2 // true

type MyFloat float64
f := new.Of[MyFloat](42)
f != nil // true
*f == 42 // true
```
