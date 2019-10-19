package colorizer

import "github.com/fatih/color"

func Red(in string) string {
	r := color.New(color.FgRed).SprintFunc()
	return r(in)
}

func PrintYellow(in string) {
	color.Yellow("%s", in)
}

func PrintBlue(in string) {
	color.Blue("%s", in)
}
