package main

import (
	"fmt"
	"os"
	"log"
)

var l_flag bool
var R_flag bool
var a_flag bool
var r_flag bool
var t_flag bool




func main(){
	pathname := validation()
	listing(pathname)
	if R_flag {
		err := listFilesRecursively(pathname, "")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func validation() (string){
	l_flag, R_flag, a_flag, r_flag, t_flag, pathname := false, false, false, false, false, "."

	if len(os.Args) < 2 || os.Args[1] != "ls" {
		fmt.Println("Error, read the manPage 1")
		os.Exit(0)
	} else if len(os.Args) == 3 && os.Args[2][0] == '-' {
			for j :=1; j<len(os.Args[2]); j++ {
				if !l_flag && os.Args[2][j] == 'l' {
					l_flag = true
				} else if !R_flag && os.Args[2][j] == 'R' {
					R_flag = true
				} else if !a_flag && os.Args[2][j] == 'a'{
					a_flag = true
				} else if !r_flag && os.Args[2][j] == 'r'{
					r_flag = true
				} else if !t_flag && os.Args[2][j] == 't'{
					t_flag = true
				} else {
					fmt.Println("Error, read the manPage 2")
					os.Exit(0)
				}
			}
	} else if len(os.Args) !=2 {
		fmt.Println("Error, read the manPage 3")
		os.Exit(0)
	}

	if len(os.Args) == 4 {
		pathname = os.Args[3]
	}

	return pathname
}

func listing(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.Name()[0] != '.' && !a_flag {
			if fileInfo.IsDir() {
				fmt.Printf("[%s]\n", fileInfo.Name())
			} else {
				fmt.Print(fileInfo.Name(), " ")
			}
		}
	}
	fmt.Println()
}

func listFilesRecursively(dirPath, indent string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()

		if fileName != "." && fileName != ".." {
			fmt.Print(indent)
			if fileInfo.IsDir() {
				fmt.Printf("[%s]\n", fileName)
				subDirPath := dirPath + "/" + fileName
				subIndent := indent + "  "
				err := listFilesRecursively(subDirPath, subIndent)
				if err != nil {
					return err
				}
			} else {
				fmt.Println(fileName)
			}
		} else {
			fmt.Print(indent)

		}
	}

	return nil
}