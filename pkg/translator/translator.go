package translator

import (
	"errors"
	"fmt"
	"github.com/RadikChernyshov/klingon/pkg/translator/dictionary"
	"strings"
)

type Transliterator struct {
	In  string
	Out string
}

func (t *Transliterator) ToKlingon() (err error) {
	if t.In == "" {
		return errors.New("blank strings not accepted")
	}
	var charsEncoded []string
	chars := strings.SplitAfter(t.In, "")
	charsLen := len(chars)
	dict := dictionary.Klingon()
	charIndex := 0
	for {
		curChar := chars[charIndex]
		nextCharIndex := charIndex + 1
		afterNextCharIndex := charIndex + 2

		if (charsLen > nextCharIndex) && (gLetter(curChar) || ngLetter(curChar, chars[nextCharIndex])) {
			curChar = curChar + chars[nextCharIndex]
			charIndex = nextCharIndex
		} else if (charsLen > afterNextCharIndex) && tlhLetter(curChar, chars[nextCharIndex], chars[afterNextCharIndex]) {
			curChar = curChar + chars[nextCharIndex] + chars[afterNextCharIndex]
			charIndex = afterNextCharIndex
		}
		if curChar != "q" {
			curChar = strings.ToUpper(curChar)
		}
		if code, exists := dict[curChar]; exists {
			charsEncoded = append(charsEncoded, code)
		} else {
			return fmt.Errorf("%s character is not translatable", curChar)
		}
		charIndex++
		if charIndex >= charsLen {
			break
		}
	}
	t.Out = strings.Join(charsEncoded[:], " ")
	return err
}

func tlhLetter(in string, in1 string, in2 string) bool {
	return strings.ToUpper(in+in1+in2) == dictionary.CapitalTlh
}

func ngLetter(in string, in1 string) bool {
	return strings.ToUpper(in+in1) == dictionary.CapitalNg
}

func gLetter(in string) bool {
	in = strings.ToUpper(in)
	return in == "C" || in == "G"
}

func New() *Transliterator {
	return &Transliterator{
		In: "",
	}
}
