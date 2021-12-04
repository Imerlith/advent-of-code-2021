package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	diagnostics, err := readDiagnostics("./input.txt")
	if err != nil {
		log.Fatalln("Error while reading the diagnostics with errors: ", err)
	}

	gammaRate, epsilonRate, err := calculateGammaAndEpsilonRates(diagnostics)
	if err != nil {
		log.Fatalln("Error while generating gamma and epsilon rates with errors: ", err)
	}
	fmt.Printf("Gamma rate: %d, Epsilon rate: %d\n", gammaRate, epsilonRate)
	fmt.Printf("And their multiplication = %d\n", gammaRate*epsilonRate)

	oxygenRating, co2Rating, err := calculateOxygenAndCo2Rating(diagnostics)
	if err != nil {
		log.Fatalln("Error while calculating oxygen and co2 rating with errors: ", err)
	}
	fmt.Printf("Oxygen rate: %d, Co2 rate: %d\n", oxygenRating, co2Rating)
	fmt.Printf("And their multiplication = %d\n", oxygenRating*co2Rating)
}

func readDiagnostics(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var diagnostics [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numberAsTable, err := transformStringToIntTableInBinary(line)
		if err != nil {
			return nil, err
		}
		diagnostics = append(diagnostics, numberAsTable)
	}
	return checkDiagnosticStructure(diagnostics)
}

func checkDiagnosticStructure(diagnostics [][]int) ([][]int, error) {
	if len(diagnostics) == 0 {
		return diagnostics, nil
	}
	firstDiagnosticLen := len(diagnostics[0])
	for i := 0; i < len(diagnostics); i++ {
		if len(diagnostics[i]) != firstDiagnosticLen {
			return nil, errors.New("diagnostic do not have the same size")
		}
	}
	return diagnostics, nil
}

func calculateGammaAndEpsilonRates(diagnostics [][]int) (int64, int64, error) {
	gammaRateInBinary, err := calculateGammaRateInBinaryForm(diagnostics)
	if err != nil {
		return -2000, -2000, err
	}
	epsilonRateInBinary, err := calculateEpsilonRateInBinaryForm(gammaRateInBinary)
	if err != nil {
		return -2000, -2000, err
	}
	fmt.Printf("Gamma Binary: %s, Epsilon Binary: %s\n", gammaRateInBinary, epsilonRateInBinary)
	gammaRateInDecimal, err := strconv.ParseInt(gammaRateInBinary, 2, 64)
	if err != nil {
		return -2000, -2000, err
	}
	epsilonRateInDecimal, err := strconv.ParseInt(epsilonRateInBinary, 2, 64)
	if err != nil {
		return -2000, -2000, err
	}
	return gammaRateInDecimal, epsilonRateInDecimal, nil
}

func transformStringToIntTableInBinary(line string) ([]int, error) {
	var table []int
	for _, char := range line {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}
		table = append(table, num)
	}
	return table, nil
}

func calculateGammaRateInBinaryForm(diagnostics [][]int) (string, error) {
	var gammaInBinary string
	for i := 0; i < len(diagnostics[0]); i++ {
		mostCommon, err := getCommonBitInColumn(diagnostics, i, most)
		if err != nil {
			return "", err
		}
		gammaInBinary += strconv.Itoa(mostCommon)
	}
	return gammaInBinary, nil
}

func getCommonBitInColumn(diagnostics [][]int, i int, typeOfCommon rating) (int, error) {
	var nOfOnes int
	var nOfZeros int
	for j := 0; j < len(diagnostics); j++ {
		numberAtPosition := diagnostics[j][i]
		switch numberAtPosition {
		case 0:
			nOfZeros++
		case 1:
			nOfOnes++
		default:
			errMsg := fmt.Sprintf("Invalid number: %d at position (%d, %d)\n", numberAtPosition, i, j)
			return -2000, errors.New(errMsg)
		}
	}
	switch typeOfCommon {
	case most:
		if nOfOnes >= nOfZeros {
			return 1, nil
		} else {
			return 0, nil
		}
	case least:
		if nOfOnes >= nOfZeros {
			return 0, nil
		} else {
			return 1, nil
		}
	default:
		return 1, nil
	}
}

func calculateEpsilonRateInBinaryForm(gammaRateInBinary string) (string, error) {
	var epsilonInBinary string
	for _, char := range gammaRateInBinary {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return "", err
		}
		if num == 1 {
			epsilonInBinary += "0"
		} else {
			epsilonInBinary += "1"
		}
	}
	return epsilonInBinary, nil
}

func calculateOxygenAndCo2Rating(diagnostics [][]int) (int64, int64, error) {
	oxygenRatingInBinary, err := calculateComplexRating(diagnostics, most)
	if err != nil {
		return -2000, -2000, err
	}
	co2RatingInBinary, err := calculateComplexRating(diagnostics, least)
	if err != nil {
		return -2000, -2000, err
	}
	oxygenRatingInDecimal, err := strconv.ParseInt(oxygenRatingInBinary, 2, 64)
	if err != nil {
		return -2000, -2000, err
	}
	co2RatingInDecimal, err := strconv.ParseInt(co2RatingInBinary, 2, 64)
	if err != nil {
		return -2000, -2000, err
	}
	return oxygenRatingInDecimal, co2RatingInDecimal, nil
}

func calculateComplexRating(diagnostics [][]int, typeOfCommon rating) (string, error) {
	reducedDiagnostics := diagnostics
	for i := 0; i < len(diagnostics[0]); i++ {
		common, err := getCommonBitInColumn(reducedDiagnostics, i, typeOfCommon)
		if err != nil {
			return "", err
		}
		reducedDiagnostics = eliminateDiagnosticsWithDifferentBitInColumn(reducedDiagnostics, common, i)
	}

	return getRatingInBinary(reducedDiagnostics[0]), nil
}

func eliminateDiagnosticsWithDifferentBitInColumn(diagnostics [][]int, mostCommon int, index int) [][]int {
	if len(diagnostics) == 1 {
		return diagnostics
	}
	var reducedDiagnostics [][]int
	for i := 0; i < len(diagnostics); i++ {
		if diagnostics[i][index] == mostCommon {
			reducedDiagnostics = append(reducedDiagnostics, diagnostics[i])
		}
	}

	return reducedDiagnostics
}

func getRatingInBinary(diagnostic []int) string {
	var diagnosticInBinary string
	for _, number := range diagnostic {
		diagnosticInBinary += strconv.Itoa(number)
	}
	return diagnosticInBinary
}

type rating int

const (
	most rating = iota
	least
)
