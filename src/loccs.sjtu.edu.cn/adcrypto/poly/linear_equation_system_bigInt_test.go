package poly

import (
	"testing"
	"math/big"
	"fmt"
)

func TestNewLinearEquationSystemBigInt(t *testing.T) {
	varCount := 6
    leq, error := NewLinearEquationSystemBigInt(varCount,big.NewInt(13))
    if (error != nil) {
    	t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s",error))
    }
	coeTest := [][]*big.Int{ {big.NewInt(2), big.NewInt(3),big.NewInt(5),big.NewInt(7), big.NewInt(9),big.NewInt(5)},
		{big.NewInt(3), big.NewInt(4),big.NewInt(3),big.NewInt(6), big.NewInt(7),big.NewInt(6)},
		{big.NewInt(7), big.NewInt(8),big.NewInt(3),big.NewInt(6), big.NewInt(2),big.NewInt(3)},
		{big.NewInt(5), big.NewInt(6),big.NewInt(4),big.NewInt(7), big.NewInt(1),big.NewInt(1)},
		{big.NewInt(6), big.NewInt(5),big.NewInt(4),big.NewInt(3), big.NewInt(9),big.NewInt(4)},
		{big.NewInt(4), big.NewInt(2),big.NewInt(9),big.NewInt(5), big.NewInt(7),big.NewInt(9)},
	}
	solTest := []*big.Int{big.NewInt(6), big.NewInt(7),big.NewInt(9),big.NewInt(10), big.NewInt(11),big.NewInt(12)}
	for i:=0;i < varCount; i++{
		 oneEquation := make([]interface{},varCount)
		 for j:=0; j < varCount ; j++ {oneEquation[j] = coeTest[i][j]}
         error := leq.AddEquation(oneEquation,solTest[i])
         if (error != nil) {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s",error))}
	}
	result, error := leq.Solve()
	if (error != nil) {t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: %s",error))}

	// true solution
	trueSolution := []*big.Int{big.NewInt(1),big.NewInt(10),big.NewInt(2),big.NewInt(12),big.NewInt(1),big.NewInt(8)}
	for i:=0;i < varCount;i++ {
		if (result[i].(*big.Int).Cmp(trueSolution[i]) != 0){
			t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: Solution is wrong"))
			break
		}
	}
	t.Log(fmt.Sprintf("Solution of the LinearEquationSystem: %s",result))
}

func TestNewLinearEquationSystemBigIntNoSolutions(t *testing.T) {
	varCount := 6
	leq, error := NewLinearEquationSystemBigInt(varCount,big.NewInt(13))
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s",error))
	}
	coeTest := [][]*big.Int{ {big.NewInt(2), big.NewInt(3),big.NewInt(5),big.NewInt(7), big.NewInt(9),big.NewInt(5)},
		{big.NewInt(3), big.NewInt(4),big.NewInt(3),big.NewInt(6), big.NewInt(7),big.NewInt(6)},
		{big.NewInt(7), big.NewInt(8),big.NewInt(3),big.NewInt(6), big.NewInt(2),big.NewInt(3)},
		{big.NewInt(5), big.NewInt(6),big.NewInt(4),big.NewInt(7), big.NewInt(1),big.NewInt(1)},
		{big.NewInt(6), big.NewInt(5),big.NewInt(4),big.NewInt(3), big.NewInt(9),big.NewInt(4)},
		{big.NewInt(10), big.NewInt(0),big.NewInt(6),big.NewInt(3), big.NewInt(2),big.NewInt(6)},
	}
	solTest := []*big.Int{big.NewInt(6), big.NewInt(7),big.NewInt(9),big.NewInt(10), big.NewInt(11),big.NewInt(12)}
	for i:=0;i < varCount; i++{
		oneEquation := make([]interface{},varCount)
		for j:=0; j < varCount ; j++ {oneEquation[j] = coeTest[i][j]}
		error := leq.AddEquation(oneEquation,solTest[i])
		if (error != nil) {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s",error))}
	}
	result, error := leq.Solve()
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: %s",error))
	} else {t.Log(fmt.Sprintf("Solution of the LinearEquationSystem: %s",result))}
}

func TestNewLinearEquationSystemBigIntInfiniteSolutions(t *testing.T) {
	varCount := 6
	leq, error := NewLinearEquationSystemBigInt(varCount,big.NewInt(13))
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s",error))
	}
	coeTest := [][]*big.Int{ {big.NewInt(2), big.NewInt(3),big.NewInt(5),big.NewInt(7), big.NewInt(9),big.NewInt(5)},
		{big.NewInt(3), big.NewInt(4),big.NewInt(3),big.NewInt(6), big.NewInt(7),big.NewInt(6)},
		{big.NewInt(7), big.NewInt(8),big.NewInt(3),big.NewInt(6), big.NewInt(2),big.NewInt(3)},
		{big.NewInt(5), big.NewInt(6),big.NewInt(4),big.NewInt(7), big.NewInt(1),big.NewInt(1)},
		{big.NewInt(6), big.NewInt(5),big.NewInt(4),big.NewInt(3), big.NewInt(9),big.NewInt(4)},
		{big.NewInt(10), big.NewInt(0),big.NewInt(6),big.NewInt(3), big.NewInt(2),big.NewInt(6)},
	}
	solTest := []*big.Int{big.NewInt(6), big.NewInt(7),big.NewInt(9),big.NewInt(10), big.NewInt(11),big.NewInt(4)}
	for i:=0;i < varCount; i++{
		oneEquation := make([]interface{},varCount)
		for j:=0; j < varCount ; j++ {oneEquation[j] = coeTest[i][j]}
		error := leq.AddEquation(oneEquation,solTest[i])
		if (error != nil) {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s",error))}
	}
	result, error := leq.Solve()
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: %s",error))
	} else {t.Log(fmt.Sprintf("Solution of the LinearEquationSystem: %s",result))}
}