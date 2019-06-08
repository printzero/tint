### Introduction

In this document demonstrates how to use `tint` go module and it's functions. Tint has some really great functions to get you started with coloring your terminal output. We first need to import
tint and initialize a pointer to tint. For that we have a `Init()` function.

Example:

```go
package main

import "github.com/printzero/tint"

func main() {
	t := tint.Init()
}
```

The next steps are written assuming you have initialized tint like the above example.

#### func Raw

This functions lets you color a string and returns a colored string.

```go
cyanStr := t.Raw("This is cyan colored string")
fmt.Println(cyanStr)
```

#### func Print

Print single line of text with specified color.

```go
t.Print("I give you a green text", tint.Green)
```

#### func Println

Println single line of text with enter character.

```go
t.Print("I give you a green text", tint.Green)
```

#### func Log

Log text with the standard lib log module.

```go
t.Log("Something went wrong", tint.Red)
```

#### func Palette

Palette lets you build a string with specific words with different background or foreground color of your choice

>NOTE: no need to specify space character at end of mixin

```go
coloredString := t.Palette(
    t.With("Welcome", tint.BgRed),
    t.With("to this", tint.Normal),
    t.With("awesome", tint.Cyan),
    t.With("module created by", tint.Normal),
    t.With("codekidX", tint.BgGreen, tint.Black),
)
```
#### func With

With is used to build a Mixin with text and color.

```go
t.With("awesome", tint.Cyan) // used in palette example above
```

#### func Swatch

Swatch helps you create your own function with specified color. 

```go
green := t.Swatch(tint.Green)
green("I give you a green text")
```

This is how you write a Swatch function with the same example used in [func Println](#func-println).
Awesome, _Right?_ 

#### func SwatchRaw

Unlike [Swatch](func-swatch), this returns a string for your own use.

```go
green := t.SwatchRaw(tint.Green)
hmm := green("I give you a green text")

fmt.Printf(hmm + " : %v", myStruct)
```
