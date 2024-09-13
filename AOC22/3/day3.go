package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func priority(s1 []string, s2 []string, m map[string]int) int {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				return m[s1[i]]
			}
		}
	}
	fmt.Println("priority err")
	return 0
}

func priority3(tmp1 []string, tmp2 []string, tmp3 []string, m map[string]int) int {
	for i := 0; i < len(tmp1); i++ {
		for j := 0; j < len(tmp2); j++ {
			if tmp1[i] == tmp2[j] {
				for k := 0; k < len(tmp3); k++ {
					if tmp1[i] == tmp3[k] {
						return m[tmp1[i]]
					}
				}
			}
		}
	}
	fmt.Println("priority3 err")
	return 0
}

func main() {
	var tmp1 []string
	var tmp2 []string
	var tmp3 []string
	var counter int
	var sum int
	var sum2 int

	//Create slice with "items"
	prio := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	prioS := strings.Split(prio, "")

	//Map "items" to values.
	m := make(map[string]int)
	for i := 0; i < len(prioS); i++ {
		m[prioS[i]] = i + 1
	}

	//Read file.
	readFile, err := os.Open("3input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	//Calculate total priority value for every row and triplet row.
	for fileScanner.Scan() {
		counter++
		s := strings.Split(fileScanner.Text(), "")
		s2 := s[(len(s) / 2):]
		s1 := s[0:(len(s) / 2)]
		sum = sum + priority(s1, s2, m)
		switch counter {
		case 1:
			tmp1 = s
		case 2:
			tmp2 = s
		case 3:
			tmp3 = s
			sum2 = sum2 + priority3(tmp1, tmp2, tmp3, m)
			counter = 0
		}
	}

	//Close file when done with it.
	readFile.Close()

	//Print results!
	fmt.Println("The sum of priorities is: " + strconv.Itoa(sum) + ".")
	fmt.Println("The sum of triplet priorities is: " + strconv.Itoa(sum2) + ".")

}
