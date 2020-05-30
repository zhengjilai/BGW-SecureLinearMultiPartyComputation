package poly

import (
	"errors"
)

/**
 * This class implements an Integer polynomial over <i>Zp</i> with single variable.
 *
 * @author 		LoCCS
 * @version		1.0
 */
type PolynomialInt struct{
	Polynomial
}
/**
 * Construct polynomialInt with degree, coefficients and modulus.
 * <p>
 * Number of coefficients must be degree + 1, coefficients[i] contains <i>a<sub>i</sub></i> (0 ≤ <i>i</i> ≤ <i>k</i>).
 *
 * @param degree Degree of the polynomial.
 * @param coefficients	Coefficients of the polynomial.
 * @param modulus Modulus p of the polynomial.
 * @throws IllegalArgumentException If any of the degree, coefficients and modulus is invalid.
 */
func NewPolynomialInt(degree int, coefficients []int, modulus int)(*PolynomialInt,error){
	polyFeedback := new(PolynomialInt)

	if (degree < 0) {
		return nil , errors.New("Invalid polynomial degree, should not be less than 0.")
	}
	if ((coefficients == nil) || len(coefficients) != (degree + 1)){
		return nil , errors.New("Number of polynomial coefficients should be degree + 1.")
	}
    if (modulus < 2){
		return nil , errors.New("Invalid polynomial modulus, should be greater than 1.")
	}
	polyFeedback.degree = degree
	polyFeedback.modulus = modulus
	polyFeedback.coefficients = make([]interface{},degree+1)

	for i:=0; i < degree+1; i++{
		if (coefficients[i] < 0 || coefficients[i] >= modulus ){
			polyFeedback.coefficients[i] = (coefficients[i] % modulus + modulus) % modulus
		} else{
			polyFeedback.coefficients[i] = coefficients[i]
		}
	}
	polyFeedback.Polynomialcal = polyFeedback
	return polyFeedback , nil
}

/**
 * Calculate the results of the polynomial by given value of variable <i>x</i>.
 *
 * @param x The value of variable <i>x</i>.
 * @return The results of the polynomial.
 * @throws IllegalArgumentException If <i>x</i> is invalid.
 */
func (poly *PolynomialInt) Calculate(x interface{}) (interface{}, error){
	xValue, ok := x.(int)
	if (!ok){
		return -1 , errors.New("Invalid type of input, should be big.Int.")
	}

	var feedback int64 = 0
	var pile int64 = 1
	var tmp int64 = 1
	for i := 0; i < poly.degree + 1; i++{
		tmp = int64(poly.coefficients[i].(int)) * pile % int64(poly.modulus.(int))
		feedback = (feedback + tmp) % int64(poly.modulus.(int))
		pile = (pile * int64(xValue)) % int64(poly.modulus.(int))
	}
	return int(feedback), nil
}

