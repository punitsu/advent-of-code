package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return data, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SafeReportCount(data string) (count int, err error) {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		nums := strings.Fields(line)

		var direction string // Tracks whether the sequence is "increasing" or "decreasing"
		valid := true

		for i := 0; i < len(nums)-1; i++ {
			num1, err := strconv.Atoi(nums[i])
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", nums[i])
			}
			num2, err := strconv.Atoi(nums[i+1])
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", nums[i+1])
			}

			diff := abs(num1 - num2)
			if diff < 1 || diff > 3 {
				valid = false
				break
			}

			if num1 < num2 {
				if direction == "" {
					direction = "increasing"
				} else if direction != "increasing" {
					valid = false
					break
				}
			} else if num1 > num2 {
				if direction == "" {
					direction = "decreasing"
				} else if direction != "decreasing" {
					valid = false
					break
				}
			}
		}
		if valid {
			count++
		}
	}
	return count, nil
}

func main() {
	data, err := ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	count, err := SafeReportCount(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("count", count)
}
