package Matrix
import (
  //"math"
  
) 

func (this *Matrix) UnitVector()(*Matrix){
  duplicate:=this.Copy()
  if(this.n==1){
    norm:=this.FrobeniusNorm();
    duplicate=duplicate.Scalar(1/norm)
  }
  return duplicate
}



