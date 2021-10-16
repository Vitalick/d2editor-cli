package main

import (
	"flag"
	"fmt"
	"github.com/vitalick/go-d2editor"
	"os"
)

type flags struct {
	toJSON  bool
	toD2s   bool
	empty   bool
	version uint
	output  string
	input   []string
}

var parsedFlags flags

func init() {
	parsedFlags.version = 97
}

func main() {
	parseArgs()

	if parsedFlags.toJSON {
		for _, f := range parsedFlags.input {
			c := &d2editor.Character{}
			var err error
			if parsedFlags.empty {
				c, err = d2editor.NewEmptyCharacter(parsedFlags.version)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error \"%v\" on create character \"%s\", skipped \n\n", err, f)
					continue
				}
				c.Name = f
			} else {
				c, err = d2editor.Open(f)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error \"%v\" on open file \"%s\", skipped \n\n", err, f)
					continue
				}
			}
			c.Header.Version = uint32(parsedFlags.version)
			err = d2editor.SaveJSON(c, parsedFlags.output)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error \"%v\" on save character \"%s\", skipped \n\n", err, c.Name)
				continue
			}
		}
	}

	if parsedFlags.toD2s {
		for _, f := range parsedFlags.input {
			c := &d2editor.Character{}
			var err error
			if parsedFlags.empty {
				c, err = d2editor.NewEmptyCharacter(parsedFlags.version)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error \"%v\" on create character \"%s\", skipped \n\n", err, f)
					continue
				}
				c.Name = f
			} else {
				c, err = d2editor.OpenJSON(f)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error \"%v\" on open file \"%s\", skipped \n\n", err, f)
					continue
				}
			}
			c.Header.Version = uint32(parsedFlags.version)
			err = d2editor.Save(c, parsedFlags.output)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error \"%v\" on save character \"%s\", skipped \n\n", err, c.Name)
				continue
			}
		}
	}
}

func parseArgs() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <input files/names>\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&parsedFlags.empty, "e", parsedFlags.empty, "Create new empty character.")
	flag.BoolVar(&parsedFlags.toJSON, "tojson", parsedFlags.toJSON, "Convert d2s to json or create json.")
	flag.BoolVar(&parsedFlags.toD2s, "tod2s", parsedFlags.toD2s, "Convert json to d2s or create d2s.")
	flag.StringVar(&parsedFlags.output, "o", parsedFlags.output, "Optional path of the output folder.")
	flag.UintVar(&parsedFlags.version, "v", parsedFlags.version, "Save with specific version (default: 97).")
	flag.Parse()

	// Make sure we have input paths.
	if flag.NArg() == 0 {
		_, err := fmt.Fprintf(os.Stderr, "Missing <input files/names>\n\n")
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
