package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//  $ tree .
//  .
//  ├── dir1
//  │   ├── dir11
//  │   │   └── file3
//  │   ├── dir12
//  │   │   └── file4
//  │   └── file2
//  └── file1

type file struct {
	name  string
	depth int
	width int
	files []file
}

const (
	charMiddle = "├── "
	charBottom = "└── "
	charFill   = "│  "
)

var (
	rows         = []string{}
	depth, width int
)

func main() {
	extract("./", ".", 0, 0, false)
	for _, row := range rows {
		fmt.Println(row)
	}
}

func extract(filePath, fileName string, depth, width int, tail bool) file {
	fullName := filePath + fileName
	n := file{name: fileName, depth: depth, width: width, files: []file{}}
	//todo: filename trim or better way
	rows = append(rows, line(fileName, depth, width, tail))
	info, err := os.Stat(fullName)
	if err != nil {
		panic(err)
	}
	if !info.IsDir() {
		return n
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
		n.files = append(n.files, extract(fullName+"/", f.Name(), depth+1, width, tail))
	}
	width -= 1
	return n
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
