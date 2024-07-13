package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func extractCalibrationValue(line string) int {
	var firstDigit, lastDigit rune
	for _, char := range line {
		if unicode.IsDigit(char) {
			if firstDigit == 0 {
				firstDigit = char
			}
			lastDigit = char
		}
	}
	if firstDigit != 0 && lastDigit != 0 {

		value, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		return value
	}
	return 0
}


func sumCalibrationValues(lines []string) int {
	totalSum := 0
	for _, line := range lines {
		totalSum += extractCalibrationValue(line)
	}
	return totalSum
}

func main() {

	file, err := os.Open("calibration_document.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	result := sumCalibrationValues(lines)
	fmt.Println(result)
}
