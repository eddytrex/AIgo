package main

import (
	"Matrix"
	//"Search"
	"ANN"
	"fmt"
	"math/cmplx"
)

func main() {
	activation := func(x complex128) complex128 {
		return 1 / (1 + cmplx.Exp(-x))
	}

	activationDerviate := func(x complex128) complex128 {
		return (1 / (1 + cmplx.Exp(-x))) * (1 - (1 / (1 + cmplx.Exp(-x))))
	}

	l := make([]int, 2)
	l[0] = 2
	l[1] = 1

	ann := ANN.CreateANN(2, l, activation, activationDerviate)

	p1 := Matrix.NullMatrix(2, 1)
	p1.SetValue(1, 1, 1.0)
	p1.SetValue(1, 2, 1.0)
	ro1 := Matrix.NullMatrix(1, 1)
	ro1.SetValue(1, 1, 0.0)
	//ro1.SetValue(2, 1, 1.0)

	p2 := Matrix.NullMatrix(2, 1)
	p2.SetValue(1, 1, 1.0)
	p2.SetValue(1, 2, 0.0)
	ro2 := Matrix.NullMatrix(1, 1)
	ro2.SetValue(1, 1, 1.0)
	//ro2.SetValue(2, 1, 0.0)

	p3 := Matrix.NullMatrix(2, 1)
	p3.SetValue(1, 1, 0.0)
	p3.SetValue(1, 2, 1.0)
	ro3 := Matrix.NullMatrix(1, 1)
	ro3.SetValue(1, 1, 1.0)
	//ro3.SetValue(2, 1, 0.0)

	p4 := Matrix.NullMatrix(2, 1)
	p4.SetValue(1, 1, 0.0)
	p4.SetValue(1, 2, 0.0)
	ro4 := Matrix.NullMatrix(1, 1)
	ro4.SetValue(1, 1, 0.0)
	//ro4.SetValue(2, 1, 1.0)

	Inputs := make([]*Matrix.Matrix, 4)
	ROutputs := make([]*Matrix.Matrix, 4)

	Inputs[0] = p1
	Inputs[1] = p2
	Inputs[2] = p3
	Inputs[3] = p4

	ROutputs[0] = ro1
	ROutputs[1] = ro2
	ROutputs[2] = ro3
	ROutputs[3] = ro4

	ann.Train(Inputs, ROutputs, 0.1, 0.3, 0.0001, 100000)

	_, _, Output := ann.ForwardPropagation(Inputs[0])
	fmt.Println(Output.ToString())

	_, _, Output = ann.ForwardPropagation(Inputs[1])
	fmt.Println(Output.ToString())
	_, _, Output = ann.ForwardPropagation(Inputs[2])
	fmt.Println(Output.ToString())
	_, _, Output = ann.ForwardPropagation(Inputs[3])
	fmt.Println(Output.ToString())

}
