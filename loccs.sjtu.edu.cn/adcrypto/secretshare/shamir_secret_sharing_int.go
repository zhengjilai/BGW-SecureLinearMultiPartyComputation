package secretshare

import (
	"math/big"
	"errors"
	"crypto/rand"
	"loccs.sjtu.edu.cn/adcrypto/poly"
)

/**
 * The class implements Shamir's secret sharing scheme on int field.
 *
 * @author 		LoCCS
 * @version		1.0
 */

type ShamirSecretSharingInt struct {
	ShamirSecretSharing
}

/**
 * Construct secret sharing scheme with the number of participants and int modulus.
 *
 * @param participantCount The number of participants that share the secret.
 * @param modulus The order the finite field used by the polynomial.
 * @return feedback the newly constructed ShamirSecretSharingBigInt
 * @return error If the number of participants or modulus is invalid.
 */
func NewShamirSecretSharingInt(participantCount int,modulus int) (*ShamirSecretSharingInt, error){
	if (participantCount < 2){
		return nil, errors.New("Invalid participant count. Should be larger than 1.")
	}
	if (modulus <= 2){
		return nil, errors.New("Invalid modulus. Should be larger than 2.")
	} else if( !big.NewInt(int64(modulus)).ProbablyPrime(20)){
		return nil, errors.New("Invalid modulus. Should be prime")
	}
	feedback := new(ShamirSecretSharingInt)
	feedback.participantCount = participantCount
	feedback.modolus = modulus
	feedback.ShamirSecretSharingITF = feedback
	feedback.SecretSharingSchemeITF = &feedback.ShamirSecretSharing
	return feedback, nil
}

/**
 * Generate<i>n</i> int random auxiliary data from each participant.
 *
 * @return Random auxiliary
 */
func (sssi *ShamirSecretSharingInt) GenerateRandomAuxiliary() []interface{}{
	feedback := make([]interface{}, sssi.participantCount)
	for i := 0 ; i < sssi.participantCount; i++{
		tag := false
		for (!tag){
			tmp, _ := rand.Int(rand.Reader,  big.NewInt( int64(sssi.modolus.(int)) ) )
			tag = (tmp.Cmp(big.NewInt(0)) > 0)
			feedback[i] = int(tmp.Int64())
		}
	}
	return feedback
}

/**
 * Get a random <i>k</i>-1 degree polynomial over <i>Zp</i>.
 * <p>
 * <i>a</i><sub>0</sub> is the int secret specified by input parameter,
 * and <i>a</i><sub>1</sub>, <i>a</i><sub>2</sub>, ..., <i>a</i><sub><i>k</i>-1</sub>
 * are chosen randomly in <i>Zp</i>.
 *
 * @param a0 Constant term of the polynomial.
 * @return The polynomial object.
 */
func (sssi *ShamirSecretSharingInt) GetRandomPolynomial(a0 interface{}) poly.PolynomialCalculator{
	degree := sssi.access.(*ThresholdAccessStructure).GetThreshold() - 1
	coeff := make([]int,degree + 1)
	coeff[0] = a0.(int)
	for i := 1 ; i < degree+1; i++{
		tag := false
		for (!tag){
			tmp, _ := rand.Int(rand.Reader,  big.NewInt( int64(sssi.modolus.(int)) ) )
			tag = (tmp.Cmp(big.NewInt(0)) > 0)
			coeff[i] = int(tmp.Int64())
		}
	}
	feedback, _ := poly.NewPolynomialInt(degree,coeff,sssi.modolus.(int))
	return feedback
}

/**
 * Create default auxiliary data (int array) from generate shares.
 * <p>
 * i.e. 1, 2, ..., <i>n</i>
 *
 * @return Default auxiliary data
 */
func (sssi *ShamirSecretSharingInt) CreateDefaultAuxiliary() []interface{}{
	feedback := make([]interface{}, sssi.participantCount)
	for i := 0 ; i < sssi.participantCount; i++{
		feedback[i] = i + 1
	}
	return feedback
}

/**
 * Get a system linear equation with <i>k</i> variables over <i>Zp</i> for calculating secret.
 *
 * @return Linear equation system object(LinearEquationSystemInt).
 */
func (sssi *ShamirSecretSharingInt) GetEquationSystem() poly.LinearEquationSystemCalculator{
	variableCount := sssi.access.(*ThresholdAccessStructure).GetThreshold()
	feedback, _ := poly.NewLinearEquationSystemInt(variableCount,sssi.modolus.(int))
	return feedback
}

/**
 * Restore the coefficients of the equation by the first element of the share.
 * <p>
 * i.e. 1, <i>x</i>, <i>x</i><sup>2</sup>, ... , <i>x</i><sup><i>k</i>-1</sup> mod <i>p</i>
 *
 * @param x The first element of the share.
 * @return The coefficient array(int array).
 */
func (sssi *ShamirSecretSharingInt) GetEquationCoefficients(x interface{}) []interface{}{
	threshold := sssi.access.(*ThresholdAccessStructure).GetThreshold()
	feedback := make([]interface{},threshold)
	pile := 1
	for i := 0; i < threshold; i++{
		feedback[i] = pile
		pile = (pile * x.(int)) % sssi.modolus.(int)
		pile = (pile + sssi.modolus.(int)) % sssi.modolus.(int)
	}
	return feedback
}

/**
 * Check if the type of input element is valid.
 *
 * @param e Element to be checked.
 * @return True if the type of input element is bigInt, otherwise return false.
 */
func (sssi *ShamirSecretSharingInt) checkElement(e interface{}) bool{
	_, ok := e.(int)
	return ok
}