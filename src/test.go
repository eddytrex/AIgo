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
     

     c,er:=Matrix.FromFile("Matrix/test/null.txt")
     d,er1:=Matrix.FromFile("Matrix/test/I.txt")     
     
     
     fmt.Println(">",Matrix.Equal(c,d))
     fmt.Println("c->",c.ToString(),"->",er,er1)
     
//      L,U,_:=c.LUDec()         
//      println("L",L.ToString())     
//      println("U",U.ToString())
     

//       Q,R:=c.QR()     
//       fmt.Println("q",Q.ToString())
//       fmt.Println("r",R.ToString())
//      
//       
//       t,_:=R.Transpose().GaussElimitation(c.Transpose())
//       t2,err:=R.GaussElimitation(t)
//       fmt.Println("m",t2.ToString(),"-<",err)

     
        
//         a,er1:=Matrix.FromFile("m.txt")
//         if(er1==nil){
// //             b:=a.EigenValues(0.0001)
// //             fmt.Println("-",b.ToString())
//                ps:=a.PInverse()
//                fmt.Println("ps",ps.ToString())
//         }
}
