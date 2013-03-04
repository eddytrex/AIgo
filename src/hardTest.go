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
     a,er1:=Matrix.FromFile("xsm.txt")   
     b,er2:=Matrix.FromFile("ysm.txt")
  
     fmt.Println(er1,"\n",er2)

     println (a.GetMRows())
     println (a.GetNColumns())
     
     println (a.GetValue(1,1))
   
     a.SetValue(1,1,1.0)
     println (a.GetValue(1,1))
     
     
     
    
}

