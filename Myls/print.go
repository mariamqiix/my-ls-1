package Myls

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"strings"
)

// type TheFiles struct

func RunLS() {
	pathname, subFile := Validation()
	if SubFile_flag {
		Print("./", subFile, Listing("./"))
	}
	pathname = SortWithTildeFirst(pathname)
	for i := 0; i < len(pathname); i++ {
		if SubFile_flag == false || pathname[0] != "./" {
			if R_flag || (len(pathname) > 1 && l_flag) || (SubFile_flag && l_flag){
				if (SubFile_flag && l_flag) {
					fmt.Println()
				}
				fmt.Println(pathname[i] + ":")
			}
			SubFile_flag = false
			var subFiles []string
			Print(pathname[i], subFiles, Listing(pathname[i]))
		}
		if len(pathname) != i+1 {
			fmt.Println()
		}
	}
}

func Print(path string, subFile []string, fileInfos []fs.FileInfo) {
	if SubFile_flag {
		fileInfos = fileFilter(subFile, fileInfos)
	}

	if len(fileInfos) == 0 && SubFile_flag {
		return
	}
	if !SubFile_flag {
	file2, err := os.Stat(path + "/..")
	if err == nil {
		fileInfos = append([]fs.FileInfo{file2}, fileInfos...)
	}

	file1, err := os.Stat(path + "/.")
	if err == nil {
		fileInfos = append([]fs.FileInfo{file1}, fileInfos...)
	}
	}

	if l_flag && !SubFile_flag {
		fmt.Println("total", math.Round(float64(TotalSize(fileInfos, path))))
	}
	fileInfos = SortByAlph(fileInfos)

	if t_flag {
		fileInfos = SortByDate(fileInfos)

	}

	max := maxSize(fileInfos)

	if !l_flag {
		Names := FormatNames(fileInfos)
		for i := 0; i < len(Names) && (len(Names) == 1 && len(Names[0]) != 0 ); i++ {
			for j := 0; j < len(Names[i]); j++ {
				fmt.Print(Names[i][j] + "  ")
			} 
			fmt.Println()
		}

	} else {

		for _, c := range fileInfos {
			if isBlockDevice(c) {
				divFile = true
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
					lFlag(path, fmt.Sprint(max), fileInfo)
				}

				if fileInfo.IsDir() {
					fmt.Print("\033[34m", fmt.Sprintf("%s  ", "\033[1m"+fileInfo.Name()), "\033[0m", "")
					// blue, bold, reset
					// fmt.Printf("\033[34,1%s\033[0m", fileInfo.Name())

				} else {
					subPath := ReturnPath(fileInfo.Name(), path)
					if CheckShortCut(subPath) && l_flag {
						shortName, _ := os.Readlink(subPath)
						// fmt.Print("\033[36m", fileInfo.Name()+" ", "\033[0m", "-> ", "\033[34m", shortName, "\033[0m")
						// Cyan, reset, blue, reset
						fmt.Printf("\033[36m%s \033[0m-> \033[34m%s\033[0m", fileInfo.Name(), shortName)

					} else {

						// fmt.Print("\033[0m", fileInfo.Name()+"  ", "\033[0m", "")
						fmt.Printf("%s ", fileInfo.Name())
					}

				}

				if l_flag {
					fmt.Println()
				}
			}
		}
	}

	if R_flag {
		Rflag(path, fileInfos)
	}
}

