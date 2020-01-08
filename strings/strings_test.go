package strings

import (
	"testing"
)

func TestCapitalize(t *testing.T) {
	test1 := Capitalize("abc")
	if test1 != "Abc" {
		t.Error("TestCapitalize abc must be Abc")
	}

	test2 := Capitalize("Abc")
	if test2 != "Abc" {
		t.Error("TestCapitalize abc must be Abc")
	}

	test3 := Capitalize("abcBH")
	if test3 != "AbcBH" {
		t.Error("TestCapitalize abcBH must be AbcBH")
	}

	test4 := Capitalize("")
	if test4 != "" {
		t.Error("TestCapitalize '' must be ''")
	}
}

func TestIsCapitalize(t *testing.T) {
	test1 := IsCapitalize("abc")
	if test1 {
		t.Error("IsCapitalize 'abc' must be false!")
	}

	test2 := IsCapitalize("Abc")
	if !test2 {
		t.Error("IsCapitalize 'Abc' must be true!")
	}

	test3 := IsCapitalize("ABC")
	if !test3 {
		t.Error("IsCapitalize 'ABC' must be true!")
	}

	test4 := IsCapitalize("")
	if test4 {
		t.Error("IsCapitalize '' must be false!")
	}
}
