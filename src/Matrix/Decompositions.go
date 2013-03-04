package Matrix


// LU Decomposition of a Matrix 
func (this *Matrix) LUDec()(L *Matrix, U *Matrix){
  if(this.m==this.n){
   U:=this.Copy()
   L:=I(this.n)	    
   
   var UAnt float64
   
   for k:=1;k<=this.m;k++{   
     
     for i:=k+1;i<=this.m;i++{
       
	L.SetValue(i,k,U.GetValue(i,k)/U.GetValue(k,k))   	
	
       for j:=1;j<=this.n;j++{
	 
	 UAnt=U.GetValue(i,j)-U.GetValue(k,j)*L.GetValue(i,k)	 
	 U.SetValue(i,j,UAnt)
	 
      }
    }
   }
   for i:=1;i<=this.m;i++{
     L.SetValue(i,i,1) 
   }
   return L,U
  }
  return nil,nil
}

// QR Decomposition using  Householder reflections

func (this *Matrix)QRDec()(Q1,R1 *Matrix){
    Q:=NullMatrixP(this.m,this.n)
    R:=NullMatrixP(this.m,this.n)
    var first=true;
    var alpha float64
    var Qp *Matrix
    Ai:=this.Copy()
    for i:=1;i<this.m;i++{
        
        X:=Ai.GetColumn(i)
        
        e:=NullMatrix(X.m,1)
        e.SetValue(i,1,1)
        
        x1:=X.GetValue(i,1)

        if(x1>0){
           alpha=-abs(X.FrobeniusNorm())
        }else{
          alpha=abs(X.FrobeniusNorm())
        }
        
        u,_:=Sustract(X,e.Scalar(alpha))        
        v:=u.UnitVector();

        Qi,_:=v.HouseholderTrasformation() 
        
        
        if(first){
            Qp=Product(Qi,this)
            Q=Qi
            
            first=false;
        }else{         
            Q=Product(Q,Qi)
            Qp=Product(Qi,Ai)
        }
        
        for l:=1;l<=i;l++{
            Qp=Qp.SubMatrix(1,1)
        }
        
        Qp=SetSubMatrixToI(this.n,i+1,Qp)        
        Ai=Qp
    }
    
    R=Product(Q.Transpose(),this)
    
    
    return Q,R
}

// Set a matrix in the position beginin in PosI,PosI to rest of matrix of NxN
func SetSubMatrixToI(n int,posI int ,pQ *Matrix)(*Matrix){
    out:=I(n);
    
    if(posI<n&&(posI+pQ.n-1)==n){
        
        if(pQ.m<n){         
            setMatrix:=NullMatrixP((n-pQ.m),pQ.n)
            pQ=pQ.AddRowsToTop(setMatrix)
            
            for i:=1;i<=pQ.n;i++{
                ci:=pQ.GetColumn(i);
                out.SetColumn(i+posI-1,*ci)
            }
            
        }else if(pQ.m==n){
            
            return pQ
        }
        return out
    }
    
    return nil
}
