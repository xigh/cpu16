package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		check(flag.Arg(i))
	}
}

func check(name string) {
	r, ok := fmap[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "%s not found\n", name)
		return
	}

	fmt.Printf("%s: %s\n", name, r)
}
