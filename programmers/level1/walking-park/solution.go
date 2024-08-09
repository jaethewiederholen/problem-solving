package walkingPark

import (
	"fmt"
	"strconv"
	"strings"
)


func Solution(park []string, routes []string) []int {
	// location[세로,가로]
	var location []int
    // make park two-dimension
    var parkTd [][]string
    for _, r := range park {
        row := strings.Split(r, "")
        parkTd = append(parkTd, row)
    }
	// x 가 세로, y 가 가로
	// find start
	for x, r := range parkTd {
		for y, c := range r {
			if c == "S" {
				location = []int{x,y}
			}
		}
	}
	fmt.Println("start : ", location)


	for _, r := range routes { 
		//var newX, newY int
		location[0], location[1] = move(location[0], location[1], r, parkTd)
		fmt.Println(location)
	}
	return location

}

func move(x, y int, order string, parkTd [][]string) (int,int){
    route := strings.Split(order, " ")
	op := route[0]
	n, _ := strconv.Atoi(route[1])

	switch op {
	case "N":
		if x - n < 0 {
			return x, y
		}
		for i:=x-1; i >= x-n ; i-- {
			if parkTd[i][y] == "X"{
				return x, y
			}
		}
		x -= n
	case "S" :
		if x + n >= len(parkTd) {
			return x, y
		}
		for i:=x+1; i <= x+n ; i++ {
			if parkTd[i][y] == "X"{
				return x, y
			}
		}
		x += n
	case "W" :
		if y - n < 0 {
			return x, y
		}
		for i:=y-1; i >= y-n ; i-- {
			if parkTd[x][i] == "X"{
				return x, y
			}
		}
		y -= n
	case "E" :
		if y + n >= len(parkTd[0]) {
			return x, y
		}
		for i:=y+1; i <= y+n ; i++ {
			if parkTd[x][i] == "X"{
				return x, y
			}
		}
		y += n
	}
	return x, y
}