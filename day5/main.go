package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StrToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

type Rule struct{
	first int
	last int
}

func (r *Rule) check(update []int) bool {
	firstIdx := 0
	firstFound := false
	lastIdx := 0
	lastFound := false

	for i, v := range update {
		if r.first == v {
			firstIdx = i
			firstFound = true
		}
		if r.last == v {
			lastIdx = i
			lastFound = true
		}
	}

	if firstFound && lastFound {
		return firstIdx < lastIdx
	} else {
		return true
	}
	
}

type RuleSet struct {
	set []Rule
}

func (s *RuleSet) AddRule(r Rule) {
	s.set = append(s.set, r)
}

func checkUpdate(s string, m map[int]*RuleSet) bool {
	pages := strings.Split(s, ",")
	res := true
	pageNum := make([]int, 0)
	for _, v := range pages {
		pageNum = append(pageNum, StrToInt(v))
	}
	
	for _, u := range pageNum[:len(pageNum)-1] {
		if m[u] != nil {
			rs := m[u].set
			for _, r := range rs {
				res = res && r.check(pageNum)
			}
		} else {
			res = false
		}
	}
	
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
	m := make(map[int]*RuleSet)

	for scanner.Scan() {
		l := scanner.Text()
		// Is Rule
		if strings.ContainsRune(l, '|') {
			s := strings.Split(l, "|")
			i := StrToInt(s[0])
			j := StrToInt(s[1])

			r := Rule{first: i, last: j }
			if m[i] == nil {
				b := []Rule{r}
				rs := RuleSet{set: b}
				m[i] = &rs
			} else {
				rs := m[i]
				rs.set = append(rs.set, r)
			}
		} else if l != "" { // Is update
			if checkUpdate(l, m) {
				lAux := strings.Split(l, ",")
				res += StrToInt(lAux[len(lAux)/2])
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

	for scanner.Scan() {
		mem = append(mem, strings.Split(scanner.Text(), ""))
	}
	return res
}

func main() {
	fmt.Println("AOC 2024 - Day 5")
	fmt.Println("")

	fmt.Printf("Part 1 solution: %d\n", part1())
	//fmt.Printf("Part 2 solution: %d\n", part2())
}
