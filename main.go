package main

import (
	"my-ls/Myls"
)

func main() {
	pathname, subFile := Myls.Validation()
	Myls.Print(pathname, subFile, Myls.Listing(pathname))
}
