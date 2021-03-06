package ANN

import (
	"Matrix"

	"fmt"

	"os"

	//"math/cmplx"

	//"math"
)

const namesufix = "_layer.csv"
const nameprefix = "f"

type ANN struct {
	Weights          []*Matrix.Matrix
	BestWeightsFound []*Matrix.Matrix

	LearningRates []*Matrix.Matrix

	Δ  []*Matrix.Matrix
	Δ1 []*Matrix.Matrix

	ð []*Matrix.Matrix

	AcumatedError     *Matrix.Matrix
	MinimumErrorFound *Matrix.Matrix

	AcumatedError1 *Matrix.Matrix

	α float64
	η float64

	Inputs  int
	Outputs int

	ActivationLayer         func(*Matrix.Matrix) *Matrix.Matrix
	DarivateActivationLayer func(*Matrix.Matrix) *Matrix.Matrix

	Activation func(complex128) complex128
	Derivate   func(complex128) complex128

	CostFunction         func(*Matrix.Matrix, *Matrix.Matrix) *Matrix.Matrix
	DerviateCostFunction func(*Matrix.Matrix, *Matrix.Matrix) *Matrix.Matrix

	PathWeightsInCSV string
}

func CreateANN(Inputs int, NeuronsByLayer []int, Act func(*Matrix.Matrix) *Matrix.Matrix, Derivate func(*Matrix.Matrix) *Matrix.Matrix, Cost func(*Matrix.Matrix, *Matrix.Matrix) *Matrix.Matrix, DCost func(*Matrix.Matrix, *Matrix.Matrix) *Matrix.Matrix, path string) ANN {

	var out ANN

	out.Weights = make([]*Matrix.Matrix, len(NeuronsByLayer), len(NeuronsByLayer))
	out.BestWeightsFound = make([]*Matrix.Matrix, len(NeuronsByLayer), len(NeuronsByLayer))
	out.LearningRates = make([]*Matrix.Matrix, len(NeuronsByLayer), len(NeuronsByLayer))

	out.Δ = make([]*Matrix.Matrix, len(NeuronsByLayer), len(NeuronsByLayer))
	out.Δ1 = make([]*Matrix.Matrix, len(NeuronsByLayer), len(NeuronsByLayer))

	out.ð = make([]*Matrix.Matrix, len(NeuronsByLayer)+1, len(NeuronsByLayer)+1)

	out.Inputs = Inputs
	out.Outputs = NeuronsByLayer[len(NeuronsByLayer)-1]

	out.ActivationLayer = Act
	out.DarivateActivationLayer = Derivate

	out.CostFunction = Cost
	out.DerviateCostFunction = DCost
	out.PathWeightsInCSV = path
	m := Inputs
	for i := 0; i < (len(NeuronsByLayer)); i++ {

		n := NeuronsByLayer[i]

		// one row extra for Bias weights, we need to change to random values for this matrixes
		//temp := Matrix.RandomRealMatrix(m+1, n)

		out.Weights[i] = Matrix.RandomRealMatrix(m+1, n, 1.2)
		out.BestWeightsFound[i] = Matrix.NullMatrixP(m+1, n)
		out.LearningRates[i] = Matrix.FixValueMatrix(m+1, n, 0.0001)

		//tempdelta := Matrix.NullMatrix(m+1, n)
		out.ð[i] = Matrix.NullMatrix(m+1, n)

		out.Δ[i] = Matrix.NullMatrixP(m+1, n)
		out.Δ1[i] = Matrix.NullMatrixP(m+1, n)
		m = n

	}

	out.AcumatedError = Matrix.NullMatrixP(m, 1)
	out.MinimumErrorFound = Matrix.NullMatrixP(m, 1)
	out.AcumatedError1 = Matrix.NullMatrixP(m, 1)
	return out
}

