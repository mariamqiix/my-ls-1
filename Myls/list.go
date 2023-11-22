package Myls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func Listing(dir string) []fs.FileInfo {
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

func ReturnPath(fileName, path string) string {
	if path != "./" {
		return path + "/" + fileName
	}
	
	return path + fileName
}

func ReturnGroupAndUSerId(path string) (string, string, string) {
	file_info, err := os.Lstat(path)
	if err != nil {
		fmt.Print(file_info)
		log.Fatal(err)
	}

	file_sys := file_info.Sys()
	flink := fmt.Sprint(file_sys.(*syscall.Stat_t).Nlink)

	grId, _ := user.LookupGroupId(fmt.Sprint(file_sys.(*syscall.Stat_t).Gid))
	usrId, _ := user.LookupId(fmt.Sprint(file_sys.(*syscall.Stat_t).Uid))

	return grId.Name, usrId.Username, flink
}

func CheckShortCut(path string) bool {
	file_info, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}

	if file_info.Mode()&os.ModeSymlink != 0 || strings.Contains(path, ".ink") {
		return true
	}

	return false
}
