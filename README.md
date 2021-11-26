# new.Of
A helper function to create a pointer to a new variable of a particular value in Go 1.18+.

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

## Installation

As of November 2021, Go 1.18 is not released, but you can install Go tip with

```
$ go install golang.org/dl/gotip@latest
$ gotip download
$ gotip init me/myproject
$ gotip get github.com/carlmjohnson/new
```

## FAQs

### [I love and hate this at the same time...](https://reddit.com/r/golang/comments/r064xk/_/hlr2bdi/?context=1)

I mean, honestly, this isn't even my worst idea for generics yet.

### What problem does this solve exactly?

In Go, you cannot return a pointer to an expression, although you can return a pointer to a struct literal. As a result, `return &struct{ 42 }` is legal, but `return &42` is not. To work around this, many popular packages include pointer helpers, such as [aws.String](https://pkg.go.dev/github.com/aws/aws-sdk-go/aws#String) and [github.Int64](https://pkg.go.dev/github.com/google/go-github/v39/github#Int64). Thanks to generics, now one function can solve this problem once and for all.

### How can you name a package `new`?

`new` is a built-in function, not a keyword, so it is legal to shadow the name. If this is a problem for your code because you are still using the built-in `new` function in legacy code, you can use an import alias for this package.

### What does test coverage look like?

100%, baby.

### Should I use this package?

Ideally, `newof` will become [a built-in function in a future version of Go](https://github.com/golang/go/issues/45624#issuecomment-927391928). Until then, [this is fine](http://gunshowcomic.com/648).

[![fine](https://user-images.githubusercontent.com/222245/142966985-627d6095-313f-475f-ba98-fa37ef892cbe.png)](https://github.com/carlmjohnson/shitpic/)
