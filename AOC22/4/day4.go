package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Converts string to integer, makes main more readable.
func str2i(str string) int {
	i, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error during conversion")
		return 0
	}

	return i
}

func main() {
	var sum int
	var sum2 int

	readFile, err := os.Open("4input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), ",")
		s1 := strings.Split(s[0], "-")
		s2 := strings.Split(s[1], "-")
		if str2i(s1[0]) >= str2i(s2[0]) && str2i(s1[1]) <= str2i(s2[1]) ||
			str2i(s2[0]) >= str2i(s1[0]) && str2i(s2[1]) <= str2i(s1[1]) {
			sum++
		}

		if str2i(s1[0]) <= str2i(s2[0]) && str2i(s1[1]) >= str2i(s2[0]) ||
			str2i(s1[0]) <= str2i(s2[1]) && str2i(s1[1]) >= str2i(s2[1]) ||
			str2i(s2[0]) <= str2i(s1[0]) && str2i(s2[1]) >= str2i(s1[0]) ||
			str2i(s2[0]) <= str2i(s1[1]) && str2i(s2[1]) >= str2i(s1[1]) {
			sum2++
		}
	}

	readFile.Close()

	fmt.Println(strconv.Itoa(sum) + " assignments fully contained by partner")
	fmt.Println(strconv.Itoa(sum2) + " assignment pairs overlap")
}
