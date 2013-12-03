package Fourier

import (
	"Matrix"
	"errors"
	"math"
	"math/cmplx"
)

// func DFT(this *Matrix.Matrix)(*Matrix.Matrix){
//
//     C:=Matrix.NullMatrixP(this.GetMRows(),1)
//
//     lengthi:=this.GetMRows()
//     lengthf:=(complex128)(lengthi)
//     for i:=1;i<=lengthi;i++{
//
//         floati:=(complex128)(i)-1
//
//         sin:=0.0
//         cos:=0.0
//
//         for j:=1;j<=lengthi;j++{
//             floatj:=(complex128)(j)-1
//             sin=sin+this.GetValue(j,1)*math.Sin((2*math.Pi*floati*floatj)/lengthf);
//             cos=cos+this.GetValue(j,1)*math.Cos((2*math.Pi*floati*floatj)/lengthf);
//         }
//         //|Ci|
//         C.SetValue(i,1,math.Sqrt(sin*sin+cos*cos));
//     }
//
//     return C
// }

func FFT(this *Matrix.Matrix, N int) (*Matrix.Matrix, error) {
	if N > this.GetMRows() {
		return nil, errors.New(" The number of Rows of the matrix (this) must be greater or equal than N ")
	}
	if N&(N-1) == 0 {
		tf := TwiddleFactors(N, false)

		Xr := FFT_ct3(this, N, 1, &tf)

		return Xr, nil
	}
	return nil, errors.New(" The N parameter has to be power of 2")
}

func IFFT(this *Matrix.Matrix, N int) (*Matrix.Matrix, error) {
	if N > this.GetMRows() {
		return nil, errors.New(" The number of Rows of the matrix (this) must be greater or equal than N ")
	}
	if N&(N-1) == 0 {
		tf := TwiddleFactors(N, true)

		Xr := FFT_ct(this, N, 1, &tf)

		Xr = Xr.Scalar(complex(float64(1)/float64(N), 0))
		return Xr, nil
	}
	return nil, errors.New(" The N parameter has to be power of 2")
}

func FFT_ct(this *Matrix.Matrix, N, skip int, tf *[]complex128) *Matrix.Matrix {

	Xr := Matrix.NullMatrixP(N, this.GetNColumns())
	RowTemp := Matrix.NullMatrixP(1, this.GetNColumns())

	FFT_aux(this, Xr, RowTemp, N, skip, tf)
	return Xr
}

func FFT_aux(this, xr, RowTemp *Matrix.Matrix, N, skip int, tf *[]complex128) {

	if N == 1 {
		xr.SetRow(1, this.GetReferenceRow(1))
		return
	}

	FFT_aux(this, xr, RowTemp, N/2, skip*2, tf)
	FFT_aux(this.MatrixWithoutFirstRows(skip), xr.MatrixWithoutFirstRows(N/2), RowTemp, N/2, skip*2, tf)

	for k := 0; k < N/2; k++ {

		xr.ScalarRowIntoRowMatrix(RowTemp, k+1+N/2, (*tf)[k*skip])

		sr, rr, _ := Matrix.Sum_Sustract(xr.GetReferenceRow(k+1), RowTemp)

		xr.SetRow(k+1, sr)
		xr.SetRow(k+1+N/2, rr)

	}
}

func FFT_ct2(this *Matrix.Matrix, N, skip int, tf *[]complex128) *Matrix.Matrix {

	Xr := Matrix.NullMatrixP(N, this.GetNColumns())
	Scratch := Matrix.NullMatrixP(N, this.GetNColumns())

	var E, D, Xp, Xstart *Matrix.Matrix
	var evenIteration bool

	if N%2 == 0 {
		evenIteration = true
	} else {
		evenIteration = false
	}

	if N == 1 {
		Xr.SetRow(1, this.GetReferenceRow(1))
	}

	E = this

	for n := 1; n < N; n *= 2 {

		if evenIteration {
			Xstart = Scratch
		} else {
			Xstart = Xr
		}

		skip := N / (2 * n)
		Xp = Xstart

		for k := 0; k != n; k++ {
			for m := 0; m != skip; m++ {
				D = E.MatrixWithoutFirstRows(skip)
				D.ScalarRow(1, (*tf)[skip*k])

				sr, rr, _ := Matrix.Sum_Sustract(E.GetReferenceRow(1), D.GetReferenceRow(1))

				Xp.SetRow(1, sr)
				Xp.SetRow(N/2+1, rr)

				Xp = Xp.MatrixWithoutFirstRows(1)
				E = E.MatrixWithoutFirstRows(1)
			}
			E = E.MatrixWithoutFirstRows(skip)
		}
		E = Xstart
		evenIteration = !evenIteration
	}
	return Scratch
}

func FFT_ct3(this *Matrix.Matrix, N, skip int, tf *[]complex128) *Matrix.Matrix {

	Xr := Matrix.NullMatrixP(N, this.GetNColumns())
	Scratch := Matrix.NullMatrixP(N, this.GetNColumns())

	var E, D, Xp, Xstart *Matrix.Matrix
	var evenIteration bool

	if N%2 == 0 {
		evenIteration = true
	} else {
		evenIteration = false
	}

	if N == 1 {
		Xr.SetRow(1, this.GetReferenceRow(1))
	}

	E = this

	for n := 1; n < N; n *= 2 {

		if evenIteration {
			Xstart = Scratch
		} else {
			Xstart = Xr
		}

		skip := N / (2 * n)
		Xp = Xstart

		for k := 0; k != n; k++ {

			var Aux = func(m0, m1 int, Xp, E, D *Matrix.Matrix) {

				println("-", m0)
				Xp = Xp.MatrixWithoutFirstRows(m0)
				E = E.MatrixWithoutFirstRows(m0)
				//D = E.MatrixWithoutFirstRows(skip)

				for m := m0; m < m1; m++ {
					D = E.MatrixWithoutFirstRows(skip)
					D.ScalarRow(1, (*tf)[skip*k])

					sr, rr, _ := Matrix.Sum_Sustract(E.GetReferenceRow(1), D.GetReferenceRow(1))

					Xp.SetRow(1, sr)
					Xp.SetRow(N/2+1, rr)

					Xp = Xp.MatrixWithoutFirstRows(1)

					println("E", E.ToString())
					E = E.MatrixWithoutFirstRows(1)

				}

			}

			mm := skip / 2
			m0 := 0
			//m1 := skip

			go Aux(m0, mm, Xp, E, D)
			//println("->E", E.ToString(), ">XP", Xp.ToString())
			//go Aux(mm, m1, Xp, E, D)

			//for m := 0; m != skip; m++ {
			//	D = E.MatrixWithoutFirstRows(skip)
			//	D.ScalarRow(1, (*tf)[skip*k])

			//	sr, rr, _ := Matrix.Sum_Sustract(E.GetReferenceRow(1), D.GetReferenceRow(1))

			//	Xp.SetRow(1, sr)
			//	Xp.SetRow(N/2+1, rr)

			//	Xp = Xp.MatrixWithoutFirstRows(1)
			//	E = E.MatrixWithoutFirstRows(1)
			//}
			E = E.MatrixWithoutFirstRows(skip)

		}
		E = Xstart
		evenIteration = !evenIteration
	}
	return Scratch
}

func TwiddleFactors(N int, ifft bool) []complex128 {
	out := make([]complex128, N)

	inv := 1
	if ifft {
		inv = -1
	}
	for i := 0; i < N/2; i++ {
		out[i] = cmplx.Rect(1, math.Pi*float64(-2*i*inv)/float64(N))
	}
	return out
}
