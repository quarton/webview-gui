package gui

/*
#cgo linux pkg-config: webkit2gtk-4.0
#cgo windows LDFLAGS: -L. -W1, -rpath\$ORIGIN -l webview_edge
#cgo CFLAGS: -w
#include "webview-sys/webview.h"

#ifdef __linux__
	#include "webview-sys/webview_gtk.c"
#endif

extern void callback(void *, void *);
extern void dispatch(void *, void *, void *);
*/
import "C"
import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"runtime"
	"unsafe"

	"github.com/quarton/webview-gui/rpc"
	"github.com/quarton/webview-gui/rpc/jsonrpc"
)

func init() {
	runtime.LockOSThread()
}

type Gui struct {
	webview  C.webview_t
	callback CallbackFunc
	dispatch chan func()
	cli      net.Conn
	srv      net.Conn
	rpc      *rpc.Server
}

// New Gui
func NewGui(title string, html string, width, height int, resizeable bool, debug bool) (g *Gui) {

	g = new(Gui)

	titlePtr := C.CString(title)
	defer C.free(unsafe.Pointer(titlePtr))

	htmlPtr := C.CString("data:text/html," + url.PathEscape(html))
	defer C.free(unsafe.Pointer(htmlPtr))

	g.webview = C.webview_new(
		titlePtr,
		htmlPtr,
		C.int(width),
		C.int(height),
		C.int(bool2int(resizeable)),
		C.int(bool2int(debug)),
		(*[0]byte)(C.callback),
		nil)

	ptrMap[uintptr(g.webview)] = g
	g.dispatch = make(chan func(), 1)

	// create JSON RPC server
	g.rpc = rpc.NewServer()

	// register exported methods
	g.rpc.RegisterName("Gui", &GuiExport{gui: g})

	return g
}

// Run Gui
func (g *Gui) Run() {

	// in-memory full duplex network connection
	g.cli, g.srv = net.Pipe()

	// start rpc server
	go g.rpc.ServeCodec(jsonrpc.NewServerCodec(g.srv))

	// write request from javascript to rpc connection
	g.callback = func(data string) {
		g.cli.Write([]byte(data))
	}

	// write responses from the rpc connection to javascript
	go func(w *Gui) {
		scanner := bufio.NewScanner(g.cli)
		for scanner.Scan() {
			g.Eval(fmt.Sprintf("_rpc.recieve('%s')", scanner.Bytes()))
		}
	}(g)

	// add code for the JS RPC client
	go g.Eval(g.rpc.JsClient())

	// loop until quit
	for C.webview_loop(g.webview, 1) == 0 {
	}
}

// Registerexports the the receiver's methods to JS
func (g *Gui) Register(rcvr interface{}) error {
	return g.rpc.Register(rcvr)
}

// RegisterName is like register but uses the provided name
func (g *Gui) RegisterName(name string, rcvr interface{}) error {
	return g.rpc.RegisterName(name, rcvr)
}

// send func() to the dispatch channel and call dispatch()
// dispatch() is an exported Go function called by C on the main thread
func (g *Gui) Dispatch(function func()) {
	g.dispatch <- function
	C.webview_dispatch(g.webview, (*[0]byte)(C.dispatch), nil)
}

// runs JS in the browser
func (g *Gui) Eval(js string) {
	g.Dispatch(func() {
		strPtr := C.CString(js)
		defer C.free(unsafe.Pointer(strPtr))
		C.webview_eval(g.webview, strPtr)
	})
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

//
//
//

func (g *Gui) SetTitle(title string) {
	g.Dispatch(func() {
		titlePtr := C.CString(title)
		defer C.free(unsafe.Pointer(titlePtr))
		C.webview_set_title(g.webview, C.CString(title))
	})
}

func (g *Gui) SetFullScreen(fullscreen bool) {
	C.webview_set_fullscreen(g.webview, C.int(bool2int(fullscreen)))
}

func (g *Gui) SetColorRGB(red, green, blue uint8) {
	C.webview_set_color(g.webview, C.uint8_t(red), C.uint8_t(green), C.uint8_t(blue), C.uint8_t(255))
}

func (g *Gui) Terminate() {
	g.Dispatch(func() {
		C.webview_terminate(g.webview)
	})
}

func (g *Gui) dialog(title string, message string, dialogType C.int, dialogFlag C.int) string {

	result := make(chan string, 1)

	g.Dispatch(func() {

		titlePtr := C.CString(title)
		defer C.free(unsafe.Pointer(titlePtr))

		resultPtr := (*C.char)(C.calloc((C.size_t)(unsafe.Sizeof((*C.char)(nil))), 4096))
		defer C.free(unsafe.Pointer(resultPtr))

		messagePtr := C.CString(message)
		defer C.free(unsafe.Pointer(messagePtr))

		C.webview_dialog(g.webview, uint32(dialogType), dialogFlag, titlePtr, messagePtr, resultPtr, 4096) /*C.PATH_MAX*/

		result <- C.GoString(resultPtr)
	})

	return <-result
}

func (g *Gui) DialogOpenDirectory(title string) (path string) {
	return g.dialog(title, "", C.WEBVIEW_DIALOG_TYPE_OPEN, C.WEBVIEW_DIALOG_FLAG_DIRECTORY)
}

func (g *Gui) DialogOpenFile(title string) (path string) {
	return g.dialog(title, "", C.WEBVIEW_DIALOG_TYPE_OPEN, C.WEBVIEW_DIALOG_FLAG_FILE)
}

func (g *Gui) DialogSaveFile(title string) (path string) {
	return g.dialog(title, "", C.WEBVIEW_DIALOG_TYPE_SAVE, C.WEBVIEW_DIALOG_FLAG_FILE)
}

func (g *Gui) DialogAlertInfo(title string, message string) {
	_ = g.dialog(title, message, C.WEBVIEW_DIALOG_TYPE_ALERT, C.WEBVIEW_DIALOG_FLAG_INFO)
	return
}

func (g *Gui) DialogAlertWarning(title string, message string) {
	_ = g.dialog(title, message, C.WEBVIEW_DIALOG_TYPE_ALERT, C.WEBVIEW_DIALOG_FLAG_WARNING)
	return
}

func (g *Gui) DialogAlertError(title string, message string) {
	_ = g.dialog(title, message, C.WEBVIEW_DIALOG_TYPE_ALERT, C.WEBVIEW_DIALOG_FLAG_ERROR)
	return
}
