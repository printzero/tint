<p align="center">
    <img src="https://raw.githubusercontent.com/printzero/tint/master/assets/tint_logo.png" width="400" height="225">
</p>

Tint aims to provide functions that you can use to color your cli text output. This project is originally written to be used 
in the project called [orb](https://github.com/orbpkg/orb) which I'm currently working on. It resolves some of my concerns
about the state of terminal styling implemented in Go.


[![GoDoc](https://godoc.org/github.com/printzero/tint?status.svg)](https://godoc.org/github.com/printzero/tint)

### Import

```bash
go get github.com/printzero/tint
```

### Getting started

The below topics will get you started with the basics of this module and also some of the features that it provides.

#### Initialization

Initialize a pointer to `Tint` struct

```go
t := tint.Init()
```

#### Functions

Tint has some functions that let you directly color your terminal output, Example:

```go
t.Println("This is an error", t.Red)
```

this function is equivalent to `fmt.Println()`, but notice the extra parameter passed. This is a color constant.

A lot of times we need to use the raw string representation of colored string for various purposes and this is where
`t.Raw` is used. Example:

```go
err := writeCode()

if err != nil {
	panic(err)
}

func writeCode() error {
	return errors.New(t.Raw("This is a very bad error and it occurred while coding", tint.Red))
}
```

#### Higher-order color functions

It is possible with this module to create higher order functions with the help of `Swatch`. A swatch returns a function for a given color variable and print as well, Example:

```go
package main

import (
	"time"
	"tint"
	)

var green func(text string)
var yellow func(text string)

func main() {
	t := tint.Init()
	
	green = t.Swatch(tint.Green)
	yellow = t.Swatch(tint.Yellow)
	
	yellow("Waiting for 2 seconds just for fun ...")
	time.Sleep(time.Second * 2)
	green("I give you a green light")
}
```

It is also possible for a swatch function to return a raw colored string using `SwatchRaw` but does not print it to stdout. So the above example becomes:

 ```go
package main

import (
	"fmt"
	"time"
	"tint"
	)

var green func(text string) string
var yellow func(text string) string

func main() {
	t := tint.Init()
	green = t.SwatchRaw(tint.Green)
	yellow = t.SwatchRaw(tint.Yellow)
	y := yellow("Waiting for 2 seconds just for fun ...")
	fmt.Printf("Yellow string: %s", y)
	time.Sleep(time.Second * 2)
	g := green("I give you a green light")
	fmt.Printf("Green string: %s", g)
}
```

### Contributing to this project

Take a look at this [CONTIBUTING.md](https://github.com/printzero/tint/blob/master/CONTIBUTING.md).

### Github Project Board

I use this github project [board](https://github.com/printzero/tint/projects/1) to track changes for this project and any new feature, bug or ideas for upcoming releases are posted here.


### License

This project is licensed under `Apache 2.0`
