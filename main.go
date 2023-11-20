package main

import "my-ls/Myls"

func main() {
	pathname := Myls.Validation()
	Myls.Print(pathname, Myls.Listing(pathname))
}
