package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//var items []int
	var inventories [][]int
	var operations [][]string
	var tests []int
	var trueActions []int
	var falseActions []int
	//var inspections []int
	inspections := []int{0, 0, 0, 0, 0, 0, 0, 0}

	//Read and parse input
	readFile, err := os.Open("11input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		if len(s) > 2 && s[2] == "Starting" {
			var items []int
			for _, sItem := range s[4:] {
				iItem, _ := strconv.Atoi(strings.Trim(sItem, ","))
				items = append(items, iItem)
			}
			inventories = append(inventories, items)
		}
		if len(s) > 2 && s[2] == "Operation:" {
			operations = append(operations, s[6:])
		}
		if len(s) > 2 && s[2] == "Test:" {
			test, _ := strconv.Atoi(s[5])
			tests = append(tests, test)
		}
		if len(s) > 5 && s[5] == "true:" {
			action, _ := strconv.Atoi(s[9])
			trueActions = append(trueActions, action)
		}
		if len(s) > 5 && s[5] == "false:" {
			action, _ := strconv.Atoi(s[9])
			falseActions = append(falseActions, action)
		}
	}

	fmt.Println(inventories)
	fmt.Println(operations)
	fmt.Println(tests)
	fmt.Println(trueActions)
	fmt.Println(falseActions)

	for i := 0; i < 20; i++ {
		for j := 0; j < len(inventories); j++ {
			for k := 0; k < len(inventories[j]); k++ {
				number, err := strconv.Atoi(operations[j][1])
				if err != nil {
					if operations[j][1] == "old" {
						number = inventories[j][k]
					} else {
						fmt.Println("Something went wrong")
					}
				}
				fmt.Println("Monkey", j, "inspects item with worry level", inventories[j][k])
				inspections[j] = inspections[j] + 1
				if operations[j][0] == "*" {
					inventories[j][k] = inventories[j][k] * number
				} else {
					inventories[j][k] = inventories[j][k] + number
				}
				fmt.Println("Worry level", operations[j], "=", inventories[j][k])

				//Relief
				inventories[j][k] = inventories[j][k] / 3
				fmt.Println("Worry level is decreased to", inventories[j][k])

				//Test worry level, yeet stuff to correct monkey
				if inventories[j][k]%tests[j] == 0 {
					fmt.Println("Worry level is divisible by", tests[j], ", yeet to monkey number:", trueActions[j])
					inventories[trueActions[j]] = append(inventories[trueActions[j]], inventories[j][k])
				} else {
					fmt.Println("Worry level is NOT divisible by", tests[j], ", yeet to monkey number:", falseActions[j])
					inventories[falseActions[j]] = append(inventories[falseActions[j]], inventories[j][k])
				}
				//Remove items when monkey has thrown everything
				if k == len(inventories[j])-1 {
					inventories[j] = nil
				}
			}
		}
	}
	sort.Ints(inspections)
	fmt.Println(inspections)
	fmt.Println("Total amount of monkey business:", inspections[len(inspections)-1]*inspections[len(inspections)-2])
}
