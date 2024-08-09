package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//wallpaper := []string{"..........", ".....#....", "......##..", "...##.....", "....#....."}
	wallpaper := []string{".#...", "..#..", "...#."}
	fmt.Println(solution(wallpaper))
}

func solution(wallpaper []string) []int {
	lux, luy := math.MaxInt64, math.MaxInt64
	rdx, rdy := 0, 0
	// (x, y ) 일때 x 가 세로 y 가 가로
	for xIdx, row := range wallpaper {
		for yIdx, file := range strings.Split(row, "") {
			if file == "#" {
				if lux > xIdx { lux = xIdx}
				if luy > yIdx { luy = yIdx}
				if rdx < xIdx { rdx = xIdx}
				if rdy < yIdx { rdy = yIdx}
			}
		}
	}
    return []int{lux, luy, rdx+1, rdy+1}
}