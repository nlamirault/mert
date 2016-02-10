// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	templateFile, err := ioutil.TempFile("", "configuration")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(templateFile.Name())
	data := []byte(`# Mert configuration file

font = "Hack 11"

[theme]
foreground = "#454545"
foreground_bold = "#989898"
background = "#000000"
cursor = "#555555"
palette = ["#666666", "#CC6699", "#99CC66", "#CC9966", "#6699CC", "#9966CC", "#66CC99", "#CCCCCC", "#999999", "#FF99CC", "#CF062B", "#CCFF99", "#FFCC99", "#99CCFF", "#CC99FF", "#99FFCC", "#FFFFFF" ]

`)
	err = ioutil.WriteFile(templateFile.Name(), data, 0700)
	if err != nil {
		t.Fatal(err)
	}
	configuration, err := Load(templateFile.Name())
	if err != nil {
		t.Fatalf("Error with configuration: %#v", err)
	}
	fmt.Printf("Configuration : %#v\n", configuration)
	if configuration.Font != "Hack 11" {
		t.Fatalf("Configuration font failed")
	}

	if configuration.Theme.Foreground != "#454545" ||
		configuration.Theme.ForegroundBold != "#989898" ||
		configuration.Theme.Background != "#000000" ||
		configuration.Theme.Cursor != "#555555" {
		t.Fatalf("Configuration theme failed: %v", configuration.Theme)
	}

}
