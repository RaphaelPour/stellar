package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadDefaultString() []string {
	return LoadString("input")
}

func LoadString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Error loading file '%s': %s\n", filename, err))
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, string(scanner.Text()))
	}

	return lines
}

func LoadDefaultInt() []int {
	return LoadInt("input")
}

func LoadInt(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Error loading file '%s': %s\n", filename, err))
	}

	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(string(scanner.Text()))
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a number\n", scanner.Text()))
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func LoadIntList(filename string) []int {
	rawContent, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("error reading file %s: %s", filename, err))
	}

	// remove new line
	content := strings.TrimSuffix(string(rawContent), "\n")

	numbers := make([]int, 0)
	for _, value := range strings.Split(content, ",") {
		num, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Sprintf("'%s' is not a number\n", value))
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func LoadDefaultDigits() []int {
	return LoadDigits("input")
}

func LoadDigits(filename string) []int {
	content := LoadString(filename)[0]
	result := make([]int, len(content))

	for i := range content {
		result[i] = int(content[i] - '0')
	}
	return result
}

func LoadDefaultIntTable() [][]int {
	return LoadIntTable("input")
}

func LoadIntTable(filename string) [][]int {
	content := LoadString(filename)
	result := make([][]int, len(content))
	for i, row := range content {
		col := strings.Fields(row)
		result[i] = make([]int, len(col))
		for j, cell := range col {
			result[i][j], _ = strconv.Atoi(strings.TrimSpace(cell))
		}
	}

	return result
}
