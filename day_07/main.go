package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	fetch "getdata"
)

var url string = "https://adventofcode.com/2022/day/7/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

func addSizeToDirectoriesPaths(fileSize int, path []string, directories map[string]int) {
	fmt.Println("ADD PATH", path)
	for i := range path {
		directoryPath := strings.Join(path[0:i+1], "/")
		directories[directoryPath] += fileSize
		fmt.Println("Adding value to => ", directoryPath, fileSize)
	}
}

func sortDirectoriesPaths(directories map[string]int) []string {
	var sortedDirectoriesPaths []string

	for key := range directories {
		fmt.Println("KEYYYYYYYY==> ", key)
		sortedDirectoriesPaths = append(sortedDirectoriesPaths, key)
		//fmt.Println("before PATH==> ", sortedDirectoriesPaths)
	}
	fmt.Println("SORTED PATH==> ", sortedDirectoriesPaths)
	sort.SliceStable(sortedDirectoriesPaths, func(i, j int) bool {
		return directories[sortedDirectoriesPaths[i]] < directories[sortedDirectoriesPaths[j]]
	})

	return sortedDirectoriesPaths
}

// part Two

func main() {
	var path []string
	directories := make(map[string]int)
	fmt.Println(directories)
	dataSplit := strings.Split(input, "\n")
	fmt.Println(dataSplit)
	for _, v := range dataSplit {
		//fmt.Println(v)
		tmp := strings.Trim(v, " ")
		tmpSplit := strings.Split(tmp, " ")
		//fmt.Println(tmpSplit, len(tmpSplit))
		if len(tmpSplit) == 3 && tmpSplit[1] == "cd" && tmpSplit[2] != ".." {
			// fmt.Println(tmpSplit)
			// folderMap.key = tmpSplit[2]
			folderName := tmpSplit[2]
			path = append(path, folderName)
			fmt.Println("FIRST PATH", path)
		} else if len(tmpSplit) == 3 && tmpSplit[1] == "cd" && tmpSplit[2] == ".." {
			path = path[0 : len(path)-1]
			fmt.Println("SECOND PATH", path)

		}
		if len(tmpSplit) == 2 && tmpSplit[0] != "dir" && tmpSplit[1] != "ls" {
			fileSizeInt, _ := strconv.Atoi(tmpSplit[0])
			fmt.Println(fileSizeInt)
			addSizeToDirectoriesPaths(fileSizeInt, path, directories)
		}
		fmt.Println(directories)
	}
	var ans int
	for _, v := range directories {
		if v <= 100000 {
			ans += v
		}
	}
	fmt.Println("Ans Part One => ", ans)
	AllFileSize := directories["/"]
	freeSpace := 70000000 - AllFileSize
	spaceToDelete := 30000000 - freeSpace

	fmt.Println("SPACE TO DELETE", spaceToDelete)

	sortedDirectoriesPaths := sortDirectoriesPaths(directories)

	fmt.Println(sortedDirectoriesPaths)

	for _, directoryPath := range sortedDirectoriesPaths {
		directorySize := directories[directoryPath]
		if directorySize >= spaceToDelete {
			fmt.Println("Part Two Answer:", directorySize)
			break
		}
	}
	fmt.Println(sortedDirectoriesPaths)

}
