package Myls

import (
	// "fmt"
	"io/fs"
	"strings"
)
func SortByAlph(filesInfos []fs.FileInfo) []fs.FileInfo {
	n := len(filesInfos)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if strings.ToLower(filesInfos[j].Name()) > strings.ToLower(filesInfos[j+1].Name()) {
				filesInfos[j], filesInfos[j+1] = filesInfos[j+1], filesInfos[j]
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
