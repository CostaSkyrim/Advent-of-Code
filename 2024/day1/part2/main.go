package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var col1, col2 []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting first column", err)
			continue
		}
		col1 = append(col1, num1)

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting second column:", err)
			continue
		}
		col2 = append(col2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(col1)
	sort.Ints(col2)

	result := 0

	if len(col1) != len(col2) {
		fmt.Println("The array of the numbers are not the same size")
	}

	similarity := make(map[int]int)
	frequencyMap := make(map[int]int)

	for _, num := range col2 {
		frequencyMap[num]++
	}

	for _, num := range col1 {
		similarity[num] = frequencyMap[num]
		result += similarity[num] * num
	}

	fmt.Printf("The similarity score is = %d", result)
}
