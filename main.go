package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/slumber1122/abs-relative-converter/converter"
)

func main() {
	source := *flag.String("s", "", "the website diretory which you want to convert")
	dest := *flag.String("d", "", "the converted website diretory")
	flag.Parse()
	fmt.Println("source: ", source, dest)
	if source == "" || dest == "" {
		fmt.Println("You need supply website source diretory and dest diretory")
		return
	}
	source = convertToAbsPath(source)
	dest = convertToAbsPath(dest)

	converter := converter.NewConverter(source, dest)
	converter.Run()
}

func convertToAbsPath(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}
	execDirAbsPath, _ := os.Getwd()
	return filepath.Join(execDirAbsPath, path)
}
