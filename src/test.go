package main

import (
  "fmt"
  "Matrix"
//"MachineLearning"
  //"ANN"
  //"math"
  
)


func main(){
 
  
     a,er1:=Matrix.FromFile("xsm.txt")
  
     b,er2:=Matrix.FromFile("ysm.txt")
  
     fmt.Println(er1,"\n",er2)
 
//   make a training set
//      ts:=MachineLearning.MakeTrainingSet(*a,*b)
     
  
//   Linear/Logistic Regression gradiente descent
//   hy:=MachineLearning.LinearRegression(0.0001,0.0001,*ts)
//   hy:=MachineLearning.LinearRegression(0.001,0.00001,*ts)
//   hy:=MachineLearning.LogisticRegression(0.001,0.001,*ts)
  
//  Normal Equation
//   hy:=MachineLearning.NormalEquation(*ts)
//   fmt.Println("--->",hy)
  
  
//valuar
//   t:=Matrix.NullMatrix(1,2)
//   t.SetValue(1,1,1)
//   t.SetValue(1,2,1)
//   ret,_:=hy.Evaluate(&t)
//   fmt.Println("h()=",ret)
     
     
     println(b.ToString())
     Q,R:=a.QRDec()
     
     println(Q.ToString(),"\n",R.ToString())
      
  
     
}
