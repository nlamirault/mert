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

package events

import (
	// "fmt"
	"log"

	"github.com/mattn/go-gtk/gdk"

	"github.com/nlamirault/mert/vte3"
)

// KeyboardHandler handle events from keyboard
func KeyboardHandler(event chan *KeyPressEvent, terminal vte3.Terminal) {
	for {
		kpe := <-event
		log.Printf("[DEBUG] KeyPressEvent : %#v", kpe)
		gdk.ThreadsEnter()
		switch kpe.KeyVal {
		case gdk.KEY_plus:
			if kpe.GetModifier() == CTRL {
				log.Printf("[DEBUG] +")
			}
			break
		case gdk.KEY_minus:
			if kpe.GetModifier() == CTRL {
				log.Printf("[DEBUG] -")
			}
			break
		case gdk.KEY_equal:
			if kpe.GetModifier() == CTRL {
				log.Printf("[DEBUG] =")
			}
			break
		case gdk.KEY_exclam:
			if kpe.GetModifier() == CTRL {
				log.Printf("[DEBUG] About")
			}
			break
		}
		gdk.ThreadsLeave()
	}
}
