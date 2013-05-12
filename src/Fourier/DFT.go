package Fourier 
import (
    "Matrix"    
    "math"
    "errors"
)

func DFT(this *Matrix.Matrix)(*Matrix.Matrix){
    
    C:=Matrix.NullMatrixP(this.GetMRows(),1)
    
    lengthi:=this.GetMRows()
    lengthf:=(float64)(lengthi)
    for i:=1;i<=lengthi;i++{
        
        floati:=(float64)(i)-1
        
        sin:=0.0
        cos:=0.0 
        
        for j:=1;j<=lengthi;j++{            
            floatj:=(float64)(j)-1            
            sin=sin+this.GetValue(j,1)*math.Sin((2*math.Pi*floati*floatj)/lengthf);
            cos=cos+this.GetValue(j,1)*math.Cos((2*math.Pi*floati*floatj)/lengthf);
        }                
        //|Ci|
        C.SetValue(i,1,math.Sqrt(sin*sin+cos*cos));
    }

    return C
}


func FFT(this *Matrix.Matrix,N int)(*Matrix.Matrix,error){
    
    if(N>this.GetMRows()){return nil,errors.New(" The number of Rows of the matrix (this) must be greater or equal than N ")}
    
    if(N&(N-1)==0){
        
        Xr,Xi:=FFT_ct(this,N,1)
         C:=Matrix.NullMatrixP(Xr.GetMRows(),Xr.GetNColumns())
         for i:=1;i<=C.GetMRows();i++{
            for j:=1;j<=C.GetNColumns();j++{
                 C.SetValue(i,j,math.Sqrt(Xi.GetValue(i,j)*Xi.GetValue(i,j)+Xr.GetValue(i,j)*Xr.GetValue(i,j)))
             }
         }
         return C,nil
    }
    return nil,errors.New(" The N parameter has to be power of 2")
}

func  FFT_ct(this *Matrix.Matrix,N, skip int )(*Matrix.Matrix,*Matrix.Matrix){
                           
        if(N==1){
            Im:=Matrix.NullMatrixP(1,this.GetNColumns())           
            return this.GetRow(1),Im
        }        
        
        //*x+skip
        xskip:=this.Copy()        
        for i:=1;i<=skip;i++{    
            xskip=xskip.MatrixWithoutRow(1)              
        }                     
         p:=Matrix.NullMatrixP(skip,this.GetNColumns());         
         xskip=xskip.AddRowsToDown(p)
         
                
        Ar,Ai:=FFT_ct(this,N/2,skip*2)
        Br,Bi:=FFT_ct(xskip,N/2,skip*2)
        
        for k:=0;k<N/2;k++{
            
                   rr:=Br.ScalarRowMatrix(k+1,math.Cos(2.0*math.Pi*float64(k)/float64(N)))
                   ii:=Bi.ScalarRowMatrix(k+1,-math.Sin(2.0*math.Pi*float64(k)/float64(N)))
                   
                   rrii,_:=Matrix.Sum(rr,ii)
                   
                   ri:=Br.ScalarRowMatrix(k+1,math.Sin(2.0*math.Pi*float64(k)/float64(N)))
                   ir:=Bi.ScalarRowMatrix(k+1,math.Cos(2.0*math.Pi*float64(k)/float64(N)))
                   
                   
                   riir,_:=Matrix.Sum(ri,ir)
                   
                   Br.SetRow(k+1,*rrii)
                   Bi.SetRow(k+1,*riir)                                        
        }
                
        Xr:=Matrix.NullMatrixP(N,this.GetNColumns())
        Xi:=Matrix.NullMatrixP(N,this.GetNColumns())
        for k:=0;k<N/2;k++{
            sr,_:=Matrix.Sum(Ar.GetRow(k+1),Br.GetRow(k+1))
            si,_:=Matrix.Sum(Ai.GetRow(k+1),Bi.GetRow(k+1))
            
            Xr.SetRow(k+1,*sr)
            Xi.SetRow(k+1,*si)
            
            rr,_:=Matrix.Sustract(Ar.GetRow(k+1),Br.GetRow(k+1));
            ri,_:=Matrix.Sustract(Ai.GetRow(k+1),Bi.GetRow(k+1));
                        
            Xr.SetRow(k+1+N/2,*rr)                       
            Xi.SetRow(k+1+N/2,*ri)            
        }          
        
        
        return Xr,Xi    
}

