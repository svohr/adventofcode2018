package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FabricClaim struct {
	id, x, y, w, h int
}

func anyOverlaps(fabricMap [][]int, claim FabricClaim) bool {
	for j := claim.y; j < claim.y+claim.h; j++ {
		for i := claim.x; i < claim.x+claim.w; i++ {
			if fabricMap[j][i] > 1 {
				return true
			}
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fabricMap := make([][]int, 1000)
	for i := range fabricMap {
		fabricMap[i] = make([]int, 1000)
	}
	fabricClaims := []FabricClaim{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fabricClaimStr := scanner.Text()

		var id, x, y, w, h int

		_, err = fmt.Sscanf(fabricClaimStr, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)

		newClaim := FabricClaim{id, x, y, w, h}
		fabricClaims = append(fabricClaims, newClaim)
		for j := y; j < y+h; j++ {
			for i := x; i < x+w; i++ {
				fabricMap[j][i] += 1
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalInTwoClaims := 0
	for j := range fabricMap {
		for i := range fabricMap[j] {
			if fabricMap[j][i] > 1 {
				totalInTwoClaims++
			}
		}
	}
	fmt.Println("Area claimed by at least 2:", totalInTwoClaims)

	for _, claim := range fabricClaims {
		if !anyOverlaps(fabricMap, claim) {
			fmt.Println("No overlaps: ", claim)
		}
	}
}
