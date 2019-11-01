package main

import (
	"flag"
	"os"
)

var reverse = flag.Bool("r", false, "reverse the result")

func run() int {
	return 0
}

func main() {
	flag.Parse()
	os.Exit(run())
}
