package Matrix

import(
  "errors"
  "strconv"
  "math/rand"
  "math"
  //"fmt"
  "time"
  "text/scanner"
  "os"
  "bufio"
 
)

type Matrix struct {
  // m rows and n columns
   m,n int
   //Values of the Matrix
   A  []float64 
}


func (this *Matrix)GetNRows()int {
  return this.m
}

func (this *Matrix)GetNColumns() int{
  return this.n
  
}



// Return the value in the Matrix position i,j
func (this *Matrix)GetValue(i,j int)float64{
  i=i-1
  j=j-1
  
  return this.A[i*this.n+j]
  
}

// Set the value (val) in the Matrix position i,j in 
func (this *Matrix)SetValue(i,j int,val float64){
  i=i-1
  j=j-1
  this.A[i*this.n+j]=val
}

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

// return the sum of the main diagonal of the square Matrix
func (this *Matrix)trace() float64{
  var out float64
  out=0
  if(this.m==this.n){
  for i:=1;i<this.m;i++{
    out=out+this.GetValue(i,i)
  }
  }
  return out
}


// return a string with the values of the Matrix
func (this *Matrix)ToString() string{
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




//Return a 1xn Matrix of a Matrix mxn
func (this *Matrix) ScalarRowMatrix(i int,  c float64)(*Matrix){
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

//Get a Matrix m rows and (n-1) columns of a Matrix mxn
func (this *Matrix)MatrixWithoutColumn(j int)*Matrix{
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

//Get a Matrix (m-1)rows and n columns of a Matrix mxn
func (this *Matrix)MatrixWithoutRow(i int)*Matrix{
    out:=NullMatrix(this.m-1,this.n)
    At:=make([]float64,len(this.A))
    copy(At,this.A)
    i=i-1
    At=append(At[:i*this.m],At[(i+1)*this.m:]...)
    out.A=At
    return &out
}


// Return the Matrix adjoint Matrix(this) of the position i,j

func (this *Matrix)AdjMatrix(i,j int) *Matrix{
  out:=this.MatrixWithoutRow(i).MatrixWithoutColumn(j)
  return out
}


// return the determinant of a square Matrix 
// O(n!) I don't think someone will use it
func (this *Matrix)Det_LapaceExpasion()(float64,error){
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
  return 0,errors.New(" the Matrix have to be square")
}


//  return  a copy of a Matrix
func (this *Matrix) Copy()(*Matrix){
   out:=NullMatrix(this.m,this.n)
   copy(out.A,this.A)
   return &out
}



// Return the determinant of a Matrix by LU  Decomposition 
func (this *Matrix) Det_LU()float64{
  _,U:=this.LUDesc()  
  
  var Det float64
  Det=1
  
  for ui:=1;ui<=U.n;ui++{
    uii:=U.GetValue(ui,ui)
    Det=Det*uii
  }
  return Det
}

// substitue Row R in the Matrix
func (this *Matrix) SetRow(i int,R Matrix){
 
 if(R.m==1&&R.n==this.n&&i>0){
 i=i-1  
 
 temp1:=this.A[:i*this.n]
 temp2:=this.A[(i+1)*this.n:]

 temp3:=append(temp1,R.A[:]...)
 this.A=append(temp3,temp2[:]...)  
 
 
 }
}

//  multiply a row of a Matrix  by a number c
func (this *Matrix) ScalarRow(i int, C float64){
  for j:=1;j<=this.n;j++{
    this.SetValue(i,j,C*this.GetValue(i,j))
  }
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

//Multiply a row i by c and adds to a row i 
func (this *Matrix) ScalarRowAndAdd(i0,i int, C float64){
      
      for j:=1;j<=this.n;j++{
	  C:=this.GetValue(i0,j)+C*this.GetValue(i,j) 	  
	  this.SetValue(i0,j,C)
      }      
}



// Return a Inverse of a Square Matrix by LU  Decomposition 
func (this *Matrix)Inverse() (*Matrix,error){
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
    return nil,errors.New(" the Matrix has to be square")
  }
  out.A=newOutA
  out=*out.Transpose()
 return &out,nil 
}

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
  for i:=1;i<=this.m;i++{
    for j:=1;j<=this.n;j++{
      out.SetValue(j,i,this.GetValue(i,j))
    }
  }
  return &out
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



// Verify if the Matrix (this) si Triangular Lower
func (this *Matrix) TriangularLower()bool{
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


// Verify if the Matrix (this) si Triangular Upper
func (this *Matrix) TriangularUpper()bool{
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



// If the Matrix (this) is Triangular Lower or Triangular Upper; return the result of it
//Back and forward substitution
func (this *Matrix) FBSubs(B Matrix)(*Matrix,error){
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
  return nil,errors.New(" The Matrix is no square")
}

func (this *Matrix) SumColum(j int)float64{
  var out float64
  out=0
  for i:=1;i<=this.m;i++{
    out=out+this.GetValue(i,j)
  }
  return out
}

//return a given row of a Matrix in Matrix 1*n
func (this *Matrix) GetRow(i int) *Matrix{
  out:=NullMatrix(1,this.n)
  for j:=1;j<=this.n;j++{
    out.SetValue(1,j,this.GetValue(i,j))
  }
  return &out
}

// return a column of Matrix in a Matrix m*1
func (this *Matrix) GetColumn(j int) *Matrix{
  out:=NullMatrix(this.m,1)
  for i:=1;i<=this.m;i++{
    out.SetValue(j,1,this.GetValue(j,i))
  }
  return &out
}



//if the Matrix is square get only the main diagonal in a Matrix n*m other is 0
func (this *Matrix) GetDiagonal() (*Matrix,error){
  if(this.n==this.m){
  out:=NullMatrix(this.n,this.m)
  for i:=1;i<=this.m;i++{
    for j:=1;j<=this.n;j++{
      out.SetValue(i,j,this.GetValue(i,j))
      }
  }
  return &out,nil
  }
  return nil,errors.New(" The Matrix is no square")
}

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
    sumR(mi,i1,A,B,C,done2)
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
  if(A.n==B.n&&A.m==B.m){
    out:=NullMatrix(A.m,A.n)
    for i:=0;i<len(A.A);i++{
      out.A[i]=A.A[i]-B.A[i]
    }
    return &out,nil
  }
  return nil,errors.New("The Matrixes don't have the same dimensions")
}

// Multiply a Matrix for a scalar   cA
func (this *Matrix) Scalar(c float64)(*Matrix){
  
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
      go out.setCValue(i,k,mult,done)
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



// Apply the function (f) to all elements of the Matrix (
func (this *Matrix) Apply(f func(float64)float64) *Matrix{
  out:=this.Copy()
  for i:=0;i<len(out.A);i++{
    newVal:=f(out.A[i])
    out.A[i]=newVal
  }
  return out
}


// Return a Matrix of m,n size and random elements 1-10
func RandomMatrix(m,n int)*Matrix{
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

func (this *Matrix) AddColumn(Ci Matrix)*Matrix{
  if(this.m==Ci.m){
    out:=NullMatrix(this.m,this.n+Ci.n)
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
    return &out
  }
  return nil
}


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


func abs(N float64 )float64 { 
  if(N>=0){
  return N  
  }else{
    return -N
  }
  return 0
}

func PatternMatrix(i0,i,di float64,f func (x float64)float64 ) *Matrix{
  
  
  rows:=(int)(abs(i0-i)/di)
  
  out:=NullMatrix(rows,2)
  
  for in:=1;in<=rows;in++{
    
    temp:=i0+float64 (in)*di
    
    out.SetValue(in,1,temp)
    
    out.SetValue(in,2,f(temp))
    
  }
  return &out
}


func FromFile(nameFile string)(*Matrix,error){
   
    var er error
    fout:=make([]float64,0)
    
    ff,errfile := os.Open(nameFile) 
    
    if(errfile!=nil){ 
     
      return nil,errfile
    }
    var state int
    state=0;
    var column int
    var columni int
    columni=0
    column=0
    var row int
    row=1
    f := bufio.NewReader(ff) 
    var s scanner.Scanner
	s.Init(f)
	s.Whitespace=1<<'\t'  | 1<<' ';
	tok:=s.Scan()
	 for (tok!=scanner.EOF){
	    if(tok==scanner.Float){
	      
	      svalue:=s.TokenText()
	      fvalue,_:=strconv.ParseFloat(svalue,64)
	      fout=append(fout,fvalue)
	      
	      columni++
	      state=1
	      
	    }else if (tok==10 && state==1){ 	      
	      if(column==0){
		column=columni;
	      }else if (column!=columni){
		er=errors.New(" Malformed File ") 
		break
	      }
	      columni=0
	      row++
	      state=0;
	     
	    }else if(tok==10 && state ==0){
	      er=errors.New(" Malformed File ") 
	      
	      break
	    }else {
	      er=errors.New(" Malformed File ") 
	      break
	    }
	    tok=s.Scan()
	  }
	  ff.Close() 
	  
	  if(er!=nil){return nil,er}
	  out:=NullMatrix(row,column)
	  out.A=fout
	  return &out,nil
}






