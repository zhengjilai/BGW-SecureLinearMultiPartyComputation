package poly

import (
	"testing"
	"math/big"
	"fmt"
)

func TestNewPolynomialBigInt(t *testing.T) {
    t.Run("TestNewPolynomialBigInt1", testNewPolynomialBigIntFunc(3,
        []*big.Int{big.NewInt(-78),big.NewInt(4),big.NewInt(71),big.NewInt(7002)},
        big.NewInt(7), false ))
	t.Run("TestNewPolynomialBigInt2", testNewPolynomialBigIntFunc(3,
		[]*big.Int{big.NewInt(-78),big.NewInt(4),big.NewInt(71)},
		big.NewInt(7), true ))
	t.Run("TestNewPolynomialBigInt3", testNewPolynomialBigIntFunc(3,
		[]*big.Int{big.NewInt(-78),big.NewInt(4),big.NewInt(71),big.NewInt(7002)},
		big.NewInt(-100), true ))
}

func testNewPolynomialBigIntFunc(degree int,coefficients []*big.Int, modulus *big.Int, errorExpected bool) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := NewPolynomialBigInt(degree,coefficients,modulus)
		if errorExpected == false && err != nil {
			t.Error(fmt.Sprintf("Error happens when constructing a PolynomialBigInt: %s", err))
		} else if errorExpected == true && err == nil {
			t.Error(fmt.Sprintf("Error happens when constructing a PolynomialBigInt: %s", err))
		}
	}
}

func TestPolynomialBigInt_GetElements(t *testing.T) {
	t.Run("TestPolynomialBigInt_GetElements1", testPolynomialBigInt_GetElementsFunc(3,
		[]*big.Int{big.NewInt(-78),big.NewInt(4),big.NewInt(71),big.NewInt(7002)},
		big.NewInt(7) ))
	t.Run("TestPolynomialBigInt_GetElements2", testPolynomialBigInt_GetElementsFunc(2,
		[]*big.Int{big.NewInt(-787),big.NewInt(40),big.NewInt(7778)},
		big.NewInt(777) ))
}

func testPolynomialBigInt_GetElementsFunc(degree int,coefficients []*big.Int, modulus *big.Int) func(t *testing.T) {
	return func(t *testing.T) {
		newPoly, err := NewPolynomialBigInt(degree,coefficients,modulus)
		if err != nil {
			t.Error(fmt.Sprintf("Error happens when constructing a PolynomialBigInt: %s", err))
		} else{
			t.Log(fmt.Sprintf("Degree: %d", newPoly.GetDegree()))
			t.Log(fmt.Sprintf("Coefficients: %s", newPoly.GetCoefficients()))
			t.Log(fmt.Sprintf("Modulus: %s", newPoly.GetModulus()))
		}
	}
}

func TestPolynomialBigInt_Calculate(t *testing.T) {
	t.Run("TestPolynomialBigInt_Calculate1", testPolynomialBigInt_CalculateFunc(3,
		[]*big.Int{big.NewInt(-78),big.NewInt(4),big.NewInt(71),big.NewInt(7007)},
		big.NewInt(7),
		big.NewInt(10),
		big.NewInt(6)))
	t.Run("TestPolynomialBigInt_Calculate2", testPolynomialBigInt_CalculateFunc(3,
		[]*big.Int{big.NewInt(-78),big.NewInt(4),big.NewInt(71),big.NewInt(7002)},
		big.NewInt(7),
		big.NewInt(10),
		big.NewInt(4)))
}

func testPolynomialBigInt_CalculateFunc(degree int,coefficients []*big.Int,
	modulus *big.Int, x *big.Int, res *big.Int) func(t *testing.T) {
	return func(t *testing.T) {
		newPoly, err := NewPolynomialBigInt(degree,coefficients,modulus)
		if err != nil {
			t.Error(fmt.Sprintf("Error happens when constructing a PolynomialBigInt: %s", err))
		} else{
			feedback, err := newPoly.Calculate(x)
			feedbackValue, ok := feedback.(*big.Int)
			if err != nil {
				t.Error(fmt.Sprintf("Error happens when calculating the result: %s", err))
			} else if !ok {
				t.Error(fmt.Sprintf("Type of the Calculated Result is False, should be %s.","*big.Int"))
			} else if res.Cmp(feedbackValue) != 0  {
			    t.Error(fmt.Sprintf("Calculated Result is False, Result:%s, Expected: %s",feedbackValue,res))
			} else {
				t.Log(fmt.Sprintf("Calculated Result is True, Result:%s, Expected %s",feedbackValue,res))
			}
		}
	}
}