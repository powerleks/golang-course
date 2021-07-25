package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func read(path string) string {
	file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()


  b, err := ioutil.ReadAll(file)
  return string(b)
}

func write(path string, text string) {
	var writer io.Writer
	if (path == "") {
		writer = os.Stdout
	} else {
		fo, err := os.Create(path)
		writer = fo
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = fo.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	_, err := io.WriteString(writer, text)
	if err != nil {
		log.Fatal(err)
	}
}

func cat(args []string) {
	text1 := read(args[0])
	var text2 string
	outputPath := ""
	if (len(args) > 1) {
		text2 = read(args[1])
	}
	if (len(args) > 2) {
		outputPath = args[2]
	}
	outputText := strings.Join([]string{text1, text2}, "")
	write(outputPath, outputText)
}

func main() {
	args := os.Args[1:]
	cat(args)
}