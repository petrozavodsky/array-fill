package main

import (
	"fmt"
	"testing"
)

func TestCreateSimpleArray(t *testing.T) {

	test := CreateSimpleArray()

	if 0 != (len(test) - intervalMax - 1) {
		t.Error("Test failed")
	}

}

func TestCreateMultiArray(t *testing.T) {

	var data []int

	for i := 0; i < rows*cols; i++ {
		data = append(data, 0)

	}

	out := CreateMultiArray(data)

	for i, _ := range out {
		row := out[i]

		for i := 0; cols > i; i++ {
			if 0 != row[i] {
				t.Error("Test failed")
			}
		}

	}

	fmt.Println(out)
}
