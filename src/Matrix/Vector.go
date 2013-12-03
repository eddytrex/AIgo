package Matrix

import (
	"math/cmplx"
	"runtime"
)

func (this *Matrix) UnitVector() *Matrix {
	duplicate := this.Copy()
	if this.n == 1 || this.m == 1 {

		norm := this.FrobeniusNorm()
		duplicate = duplicate.Scalar(complex(1/norm, 0))
	}
	return duplicate
}

func TwoVariableFuncionApply(i0, i1 int, A, B, C *Matrix, done chan<- bool, f func(a, b complex128) complex128) {
	di := (i1 - i0)

	if di >= THRESHOLD && runtime.NumGoroutine() < maxGoRoutines {
		done2 := make(chan bool, 2)
		mi := i0 + di/2
		go TwoVariableFuncionApply(i0, mi, A, B, C, done2, f)
		go TwoVariableFuncionApply(mi, i1, A, B, C, done2, f)
		<-done2
		<-done2
	} else {
		for i := i0; i < i1; i++ {
			C.A[i] = f(A.A[i], B.A[i])
		}
	}
	done <- true
}

func DotMultiplication(A, B *Matrix) *Matrix {

	if A.n == B.n && A.m == B.m {
		out := NullMatrixP(A.m, A.n)
		done := make(chan bool)
		go TwoVariableFuncionApply(0, len(A.A), A, B, out, done, func(a, b complex128) complex128 { return a * b })
		<-done
		return out
	}
	return nil
}

func DotDivision(A, B *Matrix) *Matrix {

	if A.n == B.n && A.m == B.m {
		out := NullMatrixP(A.m, A.n)
		done := make(chan bool)
		f := func(a, b complex128) complex128 {
			if cmplx.Abs(b) != 0 {
				return a / b
			}
			return 0
		}
		go TwoVariableFuncionApply(0, len(A.A), A, B, out, done, f)
		<-done
		return out
	}
	return nil
}
