/*
Copyright © 2022 Tony West

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/un4gi/adobecc-font-converter/convert"
)

var (
	fileFormat string
	outFolder  string
)

var banner string = `
                                                       
.8.          8 8888888888       ,o888888o.    
.888.         8 8888            8888     '88.  
:88888.        8 8888         ,8 8888       '8. 
. '88888.       8 8888         88 8888           
.8. '88888.      8 888888888888 88 8888           
.8'8. '88888.     8 8888         88 8888           
.8' '8. '88888.    8 8888         88 8888           
.8'   '8. '88888.   8 8888         '8 8888       .8' 
.888888888. '88888.  8 8888            8888     ,88'  
.8'       '8. '88888. 8 8888             '8888888P'    
  A d o b e  C C      F o n t      C o n v e r t e r
`
var platform string = runtime.GOOS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "afc",
	Short: "Convert your Adobe CC fonts to various font file types.",
	Long: banner + `
Adobe CC Font Converter is a simple CLI tool designed to help convert 
licensed and activated Adobe CC fonts into proper font files.

When you activate fonts in Adobe CC, the font files are not stored in
the typical system fonts directory. Instead, Adobe stores the fonts 
in a hidden folder without file extensions, making them difficult to
locate and use with other software. This tool will make simple work
of accessing your licensed and activated Adobe CC fonts in a standard
font format of your choice. Simply run the tool as follows:

afc -f [file format] -o [output file path]

That's really it! Enjoy! :)`,
	Run: func(cmd *cobra.Command, args []string) {
		switch platform {
		case "windows":
			fmt.Println(banner)
			convert.ConvertWindowsFontFiles(fileFormat, outFolder)
		case "darwin":
			fmt.Println(banner)
			convert.ConvertDarwinFontFiles(fileFormat, outFolder)
		default:
			log.Fatalf("Only Windows and MacOS are supported.")
		}

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVarP(&outFolder, "out", "o", ".", "the destination folder.")
	rootCmd.Flags().StringVarP(&fileFormat, "format", "f", "otf", `the format to convert each font file into, such as:
	[otf], [ttf], [woff], [woff2], [eot]
	`)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
