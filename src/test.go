package main

import (
  "fmt"
  "Matrix"
  "MachineLearning"
  //"ANN"
  
)


func main(){
 
  
  a,er1:=Matrix.FromFile("/home/eddytrex/go/IAgo/src/xs.txt")
  
  b,er2:=Matrix.FromFile("/home/eddytrex/go/IAgo/src/ys.txt")
  
  fmt.Println(er1,"\n",er2)
  
  c,_:=Matrix.Sum(*a,*b)
  fmt.Println(c.ToString(),">")
  
  //fmt.Println(b.ToString())
  ts:=MachineLearning.MakeTrainingSet(*a,*b)
  
  hy:=MachineLearning.GradientDescent(0.0001,0.00000001,*ts)
  
  fmt.Println("->",hy)
}
