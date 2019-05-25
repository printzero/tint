package tint

import (
	"strings"
	"testing"
)

var mod = Init()

func TestInit(t *testing.T) {
	if !tintTypeCheck(*mod) {
		t.Errorf("Init() not returning Tint struct %v", mod)
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
	brackets := strings.Split(colorMap[Green], ":")
	greenText := mod.Raw("Yes", Green)

	if !strings.HasPrefix(greenText, brackets[0]) {
		t.Error("Raw color test failed because it is not prefixed with green color from color map")
	}

	if !strings.HasSuffix(greenText, brackets[1]) {
		t.Error("Raw color test failed because it is not suffixed with green color from color map")
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
