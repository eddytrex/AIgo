package Fourier 
import (
    "Matrix"    
    "math"
    "math/cmplx"
    "errors"
)

// func DFT(this *Matrix.Matrix)(*Matrix.Matrix){
//     
//     C:=Matrix.NullMatrixP(this.GetMRows(),1)
//     
//     lengthi:=this.GetMRows()
//     lengthf:=(complex128)(lengthi)
//     for i:=1;i<=lengthi;i++{
//         
//         floati:=(complex128)(i)-1
//         
//         sin:=0.0
//         cos:=0.0 
//         
//         for j:=1;j<=lengthi;j++{            
//             floatj:=(complex128)(j)-1            
//             sin=sin+this.GetValue(j,1)*math.Sin((2*math.Pi*floati*floatj)/lengthf);
//             cos=cos+this.GetValue(j,1)*math.Cos((2*math.Pi*floati*floatj)/lengthf);
//         }                
//         //|Ci|
//         C.SetValue(i,1,math.Sqrt(sin*sin+cos*cos));
//     }
// 
//     return C
// }


func FFT(this *Matrix.Matrix,N int)(*Matrix.Matrix,error){    
    if(N>this.GetMRows()){return nil,errors.New(" The number of Rows of the matrix (this) must be greater or equal than N ")}    
    if(N&(N-1)==0){        
        Xr:=FFT_ct(this,N,1)                  
         return Xr,nil
    }
    return nil,errors.New(" The N parameter has to be power of 2")
}

func IFFT(this *Matrix.Matrix,N int)(*Matrix.Matrix,error){    
    if(N>this.GetMRows()){return nil,errors.New(" The number of Rows of the matrix (this) must be greater or equal than N ")}    
    if(N&(N-1)==0){        
        Xr:=FFT_ct(this,N,1)                  
        Xr=Xr.Scalar(complex(float64(1)/float64(N),0))
         return Xr,nil
    }
    return nil,errors.New(" The N parameter has to be power of 2")
}

func  FFT_ct(this *Matrix.Matrix,N, skip int )(*Matrix.Matrix){
                           
        if(N==1){            
            return this.GetRow(1)
        }        
        
        //*x+skip
        xskip:=this.Copy()        
        for i:=1;i<=skip;i++{    
            xskip=xskip.MatrixWithoutRow(1)              
        }                     
         p:=Matrix.NullMatrixP(skip,this.GetNColumns());         
         xskip=xskip.AddRowsToDown(p)
         
                
        Ar:=FFT_ct(this,N/2,skip*2)
        Br:=FFT_ct(xskip,N/2,skip*2)
        
        for k:=0;k<N/2;k++{
                   Br.ScalarRow(k+1,cmplx.Exp(complex(0,-2.0*math.Pi*float64(k)/float64(N))))                   
        }
                
        Xr:=Matrix.NullMatrixP(N,this.GetNColumns())
        //Xi:=Matrix.NullMatrixP(N,this.GetNColumns())
        for k:=0;k<N/2;k++{
            sr,_:=Matrix.Sum(Ar.GetRow(k+1),Br.GetRow(k+1))
            Xr.SetRow(k+1,*sr)                       
            rr,_:=Matrix.Sustract(Ar.GetRow(k+1),Br.GetRow(k+1));                                   
            Xr.SetRow(k+1+N/2,*rr)                                   
        }          
        return Xr
}