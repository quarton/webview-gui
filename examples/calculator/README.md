Calculator Example
==================

This example shows how to call a Golang function from Javascript.

This Go code creates creates and registers Calculator.Evaluate().

```
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

// in main()...

g.Register(new(Calculator))
```

The Javascript to call this function is as follows. Calculator.Evaluate() returns a promise.

```
Calculator.Evaluate(calc).then((reply) => {

    // success, answer in reply.Result

}).catch((err) => {

    // error 

})
```

Download Prebuilt Example
-------------------------

Linux
Windows


Linux Build
-----------

```
go build
```

Windows Build
-------------

```
 go build -ldflags="-H windowsgui"
```

Building Assets
---------------

The HTML, JS and CSS for this example are inlined using `inline-assets` and the resulting file stored as a byte slice using `go-asset`. You can install these two packages using npm.

```
npm install [-g] inline-assets go-asset
```


```
inline-assets assets/calculator.html | go-asset -o asset.go
```

