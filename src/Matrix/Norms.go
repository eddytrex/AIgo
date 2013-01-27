package Matrix
import (
  "math"
  //"errors"
)
func (this *Matrix) InfinityNorm()float64{
  var out float64
  out=0;
  
  
    out=this.sumColumn(1)
    for i:=2;i<this.n;i++{
	temp:=this.sumColumn(i);
	if(temp>out){out=temp}
    }
  
  return out
}


func (this *Matrix) FrobeniusNorm()float64{
  var out float64
  out=0;
  if (this.m==1||this.n==1){
    
    for i:=0;i<len(this.A);i++{
       out=out+this.A[i]*this.A[i]
    }
    
    out=math.Sqrt(out)
  }
    
  return out
} 

