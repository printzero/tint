/*
 *    Copyright 2019 Ashish Shekar a.k.a codekidX
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
// Package tint is a minimal bare-bone version of terminal styling implemented for
// Go applications with no external dependencies.
// It provides you with different types of functions that you can use to style your terminal output with ease.
//
// Tint was originally created to use in the 'orbpkg' project: https://github.com/orbpkg/orb and uses near to 0ms
// for processing color expressions: https://godoc.org/github.com/printzero/tint/#Tint.Exp. Although the time taken to process is directly propotional to the number
// of characters in the string input.

// Package tint is used to color your outputs to the standard streams
// easily and effortlessly
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

// TerminalLevel of color support for terminal and information of
// current terminal that is useful to tint.
type TerminalLevel int

// SwatchFunc is what a tint.Swatch returns after adding color sequences
type SwatchFunc func(text string)

// SwatchFuncRaw is what a tint.SwatchRaw returns which inturn returns a
// colored string output. These methods can be used as higher-order functions
// for coloring
type SwatchFuncRaw func(text string) string

// Color is a struct that holds the innate escape characters for Color variables.
type Color struct {
	open  string // opening escape character for a given color
	close string // closing escape character for a given color
}

var escSeq = "\033"

// Dim dims the foreground of terminal text
func (c Color) Dim() Color {
	return Color{
		open:  escSeq + "[2m" + c.open,
		close: c.close + escSeq + "[0m",
	}
}

// Bold applies bold style to the color you are applying to
func (c Color) Bold() Color {
	return Color{
		open:  escSeq + "[1m" + c.open,
		close: c.close + escSeq + "[0m",
	}
}

// Underline renders an underline with the target text
func (c Color) Underline() Color {
	return Color{
		open:  escSeq + "[4m" + c.open,
		close: c.close + escSeq + "[0m",
	}
}

// Italic applies italic style to your target text
func (c Color) Italic() Color {
	return Color{
		open:  escSeq + "[3m" + c.open,
		close: c.close + escSeq + "[0m",
	}
}

// Strike strikes out the target text
func (c Color) Strike() Color {
	return Color{
		open:  escSeq + "[9m" + c.open,
		close: c.close + escSeq + "[0m",
	}
}

// Add concats the color with the base color which
// can be used to apply more than one colors.
// Example: background and foreground
func (c Color) Add(this Color) Color {
	return Color{
		open:  c.open + this.open,
		close: this.close + c.close,
	}
}

// Normal equates to no style
var Normal = Color{
	open:  escSeq + "[0m",
	close: escSeq + "[0m",
}

// Black color
var Black = Color{
	open:  escSeq + "[30m",
	close: escSeq + "[39m",
}

// Red color
var Red = Color{
	open:  escSeq + "[31m",
	close: escSeq + "[39m",
}

// Green color
var Green = Color{
	open:  escSeq + "[32m",
	close: escSeq + "[39m",
}

// Yellow color
var Yellow = Color{
	open:  escSeq + "[33m",
	close: escSeq + "[39m",
}

// Blue color
var Blue = Color{
	open:  escSeq + "[34m",
	close: escSeq + "[39m",
}

// Magenta color
var Magenta = Color{
	open:  escSeq + "[35m",
	close: escSeq + "[39m",
}

// Cyan color
var Cyan = Color{
	open:  escSeq + "[36m",
	close: escSeq + "[39m",
}

// White color
var White = Color{
	open:  escSeq + "[37m",
	close: escSeq + "[39m",
}

// BgBlack applies Black Background color
var BgBlack = Color{
	open:  escSeq + "[40m",
	close: escSeq + "[49m",
}

// BgRed applies Red Background color
var BgRed = Color{
	open:  escSeq + "[41m",
	close: escSeq + "[49m",
}

// BgGreen applies Green Background color
var BgGreen = Color{
	open:  escSeq + "[42m",
	close: escSeq + "[49m",
}

// BgYellow applies Yellow Background color
var BgYellow = Color{
	open:  escSeq + "[43m",
	close: escSeq + "[49m",
}

// BgBlue applies Blue Background color
var BgBlue = Color{
	open:  escSeq + "[44m",
	close: escSeq + "[49m",
}

// BgMagenta applies Magenta Background color
var BgMagenta = Color{
	open:  escSeq + "[45m",
	close: escSeq + "[49m",
}

// BgCyan applies Cyan Background color
var BgCyan = Color{
	open:  escSeq + "[46m",
	close: escSeq + "[49m",
}

// BgLightGrey applies Light Grey Background color
var BgLightGrey = Color{
	open:  escSeq + "[47m",
	close: escSeq + "[49m",
}

// BgWhite applies White Background color
var BgWhite = Color{
	open:  escSeq + "[107m",
	close: escSeq + "[49m",
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

// Exp returns a string constructed from a series of color expressions given as an argument.
// The colors are passed as a replacement to each word that is wrapped around `@()`.
//
// The string "@(Hello), @(World)!" where 'Hello' is inside a tint color expression and 'World' inside another,
// 'Hello' will get replaced by the first color and 'World' will get replaced by the second color passed inside
// this function.
//
// Take a look at the below example.
func (t *Tint) Exp(expStr string, colors ...Color) string {
	return replaceExp(expStr, colors)
}

// Expf does to Expression replacement with formatting
func (t *Tint) Expf(expStr string, colors []Color, args ...interface{}) string {
	str := replaceExp(expStr, colors)
	return fmt.Sprintf(str, args)
}

// Raw returns the raw string with applied colors
func (t *Tint) Raw(text string, colors ...Color) string {
	return apply(text, colors)
}

// Print single line of text with specified color
func (t *Tint) Print(text string, colors ...Color) {
	fmt.Print(apply(text, colors))
}

// Println single line of text with enter character
func (t *Tint) Println(text string, colors ...Color) {
	fmt.Println(apply(text, colors))
}

// Log text with the standard lib log module
func (t *Tint) Log(text string, colors ...Color) {
	t.LogInstance.Println(apply(text, colors))
}

// Swatch will return a function for specific colors given as a parameter.
func (t *Tint) Swatch(colors ...Color) SwatchFunc {
	return func(text string) {
		t.Println(text, colors...)
	}
}

// SwatchRaw returns a functions that returns a raw colored string
func (t *Tint) SwatchRaw(colors ...Color) SwatchFuncRaw {
	return func(text string) string {
		return t.Raw(text, colors...)
	}
}

// apply the colors to the text by wrapping text with ANSI styling
func apply(text string, colors []Color) string {
	output := text
	for _, c := range colors {
		output = c.open + output + c.close
	}
	return output
}

func replaceExp(text string, colors []Color) string {
	if tc, proceed := isAtCountSame(text, len(colors)); !proceed {
		msg := "mismatching apply - Trigger count: %d, Color count: %d \n\n%sTip%s: " +
			"Does your function have colors passed as much as it has '@' character?"
		panic(fmt.Errorf(msg, tc, len(colors), Magenta.open, Magenta.close))
	}

	workingString := ""
	workingColors := colors

	// this is used as the current color on operation
	//var colorStart int
	//var colorEnd = -1
	var colorBox = Normal
	var triggered = false

	// so here is what we'll do --- we'll find the first @( which serves as opening
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
			colorBox = Normal
			workingColors = workingColors[1:]
			continue
		}

		workingString = workingString + string(c)
	}

	return workingString
}

func isAtCountSame(text string, colorCount int) (int, bool) {
	counter := 0
	for i, c := range text {
		if string(c) == "@" && string(text[i+1]) == "(" {
			counter++
		}
	}

	return counter, counter == colorCount
}
