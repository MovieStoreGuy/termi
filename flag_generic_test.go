package termi_test

import (
	"testing"

	"github.com/MovieStoreGuy/termi"
)

func TestGenericFlag(t *testing.T) {
	var (
		uid uint = 8008
	)
	defer func() {
		if r := recover(); r != nil {
			t.Fatal(r)
		}
	}()
	f := termi.Must(termi.NewFlag(&uid)).
		SetName("uid").
		SetName("user").
		SetDescription("user id to passed")
	if err := f.Set("0078"); err != nil {
		t.Fatal(err)
	}
	if uid != uint(78) {
		t.Error("Values do not match")
		t.Errorf("%v vs %v\n", uid, uint(78))
	}
}

func TestUnsupportedValue(t *testing.T) {
	var (
		steve *struct{ blank string }
	)
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("Unsupported value type should throw an error")
		}
		t.Log(r)
	}()
	termi.Must(termi.NewFlag(steve))
}
