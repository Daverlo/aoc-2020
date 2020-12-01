package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func parseInput(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	res := make([]int, 0)

	for {
		var n int
		_, err := fmt.Fscanf(f, "%d\n", &n)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		res = append(res, n)
	}

	return res, nil
}

func processExpenses(numbers []int) int {
	sort.Ints(numbers)
	present, i, j := TargetSum(2020, numbers)
	if present {
		return numbers[i] * numbers[j]
	}
	return 0
}

// Search for two numbers contained in the array that add up to the target value
// Returns whether there is a solution and the indices of the values that form it
func TargetSum(t int, a []int) (bool, int, int) {
	for i, n := range a {
		c := t - n
		present, j := BinarySearch(c, a[i+1:])
		// j i an index of the subarray used in the binSearch
		j += i + 1

		if present {
			return true, i, j
		}
	}
	return false, -1, -1
}

func BinarySearch(n int, a []int) (bool, int) {
	l := 0
	h := len(a) - 1
	for l <= h {
		i := (l + h) / 2
		if a[i] == n {
			return true, i
		}

		if a[i] < n {
			l = i + 1
		} else {
			h = i - 1
		}
	}

	return false, -1
}

func main() {
	args := os.Args[1:]
	input, err := parseInput(args[0])
	if err != nil {
		panic(err)
	}

	output := processExpenses(input)
	fmt.Println(output)
}
