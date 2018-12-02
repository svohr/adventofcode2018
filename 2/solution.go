package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countTwosAndThrees(str string) (int, int) {

	var counter map[rune]int
	counter = make(map[rune]int)
	for _, c := range str {
		counter[c]++
	}
	twos, threes := 0, 0
	for _, count := range counter {
		switch count {
		case 2:
			twos++
		case 3:
			threes++
		}
	}
	return twos, threes
}

func compareBoxIds(id1 string, id2 string) string {
	diffIdx := -1
	for i, _ := range id1 {
		if id1[i] != id2[i] {
			if diffIdx < 0 {
				diffIdx = i
			} else {
				return ""
			}
		}
	}
	if diffIdx < 0 {
		return ""
	}
	return id1[:diffIdx] + id1[diffIdx+1:]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	boxIds := []string{}

	scanner := bufio.NewScanner(file)
	sumTwos, sumThrees := 0, 0
	for scanner.Scan() {
		boxIdStr := scanner.Text()
		boxIds = append(boxIds, boxIdStr)

		twos, threes := countTwosAndThrees(boxIdStr)
		// fmt.Println(scanner.Text(), twos, threes)
		if twos > 0 {
			sumTwos++
		}
		if threes > 0 {
			sumThrees++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Checksum:", sumTwos*sumThrees)

	for i, id1 := range boxIds {
		for _, id2 := range boxIds[i+1:] {
			res := compareBoxIds(id1, id2)
			if res != "" {
				fmt.Println(id1, id2, res)
			}
		}
	}
}
