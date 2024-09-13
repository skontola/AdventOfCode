package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycle int
var reg int
var signal int
var grid [6][40]string

func peak() {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		signal = signal + cycle*reg
	}
}

func draw() {
	pixel := (cycle - 1) % 40
	if pixel == reg-1 || pixel == reg || pixel == reg+1 {
		grid[(cycle-1)/40][(cycle-1)%40] = "#"
	} else {
		grid[(cycle-1)/40][(cycle-1)%40] = "."
	}
}

func main() {

	var amt int
	reg = 1

	//Read and parse input
	readFile, err := os.Open("10input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")

		//noop is same as beginning of addx, always do this
		cycle++
		peak()
		draw()
		//addx
		if s[0] == "addx" {
			cycle++
			peak()
			draw()
			amt, _ = strconv.Atoi(s[1])
			reg = reg + amt
		}

	}

	readFile.Close()

	fmt.Println("[ - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ]")
	fmt.Println("[ - Signal strength is", signal, " - - - - - - - - - - - - - - - - - - - - - - - - - ]")
	fmt.Println("[ - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ]")
	for i := 0; i < 6; i++ {
		fmt.Println(grid[i])
	}
}
