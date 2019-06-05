package tint

import (
	"reflect"
	"strings"
	"testing"
)

var mod = Init()

func TestInit(t *testing.T) {
	if !tintTypeCheck(*mod) {
		t.Errorf("Init() not returning Tint struct %v", mod)
	}
}

func TestLoggerInstance(t *testing.T) {
	if mod.LogInstance == nil {
		t.Error("Log instance is nil on module init.")
	}
}

// In general test.Raw uses apply() which is the core function of tint module
// So in theory if we extensively test Raw() versions of all the functions
// we will gain the same results
// for all the printing functions like Println, Swatch etc ....
func TestRaw(t *testing.T) {
	sameTest := mod.Raw("Hello world")
	if sameTest != "Hello world" {
		t.Error("Raw with normal text color does not return back same string")
	}
}

func TestRawColor(t *testing.T) {
	brackets := getBrackets(Green)
	greenText := mod.Raw("Yes", Green)
	if !strings.HasPrefix(greenText, brackets[0]) {
		t.Error("Raw color test failed because it is not prefixed with green color from color map")
	}

	if !strings.HasSuffix(greenText, brackets[1]) {
		t.Error("Raw color test failed because it is not suffixed with green color from color map")
	}
}

// TestRawFandB tests foreground and background colors
func TestRawFandB(t *testing.T) {
	brackets := getBrackets(Green)
	bracketsFore := getBrackets(BgWhite)

	greenText := mod.Raw("Yes", Green, BgWhite)

	if !strings.HasPrefix(greenText, bracketsFore[0]+brackets[0]) {
		t.Error("Foreground and Background tests failed order of prefix not correct.")
	}

	if !strings.HasSuffix(greenText, brackets[1]+bracketsFore[1]) {
		t.Error("Foreground and Background tests failed order of suffix not correct.")
	}
}

func TestWith(t *testing.T) {
	m := mod.With("Yes", White)

	if !mixinTypeCheck(m) {
		t.Error("With() did not return a Mixin struct.")
	}
}

func TestPalette(t *testing.T) {
	tstr := mod.Palette(mod.With("Ashish", Blue))
	brackets := getBrackets(Blue)

	if !strings.Contains(tstr, brackets[0]) || !strings.Contains(tstr, brackets[0]) {
		t.Error("The color 'Blue' is not applied by the Palette properly.")
	}
}

func TestSwatch(t *testing.T) {
	yellowSwatchFunc := mod.Swatch(Yellow)

	swatchType := reflect.TypeOf(yellowSwatchFunc).Kind()

	if swatchType != reflect.Func {
		t.Error("Swatch did not return function.")
	}
}

func TestSwatchRawToReturnFunc(t *testing.T) {
	yellowSwatchFunc := mod.SwatchRaw(Yellow)

	swatchType := reflect.TypeOf(yellowSwatchFunc).Kind()

	if swatchType != reflect.Func {
		t.Error("SwatchRaw did not return function.")
	}
}

func TestSwatchRaw(t *testing.T) {
	brackets := getBrackets(Yellow)

	yellowSwatchFunc := mod.SwatchRaw(Yellow)

	result := yellowSwatchFunc("Yes")

	if !strings.HasPrefix(result, brackets[0]) {
		t.Error("Swatch did not have prefix of Yellow color", result)
	}

	if !strings.HasSuffix(result, brackets[1]) {
		t.Error("Swatch did not have suffix of Yellow color", result)
	}
}

func TestExp(t *testing.T) {
	estr := mod.Exp("y|Ashish|!")
	wrongEstr := mod.Exp("y|Ashish|")
	brackets := getBrackets(Yellow)
	if !strings.Contains(estr, brackets[0]) || !strings.Contains(estr, brackets[1]) {
		t.Error("Yellow color is not applied properly by tint.Expr(). Both brackets must be present.")
	}

	if strings.Contains(wrongEstr, suffixBreaker) {
		t.Errorf("When specifying an errored suffix of tint expression, this should not be present for (%s)", wrongEstr)
	}
}

func tintTypeCheck(value interface{}) bool {
	switch value.(type) {
	case Tint:
		return true
	default:
		return false
	}
}

func mixinTypeCheck(value interface{}) bool {
	switch value.(type) {
	case Mixin:
		return true
	default:
		return false
	}
}
