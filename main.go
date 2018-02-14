package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/pixelbender/go-matroska/matroska"
)

var logger *log.Logger

func main() {
	logger = log.New(os.Stderr, "[webm-reader] ", 0)
	args := os.Args[1:]
	if len(args) == 0 {
		usage(nil)
	}

	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		usage(err)
	}

	doc, err := matroska.Decode(args[0])
	if err != nil {
		logger.Println(err)
		return
	}

	spew.Dump(doc)
}

func usage(err error) {
	logger.Println("USAGE: webm-reader FILENAME")
	if err != nil {
		logger.Println(err)
	}
	os.Exit(1)
}
