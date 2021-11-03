package main

import (
	"strconv"
	"syscall/js"
)

func add(this js.Value,i []js.Value) interface{}{
	//get the value in index.html input field 1
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	//get the value in index.html input field 2
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	//convert string value to a integer
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	//set int1+int2 value in to a index.html result field
	 js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
	return nil
}

func subtract( this js.Value,i []js.Value) interface{}{
	//get the value in index.html input field 1
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	//get the value in index.html input field 2
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	//convert string value to a integer
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	//set int1-int2 value in to a index.html result field
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
	return nil
}

func divide( this js.Value,i []js.Value) interface{}{
	//get the value in index.html input field 1
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	//get the value in index.html input field 2
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	//convert string value to a integer
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	if int2 == 0 {
		js.Global().Get("document").Call("getElementById",i[2].String()).Set("value","Cannot Divide By 0")
		return nil
	}
	//set int1/int2 value in to a index.html result field
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1/int2)
	return nil
}

func multiple( this js.Value,i []js.Value) interface{}{
	//get the value in index.html input field 1
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	//get the value in index.html input field 1
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	//convert string value to a integer
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	//set int1/int2 value in to a index.html result field
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1 * int2)
	return nil
}

func registerCallbacks() {
	//return all method as a javascript global object
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("divide", js.FuncOf(divide))
	js.Global().Set("multiple", js.FuncOf(multiple))
}


func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
