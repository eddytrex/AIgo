package Matrix

import(
  "errors"
  "strconv"
  "math/rand"
  //"math"
  "math/cmplx"
  //"fmt"
  "time"
  "text/scanner"
  "os"
  "bufio"
  //"strings"
 
)

type Matrix struct {
  // m rows and n columns
   m,n int
   //Values of the Matrix
   A  []complex128 
}


func (this *Matrix)GetMRows()int {
  return this.m
}

func (this *Matrix)GetNColumns() int{
  return this.n
  
}

// Return the value in the Matrix position i,j
func (this *Matrix)GetValue(i,j int)complex128{
  i=i-1
  j=j-1
  
  return this.A[i*this.n+j]
  
}

// Set the value (val) in the Matrix position i,j in 
func (this *Matrix)SetValue(i,j int,val complex128){
  i=i-1
  j=j-1
  this.A[i*this.n+j]=val
}


//  return  a copy of a Matrix
func (this *Matrix) Copy()(*Matrix){
   out:=NullMatrixP(this.m,this.n)
   copy(out.A,this.A)
   return out
}

// Return a Matrix of m,n size and random elements 1-10
func RandomMatrix(m,n int)*Matrix{
  out:=NullMatrixP(m,n)
  rand.Seed(time.Now().UTC().UnixNano())
  for i:=1;i<=out.m;i++{
    for j:=1;j<=out.n;j++{
    
      out.SetValue(i,j,complex(rand.Float64()*10,rand.Float64()*10))
      
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
      //out=out+"\t "+strconv.FormatFloat(this.A[i*this.n+j],'f',6,64)
      var impart string 
      impart="+"+strconv.FormatFloat(imag(this.A[i*this.n+j]),'f',6,64)
      if(imag(this.A[i*this.n+j])<0){
       impart=strconv.FormatFloat(imag(this.A[i*this.n+j]),'f',6,64)
      }
      out=out+"\t "+strconv.FormatFloat(real(this.A[i*this.n+j]),'f',6,64)+impart+"i"
<<<<<<< HEAD
    }
    out=out+"\n"
  }
  }
  return out
}

