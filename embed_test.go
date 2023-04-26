package goembed

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"

	_ "embed"
)

//go:embed file.txt

var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}

//go:embed bg.jpg
var foto []byte

func TestEmbedByte(t *testing.T) {
	err := ioutil.WriteFile("fotoEbed.jpg", foto, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt

var files embed.FS

func TestEmbedMultipleFileString(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}
