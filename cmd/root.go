/*
Copyright © 2024 snamiki1212

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type arguments struct {
	// Target entity name
	entity string

	// Target slice name
	slice string

	// Input file name
	input string

	// Output file name
	output string

	// Field names to exclude
	fieldNamesToExclude []string

	// Mapping field name to accessor name
	accessors map[string]string // key: field name, value: acccessor name.
}

var accessors []string

// arguments
var args = arguments{
	accessors: map[string]string{},
}

func (a *arguments) loadAccessors(as []string) error {
	container := make([]error, 0)
	for _, ac := range as {
		pair := strings.Split(ac, ":")
		if len(pair) != 2 {
			container = append(container, fmt.Errorf("invalid accessor: %s", ac))
			continue
		}
		field, accessor := pair[0], pair[1]
		args.accessors[field] = accessor
	}
	if len(container) != 0 {
		return fmt.Errorf("%v", container)
	}
	return nil
}

func loader() error {
	// Load arguments
	if err := args.loadAccessors(accessors); err != nil {
		return fmt.Errorf("load accessor error: %w", err)
	}
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gen-slice-accessors",
	Short: "Generate accessors for each field in the slice struct.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		// Load arguments
		if err := loader(); err != nil {
			return fmt.Errorf("loader error: %w", err)
		}

		// Parse source code
		data, err := parse(args, reader)
		if err != nil {
			return fmt.Errorf("parse error: %w", err)
		}

		// Generate code
		txt, err := generate(data)
		if err != nil {
			return fmt.Errorf("generate error: %w", err)
		}

		// Write to output file
		err = write(args.output, txt)
		if err != nil {
			return fmt.Errorf("write error: %w", err)
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	// entity
	rootCmd.Flags().StringVarP(&args.entity, "entity", "e", "", "target entity name")
	_ = rootCmd.MarkFlagRequired("entity")

	// slice
	rootCmd.Flags().StringVarP(&args.slice, "slice", "s", "", "target slice name")
	_ = rootCmd.MarkFlagRequired("slice")

	// input
	rootCmd.Flags().StringVarP(&args.input, "input", "i", "", "input file name")
	_ = rootCmd.MarkFlagRequired("input")

	// output
	rootCmd.Flags().StringVarP(&args.output, "output", "o", "", "output file name")
	_ = rootCmd.MarkFlagRequired("output")

	// fieldNamesToExclude
	rootCmd.Flags().StringSliceVarP(&args.fieldNamesToExclude, "exclude", "x", []string{}, "field names to exclude")

	// accessor
	rootCmd.Flags().StringSliceVarP(&accessors, "accessor", "a", []string{}, "accessor name for field / e.g. --accessor=Name:GetName")
}
