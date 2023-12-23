package Myls

import (
	"fmt"
	"os"
	"strings"
)

var (
	l_flag       bool
	R_flag       bool
	a_flag       bool
	r_flag       bool
	t_flag       bool
	ls           bool
	SubFile_flag bool
	first        bool
)

func Validation() ([]string, []string) {
	var PhathName []string
	var subFile []string
	//chek the argument
	if len(os.Args) < 2 || os.Args[1] != "ls" {
		os.Exit(0)
	} else if len(os.Args) == 2 {
		ls = true
	} else {
		// spilt the args so maybe we have more than 1 flag
		Flags := os.Args[2:]
		// for the flags range
		for _, flag := range Flags {
			if flag[0] == '-' {
				// to check the ls flags
				if len(flag) == 2 {
					CheckFlag(rune(flag[1]))
				} else if len(flag) < 2 {
				fmt.Println("ls: cannot access '-': No such file or directory")
				os.Exit(0)
				} else {
					for i := 1; i < len(flag); i++ {
						CheckFlag(rune(flag[i]))
					}
				}
			} else {
				v := ""
				if !CheckShortCut(flag){
				if flag[len(flag)-1:] != "/" {
					v = flag + "/"
				}
				_, err := os.Stat(v)
				if err == nil {
					PhathName = append(PhathName,flag )
				} else if strings.Contains(flag, "/") || string(rune(flag[0])) == "/" || flag[0:2] == "./" {
					PhathName = append(PhathName,flag)
				} } else {
					SubFile_flag = true
					subFile = append(subFile,flag)
				}
			}
		}
		first = true
		if len(PhathName) == 0 {
			PhathName = append(PhathName,"./")
		}
		return PhathName, subFile
	}

	PhathName = append(PhathName,"./")

	return PhathName, subFile
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
		fmt.Println("ls: unvslid input -- '" + string(c) + "'")
		os.Exit(0)
	}
}
