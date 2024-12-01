package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var n1 []int
	var n2 []int
	totalItems := 0
	res := 0

	for scanner.Scan() {
		var line = scanner.Text()
		num1, e1 := strconv.Atoi(line[0:5])
		if e1 != nil {
			panic(e1)
		}
		num2, e2 := strconv.Atoi(line[len(line)-5:])
		if e2 != nil {
			panic(e2)
		}
		n1 = append(n1, num1)
		n2 = append(n2, num2)
		totalItems += 1
	}
	sort.Sort(sort.IntSlice(n1))
	sort.Sort(sort.IntSlice(n2))

	for i := 0; i < totalItems; i++ {
		if n1[i] > n2[i] {
			res += n1[i] - n2[i]
		} else {
			res += n2[i] - n1[i]
		}
	}

	return res
}

func part2() int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var n1 []int
	m := make(map[int]int)
	// var n2 []int
	totalItems := 0
	res := 0

	for scanner.Scan() {
		var line = scanner.Text()
		num1, e1 := strconv.Atoi(line[0:5])
		if e1 != nil {
			panic(e1)
		}
		num2, e2 := strconv.Atoi(line[len(line)-5:])
		if e2 != nil {
			panic(e2)
		}
		n1 = append(n1, num1)
		m[num2] = m[num2] + 1
		// n2 = append(n2, num2)
		totalItems += 1
	}

	for i := 0; i < totalItems; i++ {
		res += n1[i] * m[n1[i]]
	}

	return res
}

func main() {
	fmt.Println("AOC 2024 - Day 1")
	fmt.Println("")

	fmt.Printf("Part 1 solution: %d\n", part1())
	fmt.Printf("Part 2 solution: %d\n", part2())
}
