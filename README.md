Webview-GUI
===========

Webview GUI allows you to easily create a GUI for your Golang project using HTML and Javascript. The resulting GUI will work on Linux (WebKitGTK+2) or Windows (Edge). Functions you register in your Go code will be available to use in JavaScript.

Webview Gui is based on the C/C++ files from [Boscop/web-view](https://github.com/Boscop/web-view/) a fork of [zserge/webview](https://github.com/zserge/webview). These projects use the MIT License. The code can be found in the webview-sys directory.

Download Prebuilt Example
-------------------------

[Linux](https://github.com/quarton/webview-gui/releases/download/0.1.0/ubuntu_64_examples.tar.xz) (requires libwebkit2gtk-4.0, built on Ubuntu)
[Windows](https://github.com/quarton/webview-gui/releases/download/0.1.0/windows10_64bit_examples.zip) (Built on Windows 10 64bit)

Example Code
-------

This example shows how to call a Golang function from Javascript. This Go code creates a window and registers Calculator.Evaluate().

```
func main() {

	g := gui.NewGui("Calculator", string(asset), 400, 460, false, false)

	g.Register(new(Calculator))
	g.Run()

}

type Reply struct {
	Result string
}

type Calculator struct{}

func (t *Calculator) Evaluate(input string, reply *Reply) (err error) {

	res, err := compute.Evaluate(input)
	if err != nil {
		return err
	}
	reply.Result = strconv.FormatFloat(res, 'G', 15, 64)

	return nil
}
```

The Javascript to call this function is as follows. Calculator.Evaluate() returns a promise.

```
Calculator.Evaluate(calc).then((reply) => {

    // success, answer in reply.Result

}).catch((err) => {

    // error 

})
```

Windows DLL Build Instructions
------------------------------

Windows requires the `webview_edge.dll` to be in the same directory as the executable. You can build this using the instructions in [webview-gui/webview-sys](https://github.com/quarton/webview-gui/tree/master/webview-sys) or use the precompiled version also found in that directory.

If you wish to build the Windows DDL it requires that Visual Studio is installed on your system with the MSVC component selected. Run the following command in the 'webview-sys' directory using 'x64 Native Tools Command Prompt for VS 2019'.

`cl /LD /std:c++17 -DUNICODE -D_UNICODE webview_edge.cpp`