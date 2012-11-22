package main

import (
  "fmt"
  "Matrix"
  "MachineLearning"
  //"ANN"
  //"math"
  
)


func main(){
 
  
  a,er1:=Matrix.FromFile("/home/eddytrex/go/IAgo/src/xsl.txt")
  
  b,er2:=Matrix.FromFile("/home/eddytrex/go/IAgo/src/ysl.txt")
  
  fmt.Println(er1,"\n",er2)
  
  //c:=a.Transpose()
  //fmt.Println(c.ToString(),">")
  
  //fmt.Println(b.ToString())
  ts:=MachineLearning.MakeTrainingSet(*a,*b)
  
  //hy:=MachineLearning.GradientDescent(0.0001,0.00000001,*ts,func (a float64)float64{return a})
  
  //hy:=MachineLearning.LinearRegression(0.0001,0.0000001,*ts)
  hy:=MachineLearning.LogisticRegression(0.001,0.000001,*ts)
  
  fmt.Println("--->",hy)
  
  t:=Matrix.NullMatrix(1,2)
  
  t.SetValue(1,1,7)
  t.SetValue(1,2,7)
  
  ret,_:=hy.Evaluate(&t)
 
  fmt.Println("h()=",ret)
}
