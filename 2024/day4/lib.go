package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputMatrix(inputname string) [][]byte {
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
