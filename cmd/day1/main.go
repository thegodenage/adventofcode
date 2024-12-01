package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var fileBytes []byte

func main() {
	fmt.Println(len(fileBytes))

	scanner := bufio.NewScanner(bytes.NewReader(fileBytes))

	leftArr, rightArr, err := getSlicesFromFile(scanner)
	if err != nil {
		panic(fmt.Sprintf("get slices from file: %s", err.Error()))
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("read files using scanner: %s", err.Error()))
	}

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	log.Printf("len: %d", len(leftArr))
	log.Printf("len: %d", len(rightArr))

	sum := 0
	for i, left := range leftArr {
		diff := math.Abs((float64(rightArr[i]) - float64(left)))
		// dim := math.Dim(float64(rightArr[i]), float64(left))
		log.Printf("diff: %f", diff)
		sum += int(diff)
	}

	log.Printf("sum: %d", sum)

	// part two
	similarityScore := 0
	for _, left := range leftArr {
		count := 0
		for _, right := range rightArr {
			if right == left {
				count++
			}
		}

		similarity := left * count

		similarityScore += similarity
	}

	log.Printf("similarity score: %d", similarityScore)
}

func getSlicesFromFile(scanner *bufio.Scanner) ([]int, []int, error) {
	var left, right []int

	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, "   ")

		firstWord, err := strconv.Atoi(words[0])
		if err != nil {
			return nil, nil, fmt.Errorf("get first word from line: %w", err)
		}

		secondWord, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, nil, fmt.Errorf("get second word from line: %w", err)
		}

		left = append(left, firstWord)
		right = append(right, secondWord)
	}

	return left, right, nil
}
