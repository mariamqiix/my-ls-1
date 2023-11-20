package Myls

import (
	"fmt"
	"os"
)

var (
	l_flag bool
	R_flag bool
	a_flag bool
	r_flag bool
	t_flag bool
	ls     bool
)

func Validation() string {
	//chek the argument
	if len(os.Args) < 2 || os.Args[1] != "ls" {
		os.Exit(0)
	} else if len(os.Args) == 2 {
		ls = true
	} else {
		// spilt the args so maybe we have more than 1 flag
		Flags := os.Args[2:]
		// path name if it exist
		PhathName := "./"
		// for the flags range
		for _, flag := range Flags {
			if flag[0] == '-' {
				// to check the ls flags
				if len(flag) == 2 {
					CheckFlag(rune(flag[1]))
				} else {
					for i := 1; i < len(flag); i++ {
						CheckFlag(rune(flag[i]))
					}
				}
			} else {
				if string(rune(flag[0])) == "/" || flag[0:1] == "./" {
					PhathName = flag
				} else {
					PhathName += flag
				}
			}
		}
		return PhathName
	}
	if R_flag {
		fmt.Println(".:")
	}
	return "./"
}

// will edit the flag
func CheckFlag(c rune) {
	if c == rune('l') {
		l_flag = true
	} else if c == rune('R') {
		R_flag = true
	} else if c == rune('a') {
		a_flag = true
	} else if c == rune('r') {
		r_flag = true
	} else if c == rune('t') {
		t_flag = true
	} else {
		fmt.Println("unvslid input")
		os.Exit(0)
	}
}
