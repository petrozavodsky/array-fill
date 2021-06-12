package main

import (
	"reflect"
	"testing"
)

func TestShuffleArray(t *testing.T) {

	tests := [][7]int{
		{0, 1, 2, 3, 4, 5, 6},
		{99, 100, 101, 102, 103, 105},
		{0, 0, 0, 0, 0, 0, 0},
		{-1, 1, 1, 1, 1, 1, 1},
	}

	for i, test := range tests {

		max := test[len(test)-1]
		min := test[0]

		helpers := Helpers{max, min}

		o := helpers.ShuffleArray(test[0:len(test)])

		if false != reflect.DeepEqual(o, tests[i]) {
			t.Fatal("The array is not shuffled")
		}

	}
}
