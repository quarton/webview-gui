// +build windows

package gui

import (
	"net/url"
	"syscall"
)

// HDPI
func init() {
	shcore := syscall.NewLazyDLL("Shcore.dll")
	dpi := shcore.NewProc("SetProcessDpiAwareness")
	dpi.Call(uintptr(1))
}

// windows needs QueryEscape not PathEscape
func encode(html string) string {
	return "data:text/html," + url.QueryEscape(html)
}
