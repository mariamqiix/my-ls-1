package main

import (
	// "flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var l_flag bool
var R_flag bool
var a_flag bool
var r_flag bool
var t_flag bool
var ls bool

func main() {
	pathname := validation()
	if !l_flag && !r_flag && !R_flag && !a_flag && !t_flag {
		listing(pathname)
	}

	// listing(pathname)
	// fmt.Print(R_flag)
	// if R_flag {
	// 	err := listFilesRecursively(pathname, "")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

}

// func validation() (string,bool){
// 	l_flag, R_flag, a_flag, r_flag, t_flag, pathname := false, false, false, false, false, "."

// 	if len(os.Args) < 2 || os.Args[1] != "ls" {
// 		fmt.Println("Error, read the manPage 1")
// 		os.Exit(0)
// 	} else if len(os.Args) == 3 {
// 				if !l_flag && os.Args[2] == "-l" {
// 					l_flag = true
// 				} else if !R_flag && os.Args[2] == "-R" {
// 					R_flag = true
// 				} else if !a_flag && os.Args[2] == "-a" {
// 					a_flag = true
// 				} else if !r_flag && os.Args[2] == "-r" {
// 					r_flag = true
// 				} else if !t_flag && os.Args[2] == "-t" {
// 					t_flag = true
// 				} else {
// 					fmt.Println("Error, read the manPage 2")
// 					os.Exit(0)
// 				}
// 				fmt.Print(R_flag)
// 	} else if len(os.Args) !=2 {
// 		fmt.Println("Error, read the manPage 3")
// 		os.Exit(0)
// 	}
// 	fmt.Print(R_flag)

// 	if len(os.Args) == 4 {
// 		pathname = os.Args[3]
// 	}

// 	return pathname,R_flag
// }

func listing(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.Name()[0] != '.' && !a_flag {
			if fileInfo.IsDir() {
				fmt.Printf("%s/\n", fileInfo.Name())
			} else {
				fmt.Print(fileInfo.Name(), " ")
			}
		}
	}
	fmt.Println()
}

// func listFilesRecursively(dirPath, indent string) error {
// 	dir, err := os.Open(dirPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer dir.Close()

// 	fileInfos, err := dir.Readdir(-1)
// 	if err != nil {
// 		return err
// 	}

// 	for _, fileInfo := range fileInfos {
// 		fileName := fileInfo.Name()
// 		fileInfo.
// 		if fileName != "." && fileName != ".." {
// 			fmt.Print(indent)
// 			if fileInfo.IsDir() {
// 				fmt.Printf("%s \n ", fileName)
// 				subDirPath := dirPath + "/" + fileName
// 				subIndent := indent + "  "
// 				err := listFilesRecursively(subDirPath, subIndent)
// 				if err != nil {
// 					return err
// 				}
// 			} else {
// 				fmt.Println(fileName)
// 			}
// 		} else {
// 			fmt.Print(indent)

// 		}
// 	}

// 	return nil
// }

func validation() string {
	//chek the argument
	if len(os.Args) > 3 || len(os.Args) < 2 {
		os.Exit(0)
	} else if os.Args[1] != "ls" {
		os.Exit(0)
	} else if len(os.Args) == 2 {
		ls = true
	} else {
		// spilt the args so maybe we have more than 1 flag
		Flags := strings.Split(os.Args[2], " ")
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
				PhathName += flag
			}
		}
		return PhathName
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
		fmt.Print("unvslid input")
		os.Exit(0)
	}
}
