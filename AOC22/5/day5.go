package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Set PART2 false gives answer to part 1, true gives answer to part 2.
const PART2 = true

//I have decided to implement my own stack data structure for todays challenge

type Stack []string

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Push(str string) {
	*st = append(*st, str)
}

// Pop removes item from top of stack.
func (st *Stack) Pop() (string, bool) {
	if st.IsEmpty() {
		return "", false
	} else {
		index := len(*st) - 1
		elem := (*st)[index]
		*st = (*st)[:index]
		return elem, true
	}
}

// Popbottom removes item from bottom of stack. (useful for part 2)
func (st *Stack) Popbottom() (string, bool) {
	if st.IsEmpty() {
		return "", false
	} else {
		elem := (*st)[0]
		*st = (*st)[1:]
		return elem, true
	}
}

var first Stack
var second Stack
var third Stack
var fourth Stack
var fifth Stack
var sixth Stack
var seventh Stack
var eighth Stack
var ninth Stack

var stacks = []Stack{first, second, third, fourth, fifth, sixth, seventh, eighth, ninth}

var bufStack Stack

func main() {

	readFile, err := os.Open("5input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == " 1   2   3   4   5   6   7   8   9 " {
			break
		} else {
			s := strings.Split(fileScanner.Text(), "")
			counter := 4
			for i := 1; i < len(s); i++ {
				if counter == 4 {
					if s[i] != " " {
						stacks[i/4].Push(s[i])
					}
					counter = 0
				}
				counter++
			}
		}
	}

	//reverse stacks
	for i := 0; i < len(stacks); i++ {
		for {
			var elem, cont = stacks[i].Pop()
			if !cont {
				break
			}
			bufStack.Push(elem)
		}
		for {
			var elem, cont = bufStack.Popbottom()
			if !cont {
				break
			}
			stacks[i].Push(elem)
		}
	}

	//Remove empty row
	fileScanner.Scan()
	//Move crates following instructions
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		//Hacky way to parse the numbers from the text
		s2 := []string{s[1], s[3], s[5]}
		amount, err := strconv.Atoi(s2[0])
		from, err := strconv.Atoi(s2[1])
		to, err := strconv.Atoi(s2[2])
		if err != nil {
			fmt.Println(err)
		}
		//Part 1
		if !PART2 {
			for i := 0; i < amount; i++ {
				elem, er := stacks[from-1].Pop()
				if !er {
					fmt.Println("Popping empty stack is bad")
				}
				stacks[to-1].Push(elem)
			}
		} else { //Part 2
			for i := 0; i < amount; i++ {
				elem, er := stacks[from-1].Pop()
				if !er {
					fmt.Println("Popping empty stack is bad")
				}
				bufStack.Push(elem)
			}
			for i := 0; i < amount; i++ {
				elem, er := bufStack.Pop()
				if !er {
					fmt.Println("Popping empty stack is bad")
				}
				stacks[to-1].Push(elem)
			}
		}
	}

	readFile.Close()

	fmt.Print("The crates on top of the piles are: ")
	for i := 0; i < len(stacks); i++ {
		elem, er := stacks[i].Pop()
		if !er {
			fmt.Println("Popping empty stack is bad 2: electric bogaloo")
		}
		fmt.Print(elem)
	}
}
