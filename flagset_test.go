package termi_test

import (
	"os"
	"testing"

	"github.com/MovieStoreGuy/termi"
)

func TestFlagSet_Parse(t *testing.T) {
	var (
		creatureType = ""
		args         = []string{"hello", "--there", "human"}
		flag         = termi.NewString().SetDescription("defines creature type").SetName("there").SetValue(&creatureType)
		fs           = termi.NewFlagSet()
	)
	remainder, err := fs.Parse(args)
	if err != nil {
		t.Fatal("Experienced an issue parsing: ", err)
	}
	if len(remainder) != len(args) {
		t.Logf("%+v vs %+v", remainder, args)
		t.Fatal("Both arrays should be the same length")
	}
	remainder, err = fs.Register(flag).Parse(args)
	if len(remainder) != 1 {
		t.Logf("Remainder: %+v, creatureType: %s\n", remainder, creatureType)
		t.Fatal("Expected only one value to be remaining")
	}
	if creatureType != "human" {
		t.Fatal("Expected creatureType to be updated")
	}
}

func TestFlagSet_Description(t *testing.T) {
	const (
		description = `{{.name}}
GoVersion: {{.GoVersion}}

Environment
{{range $var, $description := .environments}} {{$var}} {{$description}}
{{end}}
Flags:
{{range $flag := .flags }} {{$flag}}
{{end}}
`
	)
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Recovered from: ", r)
		}
	}()
	fs := termi.NewFlagSet()
	fs.SetDescription(description)
	fs.Register(termi.NewString().SetName("Name").SetName("name").SetDescription("an example name"))
	if err := fs.PrintDescription(os.Stderr); err != nil {
		t.Fatal(err)
	}
}

func TestFlagSet_OptionalBooleans(t *testing.T) {
	var (
		enabled = false
		fs      = termi.NewFlagSet()
	)
	fs.Register(termi.NewBoolean().
		SetValue(&enabled).
		SetDescription("To turn on something").
		SetName("enable"))
	if _, err := fs.Parse([]string{"--enable"}); err != nil {
		t.Fatal("Should allow for option values to be parse but got ", err)
	}
	if !enabled {
		t.Fatal("Did not set enable to true")
	}
	if _, err := fs.Parse([]string{"--enable", "false"}); err != nil {
		t.Fatal("should allow for explicit setting of boolean variables but got ", err)
	}
	if enabled {
		t.Error("Did not set enable to false")
	}
	remainder, err := fs.Parse([]string{"--enable", "fortnight"})
	if err != nil {
		t.Fatal("Should not try consume fortnight as a value for a boolean flag, got ", err)
	}
	if len(remainder) > 1 && remainder[0] != "fortnight" {
		t.Error("Incorrect values returned, got ", remainder)
	}
}

func TestFlagSet_Terminator(t *testing.T) {
	var (
		fs = termi.NewFlagSet()
	)
	fs.Register(termi.NewBoolean().
		SetValue(new(bool)).
		SetName("help").
		SetDescription("shows help message"))
	remainder, err := fs.Parse([]string{"strop", "--"})
	if err != nil {
		t.Fatal("Experienced issue:", err)
	}
	if len(remainder) != 1 && remainder[0] != "strop" {
		t.Error("Incorrect level of arguments returned, got:", remainder)
	}
	remainder, err = fs.Parse([]string{"strop", "--", "--help"})
	if err != nil {
		t.Fatal("Experienced issue:", err)
	}
	if len(remainder) > 2 && remainder[1] != "--help" {
		t.Error("Processed args after the terminator, got: ", remainder)
	}
}
