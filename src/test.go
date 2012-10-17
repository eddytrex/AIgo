package main

import (
  "fmt"
  "Matrix"
  
)


func main(){
  
  a:=Matrix.I(100)
 
  a.SetValue(1,1,2)
  a.SetValue(1,2,1)
  a.SetValue(1,3,3)
  
  a.SetValue(2,1,4)
  a.SetValue(2,2,6)
  a.SetValue(2,3,5)
  
  a.SetValue(3,1,7)
  a.SetValue(3,2,8)
  a.SetValue(3,3,9)
 
   
  c:=a.Copy()
   
   i,_:=a.Inverse()
  //i,_:=a.InverseGauss()
     
   //m:=Matrix.Multiplication(*c,*i)
  m:=Matrix.Product(*c,*i)
  
 
  fmt.Println("->",m.GetValue(1,1))
   
}
