package tint_test

import (
	"testing"
	"tint"
)

func TestInit(t *testing.T) {
	mod := tint.Init()
	if !tintTypeCheck(*mod) {
		t.Errorf("Init() not returning Tint struct %v", mod)
	}
}

func tintTypeCheck(value interface{}) bool {
	switch value.(type) {
	case tint.Tint:
		return true
	default:
		return false
	}
}

func TestRaw(t *testing.T) {
	mod := tint.Init()
	sameTest := mod.Raw("Hello world")
	if sameTest != "Hello world" {
		t.Error("Raw with normal text color does not return back same string")
	}
}
