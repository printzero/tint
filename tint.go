package tint

import (
	"fmt"
	"log"
)

// Tint struct holds the whole library
type Tint struct {
	Level         TerminalLevel
	SupportsColor bool
	LogInstance   *log.Logger
}

// TerminalLevel of color support for terminal and how does tint
// works with these levels
type TerminalLevel int

// Color Default that support the terminal
type color struct {
	open  string
	close string
}

func (c color) Dim() color {
	return color{
		open:  "\u001b[2m" + c.open,
		close: c.close + "\u001b[0m",
	}
}

func (c color) Bold() color {
	return color{
		open:  "\u001b[1m" + c.open,
		close: c.close + "\u001b[0m",
	}
}

func (c color) Underline() color {
	return color{
		open:  "\u001b[4m" + c.open,
		close: c.close + "\u001b[0m",
	}
}

func (c color) Italic() color {
	return color{
		open:  "\u001b[3m" + c.open,
		close: c.close + "\u001b[0m",
	}
}

func (c color) Strike() color {
	return color{
		open:  "\u001b[9m" + c.open,
		close: c.close + "\u001b[0m",
	}
}

func (c color) Add(this color) color {
	return color{
		open:  c.open + this.open,
		close: this.close + c.close,
	}
}

// Mixin helps bind a color to a specific text
type Mixin struct {
	text   string
	colors []color
}

// Normal equates to no style
var Normal = color{
	open:  "\u001b[0m",
	close: "\u001b[0m",
}

// Black color brackets in ansi format
var Black = color{
	open:  "\u001b[30m",
	close: "\u001b[39m",
}

// Red color
var Red = color{
	open:  "\u001b[31m",
	close: "\u001b[39m",
}

// Green color
var Green = color{
	open:  "\u001b[32m",
	close: "\u001b[39m",
}

// Yellow color
var Yellow = color{
	open:  "\u001b[33m",
	close: "\u001b[39m",
}

// Blue color
var Blue = color{
	open:  "\u001b[34m",
	close: "\u001b[39m",
}

// Magenta color
var Magenta = color{
	open:  "\u001b[35m",
	close: "\u001b[39m",
}

// Cyan color
var Cyan = color{
	open:  "\u001b[36m",
	close: "\u001b[39m",
}

// White color
var White = color{
	open:  "\u001b[37m",
	close: "\u001b[39m",
}

// BgBlack color
var BgBlack = color{
	open:  "\u001b[40m",
	close: "\u001b[49m",
}

// BgRed color
var BgRed = color{
	open:  "\u001b[41m",
	close: "\u001b[49m",
}

// BgGreen color
var BgGreen = color{
	open:  "\u001b[42m",
	close: "\u001b[49m",
}

// BgYellow color
var BgYellow = color{
	open:  "\u001b[43m",
	close: "\u001b[49m",
}

// BgBlue color
var BgBlue = color{
	open:  "\u001b[44m",
	close: "\u001b[49m",
}

// BgMagenta color
var BgMagenta = color{
	open:  "\u001b[45m",
	close: "\u001b[49m",
}

// BgCyan color
var BgCyan = color{
	open:  "\u001b[46m",
	close: "\u001b[49m",
}

// BgLightGrey color
var BgLightGrey = color{
	open:  "\u001b[47m",
	close: "\u001b[49m",
}

// White background
var BgWhite = color{
	open:  "\u001b[107m",
	close: "\u001b[49m",
}

// Hyperlink text
//Hyperlink

const (
	// LevelNone for terminal that supports no colot
	LevelNone TerminalLevel = iota + 1
	// Level16bit for terminal that supports 16bit colors
	Level16bit
	// Level256 for terminal that supports 256 bit colors
	Level256
	// Level16m for terminal that supports 16 million colors (truecolor)
	Level16m
)

// Init initializes variables that tint uses and then returns the
// pointer to a Tint struct
func Init() *Tint {
	return &Tint{
		Level:         LevelNone,
		SupportsColor: false,
		LogInstance:   &log.Logger{},
	}
}

// Exp returns a string constructed from a series of color expressions given as an argument
func (t *Tint) Exp(expStr string, colors ...color) string {
	return replaceExp(expStr, colors)
}

// Raw returns the raw string with applied colors
func (t *Tint) Raw(text string, colors ...color) string {
	return apply(text, colors)
}

// Print single line of text with specified color
func (t *Tint) Print(text string, colors ...color) {
	fmt.Print(apply(text, colors))
}

// Println single line of text with enter character
func (t *Tint) Println(text string, colors ...color) {
	fmt.Println(apply(text, colors))
}

// Log text with the standard lib log module
func (t *Tint) Log(text string, colors ...color) {
	t.LogInstance.Println(apply(text, colors))
}

// Palette lets you build a string with specific words with different
// background or foreground color of your choice
// NOTE: no need to specify space character at end of mixin
func (t *Tint) Palette(mixins ...Mixin) string {
	output := ""

	// for each mixins in this palette
	for i, m := range mixins {
		if i == 0 {
			output = apply(m.text, m.colors)
		} else {
			output = output + " " + apply(m.text, m.colors)
		}
	}

	return output
}

// With is used to build a Mixin with text and color
func (t *Tint) With(text string, colors ...color) Mixin {
	return Mixin{
		text,
		colors,
	}
}

// Swatch helps you create your own function with spesified color
// Example:
//
//	green := tint.Swatch(tint.Green)
func (t *Tint) Swatch(colors ...color) func(text string) {
	return func(text string) {
		t.Println(text, colors...)
	}
}

// SwatchRaw returns a functions that returns a raw colored string
func (t *Tint) SwatchRaw(colors ...color) func(text string) string {
	return func(text string) string {
		return t.Raw(text, colors...)
	}
}

// apply the colors to the text by wrapping text with ANSI styling
func apply(text string, colors []color) string {
	output := text
	for _, c := range colors {
		output = c.open + output + c.close
	}
	return output
}

func replaceExp(text string, colors []color) string {

	if tc, proceed := isAtCountSame(text, len(colors)); !proceed {
		//tintedError := replaceExp(, []color{Magenta.Bold()})
		panic(fmt.Errorf("mismatching apply - Trigger count: %d, Color count: %d \n\n%sTip%s: Does your function have colors passed as much as it has '@' character?", tc, len(colors), Magenta.open, Magenta.close))
		return text
	}

	workingString := ""
	workingColors := colors

	// this is used as the current color on operation
	//var colorStart int
	//var colorEnd = -1
	var colorBox = White
	var triggered = false

	// so here is what we'll do --- we'll find the first #( which serves as opening
	// of a color expression, and then look for the next close with colorBox in memory
	for i, c := range text {
		if string(c) == "@" && string(text[i+1]) == "(" {
			triggered = true
			workingString = workingString + ""
			continue
		} else if triggered && string(c) == "(" {
			colorBox = workingColors[0]
			workingString = workingString + colorBox.open
			continue
		} else if triggered && string(c) == ")" {
			triggered = false
			workingString = workingString + colorBox.close
			colorBox = White
			workingColors = workingColors[1:]
			continue
		}

		workingString = workingString + colorBox.open + string(c) + colorBox.close
	}

	return workingString
}

func isAtCountSame(text string, colorCount int) (int, bool) {
	counter := 0
	for i, c := range text {
		if string(c) == "@" && string(text[i+1]) == "(" {
			counter += 1
		}
	}

	return counter, counter == colorCount
}

func indexOf(text string, char string) int {
	for i, s := range text {
		if string(s) == char {
			return i
		}
	}
	return -1
}
