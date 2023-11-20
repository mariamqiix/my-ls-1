package Myls

import (
	"io/fs"
	"os"
	"syscall"
)

func TotalSize(fileInfo []fs.FileInfo, path string) int64 {
	size := int64(0)
	for i := 0; i < len(fileInfo); i++ {
		if fileInfo[i].Name()[0] != '.' || a_flag {
			subPath := ""
			if fileInfo[i].IsDir() {
				subPath = ReturnPath(fileInfo[i].Name(), path)
			} else {
				subPath = fileInfo[i].Name()
			}

			fileInfo, err := os.Stat(subPath)

			if err == nil {
				stat, ok := fileInfo.Sys().(*syscall.Stat_t)
				if ok {
					size += stat.Blocks
				}
			}

		}
	}

	return size / 2
}
