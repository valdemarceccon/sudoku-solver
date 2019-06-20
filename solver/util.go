package solver

func isValid(val []int) bool {
	for i := 0; i < 9; i++ {
		line := getLine(val, i)
		col := getColumn(val, i)
		square := getSquare(val, i)

		if !isSetValid(line) || !isSetValid(col) || !isSetValid(square) {
			return false
		}
	}

	return true
}

func getSquare(values []int, n int) []int {
	result := make([]int, 9)
	for i := 0; i < 3; i += 3 {
		for j := 0; j < 3; j++ {
			result[j+i*3] = values[n*3+j+9*i*3]
		}
	}

	return result
}

func isSetValid(line []int) bool {
	for _, v := range line {
		if v != 0 && countValue(v, line) > 1 {
			return false
		}
	}
	return true
}

func isSolved(values []int) bool {
	return countValue(0, values) == 0
}

func countValue(value int, vals []int) int {
	result := make(map[int]int)
	for _, v := range vals {
		result[v]++
	}

	return result[value]
}

func getColumn(val []int, i int) []int {
	col := make([]int, 9)
	for j := 0; j < 9; j++ {
		col[j] = val[j*9+i]
	}
	return col
}

func getLine(val []int, i int) []int {
	line := val[9*i : 9*i+9]
	return line
}

func findNextIndexToSolve(val []int) int {
	for i, v := range val {
		if v == 0 {
			return i
		}
	}
	return -1
}
