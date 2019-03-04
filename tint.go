package tint

import (
	"fmt"
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
type Color string

const (
	// Black color brackets in ansi format
	Black Color = "[test:test]"
)

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
func (t *Tint) Print(text string, color string) {
	brackets := strings.Split(color, ":")
	fmt.Print(brackets[0] + text + brackets[1])
}

// Println single line of text with enter character
func (t *Tint) Println(text string, color string) {
	brackets := strings.Split(color, ":")
	fmt.Println(brackets[0] + text + brackets[1])
}
