package poly

import (
	"errors"
	"container/list"
	"math/big"
)

/**
 * This class implements an Integer system of linear equations over <i>Zp</i>.
 *
 * @author 		LoCCS
 * @version		1.0
 */

type LinearEquationSystemInt struct {
	LinearEquationSystem
}

/**
 * Construct a system of Int linear equation with number of variable count and modulus.
 *
 * @param variableCount Number of variables.
 * @param modulus Modules <i>p</i>
 * @return linearESFeedback The constructed LinearEquationSystemInt
 * @return error Whether number of variables or modulus is invalid.
 */
func NewLinearEquationSystemInt(variableCount int ,modulus int)(*LinearEquationSystemInt, error){
	linearESFeedback := new(LinearEquationSystemInt)
	if (variableCount < 1){
		return nil, errors.New("Number of variables should be greater than 0.");
	}
	if (modulus <= 2){
		return nil, errors.New("Modulus should be greater than 2.");
	}
	linearESFeedback.variableCount = variableCount
	linearESFeedback.modulus = modulus
	linearESFeedback.equations = list.New()
	linearESFeedback.LinearEquationSystemITF = linearESFeedback
	return linearESFeedback, nil
}

/**
 * Add a linear equation to the system.
 *
 * @param coefficients Coefficients of the equation.
 * @param constant Constant term of the equation.
 * @return error If coefficients or constant is invalid.
 */

func (les *LinearEquationSystemInt) AddEquation(coefficients []interface{}, constant interface{}) error {
	if ((coefficients == nil) || (len(coefficients) != les.variableCount)){
		return errors.New("Number of coefficients should be equals to Number of variables.")
	}
	for i := 0; i < len(coefficients); i++{
		if (!les.checkElement(coefficients[i])) {
			return errors.New("Invalid type of coefficients, should be int.")
		}
	}
	if ((constant == nil) || !les.checkElement(constant)){
		return errors.New("Invalid type of constant, should be int")
	}

	coefficientsInterface := make([]interface{},les.variableCount)
	for i:=0 ; i < les.variableCount;i++{
		// change every coefficient into zero to modulus-1
		coefficientsInterface[i] = (coefficients[i].(int) % (les.modulus.(int)) + les.modulus.(int)) % les.modulus.(int)
	}
	newLinearEquation,err := NewLinearEquation(coefficientsInterface,constant)

	if (err != nil) {
		return err
	} else{
		les.equations.PushBack(newLinearEquation)
	}
	return nil
}

/**
 * Check if the type of input element is valid.
 *
 * @param e Element to be checked.
 * @return True if the type of input element is int, otherwise return false.
 */
func (les *LinearEquationSystemInt) checkElement(e interface{}) bool{
	_,f := e.(int)
	return f
}

/**
 * Solve system of linear equation.
 *
 * @return The solution if there is single solution exists.
 * And nil if infinitely solutions exist or no solution exists.
 * @return error whether some mistakes happen, such as no solution or infinite solutions
 */