func lFlag(path, maxSize string, fileInfo fs.FileInfo) {
	Gid, UserId, filelinks, divInfo := ReturnGroupAndUSerId(path + "/" + fileInfo.Name())
	mode := fmt.Sprint(fileInfo.Mode())
	DateAndTime := fmt.Sprintf("%s %s %02d:%02d", fileInfo.ModTime().Month().String()[:3], FormatDate(fileInfo.ModTime().Day()), fileInfo.ModTime().Hour(), fileInfo.ModTime().Minute())
	size := FormatSize(fmt.Sprint(fileInfo.Size()), maxSize)
	if mode == "grwxr-xr-x" || mode == "urwxr-xr-x" {
		mode = "-rwxr-xr-x"
	}
	if strings.Contains(mode, "Drw-rw-") || strings.Contains(mode, "Dcrw--") || strings.Contains(mode, "Dcrw-") {
		if Gid == "disk" {
			mode = strings.ReplaceAll(mode, "Dc", "b")
			mode = strings.ReplaceAll(mode, "D", "b")
		} else {
			mode = strings.Replace(mode, "D", "", 1)
		}
	} else {
		if mode == "dtrwxrwxrwx" {
			mode = "drwxr-xr-x"
		}
	}
	filelinks = strings.Repeat(" ", 3-len(filelinks)) + filelinks

	if len(mode) < 10 {
		mode += strings.Repeat(" ", 10-len(mode)+1)
	}

	if len(Gid) <= 7 {
		Gid += strings.Repeat(" ", 7-len(Gid)+1)
	}

	if divFile {
		if divInfo == "0" {
			divInfo = size
		}
		if len(divInfo) < 8 {
			divInfo = strings.Repeat(" ", 8-len(divInfo)) + divInfo
		}
		size = divInfo

	}

	fmt.Print(strings.ToLower(mode) + " " + filelinks + " " + UserId + " " + Gid + " " + size + " " + DateAndTime + " ")
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
				files := Listing(subPath)
				fmt.Println("\n" + subPath + ":")
				fmt.Print("\033[97m", "")
				if len(files) != 0 || a_flag {
					var subFile []string
					Print(subPath, subFile, Listing(subPath))
				}
			}
		}
	}
}

func fileFilter(subfile []string, files []fs.FileInfo) []fs.FileInfo {
	var files2 []fs.FileInfo
	for i := 0; i < len(subfile); i++ {
		flag := false
		for _, file := range files {
			if file.Name() == subfile[i] {
				flag = true
				files2 = append(files2, file)
			}
		}
		if !flag {
			fmt.Println("ls: cannot access '" + subfile[i] + "': No such file or directory")

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

func maxNameSize(files []fs.FileInfo) int {
	max := 0
	for _, file := range files {
		if len(file.Name()) > max {
			max = len(file.Name())
		}
	}
	return max
}

func width() int {
	// cmd := exec.Command("stty", "size")
	// cmd.Stdin = os.Stdin
	// out, err := cmd.Output()
	// if err == nil {
	// 	s := string(out)
	// 	s = strings.TrimSpace(s)
	// 	sArr := strings.Split(s, " ")
	// 	width, err := strconv.Atoi(sArr[1])
	// 	fmt.Print(width)
	// 	if err == nil {
	// 		return width
	// 	}
	// }
	return 179
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func FormatNames(files []fs.FileInfo) [][]string {
	width := width()
	totalLen := 0
	for _, file := range files {
		totalLen += len(file.Name())
	}
	if totalLen < width {
		var Names [][]string
		var Name []string
		for i := 0; i < len(files); i++ {
			index := i
			if r_flag {
				index = len(files) - i - 1
			}
			if files[index].Name()[0] != '.' || a_flag {
				Name = append(Name, files[index].Name())
			}
		}
		Names = append(Names, Name)
		return Names
	}

	return Return2DArrayLen(files)

}

func ArrayWidth(files []fs.FileInfo) int {
	width := width()
	maxNameSize := maxNameSize(files)
	for i := 1; i <= 20; i++ {
		if (maxNameSize+1)*i > width {
			return i - 1

		}
	}
	return 0
}

func ArrayLenght(filesNumbers, ArrayWidth int) int {
	if ArrayWidth != 0 {
		ArrayLenght := filesNumbers / ArrayWidth
		if filesNumbers%ArrayWidth != 0 {
			ArrayLenght += 1
		}
		return ArrayLenght
	}
	return 0
}

func Return2DArrayLen(files []fs.FileInfo) [][]string {
	ArrayWidth := ArrayWidth(files)
	ArrayLenght := ArrayLenght(len(files), ArrayWidth)
	fileNum := 0
	maxNameSize := maxNameSize(files)
	Names := make([][]string, ArrayLenght)
	for i := 0; i < ArrayLenght; i++ {
		Names[i] = make([]string, ArrayWidth)
	}
	for i := 0; i < ArrayWidth; i++ {
		for j := 0; j < ArrayLenght; j++ {
			file := ""
			if r_flag {
				file = files[len(files)-fileNum-1].Name()
			} else {
				file = files[fileNum].Name()
			}
			if len(file) < maxNameSize {
				file = file + strings.Repeat(" ", maxNameSize-len(file))
			}
			if file[0] != '.' || a_flag {
				Names[j][i] = file
			} else {
				j--
			}
			fileNum++
			if fileNum == len(files)-1 {
				return Names
			}
		}
	}

	return Names
}
