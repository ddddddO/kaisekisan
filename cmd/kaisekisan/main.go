package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ddddddO/kaisekisan"
)

func main() {
	if len(os.Args) != 2 {
		exit("required file")
	}

	in, err := filepath.Abs(os.Args[1])
	if err != nil {
		exit(err)
	}

	if isNotExist(in) {
		exit("no exist")
	}

	inFile, err := os.Open(in)
	if err != nil {
		exit(err)
	}
	defer inFile.Close()

	out := fmt.Sprintf("%s.out", in)
	outFile, err := os.Create(out)
	if err != nil {
		exit(err)
	}
	defer outFile.Close()

	if err := kaisekisan.Kaiseki(inFile, outFile); err != nil {
		if err := os.Remove(out); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		exit(err)
	}

	fmt.Fprintln(os.Stdin, "succeeded!")
}

func isNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func exit(a any) {
	fmt.Fprintf(os.Stderr, "%s\n", a)
	os.Exit(1)
}
