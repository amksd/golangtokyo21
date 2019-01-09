package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

///////////// Sample Outputs
//  $ tree .
//  .
//  ├── dir1
//  │   ├── dir11
//  │   │   └── file3
//  │   ├── dir12
//  │   │   └── file4
//  │   └── file2
//  └── file1

const (
	charMiddle = "├── "
	charBottom = "└── "
	charFill   = "│   "
)

var (
	depth, width int
)

func main() {
	extract("./", os.Args[1], 0, 0, false)
}

func extract(filePath, fileName string, depth, width int, tail bool) {
	fmt.Println(line(fileName, depth, width, tail))

	fullName := filePath + fileName
	info, err := os.Stat(fullName)
	if err != nil {
		panic(err)
	}

	if !info.IsDir() {
		return
	}
	width += 1

	fs, err := ioutil.ReadDir(fullName)
	if err != nil {
		panic(err)
	}
	for i, f := range fs {
		tail := false
		if i == len(fs)-1 {
			tail = true
		}
		extract(fullName+"/", f.Name(), depth+1, width, tail)
	}
	width -= 1
	return
}

func line(name string, depth, width int, tail bool) string {
	if depth == 0 {
		return name
	}
	base := strings.Repeat(charFill, width-1)
	if tail {
		return base + charBottom + name
	}
	return base + charMiddle + name
}
