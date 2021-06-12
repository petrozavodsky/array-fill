package main

func CreateMultiArray(ar []int) [rows][cols]int {

	var resultArray [rows][cols]int
	resultArrayCols := 0
	resultArrayRows := 0

	for i := 0; i < rows*cols; i++ {

		if cols == resultArrayRows {
			resultArrayRows = 0
			resultArrayCols++
		}

		resultArray[resultArrayCols][resultArrayRows] = ar[i]

		resultArrayRows++

	}

	return resultArray
}

func CreateSimpleArray() [intervalMax + 1]int {
	var array [intervalMax + 1]int

	for i := intervalMin; i < intervalMax+1; i++ {
		array[i] = i
	}

	return array

}
