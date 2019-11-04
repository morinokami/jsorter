package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/morinokami/jsorter/jsorter"
)

var (
	reverse = flag.Bool("r", false, "reverse the result")
	indent  = flag.Int("i", 2, "the number of spaces used for indentation")
)

func run() int {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	result, err := jsorter.Sort(b, *reverse, *indent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Println(string(result))

	return 0
}

func main() {
	flag.Parse()
	os.Exit(run())
}
