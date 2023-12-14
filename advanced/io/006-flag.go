package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print newLine")

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var str string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			str += " "
			if *NewLine {
				str += "\n"
			}
		}
		str += flag.Arg(i)
	}
	str += "\n"
	_, _ = os.Stdout.WriteString(str)
}
