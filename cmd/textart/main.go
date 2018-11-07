package main

import (
	"io"
	"os"
	"strings"

	"github.com/wzshiming/textart"
)

func main() {
	str := strings.Join(os.Args[1:], " ")
	result := textart.Gen(str)
	io.WriteString(os.Stdout, result)
}
