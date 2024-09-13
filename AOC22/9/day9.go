package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type move struct {
	Dir string
	Amt int
}

var moves []move

// halp is a helper function that moves tail towards head, and returns tails new pos.
func halp(hx int, hy int, tx int, ty int) (int, int) {
	var x, y int
	x = tx
	y = ty

	if hx > tx+1 && hy > ty || hy > ty+1 && hx > tx { //diag up right
		x++
		y++
	} else if hx < tx-1 && hy > ty || hy > ty+1 && hx < tx { //diag up left
		x--
		y++
	} else if hx > tx+1 && hy < ty || hy < ty-1 && hx > tx { //diag down right
		x++
		y--
	} else if hx < tx-1 && hy < ty || hy < ty-1 && hx < tx { //diag down left
		x--
		y--
	} else if hy > ty+1 { //up
		y++
	} else if hy < ty-1 { //down
		y--
	} else if hx > tx+1 { //right
		x++
	} else if hx < tx-1 { //left
		x--
	}

	return x, y
}

func main() {
	//Read and parse input
	readFile, err := os.Open("9input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")

		tmpAmt, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
		}

		moves = append(moves, move{s[0], tmpAmt})
	}

	readFile.Close()

	//Coordinates of head, tail, and everything inbetween.
	var hx, x1, x2, x3, x4, x5, x6, x7, x8, tx, hy, y1, y2, y3, y4, y5, y6, y7, y8, ty int

	//The grid which will hold the positions that tail has been in.
	//I just doubled the starting array size value of 100 until it worked.. :)
	var secondGrid [800][800]bool
	var tailGrid [800][800]bool
	//The cooler non array-matrix approach. Could maybe use a coord struct
	//instead of a string if I ever wanted to do something cool with this data..
	m1 := make(map[string]bool)
	mTail := make(map[string]bool)

	//For each instruction,
	for i := 0; i < len(moves); i++ {
		//"move" head to new position
		switch moves[i].Dir {
		case "U":
			hy = hy + moves[i].Amt
		case "D":
			hy = hy - moves[i].Amt
		case "R":
			hx = hx + moves[i].Amt
		case "L":
			hx = hx - moves[i].Amt
		}

		//move tail one step at a time towards head
		for j := 0; j < moves[i].Amt; j++ {

			x1, y1 = halp(hx, hy, x1, y1)
			x2, y2 = halp(x1, y1, x2, y2)
			x3, y3 = halp(x2, y2, x3, y3)
			x4, y4 = halp(x3, y3, x4, y4)
			x5, y5 = halp(x4, y4, x5, y5)
			x6, y6 = halp(x5, y5, x6, y6)
			x7, y7 = halp(x6, y6, x7, y7)
			x8, y8 = halp(x7, y7, x8, y8)
			tx, ty = halp(x8, y8, tx, ty)

			//Could maybe use a map (to avoid duplicates) instead of this big array matrix.
			//"400" to start at middle of grid, since tail will sometimes be at negative coords.
			//Update: The map approach worked, but I'm keeping the old one as a monument to all of my sins.
			secondGrid[x1+400][y1+400] = true
			tailGrid[tx+400][ty+400] = true
			//"," to avoid false duplicates, e.g. 1,11 and 11,1
			m1[strconv.Itoa(x1)+","+strconv.Itoa(y1)] = true
			mTail[strconv.Itoa(tx)+","+strconv.Itoa(ty)] = true
		}
	}

	//Count how many different coords tail has been, print answer
	counter1 := 0
	counterTail := 0
	for i := 0; i < 800; i++ {
		for j := 0; j < 800; j++ {
			if secondGrid[i][j] == true {
				counter1++
			}
			if tailGrid[i][j] == true {
				counterTail++
			}
		}
	}

	fmt.Println("Answer to part 1:", counter1)
	fmt.Println("Answer to part 2:", counterTail)
	fmt.Println("Answer to part 1 and 2 again, using a different approach:", len(m1), len(mTail))

}
