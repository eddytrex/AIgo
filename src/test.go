package main

import (
	"MachineLearning"
	"Matrix"
	"fmt"
)

func main() {
	X, _ := Matrix.FromFile("MachineLearning/test/X3.txt")
	Y, _ := Matrix.FromFile("MachineLearning/test/Y3.txt")

	fmt.Println("X", X.ToString())
	fmt.Println("Y", Y.ToString())

	TraingSet := MachineLearning.MakeTrainingSet(X, Y)
	Salida := MachineLearning.LinearRegression(complex(1E-6, 0.0), complex(1E-6, 0.0), TraingSet)
	fmt.Println("Theta Parameters:", Salida.ThetaParameters.ToString())

	x1, _ := Salida.Evaluate(X.GetRow(1))
	fmt.Println("test:f(", X.GetRow(1).ToString(), ")=", x1)

	fmt.Println("Mean", TraingSet.Mean().ToString())
	//Salida2 := MachineLearning.NormalEquation(TraingSet)
	//fmt.Println("Theta Parameters:", Salida2.ThetaParameters.ToString())

	//x2, _ := Salida2.Evaluate(X.GetRow(1))
	//fmt.Println("test:f(", X.GetRow(1).ToString(), ")=", x2)

}
