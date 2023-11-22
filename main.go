package main

import "my-ls/Myls"

func main() {
	pathname,subFile := Myls.Validation()
	if l_flag && !SubFile_flag {
		fmt.println(pathname,": ")
	}
	Myls.Print(pathname,subFile, Myls.Listing(pathname))
}
