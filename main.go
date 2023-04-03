package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("logo_next.png", logo, fs.ModePerm) // proses copy file (nama file baru, file yang mau di copy, mode permission dari filesystemnya)
	if err != nil {
		panic(err)
	}

	dirEntries, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("------------------------")
			fmt.Println(entry.Name())
			content, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println("Content :", string(content))
		}
	}
}
