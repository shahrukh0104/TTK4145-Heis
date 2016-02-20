package matrix

import ("fmt")

type Matrix struct{
	col int
	row int
	matrix [][] int
}

func matrixNew(col int, row int) Matrix {
	m := new(Matrix)
	m.row = x
	m.col = y

	m.arr = make([][]int, X)

	for i := 0; i<x; i++{
		m.arr[i] = make([]int, y)
	}

	return m
}


func matrixGet(row int, col int) (m Matrix) {
	return m.arr[row][col]
}

func matrixSet(row int, col int, val int) (m Matrix){
	m.arr[row][col] = val
}