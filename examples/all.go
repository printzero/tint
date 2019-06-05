package main

import (
	"fmt"
	"tint"
)

func main() {
	t := tint.Init()
	// 1
	t.Println("Welcome to tint go module.", tint.Magenta)
	// 2
	t.Print("Welcome, ", tint.Yellow)
	t.Print("to tint go module\n", tint.Dim)
	// 3
	printThis := t.Raw("t.Raw returns a raw ansi colored string", tint.BgMagenta)
	fmt.Println(printThis)
	// 4
	printThis = t.Raw("t.Raw returns a raw ansi colored string", tint.BgMagenta, tint.Black)
	fmt.Println(printThis)
	// 5
	//t.Log("Something is awefully wrong here", tint.Red)
	// 6
	warn := t.Swatch(tint.Yellow)
	warn("This is a warning function using t.Swatch")
	// 7
	coloredString := t.Palette(
		t.With("Welcome", tint.BgRed),
		t.With("to this", tint.Normal),
		t.With("awesome", tint.Cyan),
		t.With("module created by", tint.Normal),
		t.With("codekidX", tint.BgGreen, tint.Black),
	)

	fmt.Println(coloredString)
	// 8
	fmt.Println(t.Exp("+r|Ashish|+ is a c|cool|! guy"))
	// 9
	fmt.Println(t.Exp("y|Welcome|!, to the world of i|daemons|> !!!"))
}
