package MachineLearning;
import (
  "Matrix"
//  "math"
)

type TrainingSet struct{
  Xs Matrix.Matrix   //Features    mxn   
  Y   Matrix.Matrix  //Values      mx1
}

func MakeTrainingSet(xs Matrix.Matrix,y Matrix.Matrix)(* TrainingSet){
  var out TrainingSet
  
  if(xs.GetMRows()==y.GetMRows()){
  out.Xs=xs
  out.Y=y
  return &out
  }
  return nil
}

/*func (this *TrainingSet)MeanNormalize() {
  
  sum:=Matrix.NullMatrix(1,this.Xs.GetNColumns());
  max:=Matrix.NullMatrix(1,this.Xs.GetNColumns());
  min:=Matrix.NullMatrix(1,this.Xs.GetNColumns());
  
  for i:=1;i<this.Xs.GetMRows();i++{
    xsi:=this.Xs.GetRow(i);
    sum=Matrix.Sum(xsi,sum)
  }
  sum.Scalar(1/(complex128 this.Xs.GetMRows()))  
}

func (this *TrainingSet)sumParameters(i0,i1 int, max, min ,res *Matrix.Matrix ,done chan<-bool){
  di:=i1-i0
  done2:=make(chan bool,THRESHOLD)
  if(di>=THRESHOLD)
  {
    mi:=i0+di/2
       
    res1:=Matrix.NullMatrix(1,this.Xs.GetNColumns());
    res2:=Matrix.NullMatrix(1,this.Xs.GetNColumns());
    
    go this.sumParameters(i0,mi,res1,done2)    
    
    this.sumParameters(mi,i1,res2,done2)
    
    <-done2
    <-done2
    
    res=Matrix.Sum(res1,res2) 
  }else{
    for i:=i0,i<i1;i++{
      xsi:=this.Xs.GetRow(i)
      res=Matrix.Sum(xsi,res)
    }
  }
  done<-true
}*/ 
