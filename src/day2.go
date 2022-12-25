package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// execute(score)
	execute(score_2)
}

func score_2(move string) int {
	/*
		A,X -> Rock, Lose
		B,Y -> Paper, Draw
		C,Z -> Scissor, Win
	*/
	switch move {
	case "A X":
		return 3 + 0
	case "A Y":
		return 1 + 3
	case "A Z":
		return 2 + 6
	case "B X":
		return 1 + 0
	case "B Y":
		return 2 + 3
	case "B Z":
		return 3 + 6
	case "C X":
		return 2 + 0
	case "C Y":
		return 3 + 3
	case "C Z":
		return 1 + 6
	default:
		return 0
	}
}

func execute(scoreFunc func(move string) int) {
	f, err := os.Open("../input/2.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += scoreFunc(line)
	}
	fmt.Println(sum)
}

func score(move string) int {
	/*
		A,X -> R,R
		B,Y -> P,P
		C,Z -> S,S
	*/
	switch move {
	case "A X":
		return 1 + 3
	case "A Y":
		return 2 + 6
	case "A Z":
		return 3 + 0
	case "B X":
		return 1 + 0
	case "B Y":
		return 2 + 3
	case "B Z":
		return 3 + 6
	case "C X":
		return 1 + 6
	case "C Y":
		return 2 + 0
	case "C Z":
		return 3 + 3
	default:
		return 0
	}
}
