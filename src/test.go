package main

import (
	//"MachineLearning"
	"Fourier"
	"Matrix"
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("-<", runtime.NumGoroutine(), "-<")

	X, _ := Matrix.FromFile("Fourier/test/BMFFT.txt")

	//Xn := Fourier.Order(X, 8)
	Xn, _ := Fourier.FFT(X, 32768)
	fmt.Println("X", Xn.GetMRows())

	//X, _ := Matrix.FromFile("Fourier/test/BMFFT.txt")
	//Xn := Fourier.Order(X, 32768)
	//fmt.Println("X", Xn.GetMRows())

	//fmt.Println("X", Xi.ToString())

	//fmt.Println("E:=", Matrix.AlmostEqual(X, Xi))
	//X, _ := Matrix.FromFile("MachineLearning/test/X3.txt")
	//Y, _ := Matrix.FromFile("MachineLearning/test/Y3.txt")

	//fmt.Println("X", X.ToString())
	//fmt.Println("Y", Y.ToString())

	//TraingSet := MachineLearning.MakeTrainingSet(X, Y)
	//Salida := MachineLearning.LinearRegression(complex(1E-6, 0.0), complex(1E-6, 0.0), TraingSet)
	//fmt.Println("Theta Parameters:", Salida.ThetaParameters.ToString())

	//x1, _ := Salida.Evaluate(X.GetRow(1))
	//fmt.Println("test:f(", X.GetRow(1).ToString(), ")=", x1)

	//fmt.Println("Mean", TraingSet.Mean().ToString())
	//Salida2 := MachineLearning.NormalEquation(TraingSet)
	//fmt.Println("Theta Parameters:", Salida2.ThetaParameters.ToString())

	//x2, _ := Salida2.Evaluate(X.GetRow(1))
	//fmt.Println("test:f(", X.GetRow(1).ToString(), ")=", x2)

}
