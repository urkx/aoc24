package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	res := 0

	for scanner.Scan() {
		var nums []int
		safe := true
		numsString := strings.Split(scanner.Text(), " ")
		for _, v := range numsString {
			num, e1 := strconv.Atoi(v)
			if e1 != nil {
				panic(e1)
			}
			nums = append(nums, num)
		}

		// Check if sequence is ascend or descend
		ascend := true
		if nums[0] > nums[1] {
			ascend = false
			if (nums[0]-nums[1]) < 1 || (nums[0]-nums[1]) > 3 {
				safe = false
			}
		} else {
			if (nums[1]-nums[0]) < 1 || (nums[1]-nums[0]) > 3 {
				safe = false
			}
		}

		for i := 1; i < len(nums)-1; i++ {
			if safe {
				if ascend {
					if nums[i] > nums[i+1] || ((nums[i+1]-nums[i]) < 1 || (nums[i+1]-nums[i]) > 3) {
						safe = false
					}
				} else {
					if nums[i] < nums[i+1] || ((nums[i]-nums[i+1]) < 1 || (nums[i]-nums[i+1]) > 3) {
						safe = false
					}
				}
			}
		}
		if safe {
			res += 1
		}
	}
	return res
}

func checkSafe(nu []int) int {
	l := len(nu)
	res := 0
	safeFound := false
	for i := 0; i < l; i++ {
		if !safeFound {
			n := make([]int, len(nu))
			copy(n, nu)
			safe := true
			nums := append(n[:i], n[i+1:]...)
			// Check if sequence is ascend or descend
			ascend := true
			if nums[0] > nums[1] {
				ascend = false
				if (nums[0]-nums[1]) < 1 || (nums[0]-nums[1]) > 3 {
					safe = false
				}
			} else {
				if (nums[1]-nums[0]) < 1 || (nums[1]-nums[0]) > 3 {
					safe = false
				}
			}

			for i := 1; i < len(nums)-1; i++ {
				if safe {
					if ascend {
						if nums[i] > nums[i+1] || ((nums[i+1]-nums[i]) < 1 || (nums[i+1]-nums[i]) > 3) {
							safe = false
						}
					} else {
						if nums[i] < nums[i+1] || ((nums[i]-nums[i+1]) < 1 || (nums[i]-nums[i+1]) > 3) {
							safe = false
						}
					}
				}
			}
			if safe {
				res = 1
				safeFound = true
			}
			// fmt.Printf("%d, %v, safe %t\n", i, nums, safe)
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

	for scanner.Scan() {
		var nums []int
		safe := true
		numsString := strings.Split(scanner.Text(), " ")
		for _, v := range numsString {
			num, e1 := strconv.Atoi(v)
			if e1 != nil {
				panic(e1)
			}
			nums = append(nums, num)
		}

		// Check if sequence is ascend or descend
		ascend := true
		if nums[0] > nums[1] {
			ascend = false
			if (nums[0]-nums[1]) < 1 || (nums[0]-nums[1]) > 3 {
				safe = false
			}
		} else {
			if (nums[1]-nums[0]) < 1 || (nums[1]-nums[0]) > 3 {
				safe = false
			}
		}

		for i := 1; i < len(nums)-1; i++ {
			if safe {
				if ascend {
					if nums[i] > nums[i+1] || ((nums[i+1]-nums[i]) < 1 || (nums[i+1]-nums[i]) > 3) {
						safe = false
					}
				} else {
					if nums[i] < nums[i+1] || ((nums[i]-nums[i+1]) < 1 || (nums[i]-nums[i+1]) > 3) {
						safe = false
					}
				}
			}
		}
		if safe {
			res += 1
		} else {
			res += checkSafe(nums)
		}
	}
	return res
}

func main() {
	fmt.Println("AOC 2024 - Day 2")
	fmt.Println("")

	fmt.Printf("Part 1 solution: %d\n", part1())
	fmt.Printf("Part 2 solution: %d\n", part2())
}
