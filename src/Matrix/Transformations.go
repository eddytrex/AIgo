package Matrix
import(
  "errors"
)
func (this *Matrix)HouseholderTrasformation()(*Matrix, error){
    
    if(this.n==1){
        Identity:=I(this.m)
        H,_:=Sustract(*Identity,*Product(*this,*this.Transpose()).Scalar(2.0))
        return H,nil
    }
    return nil,errors.New("the Matrix has to be a Unit Vector");
}
