package Matrix

import(
  "errors"
  "strconv"
  "math/rand"
  "math"
  "time"
 
)

type matrix struct {
  // m rows and n columns
   m,n int
   //Values of the matrix
   A  []float64 
}


// Return the value in the matrix position i,j
func (this *matrix)GetValue(i,j int)float64{
  i=i-1
  j=j-1
  
  return this.A[i*this.n+j]
  
}

// Set the value (val) in the matrix position i,j in 
func (this *matrix)SetValue(i,j int,val float64){
  i=i-1
  j=j-1
  this.A[i*this.n+j]=val
}

// return a matrix with zero  in all positions and m,n dimensions
func NullMatrix(m int, n int)matrix{
  A:=make([]float64,m*n,m*n)
  var out matrix 
  out.A=A
  out.m=m
  out.n=n
  return out
}

// return a square matrix nxn  and one's in the main diagonal 
func I(n int)*matrix{
  out:=NullMatrix(n,n)
  j:=0
  for i:=0;i<len(out.A);i=i+out.m{
    out.A[i+j]=1
    j++
  }
  return &out
}

// return the sum of the main diagonal of the square matrix
func (this *matrix)trace() float64{
  var out float64
  out=0
  if(this.m==this.n){
  for i:=1;i<this.m;i++{
    out=out+this.GetValue(i,i)
  }
  }
  return out
}


// return a string with the values of the matrix
func (this *matrix)ToString() string{
  var out string
  out=""
  if(this!=nil){
  for i:=0;i<this.m;i++{
    for j:=0;j<this.n;j++ {      
      out=out+" "+strconv.FormatFloat(this.A[i*this.n+j],'f',6,64)
    }
    out=out+"\n"
  }
  }
  return out
}




//Return a 1xn matrix of a matrix mxn
func (this *matrix) ScalarRowMatrix(i int,  c float64)(*matrix){
  out:=NullMatrix(1,this.m)
  i=i-1
  k:=0
  for j:=0;j<out.n;j++{
    pos:=i*out.n+j
    
    out.A[k]=c*this.A[pos]
    
    k++
  }
  return &out
}

// return the sum all elements of a vector a column 
func (this *matrix) SumVectorColumn()float64{
  var sum float64
  sum=0
  if(this.m==1){
    for j:=0;j<this.n;j++{
      sum=sum+this.A[j]
    }
  }
  return sum
}

//Get a matrix m rows and (n-1) columns of a matrix mxn
func (this *matrix)MatrixWithoutColumn(j int)*matrix{
    out:=NullMatrix(this.m,this.n-1)
    At:=make([]float64,len(this.A))  
    copy(At,this.A)
    err:=1  
    for i:=0;i<this.m;i++{
     var it int
     it=i*(this.n-err)+(j-err)
     At=append(At[:it],At[it+1:]...)
    }    
    out.A=At
    return &out
}

//Get a matrix (m-1)rows and n columns of a matrix mxn
func (this *matrix)MatrixWithoutRow(i int)*matrix{
    out:=NullMatrix(this.m-1,this.n)
    At:=make([]float64,len(this.A))
    copy(At,this.A)
    i=i-1
    At=append(At[:i*this.m],At[(i+1)*this.m:]...)
    out.A=At
    return &out
}


// Return the matrix adjoint matrix(this) of the position i,j

func (this *matrix)AdjMatrix(i,j int) *matrix{
  out:=this.MatrixWithoutRow(i).MatrixWithoutColumn(j)
  return out
}


// return the determinant of a square matrix 
// O(n!) I don't think someone will use it
func (this *matrix)Det_LapaceExpasion()(float64,error){
  if(this.n==this.m){
    if(this.n==1){
      return this.GetValue(1,1),nil
      
    }else{
      var sum float64
      sum=0
	for i:=1;i<=this.m;i++{
	  
	 temp,_:=this.AdjMatrix(1,i).Det_LapaceExpasion()
	 
	 if(i%2!=0){
	   
	    temp=temp*this.GetValue(i,1)	   
	  }else{
	    temp=temp*this.GetValue(i,1)*-1	   
	  }
	  sum=sum+temp 
	}
	return sum,nil
    }
    
  }
  return 0,errors.New(" the matrix have to be square")
}


//  return  a copy of a matrix
func (this *matrix) Copy()(*matrix){
   out:=NullMatrix(this.n,this.m)
   copy(out.A,this.A)
   return &out
}



// Return the determinant of a matrix by LU  Decomposition 
func (this *matrix) Det_LU()float64{
  _,U:=this.LUDesc()  
  
  var Det float64
  Det=1
  
  for ui:=1;ui<=U.n;ui++{
    uii:=U.GetValue(ui,ui)
    Det=Det*uii
  }
  return Det
}

