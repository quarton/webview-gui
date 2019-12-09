package main

import (
	"errors"

	"github.com/quarton/webview-gui"
)

func main() {

	w := gui.NewGui("Example", html(), 800, 600, true, false)

	w.Register(new(Arith))

	w.Run()
}

//---------------------------------------------------------

type Args struct {
	A, B int
}

type Reply struct {
	C float64
}

type Arith int

func (t *Arith) Div(args *Args, reply *Reply) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.C = float64(args.A) / float64(args.B)
	return nil
}

//-----------------------

func html() string {

	return `

<html>
<style>
div{margin:15px;}
div div{margin:0px;}
</style>

<div>
	<input type="text" id="title" value="Example"><input type="submit" value="Change Title" onClick="Gui.SetTitle(document.getElementById('title').value)">
</div>

<div>
<input type="number" id="rgb_r" min="0" max="255" step="5" value="255">
<input type="number" id="rgb_g" min="0" max="255" step="5" value="255">
<input type="number" id="rgb_b" min="0" max="255" step="5" value="255">
<input type="submit" value="Change Color" onClick="
	Gui.SetColorRGB(
		parseInt(document.getElementById('rgb_r').value),
		parseInt(document.getElementById('rgb_g').value),
		parseInt(document.getElementById('rgb_b').value))">
</div>

<div>
	<div><input type="submit" value="FullScreen On" onClick="Gui.SetFullScreen(true)"></div>
	<div><input type="submit" value="FullScreen Off" onClick="Gui.SetFullScreen(false)"></div>
</div>

<div>
	<div>
	<input type="submit" value="Open Directory" onClick="
		Gui.DialogOpenDirectory('Open Directory').then((result) => {document.getElementById('opendir').innerHTML=result.Path;})">
	<span id="opendir"></span>
	</div>

	<div>
	<input type="submit" value="Open File" onClick="
		Gui.DialogOpenFile('Open File').then((result) => {document.getElementById('openfile').innerHTML=result.Path})">
	<span id="openfile"></span>
	</div>

	<div>
	<input type="submit" value="Save File" onClick="
		Gui.DialogSaveFile('Save File').then((result) => {document.getElementById('savefile').innerHTML=result.Path})">
	<span id="savefile"></span>
	</div>
</div>

<div>
	<div><input type="submit" value="Info Alert" onClick="Gui.DialogAlertInfo('Info Alert', 'This is an info message!')"></div>
	<div><input type="submit" value="Warning Alert" onClick="Gui.DialogAlertWarning('Warning Alert', 'This is a warning message!')"></div>
	<div><input type="submit" value="Error Alert" onClick="Gui.DialogAlertError('Error Alert', 'This is an error message!')"></div>
</div>

<div>
	<input type="submit" value="Terminate" onClick="Gui.Terminate()">
</div>

<div>
	<div>Go Function</div>
	<input type="number" name="quantity" id="a" value="1">
	<input type="number" name="quantity" id="b" value="1">
	<input type="submit" value="Divide" onClick="divide();">
	<p id='demo'></p>
</div>


<script>

function divide () {

	var a = parseInt(document.getElementById("a").value);
	var b = parseInt(document.getElementById("b").value);

	Arith.Div(a, b).then((result) => {
	    console.log("result" + result.C)
	    document.getElementById("demo").innerHTML = result.C;
	}).catch((error) => {
		document.getElementById("demo").innerHTML = error;
	    console.log("error" + error)
	})

}

</script>

</html>`
}
