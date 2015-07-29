package config_test

import (
	"github.com/nomis43/go-tvrenamer/config"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Args = append(os.Args, "-l=fr", "-r=\"^.+[\\.]$\"", "$HOME/Media")
	f, err := ioutil.TempFile("", "testConf")
	if err != nil {
		t.Error("Unable to create temp file!")
	}

	f.Write([]byte(`[Main]
        Language=cn
        NameFormatting={{format}}
        NewPath=$HOME`))

	c := config.Load(f.Name())

	if c.Language != "fr" {
		t.Errorf("Wrong Language: %v should be %v", c.Language, "fr")
	}
	if c.NameFormatting != "{{format}}" {
		t.Errorf("Wrong NameFormatting: %v should be %v", c.NameFormatting, "{{format}}")
	}
	if c.Path != "$HOME/Media" {
		t.Errorf("Wrong Path: %v should be %v\n", c.Path, "$HOME/Media")
	}
	if c.NewPath != "$HOME" {
		t.Errorf("Wrong NewPath: %v should be %v\n", c.NewPath, "$HOME")
	}
}
