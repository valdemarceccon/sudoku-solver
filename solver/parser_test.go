package solver

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("Test that parse return error when incorrect string is passed", func(t *testing.T) {
		_, err := parse("1,2,3")

		if err == nil {
			t.Error("Error should be not nil")
		}
	})

	t.Run("Parser return a array of int if the correct string is passed", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9,
			1, 2, 3, 4, 5, 6, 7, 8, 9}

		got, err := parse("1 2 3 4 5 6 7 8 9 " +
			"1 2 3 4 5 6 7 8 9 " +
			"1 2 3 4 5 6 7 8 9 " +
			"1,2 3 4 5 6 7 8 9 " +
			"1 2 3 4 5 6 7 8 9 " +
			"1 2 3 4 5 6 7 8 9 " +
			"1 2 3 4 5 6 7 8 9 " +
			"1 2 3 4 5 6 7 8 9 " +
			"1 2 3 4 5 6 7 8 9 ")

		if err != nil {
			t.Error("Shouldn't have returned error. Returned ", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Error("String wasn't parsed correctly. wanted", want, "got", got)
		}

	})
}
