package Myls

import (
	"fmt"
	"io/fs"
	"log"
	"math"
	"os"
	"strings"
)

func Print(path, subFile string, fileInfos []fs.FileInfo) {

	fileInfos = SortByAlph(fileInfos)

	if SubFile_flag {
		fileInfos = fileFilter(subFile, fileInfos)
	}

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

		fmt.Println("total", math.Round(float64(TotalSize(fileInfos, path))))
	}

	if t_flag {
		fileInfos = SortByDate(fileInfos)

	}

	if a_flag && !r_flag && !R_flag && !l_flag {
		fmt.Print("../  ./ ")
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
				fmt.Print("\033[34m", fmt.Sprintf("%s/  ", "\033[1m"+fileInfo.Name())+"\033[0m", "\033[97m", "")

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
	if a_flag && r_flag && !l_flag {
		fmt.Println("./  ../  ")
	}

	if R_flag {
		Rflag(path, fileInfos)
	}
}

func lFlag(path string, fileInfo fs.FileInfo) {
	Gid, UserId, filelinks := ReturnGroupAndUSerId(path + "/" + fileInfo.Name())
	mode := fmt.Sprint(fileInfo.Mode())

	DateAndTime := fmt.Sprintf("%s %d %02d:%02d", fileInfo.ModTime().Month().String()[:3], fileInfo.ModTime().Day(), fileInfo.ModTime().Hour(), fileInfo.ModTime().Minute())
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
