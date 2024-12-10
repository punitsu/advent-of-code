package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FileName = "input.txt"

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return data, nil
}

func ParseData(data string) (column1, column2 []int, err error) {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		nums := strings.Fields(line)
		if len(nums) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		i, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", nums[0])
		}
		j, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", nums[1])
		}

		column1 = append(column1, i)
		column2 = append(column2, j)
	}
	sort.Ints(column1)
	sort.Ints(column2)
	return column1, column2, nil
}

func FindDifference(column1, column2 []int) int {
	diff := 0
	for i := range column1 {
		diff += abs(column1[i] - column2[i])
	}
	return diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	data, err := ReadFile(FileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	column1, column2, err := ParseData(string(data))
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}

	diff := FindDifference(column1, column2)
	fmt.Println("Difference:", diff)
}
