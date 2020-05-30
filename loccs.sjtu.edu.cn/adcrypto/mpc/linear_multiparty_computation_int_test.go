package mpc

import (
	"fmt"
	"testing"
	"math/big"
	"crypto/rand"
)

func TestNewLinearMultipartyComputationIntProcedure(t *testing.T) {
	participantCount := 15
	threshold := 6
	mpc := make([]*LinearMultipartyComputationInt, participantCount) // mpc class for every party
	max := big.NewInt(0)
	max.SetString("1000000",10)  // the max probable number for secret
	secret := make([]int,participantCount)  // secrets
	coeffcients := make([]interface{}, participantCount)  // coefficients
	var err error
	computeFrom := []int{1,4,6,8,11,14} // computing parties except zero itself

	// initialize secrets and coefficients
	for i := 0; i < participantCount; i++{
		se ,_ := rand.Int(rand.Reader,max)
		secret[i] = int(se.Int64())
		coe ,_ := rand.Int(rand.Reader,big.NewInt(50))
		coeffcients[i] = int(coe.Int64())
	}

	// constrcut class mpcs
	for i := 0; i < participantCount; i++{
		mpc[i],err = NewLinearMultipartyComputationInt(i,participantCount,threshold)
		if err != nil {t.Error(fmt.Sprintf("Error happens when constructing LinearMultipartyComputationInt: %s", err))}
	}

	modulus := 1
	for i := 0 ; i < participantCount; i++{
		// initialize mpcs
		if i == 0 {
			err = mpc[i].InitializeWithMaxValue(coeffcients,int(max.Int64()))
			if err != nil {t.Error(fmt.Sprintf("Error happens when initializing mpcs: %s", err))}
			modulus = mpc[i].GetModulus().(int)
		}else{
			err = mpc[i].InitializeWithModulus(coeffcients,modulus)
			if err != nil {t.Error(fmt.Sprintf("Error happens when initializing mpcs: %s", err))}
		}
	}

	// generate auxiliary
	auxi, err := mpc[0].GenerateInputAuxiliary()
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
		outputs[i],err = mpc[i].GenerateOutput()
		if err != nil {t.Error(fmt.Sprintf("Error happens when generating outputs: %s", err))}
	}

	// mpc[0] add other outputs (number: threshold)
	for _,from := range computeFrom{
		_ = mpc[0].AddReceivedOutput(from,outputs[from])
	}

	// mpc[0] calculate the final result
	calculatedResult,err := mpc[0].Compute()
	if err != nil {t.Error(fmt.Sprintf("Error happens when calculating the final result: %s", err))}

	// calculate the true result(never do this in a real mpc procedure)
	pile := int64(0)
	for i:=0; i<participantCount;i++{
		pile += int64(coeffcients[i].(int)) * int64(secret[i])
	}

	// test whether the mpc-calculated result is true
	if int64(calculatedResult.(int)) != pile {
		t.Error(fmt.Sprintf("Calculate Result is False, Result:%d ,Expected: %d",calculatedResult,pile))
	}else{
		t.Log(fmt.Sprintf("Calculate Result is True, Result:%d ,Expected %d",calculatedResult,pile))
	}

}