package rpc

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Is this an exported - upper case - name?
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

func (server *Server) JsClient() (rpcFuncJs string) {

	rpcFuncJs = `
	_rpc = new class {
		
	    constructor() {
	        this.pending = {};
	        this.id = 0;
	    }

	    nextId() {
	    	return this.id++;
	    }

	    recieve(data) {
	        var response = JSON.parse(data);
	        var promise = this.pending[response.id];

	        if (promise === undefined) return;

	        if (response.error) {
	            promise.reject(response.error);
	        } else {
	            promise.resolve(response.result);
	        }

	        delete this.pending[response.id];
	    }
	}

	`

	methodTemplate := `
	this.%s = function (%s) {
		var id = _rpc.nextId();
		var message = JSON.stringify({method: "%s", id: id, params: [%s]});
		window.external.invoke(message);
		return new Promise(function(resolve, reject) {
			_rpc.pending[id] = { resolve: resolve, reject: reject };
		});
	}
	`

	server.serviceMap.Range(func(serviceName, svci interface{}) bool {
		svc := svci.(*service)

		var serviceJS string

		for methodName, method := range svc.method {
			rpcMethod := fmt.Sprintf("%s.%s", serviceName, methodName)

			var fields []string

			// dereference pointer
			t := method.ArgType
			if t.Kind() == reflect.Ptr {
				t = method.ArgType.Elem()
			}

			rpcArgs := method.ArgType.Name()
			rpcParams := method.ArgType.Name()

			if t.Kind() == reflect.Struct {
				for i := 0; i < t.NumField(); i++ {
					fields = append(fields, t.Field(i).Name)
				}
				rpcArgs = strings.Join(fields, ", ")
				for i, _ := range fields {
					fields[i] = fmt.Sprintf("%s: %s", fields[i], fields[i])
				}

				rpcParams = fmt.Sprintf("{%s}", strings.Join(fields, ", "))
			}

			serviceJS += fmt.Sprintf(methodTemplate, methodName, rpcArgs, rpcMethod, rpcParams)
		}

		rpcFuncJs += fmt.Sprintf("var %s = new function(){\n%s\n}\n", serviceName, serviceJS)

		return true
	})

	return rpcFuncJs
}
