package main

import (
	"strconv"

	"github.com/alfredxing/calc/compute"
	"github.com/quarton/webview-gui"
)

func main() {

	g := gui.NewGui("Calculator", string(asset), 440, 460, true, false)

	g.Register(new(Calculator))
	g.SetColorRGB(49, 59, 72)
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
