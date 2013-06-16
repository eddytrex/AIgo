package MachineLearning

import (
	"Matrix"
	"errors"
	//"math"
	"math/cmplx"
	"runtime"
)

const THRESHOLD = 100

var maxGoRoutines = runtime.GOMAXPROCS(0) + 2

type Hypothesis struct {
	ThetaParameters *Matrix.Matrix
	M               int
	Sum             *Matrix.Matrix

	H func(complex128) complex128
}

func (this *TrainingSet) AddX0() {
	m := this.Xs.GetMRows()
	x0 := Matrix.NullMatrix(m, 1)

	for i := 1; i <= m; i++ {
		x0.SetValue(i, 1, 1.0)
	}

	this.Xs = x0.AddColumn(this.Xs)
}

func (this *Hypothesis) ApplyHypothesisToTrainingSet(Ts TrainingSet) *Matrix.Matrix {

	m := Ts.Xs.GetMRows()

	hx := Matrix.NullMatrix(m, 1)

	if this.ThetaParameters.GetNColumns() == Ts.Xs.GetNColumns() {
		for i := 1; i <= Ts.Xs.GetMRows(); i++ {
			xi := Ts.Xs.GetRow(i)

			Thi := Matrix.Product(xi, this.ThetaParameters.Transpose())

			hx.SetValue(i, 1, Thi.GetValue(1, 1))

		}
		return &hx
	}
	return nil
}

func (this *Hypothesis) Parallel_DiffH1Ys(Ts *TrainingSet) (*Matrix.Matrix, *Matrix.Matrix) {
	m := Ts.Xs.GetMRows()
	hx := Matrix.NullMatrixP(m, 1)
	hxt := Matrix.NullMatrixP(1, m)

	if this.ThetaParameters.GetNColumns() == Ts.Xs.GetNColumns() {
		done := make(chan bool)
		go this.part_DiffH1Ys(1, m, Ts, hx, hxt, done)
		<-done
	}
	return hx, hxt
}

func (this *Hypothesis) part_DiffH1Ys(i0, i1 int, Ts *TrainingSet, Ret *Matrix.Matrix, RetT *Matrix.Matrix, done chan<- bool) {
	di := i1 - i0

	if di >= THRESHOLD && runtime.NumGoroutine() < maxGoRoutines {
		done2 := make(chan bool, THRESHOLD)

		mi := i0 + di/2
		go this.part_DiffH1Ys(i0, mi, Ts, Ret, RetT, done2)
		go this.part_DiffH1Ys(mi, i1, Ts, Ret, RetT, done2)
		<-done2
		<-done2
	} else {
		for i := i0; i <= i1; i++ {
			xi := Ts.Xs.GetRow(i)

			Thi := Matrix.Product(xi, this.ThetaParameters.Transpose())
			temp := this.H(Thi.GetValue(1, 1)) - Ts.Y.GetValue(1, i)
			Ret.SetValue(i, 1, temp)
			RetT.SetValue(1, i, temp)
		}
	}
	done <- true
}

func (this *Hypothesis) DiffH1Ys(Ts TrainingSet) *Matrix.Matrix {

	m := Ts.Xs.GetMRows()

	hx := Matrix.NullMatrixP(m, 1)

	if this.ThetaParameters.GetNColumns() == Ts.Xs.GetNColumns() {
		for i := 1; i <= Ts.Xs.GetMRows(); i++ {
			xi := Ts.Xs.GetRow(i)

			Thi := Matrix.Product(xi, this.ThetaParameters.Transpose())

			hx.SetValue(i, 1, Thi.GetValue(1, 1)-Ts.Y.GetValue(1, i))

		}
		return hx
	}
	return nil
}

func LinearRegression(alpha complex128, Tolerance complex128, ts *TrainingSet) *Hypothesis {
	f := func(x complex128) complex128 { return x }
	hy := GradientDescent(alpha, Tolerance, ts, f)
	return hy
}

