package termi_test

import (
	"testing"

	"github.com/MovieStoreGuy/termi"
)

func TestUsage(t *testing.T) {
	type Settings struct {
		User string `env:"USER" description:"Obtain the current user"`
		GUID string `env:"GUID"`
	}
	s := &Settings{}
	_, err := termi.EnvironmentDescription(s)
	if err != nil {
		t.Fatal(err)
	}
}
