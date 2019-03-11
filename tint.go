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
	// Red color
	Red
	// Green color
	Green
	// Yellow color
	Yellow
	// Blue color
	Blue
	// Magenta color
	Magenta
	// Cyan color
	Cyan
	// White color
	White
	// BlackBright color
	BlackBright
	// RedBright color
	RedBright
	// GreenBright color
	GreenBright
	// YellowBright color
	YellowBright
	// BlueBright color
	BlueBright
	// MagentaBright color
	MagentaBright
	// CyanBright color
	CyanBright
	// WhiteBright color
	WhiteBright

	// BgBlack color
	BgBlack
	// BgRed color
	BgRed
	// BgGreen color
	BgGreen
	// BgYellow color
	BgYellow
	// BgBlue color
	BgBlue
	// BgMagenta color
	BgMagenta
	// BgCyan color
	BgCyan
	// BgWhite color
	BgWhite
)

var colorMap = map[color]string{
	// Normal
	0: ":",
	// Black
	1: "\u001B[30m:\u001B[39m",
	// Red
	2: "\u001B[31m:\u001B[39m",
	// Green
	3: "\u001B[32m:\u001B[39m",
	// Yellow
	4: "\u001B[33m:\u001B[39m",
	// Blue
	5: "\u001B[34m:\u001B[39m",
	// Magenta
	6: "\u001B[35m:\u001B[39m",
	// Cyan
	7: "\u001B[36m:\u001B[39m",
	// White
	8:  "\u001B[37m:\u001B[39m",
	9:  "\u001B[90m:\u001B[39m",
	10: "\u001B[91m:\u001B[39m",
	11: "\u001B[92m:\u001B[39m",
	12: "\u001B[93m:\u001B[39m",
	13: "\u001B[94m:\u001B[39m",
	14: "\u001B[95m:\u001B[39m",
	15: "\u001B[96m:\u001B[39m",
	16: "\u001B[97m:\u001B[39m",

	// backgrounds
	17: "\u001B[40m:\u001B[49m",
	18: "\u001B[41m:\u001B[49m",
	19: "\u001B[42m:\u001B[49m",
	20: "\u001B[43m:\u001B[49m",
	21: "\u001B[44m:\u001B[49m",
	22: "\u001B[45m:\u001B[49m",
	23: "\u001B[46m:\u001B[49m",
	24: "\u001B[47m:\u001B[49m",
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

// apply the colors to the text by wrapping text with ANSI styling
func apply(text string, colors []color) string {
	output := text
	for _, c := range colors {
		brackets := strings.Split(colorMap[c], ":")
		output = brackets[0] + output + brackets[1]
	}
	return output
}
