package mpc

import (
	"errors"
	"math/big"
	"loccs.sjtu.edu.cn/adcrypto/secretshare"
)

/**
 * This class implements an Int secure multi-party linear function computation.
 *
 * @author 		LoCCS
 * @version		1.0
 */
type LinearMultipartyComputationInt struct {
	LinearMultipartyComputation
}

/**
 * Construct linear function MPC scheme with number of participants, threshold and the ID of the participant.
 * <p>
 * The threshold is the max number of semi-honest adversaries, should be less than <i>n</i>/2.
 *
 * @param id ID of this participant.
 * @param participantCount Number of participants.
 * @param threshold Threshold <i>t</i>.
 * @return feedback the constructed  LinearMultipartyComputationBigInt
 * @return error IllegalArgumentException If any of ID, participantCount or threshold is invalid.
 */
func NewLinearMultipartyComputationInt(id int, participantCount int, threshold int)(*LinearMultipartyComputationInt,error){
	if (participantCount < 3){
		return nil, errors.New("Invalid participant count. Should be larger than 2.")
	}
	if (id < 0 || (id >= participantCount)){
		return nil, errors.New("Invalid id, should be between 0 and participantCount-1")
	}
	if (threshold > (participantCount / 2)){
		return nil, errors.New("Threshold should never greater than 1/2 of the participant count.")
	}
	feedback := new(LinearMultipartyComputationInt)
	feedback.id = id
	feedback.participantCount = participantCount
	feedback.threshold = threshold
	feedback.linearMultipartyComputationCalculator = feedback
	feedback.receivedInputs = make([]interface{},participantCount)
	feedback.receivedOutputs = map[int]interface{} {}
	return feedback, nil
}

/**
 * Get a Int Shamir's secret sharing object with the number of participants and the modulus.
 *
 * @param participantCount The number of the participants.
 * @param modulus The modulus <i>p</i>.
 * @return The proper Shamir's secret sharing scheme object.
 * @return error IllegalArgumentException If the number of participants or the modulus is invalid.
 */
func (lmpcb *LinearMultipartyComputationInt)getSecretSharing(participantCount int, modulus interface{})(secretshare.
ShamirSecretSharingInterface,error){
	modulusValue ,ok := modulus.(int)
	if (!ok) {return nil, errors.New("Invalid type of modulus.")}
	return secretshare.NewShamirSecretSharingInt(participantCount,modulusValue)
}

/**
 * Generate a proper Int modulus for Shamir's secret sharing from the coefficients of the linear function and the max value of the secret.
 *
 * @param coefficients The coefficients of the linear function.
 * @param max Max value of a secret.
 * @return The modulus for Shamir's secret sharing scheme.
 * @return error IllegalArgumentException If the coefficients or the max value is invalid.
 */
func (lmpcb *LinearMultipartyComputationInt) generateModulus(coefficients []interface{}, max interface{})(interface{},error){
	var pile int64 = 0  // probably max value of the sum
	for i := 0; i< len(coefficients); i++{
		pile += (int64(coefficients[i].(int)) * int64(max.(int)))
	}
	modulus := -1
	modulusList := []int{53,401,1039,7211,38923,326203,1102693,4131109,11347837,31869857,
	             54578221, 132291191, 839999029,2147483647} // a list of prime that can be represented with int
	for i := 0; i < len(modulusList) ; i++{
		if (pile < int64(modulusList[i])){ modulus = modulusList[i] }
	}
    if (modulus == -1){
    	return modulus,errors.New("Probable maximum sum is too great, cannot generate an enough modulus")
	}
	return modulus, nil
}

/**
 * Generate output during the output stage.
 *
 * @return The output value.
 */
func (lmpcb *LinearMultipartyComputationInt) generateOutputImpl() interface{}{
	modulus := int64(lmpcb.GetModulus().(int))
	var pile int64 = 0
	for i := 0; i < lmpcb.participantCount; i++{
		pile += (int64(lmpcb.coefficients[i].(int)) * int64(lmpcb.receivedInputs[i].(int)))
		pile = (pile % modulus + modulus) % modulus
	}
	return int(pile)
}
/**
 * Abstract method of checking if the coefficients and the modulus are both valid.
 *
 * @param coefficients Coefficients of the linear function.
 * @param modulus Modulus of the Shamir's scheme.
 * @throws IllegalArgumentException if the coefficients or the modulus is invalid.
 */
func (lmpcb *LinearMultipartyComputationInt)checkCoefficientsAndModulus(coefficients []interface{}, modulus interface{}) error{
	if (modulus.(int) <= 2){
		return errors.New("Modulus should be greater than 2.")
	}
	if(!big.NewInt(int64(modulus.(int))).ProbablyPrime(20)){
		return errors.New("Modulus is not a Prime.")
	}
	for i := 0; i < len(coefficients); i++{
		if (coefficients[i].(int) < 0 || coefficients[i].(int) >= modulus.(int) ){
			return errors.New("One coefficient is too great or too tiny.")
		}
	}
	return nil
}

/**
 * Get Int of value 1.
 *
 * @return Int of value 1.
 */
func (lmpcb *LinearMultipartyComputationInt) getElementOne() interface{}{
	return 1
}
/**
 * Check if the type of input element is valid.
 *
 * @param e Element to be checked.
 * @return True if the type of input element is Int, otherwise return false.
 */
func (lmpcb *LinearMultipartyComputationInt) checkElement(e interface{}) bool{
	_, ok := e.(int)
	return ok
}