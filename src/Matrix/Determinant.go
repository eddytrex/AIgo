package Matrix

import (
	//"math"
	"errors"
)

// Return the Matrix submatrix Matrix(this) of the position i,j
func (this *Matrix) SubMatrix(i, j int) *Matrix {
	out := this.MatrixWithoutRow(i).MatrixWithoutColumn(j)
	return out
}

// return the determinant of a square Matrix
// O(n!) I don't think someone will use it
func (this *Matrix) Det_LapaceExpasion() (complex128, error) {
	if this.n == this.m {
		if this.n == 1 {
			return this.GetValue(1, 1), nil

		} else {
			var sum complex128
			sum = 0
			for i := 1; i <= this.m; i++ {

				temp, _ := this.SubMatrix(1, i).Det_LapaceExpasion()

				if i%2 != 0 {

					temp = temp * this.GetValue(i, 1)
				} else {
					temp = temp * this.GetValue(i, 1) * -1
				}
				sum = sum + temp
			}
			return sum, nil
		}

	}
	return 0, errors.New(" the Matrix have to be square")
}

// Return the determinant of a Matrix by LU  Decomposition
func (this *Matrix) Det_LU() (complex128, error) {

	if this.GetMRows() == this.GetNColumns() {
		_, U, _ := this.LUDec()

		var Det complex128
		Det = 1

		for ui := 1; ui <= U.n; ui++ {
			uii := U.GetValue(ui, ui)
			Det = Det * uii
		}
		return Det, nil
	}
	return 0, errors.New(" the Matrix have to be square")
}
