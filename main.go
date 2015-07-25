package tvrenamer

import (
	"fmt"
	"github.com/nomis43/go-tvrenamer/tvdb"
	"regexp"
	"strings"
	"text/template"
)

const (
    regex string = "^(.+)[\.\ ][Ss]?(\d{2}|\d{1})[EeXx]?(\d{2}).*(\.\w{1,4})$"
)

func Rename(filepath string, nameFormatting string) (err error) {
	t := template.New("filename formatting")
	t, err = t.Parse(nameFormatting)
	if err != nil {
		return
	}
    return
}
