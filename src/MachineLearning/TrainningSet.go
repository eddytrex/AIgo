package MachineLearning

import (
	"Matrix"
	"math/cmplx"
)

type TrainingSet struct {
	Xs                 *Matrix.Matrix //Features    mxn
	Y                  *Matrix.Matrix //Values      mx1
	mean               *Matrix.Matrix
	standard_deviation *Matrix.Matrix
}

func MakeTrainingSet(xs *Matrix.Matrix, y *Matrix.Matrix) *TrainingSet {
	var out TrainingSet

	if xs.GetMRows() == y.GetMRows() {
		out.Xs = xs
		out.Y = y
		return &out
	}
	return nil
}

func (this *TrainingSet) MeanNormalize() *TrainingSet {

	variance, sustract, mean := this.Variance()
	StandardDeviation := variance.Apply(func(a complex128) complex128 { return cmplx.Sqrt(a) })

	for i := 1; i <= sustract.GetMRows(); i++ {
		Normalize := Matrix.DotDivision(sustract.GetRow(i), StandardDeviation)
		sustract.SetRow(i, Normalize)
	}

	out := MakeTrainingSet(sustract, this.Y)
	out.mean = mean
	out.standard_deviation = StandardDeviation
	return out
}

func (this *TrainingSet) Mean() *Matrix.Matrix {
	sum := Matrix.NullMatrixP(1, this.Xs.GetNColumns())

	done := make(chan bool)

	go this.sumParameters(1, this.Xs.GetMRows(), &sum, done)

	<-done

	return sum.Scalar(1.0 / (complex(float64(this.Xs.GetMRows()), 0.0)))
}

func (this *TrainingSet) sumParameters(i0, i1 int, Res **Matrix.Matrix, done chan<- bool) {
	di := i1 - i0

	if di >= THRESHOLD {
		done2 := make(chan bool, THRESHOLD)
		mi := i0 + di/2

		res1 := Matrix.NullMatrixP(1, this.Xs.GetNColumns())
		res2 := Matrix.NullMatrixP(1, this.Xs.GetNColumns())

		go this.sumParameters(i0, mi, &res1, done2)

		go this.sumParameters(mi, i1, &res2, done2)

		<-done2
		<-done2

		SP, _ := Matrix.Sum(res1, res2)

		*Res = SP

	} else {
		for i := i0; i <= i1; i++ {

			xsi := this.Xs.GetRow(i)
			SP, _ := Matrix.Sum(*Res, xsi)
			*Res = SP
		}
	}

	done <- true

}

func (this *TrainingSet) Variance() (V, Sustract, Mean *Matrix.Matrix) {
	mean := this.Mean()

	sum := Matrix.NullMatrixP(1, this.Xs.GetNColumns())
	sustract := Matrix.NullMatrixP(this.Xs.GetMRows(), this.Xs.GetNColumns())

	done := make(chan bool)
	this.Variance_sum(1, this.Xs.GetMRows(), mean, &sum, sustract, done)
	<-done

	return sum.Scalar(1 / (complex(float64(this.Xs.GetMRows()), 0) - 1.0)), sustract, mean
}

func (this *TrainingSet) Variance_sum(i0, i1 int, mean *Matrix.Matrix, res **Matrix.Matrix, sustract *Matrix.Matrix, done chan<- bool) {
	di := i1 - i0

	if di >= THRESHOLD {
		mi := i0 + di/2
		done2 := make(chan bool, THRESHOLD)

		res1 := Matrix.NullMatrixP(1, this.Xs.GetNColumns())
		res2 := Matrix.NullMatrixP(1, this.Xs.GetNColumns())

		go this.Variance_sum(i0, mi, mean, &res1, sustract, done2)
		go this.Variance_sum(mi, i1, mean, &res1, sustract, done2)

		<-done2
		<-done2

		SP, _ := Matrix.Sum(res1, res2)
		*res = SP

	} else {
		for i := i0; i <= i1; i++ {
			xsi := this.Xs.GetRow(i)
			Sustract, _ := Matrix.Sustract(mean, xsi)
			Square := Matrix.DotMultiplication(Sustract, Sustract)

			sustract.SetRow(i, Sustract)

			SP, _ := Matrix.Sum(Square, *res)
			*res = SP
		}
	}
	done <- true
}
