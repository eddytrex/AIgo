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
		Xr := FFT_ct(this, N, 1, false, tf)

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
		Xr := FFT_ct(this, N, 1, true, tf)
		Xr = Xr.Scalar(complex(float64(1)/float64(N), 0))
		return Xr, nil
	}
	return nil, errors.New(" The N parameter has to be power of 2")
}

func FFT_ct(this *Matrix.Matrix, N, skip int, ifft bool, tf []complex128) *Matrix.Matrix {

	if N == 1 {
		return this.GetRow(1)
	}

	xskip := this.SlideRows(skip)

	Ar := FFT_ct(this, N/2, skip*2, ifft, tf)
	Br := FFT_ct(xskip, N/2, skip*2, ifft, tf)

	Xr := Matrix.NullMatrixP(N, this.GetNColumns())

	for k := 0; k < N/2; k++ {

		Br.ScalarRow(k+1, tf[k*skip])

		sr, _ := Matrix.Sum(Ar.GetRow(k+1), Br.GetRow(k+1))
		Xr.SetRow(k+1, sr)
		rr, _ := Matrix.Sustract(Ar.GetRow(k+1), Br.GetRow(k+1))
		Xr.SetRow(k+1+N/2, rr)

	}
	return Xr
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
