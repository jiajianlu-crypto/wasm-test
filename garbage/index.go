package main

import (
	"syscall/js"
)

// var list [][]int32

func generateGarbage(this js.Value, args []js.Value) any {
	tempList := make([]int32, 10000000);
	// list = append(list, tempList)
	return len(tempList)
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("go webAssembly!")
	// 提供模块
	js.Global().Set("goGenGar", js.FuncOf(generateGarbage))
	
	done := make(chan int, 0)
	<-done
}