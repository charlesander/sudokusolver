package utilties

import (
	"sort"
	"strconv"
)

func ConvertStringSliceToIntSlice( stringSlice []string)  []int {
	var intSlice = []int{}

	for _, i := range stringSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, j)
	}

	return intSlice
}

func StringIntSlice(a string, list []string) bool {
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