package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	commands, err := readCommands("./input.txt")
	if err != nil {
		log.Fatalln("Error while reading commands with error: ", err)
	}
	simpleFinalPosition := calculateSimpleFinalPosition(commands)
	fmt.Printf("The final position of the submarine is horizontal: %d, depth: %d\n", simpleFinalPosition.horizontal, simpleFinalPosition.depth)
	fmt.Printf("The multiplication is %d\n", simpleFinalPosition.horizontal*simpleFinalPosition.depth)
	finalPositionWithAim := calculateFinalPositionWithAim(commands)
	fmt.Printf("The final final position of the submarine using aim is horizontal: %d, depth: %d\n", finalPositionWithAim.horizontal, finalPositionWithAim.depth)
	fmt.Printf("The multiplication is %d\n", finalPositionWithAim.horizontal*finalPositionWithAim.depth)
}

func readCommands(path string) ([]*Command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var commands []*Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		command, err := transformToCommand(line)
		if err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func transformToCommand(line string) (*Command, error) {
	split := strings.Split(line, " ")
	if len(split) != 2 {
		errMsg := fmt.Sprintf("error while reading line: %s\n", line)
		return nil, errors.New(errMsg)
	}
	var command = Command{}
	switch firstPart := strings.ToLower(split[0]); firstPart {
	case "forward":
		command.direction = forward
	case "down":
		command.direction = down
	case "up":
		command.direction = upward
	default:
		errMsg := fmt.Sprintf("Invalid command found: %s\n", firstPart)
		return nil, errors.New(errMsg)
	}
	value, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, err
	}
	command.value = value

	return &command, nil
}

func calculateSimpleFinalPosition(commands []*Command) *Position {
	var finalPosition = Position{
		horizontal: 0,
		depth:      0,
	}
	for _, command := range commands {
		switch command.direction {
		case forward:
			finalPosition.horizontal += command.value
		case upward:
			finalPosition.depth -= command.value
		case down:
			finalPosition.depth += command.value
		}
	}
	return &finalPosition
}

func calculateFinalPositionWithAim(commands []*Command) *Position {
	var finalPosition = Position{
		horizontal: 0,
		depth:      0,
		aim:        0,
	}
	for _, command := range commands {
		switch command.direction {
		case forward:
			finalPosition.horizontal += command.value
			finalPosition.depth += finalPosition.aim * command.value
		case upward:
			finalPosition.aim -= command.value
		case down:
			finalPosition.aim += command.value
		}
	}
	return &finalPosition
}

type Position struct {
	horizontal int
	depth      int
	aim        int
}
type Command struct {
	direction direction
	value     int
}
type direction int

const (
	forward direction = iota
	down
	upward
)
