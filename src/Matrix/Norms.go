package Matrix
import (
  "math"
  //"errors"
)
func (this *Matrix) InfinityNorm()float64{
  var out float64
  out=0;
  
    out=this.sumColumn(1)
    for i:=2;i<this.m;i++{
	temp:=this.sumColumn(i);
	if(temp>out){out=temp}
    }
  
  return out
}


// func (this *Matrix) FrobeniusNorm()float64{
//   var out float64
//   out=0;
//   if (this.m==1||this.n==1){
//     
//     for i:=0;i<len(this.A);i++{
//        out=out+this.A[i]*this.A[i]
//     }
//     
//     out=math.Sqrt(out)
//   }
//     
//   return out
// } 


func (this *Matrix) FrobeniusNorm()float64{
    sum:=make(chan float64,1);
    this.sumApplyFunction(0,len(this.A),sum,func (a float64)float64{return a*a;})
    v:=<-sum
    
    return math.Sqrt(v)
}

func (this *Matrix) sumApplyFunction(i0,i1 int, pSum chan<-float64,f func(float64)float64){
    sum:=0.0
    dx:=i1-i0
    xm:=i0+dx/2
    pSum2:=make(chan float64,THRESHOLD)
    if (dx>=THRESHOLD){
      
        go this.sumApplyFunction(i0,xm,pSum2,f) 
        this.sumApplyFunction(xm,i1,pSum2,f)
        p1:=<-pSum2
        p2:=<-pSum2
        sum=p1+p2
    }else{
    for i:=i0;i<i1;i++{
        sum=sum+f(this.A[i]);
    }
    }
    pSum<-sum
}
