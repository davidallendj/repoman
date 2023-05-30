package util

import "fmt"

func GetKeys(m map[string]string) (ks []string) {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func GetValues(m map[string]string) (vs []string) {
	var values []string
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func PrintMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}