//TODO the activation function and his Derviate has to be more general.. to implemente soft-max for example
func (this *ANN) ForwardPropagation(In *Matrix.Matrix) (As, AsDerviate *([]*Matrix.Matrix), Output *Matrix.Matrix) {
	if In.GetMRows() == this.Inputs && In.GetNColumns() == 1 {
		As1 := make([]*Matrix.Matrix, len(this.Weights)+1, len(this.Weights)+1)
		AsDerviate1 := make([]*Matrix.Matrix, len(this.Weights)+1, len(this.Weights)+1)

		As := &As1
		AsDerviate = &AsDerviate1

		sTemp := In.Transpose()

		//Add  a new column for a Bias Weight
		sTemp = sTemp.AddColumn(Matrix.I(1))

		holeInput := sTemp.Copy()
		As1[0] = sTemp.Transpose()

		//Derivate
		//sutract, _ := Matrix.Sustract(Matrix.OnesMatrix(As1[0].GetMRows(), 1), As1[0])
		//derivate := Matrix.DotMultiplication(As1[0], sutract)

		//derivate := holeInput.Apply(this.Derivate)
		derivate := this.DarivateActivationLayer(holeInput)

		AsDerviate1[0] = derivate.Transpose()

		for i := 0; i < len(this.Weights); i++ {
			sTemp = Matrix.Product(sTemp, (this.Weights[i]))

			//apply the activation functions
			holeInput := sTemp.Copy()
			sTemp = this.ActivationLayer(sTemp)

			//sTemp = sTemp.Apply(this.Activation)

			//Add  a new column for a Bias Weight
			sTemp = sTemp.AddColumn(Matrix.I(1))
			(*As)[i+1] = sTemp.Transpose()

			//Derivate
			//sutract, _ := Matrix.Sustract(Matrix.OnesMatrix((*As)[i+1].GetMRows(), 1), (*As)[i+1])
			//derivate := Matrix.DotMultiplication((*As)[i+1], sutract)

			derivate := this.DarivateActivationLayer(holeInput)
			//derivate := holeInput.Apply(this.Derivate)

			(*AsDerviate)[i+1] = derivate.Transpose()

		}
		Asf := sTemp.Copy()

		//Asf = Asf.AddColumn(Matrix.I(1))
		(*As)[len(As1)-1] = Asf.Transpose()
		Output = sTemp.Transpose().MatrixWithoutLastRow()
		return As, AsDerviate, Output
	}
	return nil, nil, nil
}

func (this *ANN) BackPropagation(As, AsDerviate *[](*Matrix.Matrix), ForwardOutput *Matrix.Matrix, Y *Matrix.Matrix, flen float64) {
	ð := this.DerviateCostFunction(ForwardOutput, Y)

	this.ð[len(this.ð)-1] = ð

	this.AcumatedError, _ = Matrix.Sum(this.CostFunction(ForwardOutput, Y), this.AcumatedError)

	for i := len(this.Weights) - 1; i >= 0; i-- {
		A := (*As)[i]
		Aderviate := (*AsDerviate)[i]

		var ðtemp *Matrix.Matrix
		if i == len(this.Weights)-1 {
			ðtemp = this.ð[i+1].Transpose()
		} else {
			ðtemp = this.ð[i+1].MatrixWithoutLastRow().Transpose()
		}

		//Calc ð

		//fmt.Println("ð(i+1)", this.ð[i+1].ToString())
		//fmt.Println("W(i)", this.Weights[i].ToString())

		Product := Matrix.Product(this.Weights[i], ðtemp.Transpose())
		//fmt.Println("Product", i, " ", Product.ToString())

		this.ð[i] = Matrix.DotMultiplication(Product, Aderviate.AddRowsToDown(Matrix.I(1)))

		//Calc of Derivate with respect to the Weights

		//ðtemp:= i==len(this.Weights) - 1? this.ð[i+1].Transpose() : this.ð[i+1].MatrixWithoutLastRow().Transpose()
		Dw := Matrix.Product(A, ðtemp)

		this.Δ[i], _ = Matrix.Sum(this.Δ[i], Dw)
	}

	return
}

