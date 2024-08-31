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
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gen-slice-accessor",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run")
		fmt.Println("args", entity, slice, fieldNamesToExclude, input, output)

		fmt.Println(args)
	},
	// Args: func(cmd *cobra.Command, args []string) error {

	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	// Target entity name
	entity string

	// Target slice name
	slice string

	// Field names to exclude
	fieldNamesToExclude []string

	// Input file name
	input string

	// Output file name
	output string
)

func init() {
	// entity
	_ = rootCmd.MarkFlagRequired("entity")
	rootCmd.Flags().StringVarP(&entity, "entity", "e", "", "target entity name")

	// slice
	_ = rootCmd.MarkFlagRequired("slice")
	rootCmd.Flags().StringVarP(&slice, "slice", "s", "", "target slice name")

	// input
	_ = rootCmd.MarkFlagRequired("input")
	rootCmd.Flags().StringVarP(&input, "input", "i", "", "input file name")

	// output
	_ = rootCmd.MarkFlagRequired("output")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "output file name")

	// fieldNamesToExclude
	rootCmd.Flags().StringSliceVarP(&fieldNamesToExclude, "exclude", "x", []string{}, "field names to exclude")
}
