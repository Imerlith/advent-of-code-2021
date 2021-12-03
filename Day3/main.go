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
	if !isStructureCorrect(diagnostics) {
		log.Fatalln("Diagnostic do not have the same size!")
	}
	gammaRateInBinary, err := calculateGammaRateInBinaryForm(diagnostics)
	if err != nil {
		log.Fatalln("Error while calculating Gamma Rate in binary form with errors: ", err)
	}
	epsilonRateInBinary, err := calculateEpsilonRateInBinaryForm(gammaRateInBinary)
	if err != nil {
		log.Fatalln("Error while calculating Epsilon Rate in binary form with errors: ", err)
	}
	fmt.Printf("Gamma Binary: %s, Epsilon Binary: %s\n", gammaRateInBinary, epsilonRateInBinary)
	gammaRateInDecimal, err := strconv.ParseInt(gammaRateInBinary, 2, 64)
	if err != nil {
		log.Fatalln("Error while converting gamma rate to decimal form with errors: ", err)
	}
	epsilonRateInDecimal, err := strconv.ParseInt(epsilonRateInBinary, 2, 64)
	if err != nil {
		log.Fatalln("Error while converting epsilon rate to decimal form with errors: ", err)
	}
	fmt.Printf("Gamma rate: %d, Epsilon rate: %d\n", gammaRateInDecimal, epsilonRateInDecimal)
	fmt.Printf("And their multiplication = %d\n", gammaRateInDecimal*epsilonRateInDecimal)
}

func isStructureCorrect(diagnostics [][]int) bool {
	if len(diagnostics) == 0 {
		return true
	}
	firstDiagnosticLen := len(diagnostics[0])
	for i := 0; i < len(diagnostics); i++ {
		if len(diagnostics[i]) != firstDiagnosticLen {
			return false
		}
	}
	return true
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
	return diagnostics, nil
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
				return "", errors.New(errMsg)
			}
		}
		if nOfOnes > nOfZeros {
			gammaInBinary += "0"
		} else {
			gammaInBinary += "1"
		}
	}
	return gammaInBinary, nil
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
