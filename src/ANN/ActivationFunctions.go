package ANN

import (
	"Matrix"
	"math/cmplx"
)

func Sigmoid(x complex128) complex128 {
	return 1 / (1 + cmplx.Exp(-x))
}

func DSigmoid(x complex128) complex128 {
	return (1 / (1 + cmplx.Exp(-x))) * (1 - (1 / (1 + cmplx.Exp(-x))))
}

func SigmoidLayer(X *Matrix.Matrix) *Matrix.Matrix {
	return X.Apply(Sigmoid)
}

func DSigmoidLayer(X *Matrix.Matrix) *Matrix.Matrix {
	return X.Apply(DSigmoid)
}

func Softmax(X *Matrix.Matrix) *Matrix.Matrix {
	Total := 1 / X.TaxicabNorm()
	Y := X.Scalar(complex(Total, 0))

	return Y
}

func DSoftmax(X *Matrix.Matrix) *Matrix.Matrix {
	Total := 1 / X.TaxicabNorm()
	Y := X.Scalar(complex(Total, 0))

	S, _ := Matrix.Sustract(Matrix.FixValueMatrix(X.GetNColumns(), X.GetNColumns(), 1.0), X)

	YD := Matrix.DotMultiplication(Y, S)
	return YD
}
