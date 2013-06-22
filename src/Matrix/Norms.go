package Matrix

import (
	"math"
	"math/cmplx"
	//"errors"
	"runtime"
)

func (this *Matrix) InfinityNorm() complex128 {
	var out complex128
	out = 0

	out = this.sumColumn(1)
	for i := 2; i < this.m; i++ {
		temp := this.sumColumn(i)
		if cmplx.Abs(temp) > cmplx.Abs(out) {
			out = temp
		}
	}

	return out
}

// func (this *Matrix) FrobeniusNorm()complex128{
//   var out complex128
//   out=0;
//   if (this.m==1||this.n==1){
//
//     for i:=0;i<len(this.A);i++{
//        out=out+this.A[i]*this.A[i]
//     }
//
//     out=math.Sqrt(out)
//   }
//
//   return out
// }

func (this *Matrix) FrobeniusNorm() float64 {
	sum := make(chan complex128, 1)
	this.sumApplyFunction(0, len(this.A), sum, func(a complex128) float64 { return cmplx.Abs(a) * cmplx.Abs(a) })

	v := <-sum

	return math.Sqrt(cmplx.Abs(v))
}

func (this *Matrix) sumApplyFunction(i0, i1 int, pSum chan<- complex128, f func(complex128) float64) {
	sum := complex(0.0, 0.0)
	dx := i1 - i0
	xm := i0 + dx/2
	pSum2 := make(chan complex128, THRESHOLD)
	if dx >= THRESHOLD && runtime.NumGoroutine() < maxGoRoutines {

		go this.sumApplyFunction(i0, xm, pSum2, f)
		this.sumApplyFunction(xm, i1, pSum2, f)
		p1 := <-pSum2
		p2 := <-pSum2
		sum = p1 + p2
	} else {
		for i := i0; i < i1; i++ {
			sum = sum + complex(f(this.A[i]), 0)
		}

	}
	pSum <- sum
}
