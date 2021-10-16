package main

import (
	"flag"
	"fmt"
	"github.com/vitalick/go-d2editor"
	"os"
)

type flags struct {
	d2sToJSON bool
	jsonToD2s bool
	output    string
	input     []string
}

var parsedFlags flags

func main() {
	parseArgs()

	if parsedFlags.d2sToJSON {
		for _, f := range parsedFlags.input {
			c, err := d2editor.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error \"%v\" on open file \"%s\", skipped \n\n", err, f)
				continue
			}
			err = d2editor.SaveJSON(c, parsedFlags.output)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error \"%v\" on save file \"%s\", skipped \n\n", err, f)
				continue
			}
		}
	}

	if parsedFlags.jsonToD2s {
		for _, f := range parsedFlags.input {
			c, err := d2editor.OpenJSON(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error \"%v\" on open file \"%s\", skipped \n\n", err, f)
				continue
			}
			err = d2editor.Save(c, parsedFlags.output)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error \"%v\" on save file \"%s\", skipped \n\n", err, f)
				continue
			}
		}
	}
}

func parseArgs() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <input files>\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&parsedFlags.d2sToJSON, "tojson", parsedFlags.d2sToJSON, "Convert d2s to json.")
	flag.BoolVar(&parsedFlags.jsonToD2s, "fromjson", parsedFlags.jsonToD2s, "Convert json to d2s.")
	flag.StringVar(&parsedFlags.output, "o", parsedFlags.output, "Optional path of the output folder.")

	// Make sure we have input paths.
	if flag.NArg() == 0 {
		_, err := fmt.Fprintf(os.Stderr, "Missing <input files>\n\n")
		if err != nil {
			return
		}
		flag.Usage()
		os.Exit(1)
	}

	// Create input configurations.
	parsedFlags.input = make([]string, flag.NArg())
	for i := range parsedFlags.input {
		parsedFlags.input[i] = flag.Arg(i)
	}
}
