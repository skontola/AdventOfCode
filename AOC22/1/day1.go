// day1 is a program that takes grouped values from an input file
// and returns the group with the greatest sum.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Reverse reverses a slice.
func reverse(s1 []int) []int {
	var s2 []int

	for i := len(s1) - 1; i >= 0; i-- {
		s2 = append(s2, s1[i])
	}
	return s2
}

func main() {
	var sum int
	var s []int

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	//Sum groups of ints in given file
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			val, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			sum = sum + val
		} else {
			s = append(s, sum)
			sum = 0
		}
	}

	sort.Ints(s)
	s = reverse(s)

	fmt.Println("The most calories carried by a single elf is: " + strconv.Itoa(s[0]) + ".")

	top3 := s[0] + s[1] + s[2]

	fmt.Println("The three grubbiest elves carry a total of " + strconv.Itoa(top3) + " calories.")

	readFile.Close()
}
