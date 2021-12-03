package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	measures, err := readMeasures("input.txt")
	if err != nil {
		log.Fatalln("Error while reading the file with errors: ", err)
	}
	numberOfTimesSingleDepthIncreased := countDepthIncreases(measures)
	fmt.Printf("The depth has increased %d times\n", numberOfTimesSingleDepthIncreased)
	numberOfTimesSlidingWindowDepthIncreased := countSlidingWindowDepthIncreases(measures)
	fmt.Printf("The sliding window depth has increased %d times\n", numberOfTimesSlidingWindowDepthIncreased)
}

func readMeasures(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var measures []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		measure, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		measures = append(measures, measure)
	}
	return measures, nil
}

func countDepthIncreases(measures []int) int {
	depthIncreases := -1
	var previousMeasure int
	for _, measure := range measures {
		if measure > previousMeasure {
			depthIncreases++
		}
		previousMeasure = measure
	}
	return depthIncreases
}

func countSlidingWindowDepthIncreases(measures []int) int {
	depthIncreases := -1
	var previousWindow int
	for i := 0; i < len(measures)-2; i++ {
		currentWindow := measures[i] + measures[i+1] + measures[i+2]
		if currentWindow > previousWindow {
			depthIncreases++
		}
		previousWindow = currentWindow
	}
	return depthIncreases
}
