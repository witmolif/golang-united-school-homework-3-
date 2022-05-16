package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	vals := make([]string, len(input))

	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	for i, k := range keys {
		vals[i] = input[k]
	}
	return vals
}