func LogisticRegression(alpha complex128, Tolerance complex128, ts *TrainingSet) *Hypothesis {
	f := func(x complex128) complex128 { return 1 / (1 + cmplx.Exp(-1.0*x)) }
	hy := GradientDescent(alpha, Tolerance, ts, f)
	return hy
}

func GradientDescent(alpha complex128, Tolerance complex128, ts *TrainingSet, f func(x complex128) complex128) *Hypothesis {
	n := ts.Xs.GetNColumns()
	m := ts.Xs.GetMRows()

	//Xsc:=ts.Xs.Copy()

	ts.AddX0() // add  the parametrer x0, with value 1, to all elements of the training set

	t := Matrix.NullMatrixP(1, n+1) // put 0 to the parameters theta
	thetaP := t

	//thetaP:=Matrix.RandomMatrix(1,n+1)  // Generates a random values of parameters theta

	var h1 Hypothesis

	h1.H = f
	h1.ThetaParameters = thetaP

	var Error complex128

	Error = complex(1.0, 0)

	var it = 1

	diferencia, diferenciaT := h1.Parallel_DiffH1Ys(ts)
	jt := Matrix.Product(diferenciaT, diferencia).Scalar(1/complex(2.0*float64(m), 0.0)).GetValue(1, 1)

	println("hola:)", 1/jt)
	//alpha = 1 / jt

	for cmplx.Abs(Error) >= cmplx.Abs(Tolerance) { // Until converges

		ThetaPB := h1.ThetaParameters.Copy() //for Error Calc

		//diff:=h1.DiffH1Ys(ts)
		_, diffT := h1.Parallel_DiffH1Ys(ts) //h(x)-y

		product := Matrix.Product(diffT, ts.Xs) //Sum( (hi(xi)-yi)*xij)  in matrix form

		h1.Sum = product

		alpha_it := alpha / (cmplx.Sqrt(complex(float64(it), 0.0))) // re-calc alpha

		scalar := product.Scalar(-alpha_it / complex(float64(m), 0.0))

		//println("Delta", scalar.ToString())
		ThetaTemp, _ := Matrix.Sum(h1.ThetaParameters, scalar) //Theas=Theas-alfa/m*Sum( (hi(xi)-yi)*xij)  update the parameters

		h1.ThetaParameters = ThetaTemp

		diffError, _ := Matrix.Sustract(ThetaPB, h1.ThetaParameters) //diff between theta's Vector , calc the error

		Error = complex(diffError.FrobeniusNorm(), 0) //Frobenius Norm
		//Error=diffError.InfinityNorm()              //Infinty Norm

		//println("->", h1.ThetaParameters.ToString())
		//println("Error", Error)
		/*if it > 10 {
			break
		}*/
		it++
	}
	h1.M = m
	return &h1
}

func (this *Hypothesis) Evaluate(x *Matrix.Matrix) (complex128, error) {
	x0 := Matrix.NullMatrixP(1, 1)
	x0.SetValue(1, 1, 1)
	x0 = x0.AddColumn(x)
	if x0.GetNColumns() == this.ThetaParameters.GetNColumns() {

		xt := x0.Transpose()

		res := Matrix.Product(this.ThetaParameters, xt)

		return this.H(res.GetValue(1, 1)), nil
	}
	return 0, errors.New(" The number of parameters is not equal to the parameters of the hypotesis")
}

func NormalEquation(ts *TrainingSet) *Hypothesis {
	//     n:=ts.Xs.GetNColumns()
	//     m:=ts.Xs.GetMRows().
	ts.AddX0()
	println(ts.Xs.ToString())
	Xst := ts.Xs.Transpose()
	mult := Matrix.Product(Xst, ts.Xs)

	pinv := mult.PInverse()

	xT := Matrix.Product(pinv, Xst)
	theta := Matrix.Product(xT, ts.Y)

	var h1 Hypothesis

	h1.ThetaParameters = theta
	return &h1

}
