package main

import (
	"fmt"
	"os"
)

func main() {
	// create file
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString("test")

	// open file
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println("filename in directory :", fi.Name())
	}

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
