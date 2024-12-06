package aoclib

import (
	"bufio"
	"fmt"
	"os"
)

func ReadSimpleInput(inputname string) string {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	puzzle := ""
	for scan.Scan() {
		puzzle += scan.Text()
	}

	return puzzle
}

func ReadInputMatrix(inputname string) [][]byte {
	puzzle := [][]byte{}
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		puzzle = append(puzzle, []byte(line))
	}

	return puzzle
}

func ReadInputBase(inputname string) int {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		_ = scan.Text()
	}

	return 0
}
