package main

import (
	"syscall/js"
)

func sayHello() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return "Hello from Go WebAssembly!"
	})
}

func main() {
	ch := make(chan struct{}, 0)
	js.Global().Set("goSayHello", sayHello())
	<-ch
}
