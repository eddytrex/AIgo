package Matrix

import(
  "errors"
  "strconv"
  "math/rand"
  //"math"
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


func (this *Matrix)GetMRows()int {
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


//  return  a copy of a Matrix
func (this *Matrix) Copy()(*Matrix){
   out:=NullMatrix(this.m,this.n)
   copy(out.A,this.A)
   return &out
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

//if the Matrix is square get only the main diagonal in a Matrix n*m other is 0
func (this *Matrix) GetDiagonal() (*Matrix,error){
  if(this.n==this.m){
  out:=NullMatrix(this.n,this.m)
  for i:=1;i<=this.m;i++{
      out.SetValue(i,i,this.GetValue(i,i))
      }
  return &out,nil
  }
  return nil,errors.New(" The Matrix is no square")
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


func abs(N float64 )float64 { 
  if(N>=0){
  return N  
  }else{
    return -N
  }
  return 0
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
	    if(tok==scanner.Float||tok==scanner.Int){
	      
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