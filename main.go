package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/morinokami/jsorter/jsorter"
)

var reverse = flag.Bool("r", false, "reverse the result")

func run() int {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	result, err := jsorter.Sort(b, *reverse)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Println(result)

	return 0
}

func main() {
	flag.Parse()
	os.Exit(run())
}
