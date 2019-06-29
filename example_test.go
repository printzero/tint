package tint_test

import (
	"fmt"
	"tint"
)

func Example() {
	// Initialize tint
	t := tint.Init()
	// prints output in red
	t.Println("This output is red", tint.Red)
}

// This example demonstrates how attributes are defined in colors and when and how to use those attributes.
// As of `v0.0.2`tint has 5 attributes: Dim, Bold, Italic, Underline, Strike (strikethrough)
// These attributes are functions that are exposed over the color struct.
func Example_attributes() {
	// create pointer to Tint
	t := tint.Init()
	// Dim attribute
	t.Println("Dim sentence.", tint.Yellow.Dim())
	// Bold attribute
	t.Println("Bold sentence.", tint.Yellow.Bold())
	// Italic attribute
	t.Println("Italic sentence.", tint.Yellow.Italic())
	// Underline attribute
	t.Println("Underlined sentence.", tint.Yellow.Underline())
	// Strike attribute
	t.Println("Strikeout sentence.", tint.Yellow.Strike())
}

// This example demonstrates how colors are accessed differently than tint functions.
func Example_difference() {
	// create pointer to Tint
	t := tint.Init()

	// Notice: how Raw function is accessed through pointer created by Init()
	// Notice: how color Yellow are accessed, they are static variables defined inside tint module
	// so you don't have to create any instance of a color
	t.Println("This output is red", tint.Yellow)
}

func ExampleTint_Raw() {
	// Initialize tint
	t := tint.Init()
	// get a colored string to be used as terminal output
	output := t.Raw("Rejoice fellow gopher, your test cases have passed.", tint.Green)
	// use it anywhere as input - fmt, log ...
	fmt.Println(output)
}

func ExampleTint_Exp() {
	// Initialize tint
	t := tint.Init()
	coloredString := t.Exp("@(Hello), @(World)!", tint.White.Bold(), tint.Blue)
	fmt.Println(coloredString)
}
