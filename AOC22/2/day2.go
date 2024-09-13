package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var sum int
	var sum2 int

	readFile, err := os.Open("2input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		switch fileScanner.Text() {
		case "":
			//do nothing
		case "A X":
			sum = sum + 3 + 1
		case "A Y":
			sum = sum + 6 + 2
		case "A Z":
			sum = sum + 0 + 3
		case "B X":
			sum = sum + 0 + 1
		case "B Y":
			sum = sum + 3 + 2
		case "B Z":
			sum = sum + 6 + 3
		case "C X":
			sum = sum + 6 + 1
		case "C Y":
			sum = sum + 0 + 2
		case "C Z":
			sum = sum + 3 + 3
		default:
			fmt.Println("Nu gick det fel")
		}
	}

	readFile.Close()

	//Part 2

	readFile2, err := os.Open("2input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)

	/*
		Rock 		A X 1
		Paper 		B Y 2
		Scissors 	C Z 3
	*/

	for fileScanner2.Scan() {
		switch fileScanner2.Text() {
		case "":
			//do nothing
		case "A X":
			sum2 = sum2 + 0 + 3
		case "A Y":
			sum2 = sum2 + 3 + 1
		case "A Z":
			sum2 = sum2 + 6 + 2
		case "B X":
			sum2 = sum2 + 0 + 1
		case "B Y":
			sum2 = sum2 + 3 + 2
		case "B Z":
			sum2 = sum2 + 6 + 3
		case "C X":
			sum2 = sum2 + 0 + 2
		case "C Y":
			sum2 = sum2 + 3 + 3
		case "C Z":
			sum2 = sum2 + 6 + 1
		default:
			fmt.Println("Nu gick det fel")
		}
	}

	fmt.Print("The final score for the first game was: ")
	fmt.Println(sum)
	fmt.Print("The final score for the second game was: ")
	fmt.Println(sum2)

	readFile2.Close()
}
