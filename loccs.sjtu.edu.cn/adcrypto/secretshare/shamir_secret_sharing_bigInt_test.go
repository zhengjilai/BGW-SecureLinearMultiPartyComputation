package secretshare

import (
	"testing"
	"fmt"
	"math/big"
	"crypto/rand"
)

func TestShamirSecretSharingBigIntProcedure(t *testing.T) {
	participantCount := 16
	threshold := 11
	secret := big.NewInt(381903098103891)
	modulus,_ := rand.Prime(rand.Reader,60)
	t.Run("testShamirSecretSharingBigIntProcedure",
		testShamirSecretSharingBigIntProcedure(participantCount,threshold,secret,modulus))
}

func testShamirSecretSharingBigIntProcedure(participantCount int, threshold int, secret *big.Int, modulus *big.Int) func(t *testing.T) {
	return func(t *testing.T) {
		// constructing ShamirSecretSharingBigInt
		shamirSecretSharingBigInt, err := NewShamirSecretSharingBigInt(participantCount,modulus)
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

		if secretNew.(*big.Int).Cmp(secret) != 0 {
			t.Error(fmt.Sprintf("Calculate Result is False, Result:%s ,Expected: %s",secretNew,secret))
		}else{
			t.Log(fmt.Sprintf("Calculate Result is True, Result:%s ,Expected %s",secretNew,secret))
		}
    }
}




