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
		if strings.Contains(fmt.Sprint(err),"permission denied")  {
		fmt.Print(dir)
		} else {
		log.Fatal(err)
		}
	}

	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		if strings.Contains(fmt.Sprint(err),"permission denied")  {
			fmt.Print(dir)
			} else {
			log.Fatal(err)
			}
	}

	return fileInfos
}

func ReturnPath(fileName, path string) string {
	if path != "./" && rune(path[len(path)-1]) != '/' {
		return path + "/" + fileName
	}

	return path + fileName
}

func Major(dev uint64) uint32 {
	major := uint32((dev & 0x00000000000fff00) >> 8)
	major |= uint32((dev & 0xfffff00000000000) >> 32)
	return major
}
// Minor returns the minor component of a Linux device number.
func Minor(dev uint64) uint32 {
	minor := uint32((dev & 0x00000000000000ff) >> 0)
	minor |= uint32((dev & 0x00000ffffff00000) >> 12)
	return minor
}

func isBlockDevice(fileInfo os.FileInfo) bool {
	return fileInfo.Mode()&os.ModeDevice == os.ModeDevice && fileInfo.Mode()&os.ModeCharDevice == 0
}

func ReturnGroupAndUSerId(path string) (string, string, string,string ) {
	file_info, err := os.Lstat(path)
	if err != nil {
		fmt.Print(file_info)
		log.Fatal(err)
	}

	file_sys := file_info.Sys().(*syscall.Stat_t)
	flink := fmt.Sprint(file_sys.Nlink)

    dev := file_sys.Rdev
	divInfo := ""
	if  Major(dev) != 0 {
    divInfo = fmt.Sprintf( "%d, %d", Major(dev), Minor(dev))
	} else {
		divInfo = fmt.Sprintf( "%d", Minor(dev))

	}

	grId, _ := user.LookupGroupId(fmt.Sprint(file_sys.Gid))
	usrId, _ := user.LookupId(fmt.Sprint(file_sys.Uid))

	return grId.Name, usrId.Username, flink , divInfo
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
