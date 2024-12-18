package gollections

func Any[T comparable](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

func All[T comparable](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func Equals[T comparable](slice1 []T, slice2 []T, comparator func(T, T) bool) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, item := range slice1 {
		if !comparator(item, slice2[i]) {
			return false
		}
	}
	return true
}

// 2D Arrays / Slices
func GetColumn(matrix [][]int, colNum int) []int {
	column := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		column = append(column, matrix[i][colNum])
	}

	return column
}

func GetRow(matrix [][]int, rowNum int) []int {
	row := make([]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		row = append(row, matrix[rowNum][i])
	}

	return row
}

func GetSum(matrix [][]int) int {
	sum := 0

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[1]); col++ {
			if matrix[row][col] != -1 {
				sum += matrix[row][col]
			}
		}
	}

	return sum
}

func ReplaceByMatchingValue(matrix [][]int, comparisonValue int) [][]int {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[1]); col++ {
			if matrix[row][col] == comparisonValue {
				matrix[row][col] = -1
			}
		}
	}

	return matrix
}

func RemoveRow(slice [][]int, s int) [][]int {
	return append([][]int{}, append(slice[:s], slice[s+1:]...)...)
}
