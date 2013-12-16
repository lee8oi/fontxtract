/* fontxtract project main.go

fontxtract is a simple commandline utility which parses a directory of zip
archives and extracts TTF & OTF font files to a specified directory.

USAGE

Linux:

	fontxtract path-to-zips/ path-to-fonts/

Windows:

	fontxtract.exe path-to-zips\ path-to-fonts\

*/

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func extract(file, dest string) {
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	for _, f := range r.File {
		name := f.Name
		split := strings.Split(name, ".")
		ext := strings.ToLower(split[len(split)-1])
		if ext == "ttf" || ext == "otf" {
			fmt.Printf("extracting font:%s ", name)
			src, err := f.Open()
			if err != nil {
				fmt.Println(err)
				return
			}
			defer src.Close()
			dst, err := os.Create(dest + name)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer dst.Close()
			written, err := io.Copy(dst, src)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%d bytes.\n", written)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: fontextract path-to-zips/ path-to-fonts/\n")
		os.Exit(1)
	}
	filepath.Walk(os.Args[1], func(path string, f os.FileInfo, err error) error {
		name := f.Name()
		splitup := strings.Split(name, ".")
		if f.IsDir() {
			return nil
		}
		if splitup[len(splitup)-1] == "zip" {
			fmt.Printf("Archive: %s\n", name)
			extract(os.Args[1]+"/"+name, os.Args[2])
		}
		return nil
	})
}
