package main

import (
	"flag"
	"fmt"

	"github.com/slumber1122/abs-relative-converter/converter"
)

func main() {
	source := flag.String("source", "", "the website diretory which you want to convert")
	dest := flag.String("dest", "", "the converted website diretory")
	if *source == "" || *dest == "" {
		fmt.Println("You need supply website source diretory and dest diretory")
		return
	}
	converter := converter.NewConverter(*source, *dest)
	converter.Run()
}
