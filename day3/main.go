package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"regexp"
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

func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := 0
	regex, _ := regexp.Compile("(mul[(][0-9]{1,3},[0-9]{1,3}[)])")
	regexNum, _ := regexp.Compile("([0-9]{1,3})")

	for scanner.Scan() {
		l := scanner.Text()
		matches := regex.FindAllString(l, -1)
		for _, s := range matches {
			res += mul(s, regexNum)
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
	regex, _ := regexp.Compile("(don't[(][)])|(do[(][)])|(mul[(][0-9]{1,3},[0-9]{1,3}[)])")
	regexNum, _ := regexp.Compile("([0-9]{1,3})")
	regexDont, _ := regexp.Compile("(don't[(][)])")
	regexDo, _ := regexp.Compile("(do[(][)])")
	do := true

	for scanner.Scan() {
		l := scanner.Text()
		matches := regex.FindAllString(l, -1)
		for _, s := range matches {
			if regexDont.MatchString(s) {
				do = false
			} else if regexDo.MatchString(s) && do == false {
				do = true
			}

			if do && regexNum.MatchString(s){
				res += mul(s, regexNum)
			}
		}
	}
	return res
}

func main() {
	fmt.Println("AOC 2024 - Day 3")
	fmt.Println("")

	fmt.Printf("Part 1 solution: %d\n", part1())
	fmt.Printf("Part 2 solution: %d\n", part2())
}
