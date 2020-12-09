package main

import (
	"fmt"
)

type Hourglass struct {
	top    []int32
	middle int32
	bottom []int32
}

func NewHourglass(topRow []int32, middle int32, bottomRow []int32) *Hourglass {
	return &Hourglass{
		top:    topRow,
		middle: middle,
		bottom: bottomRow,
	}
}

func (h *Hourglass) Print() {
	fmt.Println(h.top)
	fmt.Println("  ", h.middle)
	fmt.Println(h.bottom)
}

func (h *Hourglass) getScore() int32 {
	x := h.middle
	for _, v := range h.top {
		x += v
	}
	for _, v := range h.bottom {
		x += v
	}
	return x
}

func main() {

	grid := createGrid()
	printGrid(grid)
	hourglasses := generateHourglasses(grid)

	fmt.Println("Count:", len(hourglasses))

	highestScore := hourglasses[0]
	for _, hg := range hourglasses {
		if hg.getScore() > highestScore.getScore() {
			highestScore = hg
		}
	}

	fmt.Println("Highest Hourglass Score:", highestScore.getScore())
	highestScore.Print()
}

func generateHourglasses(grid [6][6]int32) []*Hourglass {
	var hourglasses []*Hourglass
	//only look at centre indexes 1-4
	innerBound, outerBound := 1, 4
	for i, v := range grid {
		if i >= innerBound && i <= outerBound {
			for j, x := range v {
				if j >= innerBound && j <= outerBound {
					topRow := grid[i-1][j-1 : j+2]
					middle := x
					bottomRow := grid[i+1][j-1 : j+2]
					hourglasses = append(hourglasses, NewHourglass(topRow, middle, bottomRow))
				}
			}
		}
	}

	return hourglasses
}

func printGrid(grid [6][6]int32) {
	for _, v := range grid {
		fmt.Println(v)
	}
}

func createGrid() [6][6]int32 {
/*
	-1 -1 0 -9 -2 -2
   -2 -1 -6 -8 -2 -5
   -1 -1 -1 -2 -3 -4
   -1 -9 -2 -4 -4 -5
   -7 -3 -3 -2 -9 -9
   -1 -3 -1 -2 -4 -5
 */
	return [6][6]int32{
		{-1,-1,0,-9,-2,-2},
		{-2,-1,-6,-8,-2,-5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
}
