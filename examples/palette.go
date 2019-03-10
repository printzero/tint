package main

import (
	"fmt"
	"tint"
)

func main() {
	t := tint.Init()
	// palette example
	fmt.Println(t.Palette(
		t.With("Ashish Shekar", tint.Magenta),
		t.With("is an", tint.Normal),
		t.With("awesome", tint.Cyan),
		t.With("guy!", tint.Normal),
	))
}
