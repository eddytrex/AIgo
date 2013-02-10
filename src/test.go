package main

import (
  "fmt"
  "Matrix"
  "MachineLearning"
  //"ANN"
  //"math"
  
)


func main(){
 
  
  a,er1:=Matrix.FromFile("xs.txt")
  
  b,er2:=Matrix.FromFile("ys.txt")
  
  fmt.Println(er1,"\n",er2)
  
  //c:=a.Transpose()
  //fmt.Println(c.ToString(),">")
  
  //fmt.Println(b.ToString())
  ts:=MachineLearning.MakeTrainingSet(*a,*b)
  
  //hy:=MachineLearning.GradientDescent(0.0001,0.00000001,*ts,func (a float64)float64{return a})
  
  hy:=MachineLearning.LinearRegression(0.0001,0.0001,*ts)
  //hy:=MachineLearning.LinearRegression(0.1,0.001,*ts)
  //hy:=MachineLearning.LogisticRegression(4,0.0001,*ts)
  
  fmt.Println("--->",hy)
  
  t:=Matrix.NullMatrix(1,2)
  
  t.SetValue(1,1,51)
  t.SetValue(1,2,51)
  
  ret,_:=hy.Evaluate(&t)
  
  
//   temp:=Matrix.NullMatrix(2,2)
//   temp2:=Matrix.NullMatrix(2,1)
//   
//   
//   temp2.SetValue(1,1,1)
//   temp2.SetValue(2,1,1)
 
  
  fmt.Println("h()=",ret)
  
//   temp.SetColumn(1,temp2)
//   
//   temp2.SetValue(1,1,3)
//   temp2.SetValue(2,1,4)
//   
//   temp.SetColumn(2,temp2)
//   
//   temp3:=temp2.UnitVector()
  
//   fmt.Println("---",temp.ToString())
//   fmt.Println("---",temp3.ToString())
}
