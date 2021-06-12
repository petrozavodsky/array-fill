package main

import (
	"math/rand"
	"time"
)

type Helpers struct {
	min int
	max int
}

func (h *Helpers) ShuffleArray(array []int) []int {

	for i := 0; i < h.min+1; i++ {
		r := h.randomInterval(0, len(array)-1)
		old := array[i]
		array[i] = array[r]
		array[r] = old
	}

	return array
}

func (h *Helpers) randomInterval(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())

	return min + rand.Intn(max-min)
}