// substitue Row R in the matrix
func (this *matrix) SetRow(i int,R matrix){
 
 if(R.m==1&&R.n==this.n&&i>0){
 i=i-1  
 
 temp1:=this.A[:i*this.m]
 temp2:=this.A[(i+1)*this.m:]

 temp3:=append(temp1,R.A[:]...)
 this.A=append(temp3,temp2[:]...)  
 
 
 }
}

//  multiply a row of a matrix  by a number c
func (this *matrix) ScalarRow(i int, C float64){
  for j:=1;j<=this.n;j++{
    this.SetValue(i,j,C*this.GetValue(i,j))
  }
}

func (this *matrix) InverseGauss()(*matrix, error){
  if(this.n==this.m){
    Aum:=I(this.n)
    for i:=1;i<=this.m;i++{
       
	j:=i
	for k:=i;k<=this.m;k++{
	  if (math.Abs(this.GetValue(k,i))>math.Abs(this.GetValue(j,i))){
	    j=k
	  }
	}
	if j!=i{
	  this.SwapRow(i,j)
	  Aum.SwapRow(i,j)
	}
	if(this.GetValue(i,i)==0){
	  
	  return nil,errors.New(" Singualr matrix")
	}	
      
	Thisii:=this.GetValue(i,i)
	Tii:=1/Thisii
	
	this.ScalarRow(i,Tii)
	Aum.ScalarRow(i,Tii)
	
	for k:=1;k<=this.m;k++{
	  
	  if( k!=i ){
	      C:=-this.GetValue(k,i);
	      this.ScalarRowAndAdd(k,i,C)
	      Aum.ScalarRowAndAdd(k,i,C)   
	  }
	}
    }
    return Aum,nil
  }
  return nil,errors.New(" the matrix is not Square ")
}

//Multiply a row i by c and adds to a row i 
func (this *matrix) ScalarRowAndAdd(i0,i int, C float64){
      
      for j:=1;j<=this.n;j++{
	  C:=this.GetValue(i0,j)+C*this.GetValue(i,j) 	  
	  this.SetValue(i0,j,C)
      }      
}



// Return a Inverse of a Square matrix by LU  Decomposition 
func (this *matrix)Inverse() (*matrix,error){
  out:=NullMatrix(this.m,this.n)
  var newOutA []float64
  if(this.n==this.m){  
  l,u:=this.LUDesc()
  
  for i:=1;i<=this.m;i++{
    column:=NullMatrix(this.m,1)
    column.SetValue(i,1,1)
    
    z:=l.fSubs(column)
    b:=u.bSubs(*z)   
    newOutA=append(newOutA,b.A[:]...)

    }
    
  }else{
    return nil,errors.New(" the matrix has to be square")
  }
  out.A=newOutA
  out=*out.Transpose()
 return &out,nil 
}

