package ANN

import (
	//"fmt"
	"Matrix"
	"math/cmplx"
)

func HalfDistance(T, O *Matrix.Matrix) *Matrix.Matrix {
	r := Matrix.DistanceSquare(O, T).Scalar(complex(0.5, 0))

	return r
}

func DerivateHalfDistance(T, O *Matrix.Matrix) *Matrix.Matrix {

	r, _ := Matrix.Sustract(T, O)

	return r
}

func CorssEntorpy(T, O *Matrix.Matrix) *Matrix.Matrix {
	log := func(x complex128) complex128 { return cmplx.Log(x) }
	return Matrix.DotMultiplication(T, O.Apply(log))

}
