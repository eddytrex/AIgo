package Matrix
import (
  //"math"
  
) 

func (this *Matrix) UnitVector()(*Matrix){
  duplicate:=this.Copy()
  if(this.n==1){
    norm:=this.FrobeniusNorm();
    duplicate=duplicate.Scalar(complex(1/norm,0))
  }
  return duplicate
}

func  DotMultiplication(a,b *Matrix)(*Matrix){
    res:=NullMatrixP(1,a.n)
    if(a.m==1&&a.n==b.n){
        for i:=1;i<=a.n;i++{
            res.SetValue(1,i,a.GetValue(1,i)*b.GetValue(1,i))
        }
    }
    return res;
}



