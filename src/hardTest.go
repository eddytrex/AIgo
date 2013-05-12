package main

import (
  "fmt"
  "Matrix"
  "MachineLearning"
  //"ANN"
  //"math"
  
)


func main(){
    
    // Load matrix from a file 
     a,er1:=Matrix.FromFile("ex1Data1X.txt")   
     b,er2:=Matrix.FromFile("ex1Data1Y.txt")
     
     if(er1==nil&&er2==nil){
        ts:=MachineLearning.MakeTrainingSet(*a,*b)
        // alfa, error, Datos
        hy:=MachineLearning.LinearRegression(0.1,0.00001,*ts)
          
        fmt.Println("hy",hy)
        //valuar
        t:=Matrix.NullMatrix(1,1)
        t.SetValue(1,1,5.5277)
        
        ret,_:=hy.Evaluate(&t)
        fmt.Println("h()=",ret)
     }
    
}

