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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gen-slice-accessors",
	Short: "Generate accessors for each field in the slice struct.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		// Load arguments
		if err := internal.Args.Load(); err != nil {
			return fmt.Errorf("loader error: %w", err)
		}

		// Parse source code
		generator, err := internal.Parse(internal.Args, internal.Reader)
		if err != nil {
			return fmt.Errorf("parse error: %w", err)
		}

		// Generate code
		txt, err := generator.Generate()
		if err != nil {
			return fmt.Errorf("generate error: %w", err)
		}

		// Write to output file
		err = internal.Write(internal.Args.Output, txt)
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
	rootCmd.Flags().StringVarP(&internal.Args.Entity, "entity", "e", "", "[required] Target entity name")
	_ = rootCmd.MarkFlagRequired("entity")

	// slice
	rootCmd.Flags().StringVarP(&internal.Args.Slice, "slice", "s", "", "[required] Target slice name")
	_ = rootCmd.MarkFlagRequired("slice")

	// input
	rootCmd.Flags().StringVarP(&internal.Args.Input, "input", "i", "", "[required] Input file name")
	_ = rootCmd.MarkFlagRequired("input")

	// output
	rootCmd.Flags().StringVarP(&internal.Args.Output, "output", "o", "", "[required] Output file name")
	_ = rootCmd.MarkFlagRequired("output")

	// fieldNamesToExclude
	rootCmd.Flags().StringSliceVarP(&internal.Args.FieldNamesToExclude, "exclude", "x", []string{}, "Field names to exclude")

	// rename
	rootCmd.Flags().StringSliceVarP(&internal.Args.RawRenames, "rename", "r", []string{}, "Rename accessor name \n e.g. --rename=Name:GetName")

	// import
	rootCmd.Flags().StringSliceVarP(&internal.Args.RawImportPaths, "import", "m", []string{}, "Import path name \n e.g. --import=time \n e.g. --import=time:aliasTime")
}
