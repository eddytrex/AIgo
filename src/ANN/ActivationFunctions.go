package ANN

import (
        "math/cmplx"
)

func Sigmoid(x complex128) complex128{
    return 1 / (1 + cmplx.Exp(-x))
}

func DSigmoid(x complex128) complex128{
    return (1 / (1 + cmplx.Exp(-x))) * (1 - (1 / (1 + cmplx.Exp(-x))))
}