func (this *Matrix)ToStringAbs() string{
  var out string
  out=""
  if(this!=nil){
  for i:=0;i<this.m;i++{
    for j:=0;j<this.n;j++ {            
      out=out+"\t "+strconv.FormatFloat(cmplx.Abs(this.A[i*this.n+j]),'f',6,64)
=======
>>>>>>> c8dd31ca064c801f714c7e09da27a197cb548ff9
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
  out:=NullMatrixP(B.m,1)
  lx:=NullMatrixP(B.m,1)
  ux:=NullMatrixP(B.m,1)
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
    return out,nil
  }
  return nil,errors.New(" The Matrix is no square")
}

func (this *Matrix) SumColum(j int)complex128{
  var out complex128
<<<<<<< HEAD
  out=0  
=======
  out=0
>>>>>>> c8dd31ca064c801f714c7e09da27a197cb548ff9
  for i:=1;i<=this.m;i++{
    out=out+this.GetValue(i,j)
  }
  return out
}

//if the Matrix is square get only the main diagonal in a Matrix n*1
func (this *Matrix) GetDiagonal() (*Matrix,error){
  if(this.n==this.m){
  out:=NullMatrixP(this.n,1)
  for i:=1;i<=this.m;i++{
      out.SetValue(i,1,this.GetValue(i,i))
      }
  return out,nil
  }
  return nil,errors.New(" The Matrix is no square")
}


// Apply the function (f) to all elements of the Matrix (
func (this *Matrix) Apply(f func(complex128)complex128) *Matrix{
  out:=this.Copy()
  done:=make(chan bool,THRESHOLD)
  applyR(0,len(out.A),this,out,f,done)
  <-done
  /*for i:=0;i<len(out.A);i++{
    newVal:=f(out.A[i])
    out.A[i]=newVal
  }*/
  return out
}

func applyR(i0,i1 int,C,out *Matrix,f func(complex128)complex128,done chan<-bool){
  di:=(i1-i0)
  done2:=make(chan bool,THRESHOLD)
  if(di>=THRESHOLD){
    mi:=i0+di/2
    go applyR(i0,mi,C,out,f,done2)
    applyR(mi,i1,C,out,f,done2)
    <-done2
    <-done2
  }else{
    for i:=i0;i<i1;i++{
      out.A[i]=f(C.A[i])
    }
  }
  done<-true 
}


func abs(N complex128 )complex128 { 
  if(cmplx.Abs(N)>=0){
  return N  
  }else{
    return -N
  }
  return 0
}


func FromFile(nameFile string)(*Matrix,error){
    ff,errfile := os.Open(nameFile) 
    
    cout:=make([]complex128,0)
    if(errfile!=nil){      
      return nil, errors.New("Error to open file: nameFile\n ")
    }
    
    f := bufio.NewReader(ff)     
    var s scanner.Scanner
        s.Init(f)        
        s.Whitespace=0
        
        sign:=1.0
        state:=0
        tok:=s.Scan()        
        
        real:=0.0
        img:=0.0
        numb:=0.0
        
         ncolumnlast:=-1
         ncolumn:=0         
         nrow:=0
        
         for (tok!=scanner.EOF){
            
          if (state==0){
                if(s.TokenText()=="-"){
                    sign=-1.0
                    state=1
                }else if(s.TokenText()=="+"){
                    state=1
                }else if(tok==scanner.Float||tok==scanner.Int){
                    t,_:=strconv.ParseFloat(s.TokenText(),64)
                    numb=t                    
                    state=2
                }else if (s.TokenText()=="\n"){
                    
                    if(ncolumnlast!=ncolumn&&ncolumnlast!=-1){
                    
                        return nil,errors.New(" Malformed File, columns don't match ");
                    }
                    
                    ncolumnlast=ncolumn
                    ncolumn=0                
                    
                    nrow++
                }
          }
          
          if(state==1){
                if(tok==scanner.Float||tok==scanner.Int){
                    t,_:=strconv.ParseFloat(s.TokenText(),64)
                    numb=sign*t                    
                    state=2
                }
          }
          
          if(state==2){
              if(tok==scanner.Ident&&s.TokenText()=="i") {
                   img=numb
                   numb=0
                   sign=1.0
                   state=3 
              }else if(s.TokenText()=="-"){
                   sign=-1.0  
                   real=numb
                   state=1
              }else if(s.TokenText()=="+"){
                   sign=1.0
                   real=numb
                   state=1
              }else if(s.TokenText()==" "||s.TokenText()=="\t"){                   
                   if(numb!=0){
                    real=numb                       
                   }                                           
                   sign=1.0   
                   
                   cout=append(cout,complex(real,img))
                   //println(real,img,"i")
                   
                   img=0
                   real=0
                   numb=0
                   
                   ncolumn++
                   
                   state=0
              }else if (s.TokenText()=="\n"){
                    if(numb!=0){
                    real=numb                       
                   }                   
                   sign=1.0
                   
                   cout=append(cout,complex(real,img))
                   //println(real,img,"i")
                   
                   img=0
                   real=0
                   numb=0
                   
                   ncolumn++
                   
                   state=0
                                      
                   if(ncolumnlast!=ncolumn&&ncolumnlast!=-1){
                        
                        return nil,errors.New(" Malformed File, columns don't match ");
                    }
                    ncolumnlast=ncolumn
                    ncolumn=0
                
                   nrow++
              }
          }          
          if(state==3){
              if(s.TokenText()=="-"){
                   sign=-1.0                   
                   numb=0
                   state=1
              }else if(s.TokenText()=="+"){
                   state=1.0                   
                   numb=0
                   state=1
              }else if(s.TokenText()==" "||s.TokenText()=="\t"){
                   
                   if(numb!=0){
                    real=numb                       
                   }                   
                   sign=1.0        
                   
                   cout=append(cout,complex(real,img))                   
                   //println(real,img,"i")
                   
                   img=0
                   real=0
                   numb=0
                   
                   ncolumn++
                   
                   state=0
              }else if (s.TokenText()=="\n"){
                    if(numb!=0){
                    real=numb                       
                   }                   
                   sign=1.0   
                   
                   cout=append(cout,complex(real,img))                   
                   //println(real,img,"i")
                   
                   img=0
                   real=0
                   numb=0
                   
                   ncolumn++                   
                   state=0
                                      
                   if(ncolumnlast!=ncolumn&&ncolumnlast!=-1){
                        
                        return nil,errors.New(" Malformed File, columns don't match ");
                    }
                   ncolumnlast=ncolumn
                   ncolumn=0
                
                   nrow++
              }
          }          
            tok=s.Scan()
            if(tok==scanner.EOF){
              if(numb!=0){
                real=numb                  
              }
              //println(real,img,"i")
              cout=append(cout,complex(real,img))
              
              ncolumn++                            
              nrow++
              state=0
                            
              if(ncolumnlast!=ncolumn&&ncolumnlast!=-1){
                        
                        return nil,errors.New(" Malformed File, columns don't match ");
                    }
              ncolumnlast=ncolumn
              
            }  
         }
         out:=NullMatrixP(nrow,ncolumn)
         out.A=cout         
         return out,nil
}


func (this *Matrix) GaussElimitation(Aum *Matrix)(*Matrix, error){
  if(this.m==Aum.m){
  //if(this.n==this.m&&Aum.m==this.m){
    //Aum:=I(this.n)
    for i:=1;i<=this.m;i++{
       
        
        j:=i
        for k:=i;k<=this.m;k++{
          if (cmplx.Abs(this.GetValue(k,i))>cmplx.Abs(this.GetValue(j,i))){
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



func Equal(A,B *Matrix)(bool){
    if(A.m==B.m&&A.n==B.n){
        done:=make(chan bool)
        go equalArray(0,len(A.A),A,B,done)        
        return <-done
    }
    return false;
}


func equalArray(i0,i1 int,A,B  *Matrix,done chan <-bool){
  di:=(i1-i0)
  out:=true;
  done2:=make(chan bool,THRESHOLD)
  if(di>=THRESHOLD){
    mi:=i0+di/2
    go equalArray(i0,mi,A,B,done2)
    go equalArray(mi,i1,A,B,done2)
    g1:=<-done2
    g2:=<-done2
    if(!g1||!g2){ out=false; }    
    
  }else{
    for i:=i0;i<i1;i++{
        if(real(A.A[i])!=real(B.A[i])||imag(A.A[i])!=imag(B.A[i])){
            out=false  
            break;            
        }
      //C.A[i]=f(A.A[i],B.A[i])
    }
  }
  done<-out
}
