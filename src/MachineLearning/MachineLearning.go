package MachineLearning;
import (
  "Matrix"
  "math"
  "errors" 
)

type Hypothesis struct{
  ThetaP Matrix.Matrix
  M int
  Sum Matrix.Matrix
  Alpah int	
  H func (float64)float64
}


func (this *TrainingSet)AddX0(){
  m:=this.Xs.GetMRows()
  x0:=Matrix.NullMatrix(m,1)
  
  for i:=1;i<=m;i++{
    x0.SetValue(i,1,1.0)
  }
  
  this.Xs=*x0.AddColumn(this.Xs)
}


func (this *Hypothesis) ApplyHypothesisToTrainingSet(Ts TrainingSet) (*Matrix.Matrix){
  
  m:=Ts.Xs.GetMRows()
  
  hx:=Matrix.NullMatrix(m,1)
  
  if(this.ThetaP.GetNColumns()==Ts.Xs.GetNColumns()){
  for i:=1;i<=Ts.Xs.GetMRows();i++{
    xi:=Ts.Xs.GetRow(i);
    
    Thi:=Matrix.Product(*xi,*this.ThetaP.Transpose())
    
    hx.SetValue(i,1,Thi.GetValue(1,1))
    
  }
  return &hx
  }
  return nil
}

func (this *Hypothesis) Parallel_DiffH1Ys(Ts TrainingSet) (*Matrix.Matrix){
  m:=Ts.Xs.GetMRows()
  hx:=Matrix.NullMatrix(m,1)
  
  if(this.ThetaP.GetNColumns()==Ts.Xs.GetNColumns()){
      done:=make(chan bool);
      go this.part_DiffH1Ys(1,m,Ts,hx,done)
      <-done
  }
 return &hx 
}

const THRESHOLD=100
func (this *Hypothesis) part_DiffH1Ys(i0,i1 int,Ts TrainingSet,Ret Matrix.Matrix,done chan<-bool){
  di:=i1-i0
  done2:=make(chan bool,THRESHOLD);

  if(di>=THRESHOLD){
    mi:=i0+di/2
    go this.part_DiffH1Ys(i0,mi,Ts,Ret,done2)
    this.part_DiffH1Ys(mi,i1,Ts,Ret,done2)
    <-done2
    <-done2
  }else{
      for i:=i0;i<i1;i++{
	xi:=Ts.Xs.GetRow(i)
	
	Thi:=Matrix.Product(*xi,*this.ThetaP.Transpose())
	
	Ret.SetValue(i,1,this.H(Thi.GetValue(1,1))-Ts.Y.GetValue(1,i))
      }
    }
    done<-true
}


func (this *Hypothesis) DiffH1Ys(Ts TrainingSet) (*Matrix.Matrix){
  
  m:=Ts.Xs.GetMRows()
  
  hx:=Matrix.NullMatrix(m,1)
  
  if(this.ThetaP.GetNColumns()==Ts.Xs.GetNColumns()){
  for i:=1;i<=Ts.Xs.GetMRows();i++{
    xi:=Ts.Xs.GetRow(i);
    
    Thi:=Matrix.Product(*xi,*this.ThetaP.Transpose())
    
    hx.SetValue(i,1,Thi.GetValue(1,1)-Ts.Y.GetValue(1,i))
    
  }
  return &hx
  }
  return nil
}


func LinearRegression(alpha float64,Tolerance float64,ts TrainingSet)( *Hypothesis){
  f:=func (x float64)float64{return x}
  hy:=GradientDescent(alpha,Tolerance,ts,f)
  return hy
}

func LogisticRegression(alpha float64,Tolerance float64,ts TrainingSet)( *Hypothesis){
  f:=func (x float64)float64{return 1/(1+math.Exp(-x))}
  hy:=GradientDescent(alpha,Tolerance,ts,f)
  return hy
}


func GradientDescent(alpha float64,Tolerance float64,ts TrainingSet,f func (x float64)float64)( *Hypothesis){
 n:=ts.Xs.GetNColumns()
 m:=ts.Xs.GetMRows()
 
 //Xsc:=ts.Xs.Copy()
 
 ts.AddX0()   // add  the parametrer x0, with value 1, to all elements of the training set
  
 t:=Matrix.NullMatrix(1,n+1) // put 0 to the parameters theta 
 thetaP:=&t  
 
 //thetaP:=Matrix.RandomMatrix(1,n+1)  // Generates a random values of parameters theta
 
 var h1 Hypothesis
 
 h1.H=f
 h1.ThetaP=*thetaP           
 
 var Error float64
 
 Error=1.0
 
 for Error>Tolerance{                        // Until converges
    
    ThetaPB:=h1.ThetaP.Copy()                //for Error Calc
    
    //diff,_:=Matrix.Sustract(*h1.ApplyHypothesisToTrainingSet(ts),ts.Y) //    h(x)-y
    //diff:=h1.DiffH1Ys(ts)
    diff:=h1.Parallel_DiffH1Ys(ts)                                            //h(x)-y
    
    diffT:=diff.Transpose();
    
    p:=Matrix.Product(*diffT,ts.Xs)                       //Sum( (hi(xi)-yi)*xij)  in vectro form 
    
    h1.Sum=*p
    
    scalar:=p.Scalar((-1)*alpha/float64(m))              //alfa/m*Sum( (hi(xi)-yi)*xij)
    
    ThetaTemp,_:=Matrix.Sum(h1.ThetaP,*scalar)           //Theas=Theas-alfa/m*Sum( (hi(xi)-yi)*xij)  update the parameters   
    
    h1.ThetaP=*ThetaTemp
 
    diffError,_:=Matrix.Sustract(*ThetaPB,h1.ThetaP)      //diff between theta's Vector , calc the error
    
    Error=diffError.EuclideanNorm()		         //Euclidean Norm 
    //Error=diffError.InfinityNorm()                     //Infinty Norm  
  
 }
 h1.M=m

 return &h1
}

func (this *Hypothesis) Evaluate(x *Matrix.Matrix) (float64,error){
  x0:=Matrix.NullMatrix(1,1)
  x0.SetValue(1,1,1);
  x0=*x0.AddColumn(*x)
  if(x0.GetNColumns()==this.ThetaP.GetNColumns()){
    
      xt:=x0.Transpose()
      
      res:=Matrix.Product(this.ThetaP,*xt);      
      
      return this.H(res.GetValue(1,1)),nil
  }  
  return 0,errors.New(" the number of parameters is not equal to the parameters of the hypotesis")
}

 

