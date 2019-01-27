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
