package gui

//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// go functions that are export to C

type CallbackFunc func(data string)

// map the callback to the correct Gui instance
var ptrMap map[uintptr]*Gui

func init() {
	ptrMap = make(map[uintptr]*Gui)
}

//export callback
func callback(w unsafe.Pointer, data unsafe.Pointer) {
	if webview, found := ptrMap[uintptr(w)]; found {
		webview.callback(C.GoString((*C.char)(data)))
	}
}

//export dispatch
func dispatch(w unsafe.Pointer, data unsafe.Pointer) {
	defer C.free(unsafe.Pointer((*C.char)(data)))
	if webview, found := ptrMap[uintptr(w)]; found {
		select {
		case function, ok := <-webview.dispatch:
			if ok {
				function()
			}
		}
	}
}
