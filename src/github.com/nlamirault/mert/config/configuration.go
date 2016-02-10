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
	"log"

	"github.com/BurntSushi/toml"
)

// Theme holds the Mert theme
type Theme struct {
	Foreground     string   `toml:"foreground"`
	ForegroundBold string   `toml:"foreground_bold"`
	Background     string   `toml:"background"`
	Cursor         string   `toml:"cursor"`
	Palette        []string `toml:"palette"`
}

// Configuration holds configuration for Mert.
type Configuration struct {
	Font  string `toml:"font"`
	Theme *Theme `toml:"theme"`
}

var defaultTheme = Theme{
	Foreground:     "#d1d1d1",
	ForegroundBold: "#ffffff",
	Background:     "#323232",
	Cursor:         "#e8e8e8",
	Palette: []string{
		"#666666", "#CC6699", "#99CC66",
		"#CC9966", "#6699CC", "#9966CC", "#66CC99",
		"#CCCCCC", "#999999", "#FF99CC", "#CF062B", "#CCFF99",
		"#FFCC99", "#99CCFF", "#CC99FF", "#99FFCC", "#FFFFFF"},
}

// New returns a Configuration with default values
func New() *Configuration {
	return &Configuration{
		Font:  "Monospace 12",
		Theme: &defaultTheme,
	}
}

// Load returns a Configuration from reading the specified file (a toml file).
func Load(file string) (*Configuration, error) {
	log.Printf("[DEBUG] Load configuration file: %s", file)
	configuration := New()
	if _, err := toml.DecodeFile(file, configuration); err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Configuration : %#v", configuration)
	return configuration, nil
}
