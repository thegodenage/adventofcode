package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var fileBytes []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(fileBytes))

	reports, err := getReports(scanner)
	if err != nil {
		panic(fmt.Sprintf("get reports: %s", err.Error()))
	}

	// part one
	validReportsCount := 0
	for _, report := range reports {
		if isValidReport(report) {
			validReportsCount++
		}
	}

	fmt.Printf("num of valid reports: %d\n", validReportsCount)

	// part two
	validReportsCount = 0
	for _, report := range reports {
		if isValidReportWithTolerance(report) {
			validReportsCount++
		}
	}

	fmt.Printf("num of valid reports with tolerance: %d\n", validReportsCount)
}

func getReports(scanner *bufio.Scanner) ([][]int, error) {
	var reports [][]int

	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, " ")

		var report []int
		for i, w := range words {
			num, err := strconv.Atoi(w)
			if err != nil {
				return nil, fmt.Errorf("get num: %d from line: %w", i, err)
			}

			report = append(report, num)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func isValidReport(report []int) bool {
	if len(report) < 2 {
		return false
	}

	isAscending := true
	isDescending := true

	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))

		if diff < 1 || diff > 3 {
			return false
		}

		if report[i] > report[i-1] {
			isDescending = false
		}

		if report[i] < report[i-1] {
			isAscending = false
		}
	}

	return isAscending || isDescending
}

func isValidReportWithTolerance(report []int) bool {
	if isValidReport(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		tempReport := append([]int{}, report[:i]...)
		tempReport = append(tempReport, report[i+1:]...)
		fmt.Println(tempReport)
		if isValidReport(tempReport) {
			return true
		}
	}

	return false
}
