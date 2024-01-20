package Myls

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func SortByAlph(filesInfos []fs.FileInfo) []fs.FileInfo {
	n := len(filesInfos)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if strings.ToLower(strings.ReplaceAll(filesInfos[j].Name(),"-","")) > strings.ToLower(strings.ReplaceAll(filesInfos[j+1].Name(),"-","")) {
				filesInfos[j], filesInfos[j+1] = filesInfos[j+1], filesInfos[j]
			}
		}
	}
	return filesInfos
}

func SortWithTildeFirst(paths []string) []string {
	var dotSlashPaths []string
	var otherPaths []string

	for _, path := range paths {
		if path == THePath() {
			dotSlashPaths = append(dotSlashPaths, path)
		} else {
			otherPaths = append(otherPaths, path)
		}
	}

	return append(dotSlashPaths, otherPaths...)
}

func THePath() string {
	homeDir := os.Getenv("HOME")
	return fmt.Sprint(homeDir)
}

// func SortByAlph(filesInfos []fs.FileInfo) []fs.FileInfo {
// 	n := len(filesInfos)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			nameI := filesInfos[j].Name()
// 			nameJ := filesInfos[j+1].Name()

// 			// Ignore the first "." in the word if it is not equal to "." or ".."
// 			if nameI != "." && nameI != ".." && strings.HasPrefix(nameI, ".") {
// 				nameI = nameI[1:]
// 			}
// 			if nameJ != "." && nameJ != ".." && strings.HasPrefix(nameJ, ".") {
// 				nameJ = nameJ[1:]
// 			}

// 			// Compare the modified names in a case-insensitive manner
// 			if strings.ToLower(nameI) > strings.ToLower(nameJ) {
// 				filesInfos[j], filesInfos[j+1] = filesInfos[j+1], filesInfos[j]
// 			}
// 		}
// 	}

// 	return filesInfos
// }

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
