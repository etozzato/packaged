package main

import "fmt"
import "gitlab.com/etozzato/packaged/tree/master/src/wt/wt"
import "gitlab.com/etozzato/packaged/tree/master/src/wc/wc"

func main() {
	fmt.Println(wt.hello())
	fmt.Println(wc.hello())
}
