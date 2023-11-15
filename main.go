package main

import (
	// "flag"
	"fmt"
	"io/fs"
	"log"
	"math"
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
	if R_flag {
		fmt.Println(".:")
	}
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

func Print(path string, fileInfos []fs.FileInfo) {
	if l_flag {
		if a_flag {
			file1, err := os.Stat(".")
			if err != nil {
				log.Fatal(err)
			}
			file2, err := os.Stat("..")
			if err != nil {
				log.Fatal(err)
			}
			fileInfos = append([]fs.FileInfo{file2}, fileInfos...)
			fileInfos = append([]fs.FileInfo{file1}, fileInfos...)

		}
		fmt.Println("total", math.Round(float64(totalSize(fileInfos, path))))
	}
	if t_flag {
		fileInfos = SortByDate(fileInfos)
	}

	if a_flag && !r_flag {
		if !R_flag && !l_flag {
			fmt.Print("../  ./ ")
		}
	}
	for i := 0; i < len(fileInfos); i++ {
		index := i
		if r_flag {
			index = len(fileInfos) - i - 1
		}
		fileInfo := fileInfos[index]
		if fileInfo.Name()[0] != '.' || a_flag {
			if l_flag {
				lFlag(path, fileInfo)
			}
			if fileInfo.IsDir() {
				fmt.Print("\033[34m", fmt.Sprintf("%s/  ", fileInfo.Name()),"\033[97m", "")
			} else {
				subPath := returnPath(fileInfo.Name(), path)

				if checkShortCut(subPath) && l_flag {
					shortName, _ := os.Readlink(subPath)
					fmt.Print("\033[36m", fileInfo.Name()+" ", "\033[97m", "-> ", "\033[34m", shortName, "\033[97m")
				} else {
					fmt.Print("\033[33m", fileInfo.Name()+"  ","\033[97m", "")
				}
			}
			if l_flag {
				fmt.Println()
			}
		}
	}

	if a_flag && r_flag && !l_flag {
		fmt.Println("./  ../  ")
	}

	if R_flag {
		Rflag(path, fileInfos)
	}

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

func Rflag(path string, fileInfos []fs.FileInfo) {
	for i := 0; i < len(fileInfos); i++ {
		index := i
		if r_flag {
			index = len(fileInfos) - i - 1
		}
		fileInfo := fileInfos[index]
		if fileInfo.Name()[0] != '.' || a_flag && fileInfo.Name() != ".." && fileInfo.Name() != "." {
			if fileInfo.IsDir() && fileInfo.Name() != "WinSAT" {
				subPath := returnPath(fileInfo.Name(), path)
				fmt.Println("\n" + subPath + ":")
				fmt.Print("\033[97m", "")
				Print(subPath, listing(subPath))
			}
		}
	}
}

func returnPath(fileName, path string) string {
	subPath := ""
	if path != "./" {
		subPath = path + "/" + fileName
	} else {
		subPath = path + fileName

	}
	return subPath
}

func validation() string {
	//chek the argument
	if len(os.Args) < 2  || os.Args[1] != "ls" {
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

func returnGroupAndUSerId(path string) (string, string, string) {
	file_info, err := os.Lstat(path)
	if err != nil {
		fmt.Print(file_info)
		log.Fatal(err)
	}
	file_sys := file_info.Sys()
	GID := fmt.Sprint(file_sys.(*syscall.Stat_t).Gid)
	UID := fmt.Sprint(file_sys.(*syscall.Stat_t).Uid)
	flink := fmt.Sprint(file_sys.(*syscall.Stat_t).Nlink)
	grId, _ := user.LookupGroupId(GID)
	usrId, _ := user.LookupId(UID)

	return grId.Name, usrId.Username, flink
	// return "", "", "1"

}

func lFlag(path string, fileInfo fs.FileInfo) {
	subpath := path + "/" + fileInfo.Name()
	Gid, UserId, filelinks := returnGroupAndUSerId(subpath)
	mode := fmt.Sprint(fileInfo.Mode())
	DateAndTime := fmt.Sprintf("%s %d %d:%d", fileInfo.ModTime().Month().String()[:3], fileInfo.ModTime().Day(), fileInfo.ModTime().Hour(), fileInfo.ModTime().Minute())
	size := fmt.Sprint(fileInfo.Size())
	if fileInfo.Size() < 10 {
		size = "   " + size
	} else if fileInfo.Size() < 100 {
		size = "  " + size
	} else if fileInfo.Size() < 1000 {
		size = "   " + size
	}
	if strings.Contains(mode, "Drw-rw-") || strings.Contains(mode, "Dcrw--") || strings.Contains(mode, "Dcrw-") {
		if Gid == "disk" {
			mode = strings.ReplaceAll(mode, "Dc", "b")
			mode = strings.ReplaceAll(mode, "D", "b")
		} else {
			mode = strings.Replace(mode, "D", "", 1)
		}
	}
	fmt.Print(mode + " " + filelinks + " " + UserId + " " + Gid + " " + size + " " + DateAndTime + " ")
}

func checkShortCut(path string) bool {
	file_info, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}
	if file_info.Mode()&os.ModeSymlink != 0 {
		// fmt.Print(os.ModeSymlink)
		return true
	}

	if strings.Contains(path, ".ink") {
		return true
	}

	return false
}

func totalSize(fileInfo []fs.FileInfo, path string) int64 {
	size := int64(0)
	for i := 0; i < len(fileInfo); i++ {
		if fileInfo[i].Name()[0] != '.' || a_flag {
			subPath := ""

			if fileInfo[i].IsDir() {
				subPath = returnPath(fileInfo[i].Name(), path)
			} else {
				subPath = fileInfo[i].Name()
			}

			fileInfo, err := os.Stat(subPath)

			if err == nil {
				stat, ok := fileInfo.Sys().(*syscall.Stat_t)
				if ok {
					blockSize := int64(stat.Blksize)
					blocks := stat.Blocks
					diskUsage := blocks * blockSize
					size += diskUsage
				}
			}
			
		}
	}

	return size/8192
}
