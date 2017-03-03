package main

import (
	"flag"
	"log"
	"os"

	tail "github.com/kaneshin/go-tail"
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(1)
	}
	filename := args[0]

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := tail.Exec(f); err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}
