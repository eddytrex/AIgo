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
//      fft,_:=Fourier.FFT(b,8)
//      println("Fourier\n",fft.ToString())
//      
//      ifft,_:=Fourier.IFFT(fft,8)
//      println("InFourier\n",ifft.ToString())
     
     c,_:=Matrix.FromFile("m.txt")     
     
     fmt.Println("c->",c.ToString())
     
//      L,U,_:=c.LUDec()         
//      println("L",L.ToString())     
//      println("U",U.ToString())
     
      Q,R:=c.QR()     
      fmt.Println("q",Q.ToString())
      fmt.Println("r",R.ToString())
     
      fmt.Println("m",Matrix.Product(Q,R).ToString())
     
        
//         a,er1:=Matrix.FromFile("m.txt")
//         if(er1==nil){
// //             b:=a.EigenValues(0.0001)
// //             fmt.Println("-",b.ToString())
//                ps:=a.PInverse()
//                fmt.Println("ps",ps.ToString())
//         }
}
