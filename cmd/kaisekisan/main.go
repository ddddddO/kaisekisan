package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ddddddO/kaisekisan"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	if len(os.Args) != 3 {
		exit("Required file and target column number.")
	}

	columnNumber, err := strconv.Atoi(os.Args[2])
	if err != nil {
		exit(err)
	}
	if columnNumber <= 0 {
		exit("Please specify 1 or more.")
	}

	in, err := filepath.Abs(os.Args[1])
	if err != nil {
		exit(err)
	}
	if isNotExist(in) {
		exit("File does not exist.")
	}
	inFile, err := os.Open(in)
	if err != nil {
		exit(err)
	}
	defer inFile.Close()

	out := fmt.Sprintf("%s_out.csv", strings.TrimSuffix(in, ".csv"))
	outFile, err := os.Create(out)
	if err != nil {
		exit(err)
	}
	defer outFile.Close()

	csvReader := csv.NewReader(inFile)
	if err := kaisekisan.Kaiseki(csvReader, outFile, columnNumber); err != nil {
		if err := os.Remove(out); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		exit(err)
	}

	fmt.Fprintf(os.Stdout, "Succeeded! Destination -> %s\n", out)
}

func isNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func exit(a any) {
	fmt.Fprintf(os.Stderr, "%s\n", a)
	os.Exit(1)
}
