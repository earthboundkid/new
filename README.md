# new
A helper function to create a pointer to a new object in Go 1.18+.

```
strptr1 := new.Of("meaning of life")
strptr2 := new.Of("meaning of life")
strptr1 != strptr2 // true
*strptr1 == *strptr2 // true

intp1 := new.Of(42)
intp2 := new.Of(42)
intp1 != intp2 // true
*intp1 == *intp2 // true

type MyFloat float64
fp := new.Of[MyFloat](42)
fp != nil // true
*fp == 42 // true
```
