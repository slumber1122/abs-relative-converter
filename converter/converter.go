package converter

import "fmt"

// Converter converter
type Converter struct {
	source string
	dest   string
}

// Run convert
func (c *Converter) Run() {
	fmt.Println("begin convert website: ", c.source)

}

// NewConverter new converter
func NewConverter(source string, dest string) *Converter {
	return &Converter{source: source, dest: dest}
}