// LU Decomposition of a matrix 
func (this *matrix) LUDesc()(L *matrix, U *matrix){
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

//Solve by forward substitution method for L matrix in Inverse
func (this *matrix) fSubs(B matrix)*matrix{
  out:=NullMatrix(B.m,1)
  lx:=NullMatrix(B.m,1)
  if(this.n==this.m&&B.m==this.m&&B.n==1){
    
    for i:=1;i<=this.n;i++{
	lx.SetValue(i,1,B.GetValue(i,1))
	for j:=1;j<i;j++{
	  
	  templx:=lx.GetValue(i,1)-this.GetValue(i,j)*lx.GetValue(j,1)
	  lx.SetValue(i,1,templx)	
	  
	}
	templx:=lx.GetValue(i,1)/this.GetValue(i,i)
	lx.SetValue(i,1,templx)
      }
      out=lx
    
  }
  return &out
}

//Solve by back substitution method for a U matrix in Inverse
func (this *matrix) bSubs(B  matrix)*matrix{
  out:=NullMatrix(B.m,1)
  ux:=NullMatrix(B.m,1)
  
  if(this.n==this.m&&B.m==this.m&&B.n==1){
   for i:=this.n;i>=1;i--{
	ux.SetValue(i,1,B.GetValue(i,1))
	for j:=i+1;j<=this.n;j++{
	  
	  tempux:=ux.GetValue(i,1)-this.GetValue(i,j)*ux.GetValue(j,1)
	  ux.SetValue(i,1,tempux)
	}
	tempux:=ux.GetValue(i,1)/this.GetValue(i,i)
	ux.SetValue(i,1,tempux)
      }
      out=ux 
  }
  return &out
}

// Return a Matrix Transpose 
func (this *matrix) Transpose() *matrix{
  out:=NullMatrix(this.n,this.m)
  for i:=1;i<=this.m;i++{
    for j:=1;j<=this.n;j++{
      out.SetValue(j,i,this.GetValue(i,j))
    }
  }
  return &out
}

// Swap Column in the position j0 with the position j
func (this *matrix)SwapColumn(j0,j int){
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

// Swap Row in the position i0 with the position i 
func (this *matrix)SwapRow(i0,i int){
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



// Verify if the matrix (this) si Triangular Lower
func (this *matrix) TriangularLower()bool{
  var out bool
  out=false
  bandera:=true
  if(this.n==this.m){
    for i:=1;i<=this.m&&bandera;i++{
      for j:=1;j<=this.n&&bandera;j++{
	if(i<j&&this.GetValue(i,j)!=0&&bandera){
	  bandera=false
	}
      }
    }
  }
  out=bandera
  return out
}


// Verify if the matrix (this) si Triangular Upper
func (this *matrix) TriangularUpper()bool{
  var out bool
  out=false
  bandera:=true
  if(this.n==this.m){
    for i:=1;i<=this.m&&bandera;i++{
      for j:=1;j<=this.n&&bandera;j++{
	if(i>j&&this.GetValue(i,j)!=0&&bandera){
	  bandera=false
	}
      }
    }
  }
  out=bandera
  return out
}



// If the matrix (this) is Triangular Lower or Triangular Upper; return the result of it
//Back and forward substitution
func (this *matrix) FBSubs(B matrix)(*matrix,error){
  out:=NullMatrix(B.m,1)
  lx:=NullMatrix(B.m,1)
  ux:=NullMatrix(B.m,1)
  if(this.n==this.m&&B.m==this.m&&B.n==1){
    if(this.TriangularLower()){
      
      for i:=1;i<=this.n;i++{
	lx.SetValue(i,1,B.GetValue(i,1))
	for j:=1;j<i;j++{
	  
	  templx:=lx.GetValue(i,1)-this.GetValue(i,j)*lx.GetValue(j,1)
	  lx.SetValue(i,1,templx)	
	  
	}
	templx:=lx.GetValue(i,1)/this.GetValue(i,i)
	lx.SetValue(i,1,templx)
      }
      out=lx
    }
    
    if(this.TriangularUpper()){
      
      for i:=this.n;i>=1;i--{
	ux.SetValue(i,1,B.GetValue(i,1))
	for j:=i+1;j<=this.n;j++{
	  
	  tempux:=ux.GetValue(i,1)-this.GetValue(i,j)*ux.GetValue(j,1)
	  ux.SetValue(i,1,tempux)
	}
	tempux:=ux.GetValue(i,1)/this.GetValue(i,i)
	ux.SetValue(i,1,tempux)
      }
      out=ux
    }
    return &out,nil
  }
  return nil,errors.New(" The matrix is no square")
}

//return a given row of a matrix in matrix 1*n
func (this *matrix) GetRow(i int) *matrix{
  out:=NullMatrix(1,this.m)
  for j:=1;j<=this.n;j++{
    out.SetValue(1,j,this.GetValue(i,j))
  }
  return &out
}

// return a column of matrix in a matrix m*1
func (this *matrix) GetColumn(j int) *matrix{
  out:=NullMatrix(this.n,1)
  for i:=1;i<=this.m;i++{
    out.SetValue(j,1,this.GetValue(j,i))
  }
  return &out
}



//if the matrix is square get only the main diagonal in a matrix n*m other is 0
func (this *matrix) GetDiagonal() (*matrix,error){
  if(this.n==this.m){
  out:=NullMatrix(this.n,this.m)
  for i:=1;i<=this.m;i++{
    for j:=1;j<=this.n;j++{
      out.SetValue(i,j,this.GetValue(i,j))
      }
  }
  return &out,nil
  }
  return nil,errors.New(" The matrix is no square")
}





// A+B  (A,B  are matrix)
func Sum(A,B *matrix)(*matrix,error){
  if(A.n==B.n&&A.m==B.m){
    
    out:=NullMatrix(A.m,A.n)
    for i:=0;i<len(A.A);i++{
      out.A[i]=A.A[i]+B.A[i]
    }
    return &out,nil
  }
  return nil,errors.New(" The matrixes don't have the same dimensions")
}

// A-B  (A,B are matrix)
func Sustract(A,B *matrix)(*matrix,error){
  if(A.n==B.n&&A.m==B.m){
    out:=NullMatrix(A.m,A.n)
    for i:=0;i<len(A.A);i++{
      out.A[i]=A.A[i]-B.A[i]
    }
    return &out,nil
  }
  return nil,errors.New("The matrixes don't have the same dimensions")
}

// Multiply a matrix for a scalar   cA
func (this *matrix) Scalar(c float64)(*matrix){
  
  if(c==0){
    out:=NullMatrix(this.n,this.m)
    return &out
  }else{
    out:=this.Copy()
  for i:=0;i<len(out.A);i++{
    out.A[i]=c*out.A[i]
  }
  return out
  }
  return nil
}





func Product(A,B matrix) *matrix{
   out:=NullMatrix(A.m,B.n)
  
   if(A.n==B.n){
     done:=make(chan bool)
      go multr(A,B,out,1,A.m,1,B.n,1,A.n,done)
      <-done
   }
  return &out
}

const THRESHOLD=100
func multr(A,B,C matrix,i0,i1,j0,j1,k0,k1 int,done chan <-bool){
  
  di:=i1-i0
  dj:=j1-j0
  dk:=k1-k0
  
  done2:=make(chan bool,THRESHOLD)
  if(di>=dj&&dj>=dk&&di>=THRESHOLD){
      mi:=i0+di/2
      go multr(A,B,C,i0,mi,j0,j1,k0,k1,done2)
       multr(A,B,C,mi,i1,j0,j1,k0,k1,done2)
      <-done2
      <-done2
  }else if ( dj>=dk&&dj>=THRESHOLD){
      mj:=j0+dj/2
      go multr(A,B,C,i0,i1,j0,mj,k0,k1,done2)
      multr(A,B,C,i0,i1,mj,j1,k0,k1,done2)
      <-done2
      <-done2
  }else if (dk>=THRESHOLD){
      mk:=k0+dk/2
      go multr(A,B,C,i0,i1,j0,j1,k0,mk,done2)
      multr(A,B,C,i0,i1,j0,j1,mk,k1,done2)
      <-done2
      <-done2
  }else{    
    for i:=i0;i<=i1;i++{
      for j:=j0;j<=j1;j++{
	var temp float64
	temp=C.GetValue(i,j)
	for k:=k0;k<=k1;k++{
	  temp=temp+A.GetValue(i,k)*B.GetValue(k,j)
	}
	C.SetValue(i,j,temp)
      }
    }
  }
  done<-true
}


// Return the AB Product
func  Multiplication(A,B matrix) *matrix{
  out:=NullMatrix(A.m,B.n)
  
    done:=make(chan bool)
    mult:=make(chan float64)
  
    for i:=1;i<=A.m;i++{
     for k:=1;k<=B.n;k++{
      
      go out.multRowColumn(i,k,A,B,mult)
      go out.setCValue(i,k,mult,done)
      <-done
    }
  }
  return &out
}


// for matrix multiplication in parallel
func (this *matrix) multRowColumn(i,k int, A,B matrix,out chan <-float64){
  var temp float64
  temp=0
  for j:=1;j<=A.n;j++{    
    temp=temp+A.GetValue(i,j)*B.GetValue(j,k)
  }  
   out<-temp
}

//for matrix multiplication in parallel
func (this *matrix) setCValue(i,k int, in <- chan  float64, done chan<- bool){ 
     for  {
      temp:=<-in
      this.SetValue(i,k,temp)
      break
    }
    done<-true
}

// In a matrix to matrix with dimensions A (nxm) and B(n1xm1) return a matrix C(n*n1xm*m1) 
// with a elements Ci=Aij*B 
func KroneckerProduct(A,B matrix)*matrix{
  out:=NullMatrix(A.m*B.m,A.n*B.n)
  for i:=1;i<=A.m;i++{
    for j:=1;j<=A.n;j++{
	Aij:=A.GetValue(i,j)
	mtemp:=B.Scalar(Aij)
	out.A=append(out.A,mtemp.A[:]...)
	
    }
  }
  return &out
}



// Apply the function (f) to all elements of the matrix (
func (this *matrix) Apply(f func(float64)float64) *matrix{
  out:=this.Copy()
  for i:=0;i<len(out.A);i++{
    newVal:=f(out.A[i])
    out.A[i]=newVal
  }
  return out
}


// Return a matrix of m,n size and random elements 1-10
func RandomMatrix(m,n int)*matrix{
  out:=NullMatrix(m,n)
  rand.Seed(time.Now().UTC().UnixNano())
  for i:=1;i<=out.m;i++{
    for j:=1;j<=out.n;j++{
      
      NumeroAleaotrio:=rand.Float64()*10
      
      out.SetValue(i,j,NumeroAleaotrio)
      
    }
  }
  return &out
}






/*func FromFile(string nameFile)(*matrix){
  
}*/






