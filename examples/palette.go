package main

import (
	"fmt"
	"tint"
)

var cyan func(s string)

func main() {
	t := tint.Init()
	cyan = t.Swatch(tint.Cyan)
	// palette example
	fmt.Println(t.Palette(
		t.With("Ashish Shekar", tint.Black, tint.BgYellow),
		t.With("is an", tint.Normal),
		t.With("awesome", tint.Cyan),
		t.With("guy!", tint.Normal),
		t.With("https://ashishshekar.com", tint.Hyperlink),
	))

	cyan("Swatch implementation with color cyan")
}
