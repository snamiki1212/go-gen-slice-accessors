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

	"github.com/snamiki1212/go-gen-slice-accessors/internal"
	"github.com/spf13/cobra"
)

type importPath struct {
	path  string
	alias string
}

// Import path name
var importPaths []string
var renames []string

// arguments
var args = Arguments{
	renames:     map[string]string{},
	importPaths: make([]importPath, 0),
}

// GenerateImportPath
func GenerateImportPath(importPaths []importPath) string {
	if len(importPaths) == 0 {
		return ""
	}

	var txt string
	for _, elem := range importPaths {
		switch elem.alias {
		case "": // no alias
			txt += fmt.Sprintf("	\"%s\"\n", elem.path)
		default:
			txt += fmt.Sprintf("	%s \"%s\"\n", elem.alias, elem.path)
		}
	}
	return "\nimport (\n" + txt + ")\n"
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
		txt, err := generate(data, args)
		if err != nil {
			return fmt.Errorf("generate error: %w", err)
		}

		// Write to output file
		err = internal.Write(args.output, txt)
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
	rootCmd.Flags().StringVarP(&args.entity, "entity", "e", "", "[required] Target entity name")
	_ = rootCmd.MarkFlagRequired("entity")

	// slice
	rootCmd.Flags().StringVarP(&args.slice, "slice", "s", "", "[required] Target slice name")
	_ = rootCmd.MarkFlagRequired("slice")

	// input
	rootCmd.Flags().StringVarP(&args.input, "input", "i", "", "[required] Input file name")
	_ = rootCmd.MarkFlagRequired("input")

	// output
	rootCmd.Flags().StringVarP(&args.output, "output", "o", "", "[required] Output file name")
	_ = rootCmd.MarkFlagRequired("output")

	// fieldNamesToExclude
	rootCmd.Flags().StringSliceVarP(&args.fieldNamesToExclude, "exclude", "x", []string{}, "Field names to exclude")

	// rename
	rootCmd.Flags().StringSliceVarP(&renames, "rename", "r", []string{}, "Rename accessor name \n e.g. --rename=Name:GetName")

	// import
	rootCmd.Flags().StringSliceVarP(&importPaths, "import", "m", []string{}, "Import path name \n e.g. --import=time \n e.g. --import=time:aliasTime")
}
