package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)


func areComplementarySubunits(a, b rune) bool {
	return (unicode.ToUpper(a) == unicode.ToUpper(b)) && (a != b)
}


func reactPolymer(poly string) string {
	parts := []string {}
	runes := []rune(poly)
	start := 0
	for i := 0; i < len(runes) - 1; i++ {
		if areComplementarySubunits(runes[i], runes[i+1]) {
			parts = append(parts, string(runes[start:i]))
			start = i + 2
			i++
		}
	}
	parts = append(parts, string(runes[start:]))
	return strings.Join(parts, "")
}

func fullyReact(polymer string) string {
	for {
		product := reactPolymer(polymer)
		if polymer == product {
			return product
		}
		polymer = product
	}
}


func dropSubunit(sub rune) func(rune) rune {
	return func (r rune) rune {
		if sub == r || unicode.ToLower(sub) == r {
			return rune(-1)
		}
		return r
	}
}


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	polymer, err := reader.ReadString('\n')
	polymer = strings.TrimSpace(polymer)

	product := fullyReact(polymer)

	fmt.Println("product:", len(product))

	subunits := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	best := len(polymer)
	for _, subunit := range subunits {
		newPolymer := strings.Map(dropSubunit(subunit), polymer)
		product = fullyReact(newPolymer)
		if len(product) < best {
			best = len(product)
		}
	}
	fmt.Println("best:", best)
}
