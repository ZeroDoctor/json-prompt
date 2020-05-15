package jsontype

import "sort"

// SortMap : sort a map based on its keys
func SortMap(targetMap map[string]interface{}) []string {
	var keys []string

	for k := range targetMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
