package Matrix


// LU Decomposition of a Matrix 
func (this *Matrix) LUDesc()(L *Matrix, U *Matrix){
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

