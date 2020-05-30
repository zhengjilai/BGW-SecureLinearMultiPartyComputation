package secretshare

import (
	"testing"
	"crypto/rand"
	"fmt"
)


func TestShamirSecretSharingIntProcedure(t *testing.T) {
	participantCount := 16
	threshold := 11
	secret := 218932111
	modulusBigInt,_ := rand.Prime(rand.Reader,30)
	modulus := int(modulusBigInt.Int64())
	t.Run("testShamirSecretSharingBigIntProcedure",
		testShamirSecretSharingIntProcedure(participantCount,threshold,secret,modulus))
}

func testShamirSecretSharingIntProcedure(participantCount int, threshold int, secret int, modulus int) func(t *testing.T) {
	return func(t *testing.T) {
		// constructing ShamirSecretSharingBigInt
		shamirSecretSharingBigInt, err := NewShamirSecretSharingInt(participantCount,modulus)
		if err != nil {t.Error(fmt.Sprintf("Error happens when constructing ShamirSecretSharingBigInt: %s", err))}
		// constructing ThresholdAccessStructure
		threAccessStruct, err := NewThresholdAccessStructure(participantCount,threshold)
		if err != nil {t.Error(fmt.Sprintf("Error happens when constructing ThresholdAccessStructure: %s", err))}
		// adding ThresholdAccessStructure
		err = shamirSecretSharingBigInt.SetAccessStructure(threAccessStruct)
		if err != nil {t.Error(fmt.Sprintf("Error happens when adding ThresholdAccessStructure: %s", err))}
		// generating shares
		//auxi := shamirSecretSharingBigInt.CreateDefaultAuxiliary()
		auxi := shamirSecretSharingBigInt.GenerateRandomAuxiliary()
		shares, err:= shamirSecretSharingBigInt.GenerateShares(secret,auxi)
		if err != nil {t.Error(fmt.Sprintf("Error happens when generating shares: %s", err))}

		// calculating shares
		secretNew, err := shamirSecretSharingBigInt.CalculateSecret(shares)
		if err != nil {t.Error(fmt.Sprintf("Error happens when calculating secret: %s", err))}

		if secretNew.(int) != secret {
			t.Error(fmt.Sprintf("Calculate Result is False, Result:%d ,Expected: %d",secretNew,secret))
		}else{
			t.Log(fmt.Sprintf("Calculate Result is True, Result:%d ,Expected %d",secretNew,secret))
		}
	}
}
