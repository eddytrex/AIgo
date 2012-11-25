package Matrix
import (
  "math"
  //"errors"
)

func (this *Matrix) InfinityNorm()float64{
  var out float64
  out=0;
  
  if (this.m==1||this.n==1){
    out=this.A[0]
    for i:=1;i<len(this.A);i++{
      if(this.A[i]>out){out=this.A[i]}
    }
  }
  return out
}


func (this *Matrix) EuclideanNorm()float64{
  var out float64
  out=0;
  if (this.m==1||this.n==1){
    
    for i:=1;i<len(this.A);i++{
       out=out+this.A[i]*this.A[i]
    }
    out=math.Sqrt(out)
  }
  
  return out
} 


