package sort

import (
	"fmt"
	"regexp"
	"strings"
)

type Charsets struct {
	options  []string
	filename string
}

func CharsetsNew(options []string, filename string) *Charsets {
	return &Charsets{
		options:  options,
		filename: filename,
	}
}

func (ch *Charsets) Manage() {
	for _, value := range ch.options {
		switch value {
		case "LowerEmphasis":
			ch.RemoveCharset("_")
		case "Dash":
			ch.RemoveCharset("-")
		case "Number":
			ch.RemoveNumber()
		case "OnlyLowerCase":
			ch.filename = strings.ToLower(ch.filename)
		}
	}
}

func (ch *Charsets) RemoveCharset(char string) {
	ch.filename = strings.ReplaceAll(ch.filename, char, "")
}

func (ch *Charsets) RemoveNumber() {
	re := regexp.MustCompile("[0-9]")
	for _, value := range ALLOWED_IMAGES {
		if re.ReplaceAllString(ch.filename, "") == value {
			fmt.Println("When you remove the numbers, only the extension will remain!")
			return
		}
	}
	ch.filename = re.ReplaceAllString(ch.filename, "")
}
