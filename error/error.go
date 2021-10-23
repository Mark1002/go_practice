package error

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// if not convert e to float, it will cause nested recursive to infinite loop
	// because fmt call e.Error()
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func ExecuteCustomError() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
