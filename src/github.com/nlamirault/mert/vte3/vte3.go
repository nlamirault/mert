package vte3

// #include "vte3.go.h"
// #cgo pkg-config: vte
import "C"

import (
	"unsafe"

	"github.com/mattn/go-gtk/gtk"
)

func (t Terminal) VteToGtk() *gtk.Widget {
	return gtk.WidgetFromNative(unsafe.Pointer(t.Widget))
}

type Terminal struct {
	Widget *C.GtkWidget
}

func NewTerminal() Terminal {
	//new_term := unsafe.Pointer(C.vte_terminal_new())
	//widget := gtk3.WidgetFromNative(new_term)
	return Terminal{C.vte_terminal_new()}
}

func (t Terminal) Fork(args []string) {
	argv := C.make_strings(C.int(len(args) + 1))
	defer C.free(unsafe.Pointer(argv))
	for i, arg := range args {
		ptr := C.CString(arg)
		defer C.free(unsafe.Pointer(ptr))
		C.set_string(argv, C.int(i), ptr)
	}
	C.set_string(argv, C.int(len(args)), nil)
	C.vte_terminal_fork_command_full(C.toVTerminal(t.Widget), C.VTE_PTY_DEFAULT,
		nil, argv, nil, C.G_SPAWN_SEARCH_PATH, nil, nil, nil, nil)
}

func (t Terminal) GetIconTitle() string {
	return C.GoString(C.vte_terminal_get_icon_title(C.toVTerminal(t.Widget)))
}

func (t Terminal) SetFont(font string) {
	C.vte_terminal_set_font_from_string(C.toVTerminal(t.Widget), C.CString(font))
}

func createColor(pattern string) C.GdkColor {
	var color C.GdkColor
	ptr := C.CString(pattern)
	defer C.free(unsafe.Pointer(ptr))
	C.gdk_color_parse(C.toGstr(ptr), &color)
	return color
}

func (t Terminal) SetColors(foreground string, background string, palette []string) {
	fColor := createColor(foreground)
	bColor := createColor(background)
	pColors := new([16]C.GdkColor)
	for i := 0; i < len(pColors); i++ {
		C.gdk_color_parse((*C.gchar)(C.CString(palette[i])), &pColors[i])
	}
	C.vte_terminal_set_colors(C.toVTerminal(t.Widget), &fColor, &bColor, (*C.GdkColor)(unsafe.Pointer(pColors)), 16)
}

func (t Terminal) SetColorCursor(pattern string) {
	color := createColor(pattern)
	C.vte_terminal_set_color_cursor(C.toVTerminal(t.Widget), &color)
}

// func (t Terminal) SetColorForeground(pattern string) {
// 	color := createColor(pattern)
// 	C.vte_terminal_set_color_foreground(C.toVTerminal(t.Widget), &color)
// }

// func (t Terminal) SetColorBackground(pattern string) {
// 	color := createColor(pattern)
// 	C.vte_terminal_set_color_background(C.toVTerminal(t.Widget), &color)
// }
