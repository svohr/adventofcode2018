package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type LogEntry struct {
	time time.Time
	note []string
}

func SleepiestGuard(guardMinsSlept map[int]int) int {
	mostMinutes := -1
	guard := -1
	for g, minutesSlept := range guardMinsSlept {
		if mostMinutes < minutesSlept {
			guard = g
			mostMinutes = minutesSlept
		}
	}
	return guard
}

func SleepiestMinute(minSleptThru []int) int {
	minute := -1
	timesSlept := -1
	for m, s := range minSleptThru {
		if timesSlept < s {
			minute = m
			timesSlept = s
		}
	}
	return minute
}

func ConsistentSleeper(guardMinsSlept map[int][]int) (int, int) {
	guard := -1
	mostSlept := -1
	timesSlept := -1
	for g, sleepByMin := range guardMinsSlept {
		m := SleepiestMinute(sleepByMin)
		if sleepByMin[m] > timesSlept {
			guard = g
			mostSlept = m
			timesSlept = sleepByMin[m]
		}
	}
	return guard, mostSlept
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	guardLog := make([]LogEntry, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		items := strings.Split(line, " ")
		fmt.Println(items)

		t, _ := time.Parse("[2006-01-02 15:04]", strings.Join(items[:2], " "))
		guardLog = append(guardLog, LogEntry{t, items[2:]})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(guardLog, func(i, j int) bool {
		return guardLog[i].time.Before(guardLog[j].time)
	})

	guardAsleepByMin := make(map[int][]int)
	guardAsleepTotal := make(map[int]int)
	var curGuard, sleepMinute, wakeMinute int
	for _, entry := range guardLog {
		fmt.Println(entry.time, entry.note)
		switch entry.note[0] {
		case "Guard":
			fmt.Sscanf(entry.note[1], "#%d", &curGuard)
		case "falls":
			sleepMinute = entry.time.Minute()
		case "wakes":
			wakeMinute = entry.time.Minute()
			if _, ok := guardAsleepTotal[curGuard]; !ok {
				guardAsleepTotal[curGuard] = 0
				guardAsleepByMin[curGuard] = make([]int, 60)
			}
			guardAsleepTotal[curGuard] += wakeMinute - sleepMinute
			for i := sleepMinute; i < wakeMinute; i++ {
				guardAsleepByMin[curGuard][i]++
			}
		}
	}

	sleepyGuard := SleepiestGuard(guardAsleepTotal)
	sleepyMinute := SleepiestMinute(guardAsleepByMin[sleepyGuard])
	fmt.Println("Sleepiest Guard:", sleepyGuard)
	fmt.Println("Sleepiest Minute:", sleepyMinute)
	fmt.Println(sleepyGuard * sleepyMinute)

	consistentGuard, consistentMinute := ConsistentSleeper(guardAsleepByMin)
	fmt.Println("Consistent Guard:", consistentGuard)
	fmt.Println("Consistent Minute:", consistentMinute)
	fmt.Println(consistentGuard * consistentMinute)

}
