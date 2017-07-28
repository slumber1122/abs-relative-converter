package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Converter converter
type Converter struct {
	source string
	dest   string
}

// Run convert
func (c *Converter) Run() {
	fmt.Println("begin convert website: ", c.source)
	c.convert(0, c.source, c.dest)
}

func (c *Converter) genNewDestDir(dest string) {

}

func (c *Converter) genNewRelativeFile(depth int, path string, dest string) {
	ss := strings.Split(path, "/")
	destfile := filepath.Join(dest, ss[len(ss)-1])
	if !isFileExist(destfile) {
		newfile, err := os.Create(destfile)
		if err != nil {

		}

	}
}

func (c *Converter) convert(depth int, source string, dest string) {
	c.genNewDestDir(dest)
	walkFunc := func(source string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			dest = filepath.Join(dest, f.Name())
			c.convert(depth+1, source, dest)
		}
		if isHtmlFile(f.Name()) {
			c.genNewRelativeFile(depth, source, dest)
		}
		return nil
	}
	err := filepath.Walk(c.source, walkFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() return %v\n", err)
	}
}

func isHtmlFile(path string) bool {
	path = strings.ToLower(path)
	if strings.HasSuffix(path, "html") || strings.HasSuffix(path, "htm") {
		return true
	}
	return false
}

func isFileExist(path string) bool {
	return false
}

// NewConverter new converter
func NewConverter(source string, dest string) *Converter {
	return &Converter{source: source, dest: dest}
}
