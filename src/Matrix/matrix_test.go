package Matrix

import (
    "testing"
    )


func TestNullMatrix(t *testing.T){
    
    a:=NullMatrixP(2,2)
    b,_:=FromFile("test/null.txt")
    
    if(!Equal(a,b)){
        t.Errorf("NullMatrix of 2x2 has to be %v like this: ",b.ToString())
    }    
}

func TestRows(t *testing.T){
    a,_:=FromFile("test/null.txt")    
    if(a.GetMRows()!=2){
        t.Errorf(" has to be 2 rows, not ",a.GetMRows())
    }    
}

func TestColumns(t *testing.T){
    a,_:=FromFile("test/null.txt")    
    if(a.GetNColumns()!=2){
        t.Errorf(" has to be 2 columns, not ",a.GetNColumns())
    }    
}

func TestGetValue(t *testing.T){
    a,_:=FromFile("test/a.txt")        
    if(a.GetValue(4,3)!=12){
        t.Errorf(" hast to be 12 not ",a.GetValue(4,3))
    }
}

func TestSetValue(t *testing.T){
    a,_:=FromFile("test/a.txt")
    a.SetValue(2,3,45)
    
    if(a.GetValue(2,3)!=45){
        t.Errorf(" has to be 45 not ",a.GetValue(4,3))
    }
}

func TestCopy(t *testing.T){
   a,_:=FromFile("test/a.txt")
   b:=a.Copy()
   if(!Equal(a,b)){
        t.Errorf(" has to be equal to ",a.ToString())
   }
}

func TestTriangularLower(t *testing.T){
    a,_:=FromFile("test/a.txt")
    b,_:=FromFile("test/tLower.txt")    
    if (!a.TriangularLower()&&b.TriangularLower()){
         t.Errorf(" the first matrix is not a TriangularLower, but a second it is")
    }
}


func TestTriangularUpper(t *testing.T){
    a,_:=FromFile("test/a.txt")
    b,_:=FromFile("test/tUpper.txt")
    if (!a.TriangularUpper()&&b.TriangularUpper()){
        t.Errorf(" the first matrix is not a TriangularUpper, but a second it is")
    }        
}


//TODO FBSubstitution


func TestSumColum(t *testing.T){
    a,_:=FromFile("test/a.txt")
    
    if(real(a.SumColum(1))!=22){
        t.Errorf(" the sum hast to be 22 not ")
    }
}

func TestGetDiagonal(t *testing.T){
    b,_:=FromFile("test/tUpper.txt")
    diag,_:=FromFile("test/bdiagonal.txt")
    diagc,er1:=b.GetDiagonal()
    if(er1!=nil){
        t.Errorf(" erro to get Diagonal ",er1)
    }else{
    if(!Equal(diagc,diag)){
        t.Errorf(" the diagonal has to be ",diagc.ToString())
    }    
        
    }
}


func TestApply(t *testing.T){
    a,_:=FromFile("test/a.txt")
    r,_:=FromFile("test/apply.txt")
    f:=func(a complex128)(complex128){
        return a*a
    }
    b:=a.Apply(f)
    if(!Equal(r,b)){
        t.Errorf("the result has to be ",r.ToString())
    }
}


//TODO TestGaussElimitation


// Arithmetic

func TestScalar(t *testing.T){
    b,_:=FromFile("test/b.txt")
    r,_:=FromFile("test/scalarb.txt")
    
    c:=b.Scalar(complex(10,0))
    
    if(!Equal(c,r)){
        t.Errorf("the result has to be ",r.ToString())
    }
}


func TestSum(t *testing.T){
    a,_:=FromFile("test/a.txt")
    b,_:=FromFile("test/b.txt")
    r,_:=FromFile("test/sum.txt")
    
    c,err:=Sum(a,b)
    if(err!=nil){
        t.Errorf(" A and B are not the same size") 
    }else{
    if(!Equal(c,r)){
        t.Errorf("the result has to be ",r.ToString())
    }
    }
}

func TestSustract(t *testing.T){
    a,_:=FromFile("test/a.txt")
    b,_:=FromFile("test/b.txt")
    r,_:=FromFile("test/sustract.txt")
    
    c,err:=Sustract(a,b)
    if(err!=nil){
        t.Errorf(" A and B are not the same size")
    }else{
        if(!Equal(c,r)){
            t.Errorf("the result has to be ",r.ToString())
        }        
    }
}


