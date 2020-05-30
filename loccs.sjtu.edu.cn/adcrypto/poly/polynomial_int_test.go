package poly
import (
	"testing"
	"fmt"
)

func TestNewPolynomiaInt(t *testing.T) {
	t.Run("TestNewPolynomialInt1", testNewPolynomialIntFunc(3, []int{-78,4,71,7002}, 7, false ))
	t.Run("TestNewPolynomialInt2", testNewPolynomialIntFunc(3, []int{-78,4,71}, 7, true ))
	t.Run("TestNewPolynomialInt3", testNewPolynomialIntFunc(3, []int{-78,4,71,7002}, -100, true ))
}

func testNewPolynomialIntFunc(degree int,coefficients []int, modulus int, errorExpected bool) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := NewPolynomialInt(degree,coefficients,modulus)
		if err != nil && !errorExpected{
			t.Error(fmt.Sprintf("Error happens when constructing a PolynomialInt: %s", err))
		}
	}
}

func TestPolynomialInt_GetElements(t *testing.T) {
	t.Run("TestPolynomialInt64_GetElements1", testPolynomialInt_GetElementsFunc(3, []int{-78,4,71,7002}, 7 ))
	t.Run("TestPolynomialInt64_GetElements2", testPolynomialInt_GetElementsFunc(2,[]int{-787,40,7778},777))
}

func testPolynomialInt_GetElementsFunc(degree int,coefficients []int, modulus int) func(t *testing.T) {
	return func(t *testing.T) {
		newpoly , err := NewPolynomialInt(degree,coefficients,modulus)
		if err != nil {
			t.Error(fmt.Sprintf("Error happens when constructing a PolynomialBigInt: %s", err))
		} else{
			t.Log(fmt.Sprintf("Degree: %d",newpoly.GetDegree()))
			t.Log(fmt.Sprintf("Coefficients: %d",newpoly.GetCoefficients()))
			t.Log(fmt.Sprintf("Modulus: %d",newpoly.GetModulus()))
		}
	}
}

func TestPolynomialInt_Calculate(t *testing.T) {
	t.Run("TestPolynomialInt64_Calculate1", testPolynomialInt_CalculateFunc(3,
		[]int{-78,4,71,7007}, 7,10,6))
	t.Run("TestPolynomialInt64_Calculate1", testPolynomialInt_CalculateFunc(3,
		[]int{-78,4,71,7002}, 7,10,4))
}

func testPolynomialInt_CalculateFunc(degree int,coefficients []int, modulus int, x int, res int) func(t *testing.T) {
	return func(t *testing.T) {
		newPoly, err := NewPolynomialInt(degree,coefficients,modulus)
		if err != nil {
			t.Error(fmt.Sprintf("Error happens when constructing a PolynomialBigInt: %s", err))
		} else{
			feedback, err := newPoly.Calculate(x)
			feedbackValue, ok := feedback.(int)
			if err != nil {
				t.Error(fmt.Sprintf("Error happens when calculating the result: %s", err))
			} else if !ok {
				t.Error(fmt.Sprintf("Type of the Calculated Result is False, should be %s.","int"))
			}else if res != feedback {
				t.Error(fmt.Sprintf("Calculate Result is False, Result:%d ,Expected: %d",feedbackValue,res))
			} else{
				t.Log(fmt.Sprintf("Calculate Result is True, Result:%d ,Expected %d",feedbackValue,res))
			}
		}
	}
}