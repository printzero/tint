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

const (
	// Black color brackets in ansi format
	Black color = iota
	// Cyan color
	Cyan
	// BgWhite background
	BgWhite
)

var colorMap = map[color]string{
	// Black
	0: "\u001B[30m:\u001B[39m",
	// Cyan
	1: "\u001B[36m:\u001B[39m",
	2: "\u001B[47m:\u001B[49m",
}

const (
	// None for terminal that supports no colot
	None TerminalLevel = iota + 1
	// Color16bit for terminal that supports 16bit colors
	Color16bit
	// Color256 for terminal that supports 256 bit colors
	Color256
	// Color16m for terminal that supports 16 million colors (truecolor)
	Color16m
)

// Init initializes variables that tint uses and then returns the
// pointer to a Tint struct
func Init() *Tint {
	return &Tint{
		Level:         None,
		SupportsColor: false,
	}
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

func apply(text string, colors []color) string {
	output := text
	for _, c := range colors {
		brackets := strings.Split(colorMap[c], ":")
		output = brackets[0] + output + brackets[1]
	}
	return output
}
