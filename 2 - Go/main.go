package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args
	DEBUG := false
	PART := 1

	if len(args) > 1 && args[1] == "true" {
		DEBUG = true
	}

	if len(args) > 2 {
		convertedNum, err := strconv.Atoi(args[2])
		check(err)
		PART = convertedNum
	}

	var file *os.File
	var err error

	var file_path string

	if DEBUG {
		file_path = fmt.Sprintf("2 - Go/input/t_input_%s.txt", map[int]string{
			1: "one",
			2: "two",
		}[PART])
	} else {
		file_path = "2 - Go/input/input.txt"
	}

	file, err = os.Open(file_path)
	check(err)

	if PART == 1 {
		part_1(file)
	} else {
		part_2(file)
	}
}

func part_1(file *os.File) {
	safeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		safe := checkLine(line, 0)

		fmt.Println(safe)
		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func part_2(file *os.File) {
	safeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		safe := checkLine(line, 1)

		fmt.Println(safe)
		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func checkLine(line string, tolerance int) (safe bool) {
	sort := ""
	previousNumber := 0
	badLevels := 0
	initialized := false

	for _, number := range strings.Split(line, " ") {
		convertedNumber, err := strconv.Atoi(number)
		check(err)

		if initialized == false {
			previousNumber = convertedNumber
			initialized = true
			continue
		}

		if sort == "" && convertedNumber > previousNumber {
			sort = "asc"
		}

		if sort == "" && convertedNumber < previousNumber {
			sort = "desc"
		}

		if convertedNumber == previousNumber {
			badLevels++
			continue
		}

		if sort == "asc" && ((convertedNumber-previousNumber) > 3 || convertedNumber < previousNumber) {
			badLevels++
			continue
		}

		if sort == "desc" && ((previousNumber-convertedNumber) > 3 || convertedNumber > previousNumber) {
			badLevels++
			continue
		}
		previousNumber = convertedNumber

		if badLevels > tolerance {
			return false
		}
	}

	return true
}
