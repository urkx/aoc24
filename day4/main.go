package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mul(s string, regexNum *regexp.Regexp) int {
	nums := regexNum.FindAllString(s, -1)
	n1, e := strconv.Atoi(nums[0])
	if e != nil {
		panic(e)
	}
	n2, e := strconv.Atoi(nums[1])
	if e != nil {
		panic(e)
	}
	return n1 * n2
}

func processXmas(x int, y int, w int, h int, mem [][]string, dir []int) int {
	res := 1
	ex := "XMAS"
	buff := ""
	for i, v := range ex {
		xx := x + i * dir[0]
		yy := y + i * dir[1]
		if xx >= 0 && xx < h && yy >= 0 && yy < w {
			if mem[xx][yy] != string(v) {
				res = 0
			} else {
				buff += mem[xx][yy]
			}
		}  
	}
	if ex != buff {
		res = 0
	}
	return res
}

func processMas(x int, y int, w int, h int, mem [][]string, dir []int) int {
	res := 1
	ex := "MAS"
	buff := ""
	
	fmt.Printf("Direction: %v\n", dir)
	for i, v := range ex {
		xx := x + i * dir[0]
		yy := y + i * dir[1]
		if xx >= 0 && xx < h && yy >= 0 && yy < w {
			fmt.Printf("%s\n", mem[xx][yy])
			if mem[xx][yy] != string(v) {
				res = 0
			} else {
				buff += mem[xx][yy]
			}
		}  
	}
	if ex != buff {
		res = 0
	}
	fmt.Printf("Result: %d\n", res)
	return res
}


func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := 0
	var mem [][]string
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, 
					{0, -1}, {0, 0}, {0, 1}, 
					{1, -1}, {1, 0}, {1, 1}}

	for scanner.Scan() {
		mem = append(mem, strings.Split(scanner.Text(), ""))
	}

	x := len(mem)
	y := len(mem[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			for d := 0; d < 9; d++ {
				res += processXmas(i, j, x, y, mem, dirs[d])
			}
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
	res := 0
	var mem [][]string
	dirs := [][]int{{1, 1}, {1, -1}, 
					       
					{-1, 1}, {-1, -1}}
	ini := [][]int{{-1, -1}, {-1, 1}, 
					       
					{1, -1}, {1, 1}}

	for scanner.Scan() {
		mem = append(mem, strings.Split(scanner.Text(), ""))
	}

	x := len(mem)
	y := len(mem[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if mem[i][j] == "A" {
				c := 0
				for d := 0; d < 4; d++ {
					c += processMas(i+ini[d][0], j+ini[d][1], x, y, mem, dirs[d])
				}
				if c == 2 {
					res += 1
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println("AOC 2024 - Day 4")
	fmt.Println("")

	fmt.Printf("Part 1 solution: %d\n", part1())
	fmt.Printf("Part 2 solution: %d\n", part2())
}
