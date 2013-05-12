package Matrix
import(
  "errors"
)
func (this *Matrix)HouseholderTrasformation()(*Matrix, error){
    
    if(this.n==1){
        Identity:=I(this.m)
        H,_:=Sustract(Identity,Product(this,this.Transpose()).Scalar(2.0))
        return H,nil
    }
    return nil,errors.New("the Matrix has to be a Column Vector");
}

func (this *Matrix)Orthogonalized(j int,xj *Matrix)(*Matrix){
    rj:=Product(this.GetColumn(j).Transpose(),xj)
    qj,_:=Sustract(xj,Product(this.GetColumn(j),rj))
    return qj
}
