package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

type Coord struct {
	x int
	y int
}

type PointDist struct {
	label int
	dist int
}

func distance(a, b Coord) int {
	return abs(a.x - b.x) + abs(a.y - b.y)
}

func closestWaypoint(c Coord, waypoints []Coord) int {
	dist := make([]PointDist, len(waypoints))
	for i, w := range waypoints {
		dist[i] = PointDist {i, distance(c, w)}
	}
	sort.Slice(dist, func (i, j int) bool {return dist[i].dist < dist[j].dist})
	if dist[0].dist < dist[1].dist {
		return dist[0].label
	}
	return -1
}

func mapTerritories(points []Coord) [][]int {
	territories := make([][]int, 400)
	for i := range territories {
		territories[i] = make([]int, 400)
		for j := range territories[i] {
			territories[i][j] = closestWaypoint(Coord {j, i}, points)
		}
	}
	return territories
}

func findInfinitesTerritories(territories [][]int) map[int] bool {
	infs := make(map[int] bool)

	for i := range territories {
		infs[territories[i][0]] = true
		infs[territories[i][len(territories[i]) - 1]] = true
	}
	for i := range territories[0] {
		infs[territories[0][i]] = true
		infs[territories[len(territories) - 1][i]] = true
	}
	return infs
}

func areasByTerritory(territories [][]int) map[int]int {
	areas := make(map[int]int)
	for i := range(territories) {
		for j := range(territories) {
			if _, ok := areas[territories[i][j]]; !ok {
				areas[territories[i][j]] = 0
			}
			areas[territories[i][j]]++
		}
	}
	return areas
}

func sumDistances(c Coord, waypoints []Coord) int {
	sum := 0
	for _, w := range waypoints {
		sum += distance(c, w)
	}
	return sum
}

func mapDistances(waypoints []Coord, min int) int {
	areaInProximity := 0
	territories := make([][]int, 400)
	for i := range territories {
		territories[i] = make([]int, 400)
		for j := range territories[i] {
			if sumDistances(Coord{j,i}, waypoints) < min {
				areaInProximity += 1
			}
		}
	}
	return areaInProximity
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := make([]Coord, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y int
		line := scanner.Text()
		_, err = fmt.Sscanf(line, "%d, %d", &x, &y)
		point := Coord {x, y}
		points = append(points, point)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	territories := mapTerritories(points)
	areas := areasByTerritory(territories)
	infs := findInfinitesTerritories(territories)

	maxArea := -1
	for t, a := range areas {
		if _, ok := infs[t]; !ok {
			fmt.Println(t, a, infs[t])
			if maxArea < a {
				maxArea = a
			}
		}
	}
	fmt.Printf("Area: %d\n", maxArea)
	fmt.Println(infs)
	fmt.Println(points)
	fmt.Println("Area in proximity (10000):", mapDistances(points, 10000))
}
