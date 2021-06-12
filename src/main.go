package main

import (
	"errors"
	"log"
)

const (
	rows        = 5
	cols        = 3
	intervalMin = 0
	intervalMax = 20
)

func main() {

	if err := validate(intervalMin, intervalMax); err != nil {
		log.Fatal(err)
	}

	helpers := Helpers{intervalMin,intervalMax }

	data := CreateSimpleArray()
	shuffled := helpers.ShuffleArray(data[0 : intervalMax+1])
	View(CreateMultiArray(shuffled))
}

func validate(min int, max int) error {

	if max <= min {
		return errors.New("min must be strictly less than max")
	}

	if 0 > max || 0 > min {
		return errors.New("arguments are required greater than zero ")
	}

	return nil
}