func (this *ANN) CleanΔ() {
	for i := 0; i < len(this.Weights); i++ {

		this.Δ1[i] = this.Δ[i].Copy()
		this.Δ[i] = this.Δ[i].ZeroMatrix()

	}
}
func (this *ANN) UpdateWeights(length float64, changeBeasWeights bool) {

	for i := 0; i < len(this.Weights); i++ {

		if changeBeasWeights {
			this.BestWeightsFound[i] = this.Weights[i]
		}

		D, _ := Matrix.Sum(this.Δ[i].Scalar(complex(-this.η, 0)), this.Δ1[i].Scalar(complex(this.α, 0)))

		this.Weights[i], _ = Matrix.Sum(this.Weights[i], D)

	}
}

func (this *ANN) RevertWeithgs() {
	for i := 0; i < len(this.Weights); i++ {
		this.Weights[i] = this.BestWeightsFound[i].Copy()
	}
}

func (this *ANN) Train(Patters []*Matrix.Matrix, Results []*Matrix.Matrix, α, η, Tolerance float64, MaxIteration int) float64 {
	if len(Patters) != len(Results) {
		return 1.0
	}

	if this.PathWeightsInCSV != "" {
		this.NewFiles()
	}

	this.α = α
	this.η = η

	Error := 1.0
	//LastError := 0.0

	flen := float64(len(Patters))

	for iteration := 1; iteration <= MaxIteration && Error > Tolerance; iteration++ {

		this.AcumatedError1 = this.AcumatedError.Copy()
		this.AcumatedError = this.AcumatedError.ZeroMatrix()

		for i := 0; i < len(Patters); i++ {

			if Patters[i].GetMRows() != this.Inputs || Patters[i].GetNColumns() != 1 && Results[i].GetMRows() != this.Outputs || Results[i].GetNColumns() != 1 {
				return 1.0
			}

			As, AsDerviate, Output := this.ForwardPropagation(Patters[i])
			this.BackPropagation(As, AsDerviate, Output, Results[i], flen)

		}

		Error = this.AcumatedError.TaxicabNorm()

		flag := false
		if iteration == 1 {
			this.MinimumErrorFound = this.AcumatedError.Copy()
		} else {
			ActualError := this.AcumatedError.TaxicabNorm()
			BError := this.MinimumErrorFound.TaxicabNorm()

			if ActualError < BError {
				this.MinimumErrorFound = this.AcumatedError.Copy()
				flag = true
			}
		}

		this.UpdateWeights(flen, flag)

		this.CleanΔ()

		if this.PathWeightsInCSV != "" {
			for l := 0; l < len(this.Weights); l++ {
				this.SaveWeightsInPlainText(l, iteration)
			}
		}

		fmt.Println("i:", iteration, Error)
	}
	return Error
}

func (this *ANN) NewFiles() {
	for i := 0; i < len(this.Weights); i++ {
		err := os.Remove(this.PathWeightsInCSV + "/" + nameprefix + fmt.Sprint(i) + namesufix)
		fmt.Println(err)
	}

}

func (this *ANN) SaveWeightsInPlainText(layerN int, ite int) {
	layer := this.Weights[layerN]
	m := layer.GetMRows()
	n := layer.GetNColumns()

	file, err := os.OpenFile(this.PathWeightsInCSV+"/"+nameprefix+fmt.Sprint(layerN)+namesufix, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		file, _ = os.Create(this.PathWeightsInCSV + "/" + nameprefix + fmt.Sprint(layerN) + namesufix)
	}

	defer file.Close()

	weigth := ""
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {

				weigth = "" + fmt.Sprintf("%d", ite) + "," + fmt.Sprintf("%.6f", real(layer.GetValue(i, j)))
			} else {
				weigth = weigth + " , " + fmt.Sprintf("%.6f", real(layer.GetValue(i, j)))
			}
		}
	}
	weigth = weigth + "\n"
	if _, err = file.WriteString(weigth); err != nil {
		panic(err)
	}
}
