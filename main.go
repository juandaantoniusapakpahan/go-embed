package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed file.txt
var version string

//go:embed bg.jpg
var foto []byte

//go:embed files/*.txt
var folder embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("fotoEbed.jpg", foto, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	defer fmt.Println("Success")

	fileNames, _ := folder.ReadDir("files")

	for _, entry := range fileNames {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := folder.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}

}
