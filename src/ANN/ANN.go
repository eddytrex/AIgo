package ANN
import(
  "Matrix"
)

type ANN struct{
  Layers [] Matrix.Matrix
  Inputs int
  Activation func(complex128)complex128
}


func CreateANN(NeuronsByLayer []int,Inputs int,Act func(complex128)complex128) ANN{
  var out ANN
  out.Layers=make([]Matrix.Matrix,len(NeuronsByLayer),len(NeuronsByLayer))
  
  out.Inputs=Inputs
  out.Activation=Act
  
  m:=Inputs
  for i:=0;i<(len(NeuronsByLayer));i++{
    
    n:=NeuronsByLayer[i]
    temp:=Matrix.NullMatrix(m+1,n)
    out.Layers[i]=temp
    m=n
    
  }
  return out
}



func (this *ANN)ForwardPropagation(In Matrix.Matrix)(Output *Matrix.Matrix){
  if(In.GetNRows()==1&&In.GetNRows()==this.Inputs){
	
	sTemp:=In.Copy()
	
	sTemp=sTemp.AddColumn(*Matrix.I(1))//Add  a new row for a Gain Weight 
	
	for i:=0;i<len(this.Layers);i++{
	  
	  sTemp=Matrix.Product(*sTemp,(this.Layers[i]))
	  
	  if(i<len(this.Layers)-1){
	  sTemp=sTemp.AddColumn(*Matrix.I(1))	//Add  a new row for a Gain Weight 
	  }
	 
	  sTemp=sTemp.Apply(this.Activation) //apply the functions of activation
	}
	Output=sTemp
	return
  }	  
  return nil
}


