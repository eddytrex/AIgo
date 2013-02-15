package Matrix

import(
  "errors"
  "math"
  "fmt"
)

// return a Matrix with zero  in all positions and m,n dimensions
func NullMatrix(m int, n int)Matrix{
  A:=make([]float64,m*n,m*n)
  var out Matrix 
  out.A=A
  out.m=m
  out.n=n
  return out
}

// return a square Matrix nxn  and one's in the main diagonal 
func I(n int)*Matrix{
  out:=NullMatrix(n,n)
  j:=0
  for i:=0;i<len(out.A);i=i+out.m{
    out.A[i+j]=1
    j++
  }
  return &out
}

// Multiply a Matrix for a scalar   cA
func (this *Matrix)Scalar(c float64)(*Matrix){
    out:=NullMatrix(this.m,this.n)
    if(c==0){ 
      return &out
    }
    done:=make(chan bool)
    go scalarR(0,len(this.A),c,*this,out,done)
    <-done
    return &out 
}

func scalarR(i0,i1 int,c float64,A,C Matrix,done chan <-bool ){
  di:=(i1-i0)
  done2:=make(chan bool,THRESHOLD)
  if(di>=THRESHOLD){
    mi:=i0+di/2
    go scalarR(i0,mi,c,A,C,done2)
     scalarR(mi,i1,c,A,C,done2)
    <-done2
    <-done2
  }else{
    for i:=i0;i<i1;i++{
      C.A[i]=c*A.A[i]
    }
  }
  done<-true
}

/*
// Multiply a Matrix for a scalar   cA
func (this *Matrix) Scalar(c float64)(*Matrix){
  
  if(c==0){
    out:=NullMatrix(this.m,this.n)
    return &out
  }else{
    out:=this.Copy()
  for i:=0;i<len(out.A);i++{
    out.A[i]=c*out.A[i]
  }
  return out
  }
  return nil
}*/



// A+B  (A,B  are Matrix)
func Sum(A,B Matrix)(*Matrix,error){
  if(A.m==B.m&&A.n==B.n){
    
    out:=NullMatrix(A.m,A.n)
    done:=make(chan bool)
    go sumR(0,len(A.A),A,B,out,done)
    <-done
    return &out,nil
  }
  return nil,errors.New(" The Matrixes don't have the same dimensions")
}

func sumR(i0,i1 int,A,B,C Matrix,done chan <-bool ){
  di:=(i1-i0)
  done2:=make(chan bool,THRESHOLD)
  if(di>=THRESHOLD){
    mi:=i0+di/2
    go sumR(i0,mi,A,B,C,done2)
    go sumR(mi,i1,A,B,C,done2)
    <-done2
    <-done2
  }else{
    for i:=i0;i<i1;i++{
      C.A[i]=A.A[i]+B.A[i]
    }
  }
  done<-true
}

// A+B  (A,B  are Matrix)
/*func Sum(A,B Matrix)(*Matrix,error){
  if(A.n==B.n&&A.m==B.m){
    
    out:=NullMatrix(A.m,A.n)
    for i:=0;i<len(A.A);i++{
      out.A[i]=A.A[i]+B.A[i]
    }
    return &out,nil
  }
  return nil,errors.New(" The Matrixes don't have the same dimensions")
}*/

// A-B  (A,B are Matrix)
func Sustract(A,B Matrix)(*Matrix,error){
  if(A.m==B.m&&A.n==B.n){
    
    out:=NullMatrix(A.m,A.n)
    done:=make(chan bool)
    go sustractR(0,len(A.A),A,B,out,done)
    <-done
    return &out,nil
  }
  return nil,errors.New(" The Matrixes don't have the same dimensions")
}

func sustractR(i0,i1 int,A,B,C Matrix,done chan <-bool ){
  di:=(i1-i0)
  done2:=make(chan bool,THRESHOLD)
  if(di>=THRESHOLD){
    mi:=i0+di/2
    go sustractR(i0,mi,A,B,C,done2)
    sustractR(mi,i1,A,B,C,done2)
    <-done2
    <-done2
  }else{
    for i:=i0;i<i1;i++{
      C.A[i]=A.A[i]-B.A[i]
    }
  }
  done<-true
}
/*
// A-B  (A,B are Matrix)
func Sustract(A,B Matrix)(*Matrix,error){
  if(A.n==B.n&&A.m==B.m){
    out:=NullMatrix(A.m,A.n)
    for i:=0;i<len(A.A);i++{
      out.A[i]=A.A[i]-B.A[i]
    }
    return &out,nil
  }
  return nil,errors.New("The Matrixes don't have the same dimensions")
}*/



