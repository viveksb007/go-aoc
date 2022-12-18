package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// part1()
	part2()
}

func part2() {
	f, err := os.Open("../input/1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int = 0
	clubbedInput := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			clubbedInput = append(clubbedInput, sum)
			sum = 0
		} else {
			val, _ := strconv.Atoi(line)
			sum += val
		}
	}

	sort.Slice(clubbedInput, func(i, j int) bool {
		return clubbedInput[i] > clubbedInput[j]
	})

	sol := 0
	for i := 0; i < 3; i++ {
		sol += clubbedInput[i]
	}
	fmt.Println(sol)
}

func part1() {
	f, err := os.Open("../input/1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sol int64 = 0
	var count int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if count > sol {
				sol = count
			}
			count = 0
		} else {
			val, _ := strconv.Atoi(line)
			count += int64(val)
		}
	}

	fmt.Println(sol)
}
