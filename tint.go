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
	LogInstance   *log.Logger
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
	// Hyperlink text
	Hyperlink
	// Dim text
	Dim

	// internal constants
	suffixBreaker     = "\u001B[39m"
	suffixBgBreaker   = "\u001B[49m"
	suffixAttrBreaker = "\u001B[0m"
)

var colorMap = []string{
	// Normal
	":=def|=|",
	// Black
	"\u001B[30m:\u001B[39m=bl|=|!",
	// Red
	"\u001B[31m:\u001B[39m=r|=|!",
	// Green
	"\u001B[32m:\u001B[39m=g|=|!",
	// Yellow
	"\u001B[33m:\u001B[39m=y|=|!",
	// Blue
	"\u001B[34m:\u001B[39m=b|=|!",
	// Magenta
	"\u001B[35m:\u001B[39m=m|=|!",
	// Cyan
	"\u001B[36m:\u001B[39m=c|=|!",
	// White
	"\u001B[37m:\u001B[39m=w|=|!",
	"\u001B[90m:\u001B[39m=*bl|=|!",
	"\u001B[91m:\u001B[39m=*r|=|!",
	"\u001B[92m:\u001B[39m=*g|=|!",
	"\u001B[93m:\u001B[39m=*y|=|!",
	"\u001B[94m:\u001B[39m=*b|=|!",
	"\u001B[95m:\u001B[39m=*m|=|!",
	"\u001B[96m:\u001B[39m=*c|=|!",
	"\u001B[97m:\u001B[39m=*w|=|!",

	// backgrounds
	"\u001B[40m:\u001B[49m=+bl|=|+",
	"\u001B[41m:\u001B[49m=+r|=|+",
	"\u001B[42m:\u001B[49m=+g|=|+",
	"\u001B[43m:\u001B[49m=+y|=|+",
	"\u001B[44m:\u001B[49m=+b|=|+",
	"\u001B[46m:\u001B[49m=+m|=|+",
	"\u001B[47m:\u001B[49m=+c|=|+",
	"\u001B[45m:\u001B[49m=+w|=|+",
	"\u001B]8;;:\u0007link\u001B]8;;\u0007=L|=|",

	// attributes
	"\u001B[2m:\u001B[0m=i|=|>",
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
		LogInstance:   &log.Logger{},
	}
}

// Exp returns a string constructed from a series of color expressions given as an argument
func (t *Tint) Exp(expStr string) string {
	return replaceExp(expStr)
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
		brackets := getBrackets(c)
		output = brackets[0] + output + brackets[1]
	}
	return output
}

func getBrackets(c color) []string {
	comp := strings.Split(colorMap[c], "=")
	return strings.Split(comp[0], ":")
}

func getColorParanthesis(c color) (string, string) {
	comp := strings.Split(colorMap[c], "=")
	return comp[1], comp[2]
}

func replaceExp(text string) string {
	workingString := text

	// first lets take care of all the suffixes
	workingString = strings.ReplaceAll(workingString, "|!", suffixBreaker)
	workingString = strings.ReplaceAll(workingString, "|+", suffixBgBreaker)
	workingString = strings.ReplaceAll(workingString, "|>", suffixAttrBreaker)

	// lets deal with the prefixes
	for i, _ := range colorMap {
		brackets := getBrackets(color(i))
		pre, _ := getColorParanthesis(color(i))

		// optimization: if a prefix contains a + before them it denotes background color
		// we can continue without applying colors as foregrounds are first in our slice
		if strings.Contains(workingString, "+"+pre) || strings.Contains(workingString, "*"+pre) {
			continue
		}

		if strings.Contains(workingString, pre) {
			workingString = strings.Replace(workingString, pre, brackets[0], 1)
		}
	}
	return workingString
}