func (les *LinearEquationSystemInt) Solve()([]interface{}, error){
	if (les.equations.Len() < les.variableCount) {
		return nil, errors.New("Linear Equations are not enough for solving")
	}

	var i,j,k int // indexes
	var validTag bool
	var tmp64 int // a tmp variable for changing int to int64

	// copy the LinearEquations
	coeffMatrix := make([]interface{},les.variableCount)
	solMatrix := make([]int64,les.variableCount)
	// int64 representation of modulus
	var modulus int64
	modulus = int64(les.modulus.(int))

	// start from the front of LinearEquations
	countLinearEquations := les.equations.Front()
	for i = 0; i < les.variableCount ; i++{
		if (len(countLinearEquations.Value.(*LinearEquation).coefficients) != les.variableCount){
			return nil, errors.New("Length of Linear Equations doesn't fit the number of Variables")
		}
		tmp := make([]int64, les.variableCount) // one row of coeffMatrix
		for j = 0; j < les.variableCount; j++{
			tmp64, validTag = countLinearEquations.Value.(*LinearEquation).coefficients[j].(int)
			tmp[j] = int64(tmp64)
			if (!validTag) {return nil, errors.New("Invalid type in Linear Equations, should be int")}
		}
		coeffMatrix[i] = tmp
		tmp64,validTag = countLinearEquations.Value.(*LinearEquation).constant.(int)
		if (!validTag) {return nil, errors.New("Invalid type in Linear Equations, should be int")}
		solMatrix[i] = int64(tmp64)
		countLinearEquations = countLinearEquations.Next() // get the next LinearEquation
	}

	// Gauss Elimination
	for i = 0;i < les.variableCount; i++{ // i is the row
		// first let mat[i][i] a non-zero number
		if (coeffMatrix[i].([]int64)[i] == 0){
			for k= i+1; k < les.variableCount; k++{
				if (coeffMatrix[k].([]int64)[i] != 0){
					// change row i and row k
					for j=0;j < les.variableCount;j++{
						tmp := coeffMatrix[k].([]int64)[j]
						coeffMatrix[k].([]int64)[j] = coeffMatrix[i].([]int64)[j]
						coeffMatrix[i].([]int64)[j] = tmp
					}
					// solMatrix should not be forgotten
					tmp := solMatrix[k]
					solMatrix[k] = solMatrix[i]
					solMatrix[i] = tmp
					break
				}
			}
			if k == les.variableCount {
				continue
			} // k == t+1 only if this col has all zero
		}

		invBigInt := big.NewInt(1)
		invBigInt.ModInverse( big.NewInt(coeffMatrix[i].([]int64)[i]),big.NewInt(modulus) ) // inverse of coeff[i][i]
		validTag = invBigInt.IsInt64()
		if (validTag == false) {
			return nil, errors.New("Error happens when calculating an inverse.")
		}
		inv := invBigInt.Int64()
		coeffMatrix[i].([]int64)[i] = (coeffMatrix[i].([]int64)[i] * inv) % modulus // make coeff[i][i] = 1
		// for row i , multiply every number with inv
		for j = i+1; j < les.variableCount; j++{
			coeffMatrix[i].([]int64)[j] = (coeffMatrix[i].([]int64)[j] * inv) % modulus
		}
		// solMatrix of row i should not be forgotten
		solMatrix[i] = (solMatrix[i] * inv) % modulus
		// for row below i, substract to make coe[k][i] = 0
		for k= i+1; k < les.variableCount; k++ {
			multiple := coeffMatrix[k].([]int64)[i]// the multiple of row k substract
			for j = i; j < les.variableCount; j++{
				tmp := (multiple * coeffMatrix[i].([]int64)[j]) % modulus
				coeffMatrix[k].([]int64)[j] = (coeffMatrix[k].([]int64)[j] - tmp) % modulus
			}
			// solMatrix should also be refreshed
			tmp := (multiple * solMatrix[i]) % modulus
			solMatrix[k] = (solMatrix[k] - tmp) % modulus
		}
	}

	// judge whether the equation system has zero/infinite solutions
	for i =0 ; i < les.variableCount; i++{
		// if not all numbers in the last row is azero
		if (coeffMatrix[les.variableCount-1].([]int64)[i] != 0) {break;}
		// all-zero processing
		if (i== les.variableCount-1 && solMatrix[les.variableCount-1] == 0){
			return nil, errors.New("Infinite solutions for this LinearEquationSystem.")
		}
		if (i== les.variableCount-1 && solMatrix[les.variableCount-1] != 0){
			return nil, errors.New("No solutions for this LinearEquationSystem.")
		}
	}

	//fmt.Println(solMatrix)
	feedback := make([]interface{},les.variableCount) // the solution vector
	feedback[les.variableCount - 1] = solMatrix[les.variableCount - 1]
	// back calculation from the t+1 row
	for i = les.variableCount - 2; i >= 0; i-- {
		tmp := int64(0)
		for j = i + 1; j < les.variableCount; j++{
			tmp2 := (feedback[j].(int64) * coeffMatrix[i].([]int64)[j]) % modulus
			tmp += tmp2
		}
		tmp3 := ((solMatrix[i] - tmp) % modulus + modulus) % modulus
		invBigInt := big.NewInt(1)
		invBigInt.ModInverse( big.NewInt(coeffMatrix[i].([]int64)[i]),big.NewInt(modulus) ) // inverse of coeff[i][i]
		validTag = invBigInt.IsInt64()
		if (validTag == false) {
			return nil, errors.New("Error happens when calculating an inverse.")
		}
		inv2 := invBigInt.Int64()
		feedback[i] = ((tmp3 * inv2) % modulus + modulus) % modulus
	}
	for i = 0; i < len(feedback); i++{
		feedback[i] = int( (feedback[i].(int64) % modulus + modulus) % modulus )
	}
	return feedback, nil
}
