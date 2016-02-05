package main

import "C"
import wc "./wc"
import wt "./wt"

//export wthello
func wthello(name string) string {
	return wt.Hello(name)
}

//export wchello
func wchello(name string) string {
	return wc.Hello(name)
}

func main() {}
