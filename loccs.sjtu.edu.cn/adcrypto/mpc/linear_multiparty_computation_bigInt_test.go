package mpc

import (
	"testing"
	"fmt"
	"math/big"
	"crypto/rand"
)

func TestNewLinearMultipartyComputationBigIntProcedure(t *testing.T) {
	participantCount := 16
	threshold := 7
	mpc := make([]*LinearMultipartyComputationBigInt, participantCount) // mpc class for every party
	max := big.NewInt(0)
	max.SetString("100000000000000000",10)  // the max probable number for secret
	secret := make([]*big.Int,participantCount)  // secrets
	coeffcients := make([]interface{}, participantCount)  // coefficients
	var err error
	computeFrom := []int{1,4,6,8,10,11,15} // computing parties except zero itself

	// initialize secrets and coefficients
	for i := 0; i < participantCount; i++{
		secret[i],err = rand.Int(rand.Reader,max)
		coeffcients[i],err = rand.Int(rand.Reader,big.NewInt(1000))
	}

    // constrcut class mpcs
	for i := 0; i < participantCount; i++{
		mpc[i],err = NewLinearMultipartyComputationBigInt(i,participantCount,threshold)
		if err != nil {t.Error(fmt.Sprintf("Error happens when constructing LinearMultipartyComputationBigInt: %s", err))}
	}

	modulus := big.NewInt(1)
	for i := 0 ; i < participantCount; i++{
		// initialize mpcs
		if i == 0 {
			err = mpc[i].InitializeWithMaxValue(coeffcients,max)
			if err != nil {t.Error(fmt.Sprintf("Error happens when initializing mpcs: %s", err))}
			modulus = mpc[i].GetModulus().(*big.Int)
		}else{
			err = mpc[i].InitializeWithModulus(coeffcients,modulus)
			if err != nil {t.Error(fmt.Sprintf("Error happens when initializing mpcs: %s", err))}
		}
	}

	// generate auxiliary
	auxi,err := mpc[0].GenerateInputAuxiliary()
	if err != nil {t.Error(fmt.Sprintf("Error happens when constructing Auxiliary: %s", err))}

	for i := 0 ; i <participantCount; i++{
		// every party generate inputs
		inputs, err := mpc[i].GenerateInputs(secret[i],auxi)
		if err != nil {t.Error(fmt.Sprintf("Error happens when generating inputs: %s", err))}
		// every party receive and add inputs
		for j := 0; j < participantCount ; j++{
			err = mpc[j].AddReceivedInput(i,inputs[j])
			if err != nil {t.Error(fmt.Sprintf("Error happens when adding received inputs: %s", err))}
		}
	}

	// every party generate outputs
	outputs := make([]interface{},participantCount)
	for i := 0; i<participantCount; i++{
		outputs[i], err = mpc[i].GenerateOutput()
		if err != nil {t.Error(fmt.Sprintf("Error happens when generating outputs: %s", err))}
	}

	// mpc[0] add other outputs(number: threshold)
	for _,from := range computeFrom{
		err = mpc[0].AddReceivedOutput(from,outputs[from])
		if err != nil {
			t.Error(fmt.Sprintf("Error happens after adding received output: %s", err))
		}
	}

	// mpc[0] calculate the final result
	calculatedResult,err := mpc[0].Compute()
	if err != nil {t.Error(fmt.Sprintf("Error happens when calculating the final result: %s", err))}

	// calculate the true result(never do this in a real mpc procedure)
	pile := big.NewInt(0)
	for i:=0; i<participantCount;i++{
		tmp := big.NewInt(0)
		tmp.Mul(coeffcients[i].(*big.Int), secret[i])
		pile.Add(pile,tmp)
	}

	// test whether the mpc-calculated result is true
	if calculatedResult.(*big.Int).Cmp(pile) != 0 {
		t.Error(fmt.Sprintf("Calculate Result is False, Result:%s ,Expected: %s",calculatedResult,pile))
	}else{
		t.Log(fmt.Sprintf("Calculate Result is True, Result:%s ,Expected %s",calculatedResult,pile))
	}

}