package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sudoku-solver/solver"
)

func main() {
	filePath := os.Args[1]
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(bytes)

	fmt.Println(solver.Solve(fileContent))
}
