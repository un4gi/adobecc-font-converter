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
package convert

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ConvertWindowsFontFiles(format, dst string) {
	homepath, _ := os.UserHomeDir()
	adobeccFontPath := homepath + "\\AppData\\Roaming\\Adobe\\CoreSync\\plugins\\livetype\\r\\"
	fmt.Printf("Converting font files...\r\n")
	files, _ := ioutil.ReadDir(adobeccFontPath)
	for _, file := range files {
		sourceFile := adobeccFontPath + file.Name()
		fullDest := dst + file.Name() + "." + format
		copyFile(sourceFile, fullDest)
	}
}

func ConvertDarwinFontFiles(format, dst string) {
	homepath, _ := os.UserHomeDir()
	adobeccFontPath := homepath + "/Library/Application Support/Adobe/CoreSync/plugins/livetype/.r/"
	fmt.Printf("Converting font files...\r\n")
	files, _ := ioutil.ReadDir(adobeccFontPath)
	for _, file := range files {
		sourceFile := adobeccFontPath + file.Name()
		fullDest := dst + file.Name() + "." + format
		copyFile(sourceFile, fullDest)
	}
}

func copyFile(src, dst string) (int64, error) {
	in, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer out.Close()
	fmt.Printf("Copying %s to %s.\r\n", src, dst)
	return io.Copy(out, in)
}
