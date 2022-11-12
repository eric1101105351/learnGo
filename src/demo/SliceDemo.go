package demo

import (
	"fmt"
)

func MySlice(length int) []int {
	array := make([]int, length)

	for i := 0; i < length; i++ {
		array[i] = i
	}

	for i, v := range array {
		fmt.Printf("index:%v, value:%v\n", i, v)
	}
	return array
}

func MyAppendSlice(slice []int, length int) {

	for i := 0; i < length; i++ {
		slice = append(slice, i+10)
	}

	for i, v := range slice {
		fmt.Printf("index:%v, value:%v\n", i, v)
	}

}
