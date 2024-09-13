package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dirVis struct {
	visTop bool //Visible from the top
	visLef bool
	visBot bool
	visRig bool
}

func main() {

	var grid [][]int
	vis := [99][99]dirVis{}
	var answer int
	score := [99][99]int{}

	// Read input, handle errors.
	readFile, err := os.Open("8input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		var ints []int
		var tree int

		s := strings.Split(fileScanner.Text(), "")

		for i := 0; i < len(s); i++ {
			tree, _ = strconv.Atoi(s[i])
			ints = append(ints, tree)
		}
		grid = append(grid, ints)
	}

	readFile.Close()

	// Make trees visible
	for i := 0; i < len(grid); i++ {
		max1 := -1
		max2 := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > max1 {
				max1 = grid[i][j]
				vis[i][j].visTop = true
			}
		}
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] > max2 {
				max2 = grid[i][j]
				vis[i][j].visBot = true
			}
		}
	}
	for j := 0; j < len(grid[0]); j++ {
		maxLef := -1
		maxRig := -1
		for i := 0; i < len(grid); i++ {
			if grid[i][j] > maxLef {
				maxLef = grid[i][j]
				vis[i][j].visLef = true
			}
		}
		for i := len(grid) - 1; i >= 0; i-- {
			if grid[i][j] > maxRig {
				maxRig = grid[i][j]
				vis[i][j].visRig = true
			}
		}
	}

	// Count the amount of visible trees.
	invis := 0
	top := 0
	bot := 0
	rig := 0
	lef := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if vis[i][j].visTop || vis[i][j].visLef || vis[i][j].visBot || vis[i][j].visRig {
				answer++
				if vis[i][j].visTop {
					top++
				}
				if vis[i][j].visLef {
					lef++
				}
				if vis[i][j].visBot {
					bot++
				}
				if vis[i][j].visRig {
					rig++
				}
			} else {
				invis++
			}
		}
	}

	graphic := [99][99]string{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if vis[i][j] == vis[50][50] {
				graphic[i][j] = "."
			} else {
				graphic[i][j] = strconv.Itoa(grid[i][j])
			}
		}
	}

	var up, down, left, right int

	//Calculate score
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			up = 0
			down = 0
			left = 0
			right = 0
			var s1, s2, s3, s4 []int
			//up
			for k := 1; i-k >= 0; k++ {
				s1 = append(s1, grid[i-k][j])
			}
			//fmt.Println(s)
			for k := 0; k < len(s1); k++ {
				if grid[i][j] > s1[k] { //&& s1[k] >= maxUp {
					up++
				} else {
					up++
					break
				}
			}
			//down
			for k := 1; i+k < len(grid); k++ {
				s2 = append(s2, grid[i+k][j])
			}
			//fmt.Println(s2)
			for k := 0; k < len(s2); k++ {
				if grid[i][j] > s2[k] { //&& s2[k] >= maxDown {
					down++
				} else {
					down++
					break
				}
			}
			//left
			for k := 1; j-k >= 0; k++ {
				s3 = append(s3, grid[i][j-k])
			}
			//fmt.Println(s3)
			for k := 0; k < len(s3); k++ {
				if grid[i][j] > s3[k] { //&& s3[k] >= maxLeft {
					left++
				} else {
					left++
					break
				}
			}
			//right
			for k := 1; j+k < len(grid); k++ {
				s4 = append(s4, grid[i][j+k])
			}
			//fmt.Println(s4)
			for k := 0; k < len(s4); k++ {
				if grid[i][j] > s4[k] { //&& s4[k] >= maxRight {
					right++
				} else {
					right++
					break
				}
			}
			score[i][j] = up * down * left * right
		}
	}

	for i := 0; i < len(grid); i++ {
		fmt.Println(graphic[i])
	}

	fmt.Println("top", top, "bot", bot, "lef", lef, "rig", rig, "invis", invis)
	fmt.Println("There are", answer, "trees visible from outside the grid")

	maxI := 0
	maxJ := 0
	maxScore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if score[i][j] > maxScore {
				maxScore = score[i][j]
				maxI = i
				maxJ = j
			}
		}
	}

	fmt.Println("Maximum score is:", maxScore, "at coordinates:", maxI, maxJ)
}
