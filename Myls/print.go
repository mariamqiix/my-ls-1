package Myls

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"strings"
)

func Print(path, subFile string, fileInfos []fs.FileInfo) {
	if l_flag && (path == "../" || path == "./") && !SubFile_flag {
		fmt.Println(path, ": ")
	}
	fileInfos = SortByAlph(fileInfos)

	if SubFile_flag {
		fileInfos = fileFilter(subFile, fileInfos)
	}

	if len(fileInfos) == 0 {
		fmt.Println("ls: cannot access '" + subFile + "': No such file or directory")
		os.Exit(0)

	}

	file2, err := os.Stat("..")
	if err == nil {
		fileInfos = append([]fs.FileInfo{file2}, fileInfos...)
	}

	file1, err := os.Stat(".")
	if err == nil {
		fileInfos = append([]fs.FileInfo{file1}, fileInfos...)
	}

	if l_flag {
		fmt.Println("total", math.Round(float64(TotalSize(fileInfos, path))))
	}

	if t_flag {
		fileInfos = SortByDate(fileInfos)

	}

	max := maxSize(fileInfos)

	for i := 0; i < len(fileInfos); i++ {

		index := i
		if r_flag {
			index = len(fileInfos) - i - 1
		}

		fileInfo := fileInfos[index]

		if fileInfo.Name()[0] != '.' || a_flag {

			if l_flag {
				lFlag(path, fmt.Sprint(max), fileInfo)
			}

			if fileInfo.IsDir() {
				fmt.Print("\033[34m", fmt.Sprintf("%s  ", "\033[1m"+fileInfo.Name())+"\033[0m", "\033[97m", "")

			} else {
				subPath := ReturnPath(fileInfo.Name(), path)
				if CheckShortCut(subPath) && l_flag {
					shortName, _ := os.Readlink(subPath)
					fmt.Print("\033[36m", fileInfo.Name()+" ", "\033[97m", "-> ", "\033[34m", shortName, "\033[97m")

				} else {

					fmt.Print("\033[97m", fileInfo.Name()+"  ", "\033[97m", "")
				}

			}

			if l_flag {
				fmt.Println()
			}
		}
	}
	if !l_flag {
		fmt.Println()
	}

	if R_flag {
		Rflag(path, fileInfos)
	}
}

func lFlag(path, maxSize string, fileInfo fs.FileInfo) {
	Gid, UserId, filelinks := ReturnGroupAndUSerId(path + "/" + fileInfo.Name())
	mode := fmt.Sprint(fileInfo.Mode())
	DateAndTime := fmt.Sprintf("%s %s %02d:%02d", fileInfo.ModTime().Month().String()[:3], FormatDate(fileInfo.ModTime().Day()), fileInfo.ModTime().Hour(), fileInfo.ModTime().Minute())
	size := FormatSize(fmt.Sprint(fileInfo.Size()), maxSize)
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

func Rflag(path string, fileInfos []fs.FileInfo) {
	for i := 0; i < len(fileInfos); i++ {
		index := i
		if r_flag {
			index = len(fileInfos) - i - 1
		}
		fileInfo := fileInfos[index]

		if fileInfo.Name()[0] != '.' || a_flag && fileInfo.Name() != ".." && fileInfo.Name() != "." {
			if fileInfo.IsDir() && fileInfo.Name() != "WinSAT" {
				subPath := ReturnPath(fileInfo.Name(), path)
				fmt.Println("\n" + subPath + ":")
				fmt.Print("\033[97m", "")
				Print(subPath, "", Listing(subPath))
			}
		}
	}
}

func fileFilter(subfile string, files []fs.FileInfo) []fs.FileInfo {
	var files2 []fs.FileInfo

	for _, file := range files {
		if file.Name() == subfile {
			files2 = append(files2, file)
		}
	}

	return files2
}

func maxSize(files []fs.FileInfo) int {
	max := 0
	for _, file := range files {
		if int(file.Size()) > max {
			max = int(file.Size())
		}
	}
	return max
}

func FormatSize(size, MaxSize string) string {
	if len(size) < len(MaxSize) {
		size = strings.Repeat(" ", len(MaxSize)-len(size)) + size
	}
	return size
}

func FormatDate(date int) string {
	if date < 10 {
		return " " + fmt.Sprint(date)
	}
	return fmt.Sprint(date)
}
