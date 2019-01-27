package termi_test

import (
	"testing"

	"github.com/MovieStoreGuy/termi"
)

func TestInteger_Set(t *testing.T) {
	var (
		value   = 10
		integer = termi.NewInteger().SetValue(&value)
	)
	if value != 10 {
		t.Fatal("Value must remain unmodified until integer has parsed a string")
	}
	if err := integer.Set("100"); err != nil {
		t.Fatal("Unable to parse integer string", err)
	}
	if value != 100 {
		t.Fatal("The original value variable should be updated")
	}
}

func TestInteger_IsFlag(t *testing.T) {
	var (
		integer = termi.NewInteger().
			SetName("linus").
			SetName("--torvalds").
			SetDescription("Primary maintainer of linux kernel")
	)
	if !(integer.IsFlag("linus") && integer.IsFlag("torvalds")) {
		t.Fatal("Incorrectly validates the flag names")
	}
	t.Log("Current flag output: ", integer)
}

func TestIncorrectlyInitialised(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Had to recover from a dumb panic", r)
		}
	}()
	var i = termi.Integer{}
	i.SetName("test-help")
	if !i.IsFlag("test-help") {
		t.Fatal("Didn't keep the dash in the middle")
	}
}
