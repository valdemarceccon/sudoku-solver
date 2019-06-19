package solver

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func parse(content string) ([]int, error) {
	fields := strings.FieldsFunc(content, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	if len(fields) != 81 {
		return nil, fmt.Errorf("Incorrect string parameter, size %d, should be 81", len(fields))
	}
	result := make([]int, 81)
	for index, val := range fields {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		result[index] = i
	}
	return result, nil
}
