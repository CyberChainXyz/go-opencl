package main

import (
	"github.com/kr/pretty"
	cl "github.com/CyberChainXyz/go-opencl"
)

func main() {
	info, _ := cl.Info()
	pretty.Println(info)
}
