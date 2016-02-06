package main

import "C"
import wc "./wc"
import wt "./wt"

//export wthello
func wthello(name *C.char) *C.char {
	return C.CString(wt.Hello(C.GoString(name)))
}

//export wchello
func wchello(name *C.char) *C.char {
	return C.CString(wc.Hello(C.GoString(name)))
}

func main() {}
