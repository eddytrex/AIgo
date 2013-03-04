package main

import (
  "fmt"
  "Matrix"
  "MachineLearning"
  //"ANN"
  //"math"
  
)


func main(){
 
  
     a,er1:=Matrix.FromFile("xsm.txt")
  
     b,er2:=Matrix.FromFile("ysm.txt")
  
     fmt.Println(er1,"\n",er2)
 
//   make a training set
     ts:=MachineLearning.MakeTrainingSet(*a,*b)
     
  
//   Linear/Logistic Regression gradiente descent
//   hy:=MachineLearning.LinearRegression(0.0001,0.0001,*ts)
//   hy:=MachineLearning.LinearRegression(0.001,0.00001,*ts)
//   hy:=MachineLearning.LogisticRegression(0.001,0.001,*ts)
  
//  Normal Equation
//   fmt.Println("->",a.PInverse().ToString())   
//   hy:=MachineLearning.NormalEquation(*ts)
//   fmt.Println("--->",hy)
//   
//   
// //valuar
//   t:=Matrix.NullMatrix(1,2)
//   t.SetValue(1,1,1)
//   t.SetValue(1,2,1)
//   ret,_:=hy.Evaluate(&t)
//   fmt.Println("h()=",ret)
     
     ei:=ts.Xs.EigenValues(0.000001)
     
     ev:=ts.Xs.EigenVector(ei.GetValue(1,1))
     
     println (ei.ToString(),"\n",ev.ToString())
//      println(b.ToString())
//      Q,R:=a.QRDec()
//      
//      pi:=a.PInverse();
//      
//      println(Q.ToString(),"\n",R.ToString())
//      
//      println ("PI",pi.ToString())
      
  
     
}
