package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// part1()
	part2()
}

func part2() {
	f, err := os.Open("../input/3.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var sum int = 0
	for i := 0; i < len(lines); i += 3 {
		sum += process_2(lines[i : i+3])
	}
	fmt.Println(sum)
}

func process_2(lines []string) int {
	s1 := stringToSet(lines[0])
	s2 := stringToSet(lines[1])
	str := lines[2]
	for i := 0; i < len(str); i++ {
		_, ok1 := s1[str[i]]
		_, ok2 := s2[str[i]]
		if ok1 && ok2 {
			return mapToNumber(str[i])
		}
	}
	return 0
}

func part1() {
	f, err := os.Open("../input/3.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int = 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += process(line)
	}
	fmt.Println(sum)
}

type void struct{}

var member void

func process(line string) int {
	length := len(line)
	set := make(map[byte]void)
	i := 0
	for ; i < length/2; i++ {
		set[line[i]] = member
	}
	for ; i < length; i++ {
		if _, ok := set[line[i]]; ok {
			break
		}
	}

	return mapToNumber(line[i])
}

func mapToNumber(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b-'a') + 1
	} else {
		return int(b-'A') + 1 + 26
	}
}

func stringToSet(str string) map[byte]void {
	set := make(map[byte]void)
	for i := 0; i < len(str); i++ {
		set[str[i]] = member
	}
	return set
}
