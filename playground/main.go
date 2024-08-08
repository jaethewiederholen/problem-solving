package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	
	park := []string{"SOO","OXX","OOO"}
	routes := []string {"E 2","S 2","W 1"}
	fmt.Println(solution(park, routes))
}

func solution(park []string, routes []string) []int {
	// location[세로,가로]
	var location []int
    // make park two-dimension
    var parkTd [][]string
    for _, r := range park {
        row := strings.Split(r, "")
        parkTd = append(parkTd, row)
    }
	fmt.Println(parkTd)
	// find start
	for y, r := range parkTd {
		for x, c := range r {
			if c == "S" {
				location = []int{x,y}
			}
		}
	}

	for _, r := range routes {
		fmt.Println(location)
		var newX, newY int
		newX, newY = move(location[0], location[1], r)
		fmt.Println(newX, newY)
		if newX >= 0 && newY >= 0 && parkTd[newX][newY] != "X" {
			location[0], location[1] = newX, newY 	
		}
	}
	
	return location
}

func move(x, y int, order string) (int,int){
    route := strings.Split(order, " ")
	op := route[0]
	n, _ := strconv.Atoi(route[1])

	switch op {
	case "N":
		x -= n
	case "S" :
		x += n
	case "W" :
		y -= n
	case "E" :
		y += n
	}

	return x, y
}