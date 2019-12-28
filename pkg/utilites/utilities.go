package utilties

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func ExplodeToIntSlice(sudokuCellValues string) []int {
	tmp := strings.Split(sudokuCellValues, ",")
	values := make([]int, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		values = append(values, v)
	}

	return values
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Disregard 0 values as we're not comparing them
func AreSudokuValuesUnique(input []int) bool {

	checkAgainst := []int{}
	for _, value := range input {
		if value != 0 {
			checkAgainst = append(checkAgainst, value)
		}
	}

	sort.Ints(checkAgainst[:])

	len := len(checkAgainst)
	i := 0
	if (len > 1) {
		for {

			if (i == len-1) {
				break
			}
			if checkAgainst[i] == checkAgainst[i+1] {
				return false
			}

			i++

		}
	}

	return true
}