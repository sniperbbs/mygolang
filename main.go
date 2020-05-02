// testAll project main.go
package main

import (
	"fmt"
	//"golang.org/x/net/html"
)

func main() {
	//tm := make(map[string]string)
	var tt bool
	tm := map[string]string{"a": "name", "b": "address", "c": "tel"}
	for _, v := range tm {
		fmt.Println("vv=", v)
	}
	fmt.Println("tt=", tt)
	fmt.Println("map-nil:", tm["d"] == "")

}
