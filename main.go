package main

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ideahitme/define/dictionary"
)

var word string

func init() {
	kingpin.Arg("word", "Word to be defined").Required().StringVar(&word)
	kingpin.Parse()
}

func main() {
	dict := dictionary.Babla{}
	dict.Print(word)
}
