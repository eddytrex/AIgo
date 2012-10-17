package ANN
import(
  "Matrix"
)

type ANN struct{
  Layers []Matrix.matrix
 
}

func  InitANN() *ANN{
  
  return &ANN
}

func (this *ANN)ForwardPropagation(in matrix)(*matrix){
  
}
