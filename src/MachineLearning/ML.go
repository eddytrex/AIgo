package ML;
import (
  "Matrix"
  "math"
  
)


type TrainingSet struct{
  Xs Matrix.Matrix   //Features    mxn   
  Y   Matrix.Matrix  //Values      mx1
}

type Hypothesis struct{
  ThetaP Matrix.Matrix
}


func (this *TrainingSet)AddX0(){
  m:=this.Xs.GetNRows()
  x0:=Matrix.NullMatrix(m,1)
  
  for i:=1;i<=m;i++{
    x0.SetValue(j,1,1.0)
  }
  
  Xs=x0.AddColumns(this.Xs)
}


func (this *Hypothesis) ApplyHypothesis(Ts TrainingSet) (*Matrix.Matrix){
  
  m:=Ts.Xs.GetNRows()
  
  hx:=Matrix.NullMatrix(m,1)
  
  if(this.ThetaP.GetNColumns()==Ts.Xs.GetNColumns()){
  for i:=1;i<=Ts.Xs.GetNRows();i++{
    xi:=Ts.Xs.GetRow(i);
    
    Thi:=Matrix.Product(this.ThetaP.Traspose(),xi)
    
    hx.SetValue(i,1,Thi.GetValue(1,1))
    
  }
  return hx
  }
}


func (this *Hypothesis) ApplyHSigmoid(Ts TrainingSet)(*Matrix.Matrix)
{
  out:=this.ApplyHypothesis(Ts)
  
  sigmoid:=func(z float64)float64{ return 1/(1+math.Exp(-z))}
  
  out=Matrix.Apply(sigmoid)
  
  return out
}



func GradientDescent(alpha float64,Tolerance float64,ts TrainingSet)( *Hypothesis){
 n:=ts.Xs.GetNColumns()
 m:=ts.Xs.GetNRows()
 
 ts=ts.AddX0()   // add  the parametrer x0, with value 1, to all elements of the training set
 
 //thetaP:=Matrix.NullMatrix(n+1,1)  // put 0 to the parameters theta 
 thetaP:=Matrix.RandomMatrix(n+1,1)  // Generates a random values of parameters theta
 
 var h1 Hypothesis
 
 h1.ThetaP=thetaP           
 
 var Error float64
 
 Error=1.0
 
 for Error<=Tolerance{          // Until converges
   
    diff:=Matrix.Sustract(h1.ApplyHypothesis(ts),ts.Y) //    h(x)-y
 
    p:=Matrix.Product(diff.Traspose(),ts.Xs) //    Sum( (hi(xi)-yi)*xij)  in vectro form 
 
    scalar:=p.Scalar(1/m*alpha)              //    alfa/m*Sum( (hi(xi)-yi)*xij)
 
    ThetaPB:=h1.ThetaP.Copy()                      //for Error Calc
 
    h1.ThetaP=Matrix.Sustract(h1.ThetaP,scalar.Traspose())    //  Theas=Theas-alfa/m*Sum( (hi(xi)-yi)*xij)  update the parameters
 
    diffError:=Matrix.Sustract(ThetaPB,h1.ThetaP)   // diff between theta's Vector , calc the error
 
    Error=diffError.EuclideanNorm()		//Euclidean Norm 
    // Error=diffError.InfinityNorm()             // Infinty Norm
 }

} 


