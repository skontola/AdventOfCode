package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
comparator checks if all elements in a slice are unique by
converting them into a map (which removes duplicates) and
comparing the length of the map with the length of the original slice.
*/
func comparator(s []string) bool {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	if len(m) == len(s) {
		return true
	}
	return false
}

func main() {
	var a1 = [4]string{}
	var a2 = [14]string{}
	var answ1 int
	var answ2 int

	// Read input, handle errors.
	readFile, err := os.Open("6input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	s := strings.Split(fileScanner.Text(), "")

	for i := 0; i < len(s); i++ {
		a1[i%4] = s[i]
		a2[i%14] = s[i]

		// Cant use arrays of different size as argument to
		// the same function, solution: slice it!
		if comparator(a1[:]) && answ1 == 0 {
			answ1 = i + 1
		}
		if comparator(a2[:]) {
			answ2 = i + 1
			break
		}
	}

	// Close filereader
	readFile.Close()

	// Print answers!
	fmt.Println("The first 4 unique characters appear after", answ1, "characters.")
	fmt.Println("The first 14 unique characters appear after", answ2, "characters.")

}
