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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"

	"github.com/nlamirault/mert/config"
	"github.com/nlamirault/mert/version"
	"github.com/nlamirault/mert/vte3"
)

const (
	application      = "Mert"
	defaultWinWidth  = 1024
	defaultWinHeight = 768

	defaultConfigurationFile = ".config/mert/mert.toml"
)

var (
	port  string
	debug bool
)

func init() {
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.Parse()
}

func getApplicationTitle() string {
	return fmt.Sprintf("%s - v%s", application, version.Version)
}

func getConfigurationFile() string {
	return filepath.Join(os.Getenv("HOME"), defaultConfigurationFile)
}

func configure(t vte3.Terminal, conf *config.Configuration) {
	t.SetFont(conf.Font)
	t.SetColors(conf.Theme.Foreground, conf.Theme.Background, conf.Theme.Palette)
	t.SetColorCursor(conf.Theme.Cursor)
}

func runGUI(argv []string) {
	conf, err := config.Load(getConfigurationFile())
	if err != nil {
		log.Printf("[WARN] No configuration file or invalid configuration. Use default settings")
		conf = config.New()
	}

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle(getApplicationTitle())
	terminal := vte3.NewTerminal()
	widget := terminal.VteToGtk()

	// Events
	window.Connect("destroy", gtk.MainQuit)
	widget.Connect("child-exited", gtk.MainQuit)

	window.Add(widget)
	window.SetSizeRequest(defaultWinWidth, defaultWinHeight)
	window.ShowAll()

	configure(terminal, conf)

	terminal.Fork(argv)
	gtk.Main()
}

func main() {
	glib.ThreadInit(nil)
	gdk.ThreadsInit()
	gtk.Init(nil)
	runGUI([]string{os.Getenv("SHELL")})

}
