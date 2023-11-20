package Myls

import (
	"fmt"
	"io/fs"
)

func SortByAlph(filesInfos []fs.FileInfo) []fs.FileInfo {
	for i := 0; i < len(filesInfos); i++ {
		for j := i + 1; j < len(filesInfos); j++ {
			if fmt.Sprintf("%s/  ", filesInfos[j].Name()) < fmt.Sprintf("%s/  ", filesInfos[i].Name()) {
				filesInfos[i], filesInfos[j] = filesInfos[j], filesInfos[i]
			}
		}
	}
	return filesInfos
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
