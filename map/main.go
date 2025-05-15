package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{}
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)
	// delete
	delete(m, "a")
	fmt.Println(m)
	// check if key exists
	if v, ok := m["b"]; ok {
		fmt.Println("key exists: ", v)
	} else {
		fmt.Println("key does not exist")
	}
	// iterate over map
	for k, v := range m {
		fmt.Println(k, v)
	}
	// map of maps
	m2 := map[string]map[string]int{}
	m2["a"] = map[string]int{"b": 1, "c": 2}
	m2["b"] = map[string]int{"d": 3, "e": 4}
	m2["c"] = map[string]int{"f": 5, "g": 6}
	fmt.Println(m2)
	// iterate over map of maps
	for k, v := range m2 {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}
	fmt.Println("===================================")
	// map of slices
	m3 := map[string][]int{}
	m3["a"] = []int{1, 2, 3}
	m3["b"] = []int{4, 5, 6}
	m3["c"] = []int{7, 8, 9}
	// fmt.Println(m3)
	// iterate over map of slices
	for k, v := range m3 {
		fmt.Println(k)
		for k, v2 := range v {
			fmt.Println(k, v2)
			// fmt.Println(v2)
		}
	}
	fmt.Println("===================================")
	fmt.Println(maps.Equal(m, m2["a"]))
}
