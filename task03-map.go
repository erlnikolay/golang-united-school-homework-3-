package homework

import (
	"sort"
)

func getKeys(input map[int]string) (result []int) {
	var resKeys []int
	for mKey, _ := range input {
		//fmt.Println("Key:", mKey)
		resKeys = append(resKeys, mKey)
	}
	sort.Ints(resKeys)
	return resKeys
}

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var resSlice []string
	sortedKeys := getKeys(input)
	//fmt.Println(sortedKeys)
	for iter := 0; iter < len(sortedKeys); iter++ {
		resSlice = append(resSlice, input[sortedKeys[iter]])
	}
	return resSlice
}

//func main() {
//	sourceVal := map[int]string{10: "aa", 1: "bb", 500: "cc"}
//	fmt.Println(sortMapValues(sourceVal))
//}
