package main

import (
	// "flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
)

var l_flag bool
var R_flag bool
var a_flag bool
var r_flag bool
var t_flag bool
var ls bool

func main() {
	pathname := validation()
	x := listing(pathname)
	Print(pathname, x)
}

func listing(dir string) []fs.FileInfo {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfos

}

func SortByDate(filesInfos []fs.FileInfo) []fs.FileInfo {
	n := len(filesInfos)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if filesInfos[j].ModTime().Before(filesInfos[j+1].ModTime()) {
				filesInfos[j], filesInfos[j+1] = filesInfos[j+1], filesInfos[j]
			}
		}
	}
	return filesInfos

}

func Print(path string, fileInfos []fs.FileInfo) {
	if t_flag {
		fileInfos = SortByDate(fileInfos)
	}

	if a_flag && !r_flag {
		fmt.Print("./  ../  ")
	}

	for i := 0; i < len(fileInfos); i++ {
		index := i
		if r_flag {
			index = len(fileInfos) - i - 1
		}
		fileInfo := fileInfos[index]
		if fileInfo.Name()[0] != '.' || a_flag {
			if l_flag {
				subpath := path + fileInfo.Name()
				Gid, UserId , filelinks := returnGroupAndUSerId(subpath)
				mode := fmt.Sprint(fileInfo.Mode())
				time := fmt.Sprint(fileInfo.ModTime())
				size := fmt.Sprint(fileInfo.Size())
				fmt.Print(mode + " " + filelinks + " " + UserId + " " + Gid + " " +size+" "+ time + " ")
			}
			if fileInfo.IsDir() {
				fmt.Printf("%s/  ", fileInfo.Name())
				if R_flag {
					subPath := path + "/" + fileInfo.Name()
					fmt.Println("\n" + subPath + ":")
					Print(subPath, listing(subPath))
				}
			} else {
				fmt.Print(fileInfo.Name() + "  ")
			}
			if l_flag {
				fmt.Println()
			}
		}
	}
	if a_flag && r_flag {
		fmt.Println("./  ../  ")
	}
	fmt.Println()

}

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
				if string(rune(flag[0])) == "/" || flag[0:1] == "./" {
					PhathName = flag
				} else {
					PhathName += flag
				}
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

func returnGroupAndUSerId(path string) (string, string, string) {
	file_info, _ := os.Stat(path)
	file_sys := file_info.Sys()
	GID := fmt.Sprint(file_sys.(*syscall.Stat_t).Gid)
	UID := fmt.Sprint(file_sys.(*syscall.Stat_t).Uid)
	flink := fmt.Sprint(file_sys.(*syscall.Stat_t).Nlink)
	grId, _ := user.LookupGroupId(GID)
	usrId, _ := user.LookupId(UID)

	return grId.Name, usrId.Username , flink

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
