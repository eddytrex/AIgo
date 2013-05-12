package Matrix


//return a given row of a Matrix in Matrix 1*n
func (this *Matrix) GetRow(i int) *Matrix{
  out:=NullMatrixP(1,this.n)
  for j:=1;j<=this.n;j++{
    out.SetValue(1,j,this.GetValue(i,j))
  }
  return out
}

// return a column of Matrix in a Matrix m*1
func (this *Matrix) GetColumn(j int) *Matrix{    
  out:=NullMatrixP(this.m,1)
  if(j<=this.n){
  for i:=1;i<=this.m;i++{
      //println(this.GetValue(i,j))
    out.SetValue(i,1,this.GetValue(i,j))
  }
  }
  return out
}

// substitue Row,R, in the Matrix this
func (this *Matrix) SetRow(i int,R Matrix){
   
 if(R.m==1&&R.n==this.n&&i>0&&i-1<=this.m){
    i=i-1  
    temp1:=this.A[:i*this.n]
    temp2:=this.A[(i+1)*this.n:]

    temp3:=append(temp1,R.A[:]...)
    this.A=append(temp3,temp2[:]...)  
 
 
 }
}

//substitue Column,. C, in the Matrix this
func (this *Matrix) SetColumn(j int, C Matrix){
  if(C.m==this.m&&C.n==1&&j>0&&j<=this.n){
    for i:=1;i<=this.m;i++{
      this.SetValue(i,j,C.GetValue(i,1))
    }
  }
}


// return the sum of the main diagonal of the square Matrix
func (this *Matrix)Trace() float64{
  var out float64
  out=0
  if(this.m==this.n){
  for i:=1;i<this.m;i++{
    out=out+this.GetValue(i,i)
  }
  }
  return out
}



//Multiply a scalar(c) by a row(i) and Return a 1xn Matrix of a Matrix mxn
func (this *Matrix) ScalarRowMatrix(i int,  c float64)(*Matrix){
  out:=NullMatrixP(1,this.n)
  i=i-1
  k:=0
  for j:=0;j<out.n;j++{
    pos:=i*out.n+j
    
    out.A[k]=c*this.A[pos]
    
    k++
  }
  return out
}


//Multiply a row i by c and adds to  row i0 
func (this *Matrix) ScalarRowAndAdd(i0,i int, C float64){
      
      for j:=1;j<=this.n;j++{
	  NV:=this.GetValue(i0,j)+C*this.GetValue(i,j) 	  
	  this.SetValue(i0,j,NV)
      }      
}


//  multiply a row of a Matrix  by a number c
func (this *Matrix) ScalarRow(i int, C float64){
  for j:=1;j<=this.n;j++{
      
    this.SetValue(i,j,C*this.GetValue(i,j))    
  }
}


//Get a Matrix (m-1)rows and n columns of a Matrix mxn
func (this *Matrix)MatrixWithoutRow(i int)*Matrix{
    out:=NullMatrixP(this.m-1,this.n)
    At:=make([]float64,len(this.A))
    copy(At,this.A)
    i=i-1
    
    At=append(At[:i*this.n],At[(i+1)*this.n:]...)
    //println("at",this.m-1)
    out.A=At
    return out
}

//Get a Matrix m rows and (n-1) columns of a Matrix mxn
func (this *Matrix)MatrixWithoutColumn(j int)*Matrix{
    out:=NullMatrixP(this.m,this.n-1)
    At:=make([]float64,len(this.A))  
    copy(At,this.A)
    err:=1  
    for i:=0;i<this.m;i++{
     var it int
     it=i*(this.n-err)+(j-err)
     At=append(At[:it],At[it+1:]...)
    }    
    out.A=At
    return out
}






// Swap Row in the position i0 with the position i 
func (this *Matrix)SwapRow(i0,i int){
  if(i0!=i){
    i=i-1
    i0=i-1
    for j:=0;j<this.n;j++{
      Posi0:=(i0)*this.n+j
      Posi:=(i)*this.n+j
      
      temp:=this.A[Posi0]
      this.A[Posi0]=this.A[Posi]
      this.A[Posi]=temp
    }
  }
}

// Swap Column in the position j0 with the position j
func (this *Matrix)SwapColumn(j0,j int){
  if(j0!=j){
    j=j-1
    j0=j0-1
    for i:=0;i<this.m;i++{
      Posj0:=i*this.n+j0
      Posj:=i*this.n+j
    
      temp:=this.A[Posj0]
      this.A[Posj0]=this.A[Posj]
      this.A[Posj]=temp

    }
  }
}

//append the columns of matrix Ci to this
func (this *Matrix) AddColumn(Ci Matrix)*Matrix{
  if(this.m==Ci.m){
    out:=NullMatrixP(this.m,this.n+Ci.n)
    var newA []float64
    for i:=0;i<this.m;i++{
      
      rowTempThis:=make([]float64,this.n)
      rowTempCi:=make([]float64,Ci.n)
      
      copy(rowTempThis,this.A[i*this.n:(i+1)*this.n])
      copy(rowTempCi,Ci.A[i*Ci.n:(i+1)*Ci.n])
           	
      newRow:=append(rowTempThis,rowTempCi[:]...)
      newA=append(newA,newRow[:]...)
      
    }
    copy(out.A,newA)
    return out
  }
  return nil
}


func (this *Matrix) AddRowsToTop(Cj *Matrix)(*Matrix){
    
    if(Cj.n==this.n){
        
        out:=NullMatrixP(this.m+Cj.m,this.n)
        var newA []float64;
        newA=append(newA,Cj.A[:]...)
        newA=append(newA,this.A[:]...)
        
        out.A=newA;
  
        return out
        
    }
    return nil
}


func (this *Matrix) AddRowsToDown(Cj *Matrix)(*Matrix){
    
    if(Cj.n==this.n){
        
        out:=NullMatrixP(this.m+Cj.m,this.n)
        var newA []float64;        
        newA=append(newA,this.A[:]...)
        newA=append(newA,Cj.A[:]...)
        
        out.A=newA;
  
        return out
        
    }
    return nil
}

// return the sum all elements of a vector a column 
func (this *Matrix) SumVectorColumn()float64{
  var sum float64
  sum=0
  if(this.m==1){
    for j:=0;j<this.n;j++{
      sum=sum+this.A[j]
    }
  }
  return sum
}


func (this *Matrix) sumColumn(i int) float64{
   var sum float64  
   sum=0
   for j:=1; j<this.n;j++{
     sum=sum+this.GetValue(i,j)
  }
  return sum
}
