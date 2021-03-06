package Matrix

import (
	"math/cmplx"
)

//QR algorithm for EigenValues
func (this *Matrix) EigenValues(Tol complex128) *Matrix {

	Ai := this.Copy()
	Error := 1.0

	Qi, Ri := Ai.QRDec()
	Ai = Product(Ri, Qi)

	xi, _ := Ai.GetDiagonal()

	for Error > cmplx.Abs(Tol) {

		Qi, Ri := Ai.QRDec()
		Ai = Product(Ri, Qi)

		xi1, _ := Ai.GetDiagonal()
		diff, _ := Sustract(xi, xi1)
		Error = diff.FrobeniusNorm()
		xi = xi1

	}

	Eig := NullMatrixP(this.n, 1)
	for i := 1; i <= this.n; i++ {
		Eig.SetValue(i, 1, Ai.GetValue(i, i))
	}
	return Eig
}

func (this *Matrix) EigenVector(eigenV complex128) *Matrix {
	Id := I(this.n)
	//Z:=NullMatrixP(this.n,1)

	S, _ := Sustract(this, Id.Scalar(eigenV))

	println(S.ToString(), "<solve")

	return S.Transpose()
}
