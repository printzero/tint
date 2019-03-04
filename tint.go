package tint

import (
	"fmt"
	"log"
	"strings"
)

// Tint struct holds the whole library
type Tint struct {
	Level         TerminalLevel
	SupportsColor bool
}

// TerminalLevel of color support for terminal and how does tint
// works with these levels
type TerminalLevel int

// Color Default that support the terminal
type color int

// Mixin helps bind a color to a specific text
type Mixin struct {
	text   string
	colors []color
}

const (
	// Normal equates to no style
	Normal color = iota
	// Black color brackets in ansi format
	Black
	// Cyan color
	Cyan
	// BgWhite background
	BgWhite
)

var colorMap = map[color]string{
	// Normal
	0: ":",
	// Black
	1: "\u001B[30m:\u001B[39m",
	// Cyan
	2: "\u001B[36m:\u001B[39m",
	3: "\u001B[47m:\u001B[49m",
}

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
	}
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
	log.Print(apply(text, colors))
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

func apply(text string, colors []color) string {
	output := text
	for _, c := range colors {
		brackets := strings.Split(colorMap[c], ":")
		output = brackets[0] + output + brackets[1]
	}
	return output
}
