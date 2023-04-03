package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// --------------------------------------------------------------------------------------------------------------------
// contoh embed ke string
//
//go:embed version.txt
var version_test string // variabelnya harus di luar function

func TestEmbedString(t *testing.T) {
	fmt.Println(version_test)
}

// --------------------------------------------------------------------------------------------------------------------
// contoh embed ke Byte[]
//
//go:embed logo.png
var logo_test []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("logo_next.png", logo_test, fs.ModePerm) // proses copy file (nama file baru, file yang mau di copy, mode permission dari filesystemnya)
	if err != nil {
		panic(err)
	}
}

// --------------------------------------------------------------------------------------------------------------------
// contoh embed multiple file
//
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
//go:embed files/d.txt
var fils embed.FS

func TestMultiplesFiles(t *testing.T) {
	a, err := fils.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))

	b, err := fils.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	c, err := fils.ReadFile("files/c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(c))

	d, err := fils.ReadFile("files/d.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
}

// --------------------------------------------------------------------------------------------------------------------
// contoh embed multiple file menggunakan path matcher
//
//go:embed files/*.txt
var path_test embed.FS

func TestMultiplesFilesUsingPathMatcher(t *testing.T) {
	dirEntries, err := path_test.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("------------------------")
			fmt.Println(entry.Name())
			content, err := path_test.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println("Content :", string(content))
		}
	}
}
