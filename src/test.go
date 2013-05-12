package main

import (
  "fmt"
  "Matrix"
  //"MachineLearning"
  //"Fourier"
  //"ANN"
  //"math"
  
)


func main(){
    
//      b,er2:=Matrix.FromFile("tDFT.txt")     
//      fmt.Println("->",er2)
//      fft:=Fourier.FFT(b,8)
//      println("Fourier\n",fft.ToString())
        
        a,er1:=Matrix.FromFile("m.txt")
        if(er1==nil){
//             b:=a.EigenValues(0.0001)
//             fmt.Println("-",b.ToString())
               ps:=a.PInverse()
               fmt.Println("ps",ps.ToString())
        }
}