func Product(A,B Matrix) *Matrix{
   out:=NullMatrix(A.m,B.n)
  
   if(A.n==B.m){
     
     done:=make(chan bool)
      go multr(A,B,out,1,A.m,1,B.n,1,A.n,done)
      <-done
   }
  return &out
}

const THRESHOLD=100
func multr(A,B,C Matrix,i0,i1,j0,j1,k0,k1 int,done chan <-bool){
  
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
func  Multiplication(A,B Matrix) *Matrix{
  out:=NullMatrix(A.m,B.n)
  
    done:=make(chan bool)
    mult:=make(chan float64)
  
    for i:=1;i<=A.m;i++{
     for k:=1;k<=B.n;k++{
      
      go out.multRowColumn(i,k,A,B,mult)
       out.setCValue(i,k,mult,done)
      <-done
    }
  }
  return &out
} 

// for Matrix multiplication in parallel
func (this *Matrix) multRowColumn(i,k int, A,B Matrix,out chan <-float64){
  var temp float64
  temp=0
  for j:=1;j<=A.n;j++{    
    temp=temp+A.GetValue(i,j)*B.GetValue(j,k)
  }  
   out<-temp
}

//for Matrix multiplication in parallel
func (this *Matrix) setCValue(i,k int, in <- chan  float64, done chan<- bool){ 
     for  {
      temp:=<-in
      this.SetValue(i,k,temp)
      break
    }
    done<-true
}


// Return a Matrix Transpose 
func (this *Matrix) Transpose() *Matrix{
  
  if(this.m==1||this.n==1){    
    c:=this.Copy()
    t:=c.m
    c.m=c.n
    c.n=t;
    return c
  
  }
  out:=NullMatrix(this.n,this.m)
  done:=make(chan bool)
  go this.parallel_Traspose(1,this.m,1,this.n,out,done)
  <-done
  return &out
}

func (this *Matrix) parallel_Traspose(i0,i1,j0,j1 int, res Matrix,done chan<-bool){
  di:=i1-i0
  dj:=j1-j0
  done2:=make(chan bool,THRESHOLD)
  if(di>=dj&&di>=THRESHOLD){
      mi:=i0+di/2
      go this.parallel_Traspose(i0,mi,j0,j1,res,done2)
       this.parallel_Traspose(mi,i1,j0,j1,res,done2)
      <-done2
      <-done2
  }else if (dj>=THRESHOLD){
      mj:=j0+dj/2
      
      go this.parallel_Traspose(i0,i1,j0,mj,res,done2)
       this.parallel_Traspose(i0,i1,mj,i1,res,done2)
      <-done2
      <-done2
  }else{
      for i:=i0;i<=i1;i++{
	for j:=j0;j<=j1;j++{
	  res.SetValue(j,i,this.GetValue(i,j))
	}
      }
  }
  done<-true
}

func (this *Matrix) InverseGauss()(*Matrix, error){
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
	  
	  return nil,errors.New(" Singualr Matrix")
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
  return nil,errors.New(" the Matrix is not Square ")
}

// Return a Inverse of a Square Matrix by LU  Decomposition 
func (this *Matrix)Inverse() (*Matrix,error){
  out:=NullMatrix(this.m,this.n)
  var newOutA []float64
  if(this.n==this.m){  
  l,u:=this.LUDec()
  
  for i:=1;i<=this.m;i++{
    column:=NullMatrix(this.m,1)
    column.SetValue(i,1,1)
    
    z:=l.fSubs(column)
    b:=u.bSubs(*z)   
    newOutA=append(newOutA,b.A[:]...)
    
    }
    
  }else{
    return nil,errors.New(" the Matrix has to be square")
  }
 
  out.A=newOutA
  out=*out.Transpose()
 return &out,nil 
}

func (this *Matrix) PInverse()(*Matrix) {
    if(this.n==this.m){
        _,R:=this.QRDec()
        
        println("temp",this.ToString(),R.Transpose().ToString())
        
        temp1,err:=R.Transpose().GaussElimitation(this.Transpose())
        fmt.Println("Error joder")
        if(err==nil){
        temp2,_:=R.GaussElimitation(temp1)
        return temp2
        }
        
        
        
    }
    return nil
}



//Solve by forward substitution method for L Matrix in Inverse
func (this *Matrix) fSubs(B Matrix)*Matrix{
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

//Solve by back substitution method for a U Matrix in Inverse
func (this *Matrix) bSubs(B  Matrix)*Matrix{
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


// In a Matrix to Matrix with dimensions A (nxm) and B(n1xm1) return a Matrix C(n*n1xm*m1) 
// with a elements Ci=Aij*B 
func KroneckerProduct(A,B Matrix)*Matrix{
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