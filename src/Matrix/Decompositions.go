package Matrix

import (
	"math"
	"math/cmplx"
)

//Pivot
func (this *Matrix) Pivot() (*Matrix, *Matrix) {
	pivot := this.Copy()
	if this.m == this.n {
		p := I(this.m)
		for j := 1; j <= this.m; j++ {
			max := cmplx.Abs(pivot.GetValue(j, j))
			row := j
			for i := j; i <= pivot.m; i++ {
				pvalue := pivot.GetValue(i, j)
				if cmplx.Abs(pvalue) > max {
					max = cmplx.Abs(pvalue)
					row = i
				}
			}
			if j != row {
				tj := this.GetRow(j)
				trow := this.GetRow(row)

				pivot.SetRow(j, trow)
				pivot.SetRow(row, tj)

				pj := p.GetRow(j)
				prow := p.GetRow(row)

				p.SetRow(j, prow)
				p.SetRow(row, pj)
			}
		}
		return p, pivot
	}
	return nil, nil
}

//LU Decomposition of a Matrix implementation from gomatrix
func (this *Matrix) LUDec() (L *Matrix, U *Matrix, P *Matrix) {
	L = NullMatrixP(this.n, this.m)
	U = NullMatrixP(this.n, this.m)
	P, thisi := this.Pivot()
	for j := 1; j <= thisi.m; j++ {
		L.SetValue(j, j, complex(1.0, 0.0))
		for i := 1; i <= j; i++ {
			sum := complex(0.0, 0.0)
			for k := 1; k <= i; k++ {
				sum = sum + U.GetValue(k, j)*L.GetValue(i, k)
			}
			U.SetValue(i, j, thisi.GetValue(i, j)-sum)
		}
		for i := j; i <= thisi.m; i++ {
			sum := complex(0.0, 0.0)
			for k := 1; k < j; k++ {
				sum = sum + U.GetValue(k, j)*L.GetValue(i, k)
			}
			L.SetValue(i, j, (thisi.GetValue(i, j)-sum)/U.GetValue(j, j))
		}
	}
	return
}

//QR Decomposition using  Householder reflections
func (this *Matrix) QR() (Q1, R1 *Matrix) {
	n := this.n //rows
	m := this.m //columns

	last := n - 1

	var alpha complex128
	Q := I(m)
	if m == n {
		last--
	}
	Ai := this.Copy()

	for i := 0; i <= last; i++ {
		b := Ai.GetSubMatrix(i+1, i+1, m-i+1, n-i+1)
		x := b.GetColumn(1)

		e := NullMatrix(x.m, 1)
		e.SetValue(i+1, 1, 1)

		x1 := x.GetValue(i+1, 1)

		alpha = cmplx.Exp(complex(0, -math.Atan2(imag(x1), real(x1)))) * complex(x.FrobeniusNorm(), 0)

		u, _ := Sustract(x, e.Scalar(alpha))
		v := u.UnitVector()

		hht, _ := v.HouseholderTrasformation()

		h := SetSubMatrixToI(m, i+1, hht)

		Q = Product(Q, h)

		Ai = Product(h, Ai)

	}
	return Q, Ai
}

// QR Decomposition using  Householder reflections
func (this *Matrix) QRDec() (Q1, R1 *Matrix) {
	Q := NullMatrixP(this.m, this.n)
	R := NullMatrixP(this.m, this.n)
	var first = true
	var alpha complex128
	var Qp *Matrix
	Ai := this.Copy()
	for i := 1; i < this.m; i++ {

		X := Ai.GetColumn(i)

		e := NullMatrix(X.m, 1)
		e.SetValue(i, 1, 1)

		x1 := X.GetValue(i, 1)
		if real(x1) > 0 {
			alpha = complex(-X.FrobeniusNorm(), 0)
		} else {
			alpha = complex(X.FrobeniusNorm(), 0)
		}

		u, _ := Sustract(X, e.Scalar(alpha))
		v := u.UnitVector()

		Qi, _ := v.HouseholderTrasformation()

		if first {
			Qp = Product(Qi, this)
			Q = Qi
			first = false
		} else {
			Q = Product(Q, Qi)
			Qp = Product(Qi, Ai)
		}

		for l := 1; l <= i; l++ {
			Qp = Qp.SubMatrix(1, 1)
		}

		Qp = SetSubMatrixToI(this.n, i+1, Qp)
		Ai = Qp
	}

	R = Product(Q.Transpose(), this)
	return Q, R
}

//TODO this can be improve
// Set a matrix in the position beginin in PosI,PosI on a Identity Matrix

func SetSubMatrixToI(n int, posI int, pQ *Matrix) *Matrix {
	out := I(n)
	if posI < n && (posI+pQ.n-1) == n {
		if pQ.m < n {
			setMatrix := NullMatrixP((n - pQ.m), pQ.n)
			pQ = pQ.AddRowsToTop(setMatrix)

			for i := 1; i <= pQ.n; i++ {
				ci := pQ.GetColumn(i)
				out.SetColumn(i+posI-1, *ci)
			}

		} else if pQ.m == n {

			return pQ
		}
		return out
	}
	return nil
}
