package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	name  string
	files []*File
	size  int

	parent      *Directory
	directories []*Directory
}

func cd(current *Directory, path string) *Directory {
	if path == ".." {
		return current.parent
	}
	for _, dir := range current.directories {
		if dir.name == path {
			return dir
		}
	}
	log.Fatalf("cd: directory %s not found\n", path)
	return nil
}

func addDir(current *Directory, name string) {
	if current.directories == nil {
		directories := make([]*Directory, 0, 8)
		current.directories = directories
	}
	current.directories = append(current.directories, &Directory{
		name:        name,
		parent:      current,
		directories: []*Directory{},
	})
}

func addFile(current *Directory, name, size string) {
	if current.files == nil {
		files := make([]*File, 0, 8)
		current.files = files
	}
	if s, err := strconv.Atoi(size); err == nil {
		current.files = append(current.files, &File{
			name: name,
			size: s,
		})
	} else {
		fmt.Printf("Could not convert %s to int\n", size)
	}
}

// sizeDirectoryTree appends the size to each directory struct and then appends a pointer to that struct to the passed in array
func sizeDirectoryTree(current *Directory, dirArray *[]*Directory) {
	if current.directories != nil {
		for _, dir := range current.directories {
			sizeDirectoryTree(dir, dirArray)
		}
	}
	if current.files != nil {
		for _, f := range current.files {
			current.size += f.size
		}
	}
	if current.directories != nil {
		for _, d := range current.directories {
			current.size += d.size
		}
	}
	*dirArray = append(*dirArray, current)
}

func getArrayKeys(arr []*Directory) []int {
	var ret []int
	for i := range arr {
		ret = append(ret, i)
	}
	return ret
}

type File struct {
	name string
	size int
}

const input = "./7/input.txt"
const totalDiskSpace = 70_000_000
const neededSpace = 30_000_000

func main() {
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	inputLines := strings.Split(bytes.NewBuffer(data).String(), "\n")
	root := &Directory{
		name:        "/",
		files:       nil,
		parent:      nil,
		directories: nil,
	}

	current := root

	for _, line := range inputLines {
		if strings.HasPrefix(line, "$ cd /") {
			current = root
			continue
		}

		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "$":
			switch tokens[1] {
			case "cd":
				if len(tokens) < 3 {
					log.Fatalf("No dir name after command")
				}
				current = cd(current, tokens[2])
			case "ls":
				continue
			}
		case "dir":
			if len(tokens) < 2 {
				log.Fatalf("No dir name after command")
			}
			addDir(current, tokens[1])
		default:
			if len(tokens) < 2 {
				log.Fatalf("No file name after size")
			}
			addFile(current, tokens[1], tokens[0])
		}
	}
	//printTree(root)
	dirs := make([]*Directory, 0, 64)
	sizeDirectoryTree(root, &dirs)
	partOne := 0
	var totalUsed int
	for _, d := range dirs {
		if d.name == "/" {
			totalUsed = d.size
		}
		if d.size <= 100_000 {
			partOne += d.size
		}
	}
	fmt.Printf("solution to part 1: %d\n", partOne)

	// Part 2
	if totalUsed > totalDiskSpace {
		log.Fatalf("something is horribly wrong. Total:: %d\t\tTotal Used :: %d\n", totalDiskSpace, totalUsed)
	}
	neededDeleteSize := neededSpace - (totalDiskSpace - totalUsed)
	partTwoDirs := make([]*Directory, 0, 64)
	for _, d := range dirs {
		if d.size > neededDeleteSize {
			partTwoDirs = append(partTwoDirs, d)
		}
	}
	keys := getArrayKeys(partTwoDirs)
	sort.SliceStable(keys, func(i, j int) bool {
		return partTwoDirs[keys[i]].size < partTwoDirs[keys[j]].size
	})

	var smallestToDelete *Directory
	for _, k := range keys {
		if partTwoDirs[k].size > neededDeleteSize {
			smallestToDelete = partTwoDirs[k]
			break
		}
	}
	fmt.Printf("Size of smallest needed file for free enough space is: \ndir :: %s\t\tsize :: %d\n", smallestToDelete.name, smallestToDelete.size)
}
