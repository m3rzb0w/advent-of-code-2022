package main

import (
	"fmt"
	fetch "getdata"
	"strings"
)

var url string = "https://adventofcode.com/2022/day/12/input"

var session = fetch.Grabsession()

var input = fetch.Getdata(url, session)

//var ex string = "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"

type Coord struct {
	x, y int
}

type Item struct {
	Coord
	val int
}

type Queue struct {
	items  []Item
	length int
}

func (q *Queue) Enqueue(i Item) {
	q.items = append(q.items, i)
	q.length++
}

func (q *Queue) Dequeue() Item {
	toRemove := q.items[0]
	q.items = q.items[1:]
	q.length--
	return toRemove
}

func (q *Queue) Front() Item {
	front := q.items[0]
	return front
}

func (q *Queue) Len() int {
	return q.length
}

var start, end Coord
var ans, ansTwo int
var maze = make([][]rune, 0)

func main() {
	inputTrim := strings.Trim(input, "\n")
	dataSplit := strings.Split(inputTrim, "\n")

	for r, row := range dataSplit {
		var line []rune
		for c, item := range row {
			if item == 'S' {
				start = Coord{c, r}
				item = 'a'
			}
			if item == 'E' {
				end = Coord{c, r}
				item = 'z'
			}
			line = append(line, item)
		}
		maze = append(maze, line)
	}
	for _, v := range maze {
		fmt.Println(string(v))
	}
	visited := make(map[Coord]bool)
	visited[start] = true
	q := Queue{}
	q.Enqueue(Item{start, 0})
	// fmt.Println("Q state =>", q)
	for q.Len() != 0 {
		fmt.Println("Q state =>", q)
		current := q.Front()
		currVal := current.val
		q.Dequeue()
		for _, near := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			j, i := near[1], near[0]
			nextNode := Coord{current.x + j, current.y + i}
			if nextNode.x < 0 || nextNode.y < 0 || nextNode.x >= len(maze[0]) || nextNode.y >= len(maze) {
				continue
			}
			if visited[nextNode] {
				continue
			}
			if maze[nextNode.y][nextNode.x]-maze[current.y][current.x] > 1 {
				continue
			}
			if nextNode.x == end.x && nextNode.y == end.y {

				// fmt.Println(current)
				// fmt.Println("ANS=>", currVal+1)
				// fmt.Println(start, end)
				ans = currVal + 1
				//return
			}
			fmt.Println("NEXTNODE=>", nextNode)
			visited[nextNode] = true
			q.Enqueue(Item{nextNode, currVal + 1})
		}
	}
	//part2
	visited = make(map[Coord]bool)
	visited[end] = true
	q = Queue{}
	q.Enqueue(Item{end, 0})
	for q.Len() != 0 {
		fmt.Println("Q state =>", q)
		current := q.Front()
		currVal := current.val
		q.Dequeue()
		for _, near := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			j, i := near[1], near[0]
			nextNode := Coord{current.x + j, current.y + i}
			if nextNode.x < 0 || nextNode.y < 0 || nextNode.x >= len(maze[0]) || nextNode.y >= len(maze) {
				continue
			}
			if visited[nextNode] {
				continue
			}
			if maze[nextNode.y][nextNode.x]-maze[current.y][current.x] < -1 {
				continue
			}
			if maze[nextNode.y][nextNode.x] == 'a' && ansTwo == 0 {

				// fmt.Println(current)
				// fmt.Println("ANS=>", currVal+1)
				// fmt.Println(start, end)
				//return
				ansTwo = currVal + 1
			}
			fmt.Println("NEXTNODE=>", nextNode)
			visited[nextNode] = true
			q.Enqueue(Item{nextNode, currVal + 1})
		}
	}
	fmt.Println("part1", ans)
	fmt.Println("part2", ansTwo)

}
