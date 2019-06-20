package solver

import (
	"errors"
	"fmt"
)

func Solve(rawValues string) ([]int, error) {
	parsedValues, err := parseAndValidate(rawValues)

	if err != nil {
		return nil, err
	}

	solvedPuzzle := solve(parsedValues)

	return solvedPuzzle, nil
}

func solve(values []int) []int {
	fmt.Printf("amount to solve %d\n", countValue(0, values))
	var solvedPuzzle []int = nil

	if isSolved(values) && isValid(values) {
		return values
	}

	nextIndexToSolve := findNextIndexToSolve(values)

	if nextIndexToSolve >= 0 {
		nextStep := make([]int, 81)
		copy(nextStep, values)
		for tentativa := 1; tentativa <= 9; tentativa++ {
			nextStep[nextIndexToSolve] = tentativa
			if isValid(nextStep) {
				if isSolved(nextStep) {
					return nextStep
				}
				solvedPuzzle = solve(nextStep)
				if solvedPuzzle != nil && isValid(solvedPuzzle) && isSolved(solvedPuzzle) {
					break
				}
			}
		}
	}

	return solvedPuzzle
}

func parseAndValidate(rawValues string) ([]int, error) {
	parsedValues, e := parse(rawValues)

	if e != nil {
		return nil, e
	}

	if !isValid(parsedValues) {
		return nil, errors.New("the passed puzzle is not valid")
	}

	return parsedValues, nil
}
