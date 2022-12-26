package main

import (
	"fmt"
	fetch "getdata"
)

var url string = "https://adventofcode.com/2022/day/13/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

func main() {
	fmt.Println(input)
}
