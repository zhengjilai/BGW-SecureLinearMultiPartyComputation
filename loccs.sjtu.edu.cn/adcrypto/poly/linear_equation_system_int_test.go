package poly

import (
	"testing"
	"fmt"
)

func TestNewLinearEquationSystemInt(t *testing.T) {
	varCount := 6
	leq, err := NewLinearEquationSystemInt(varCount,13)
	if err != nil {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s", err))
	}
	coeTest := [][]int{  {2, 3, 5, 7, 9, 5},
		                 {3, 4, 3, 6, 7, 6},
		                 {7, 8, 3, 6, 2, 3},
		                 {5, 6, 4, 7, 1, 1},
		                 {6, 5, 4, 3, 9, 4},
		                 {4, 2, 9, 5, 7, 9},
	}
	solTest := []int{6, 7, 9, 10, 11, 12}
	for i:=0;i < varCount; i++{
		oneEquation := make([]interface{},varCount)
		for j:=0; j < varCount ; j++ {oneEquation[j] = coeTest[i][j]}
		err := leq.AddEquation(oneEquation,solTest[i])
		if err != nil {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s", err))}
	}
	result, err := leq.Solve()
	if err != nil {t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: %s", err))}

	// true solution
	trueSolution := []int{1,10,2,12,1,8}
	for i:=0;i < varCount;i++ {
		if result[i].(int) != (trueSolution[i]) {
			t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: Solution is wrong"))
			break
		}
	}
	t.Log(fmt.Sprintf("Solution of the LinearEquationSystem: %d",result))
}

func TestNewLinearEquationSystemIntNoSolutions(t *testing.T) {
	varCount := 6
	leq, err := NewLinearEquationSystemInt(varCount,13)
	if err != nil {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s", err))
	}
	coeTest := [][]int{  {2, 3, 5, 7, 9, 5},
		                 {3, 4, 3, 6, 7, 6},
		                 {7, 8, 3, 6, 2, 3},
		                 {5, 6, 4, 7, 1, 1},
		                 {6, 5, 4, 3, 9, 4},
		                  {10, 0, 6, 3, 2, 6},
	}
	solTest := []int{6, 7, 9, 10, 11, 12}
	for i:=0;i < varCount; i++{
		oneEquation := make([]interface{},varCount)
		for j:=0; j < varCount ; j++ {oneEquation[j] = coeTest[i][j]}
		err := leq.AddEquation(oneEquation,solTest[i])
		if err != nil {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s", err))}
	}
	result, err := leq.Solve()
	if err != nil {
		t.Log(fmt.Sprintf("Expected error happens when solving LinearEquationSystem: %s", err))
	} else {t.Error(fmt.Sprintf("Solution of the LinearEquationSystem: %d",result))}
}

func TestNewLinearEquationSystemIntInfiniteSolutions(t *testing.T) {
	varCount := 6
	leq, err := NewLinearEquationSystemInt(varCount,13)
	if err != nil {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s", err))
	}
	coeTest := [][]int{  {2, 3, 5, 7, 9, 5},
		{3, 4, 3, 6, 7, 6},
		{7, 8, 3, 6, 2, 3},
		{5, 6, 4, 7, 1, 1},
		{6, 5, 4, 3, 9, 4},
		{10, 0, 6, 3, 2, 6},
	}
	solTest := []int{6, 7, 9, 10, 11, 4}
	for i:=0;i < varCount; i++{
		oneEquation := make([]interface{},varCount)
		for j:=0; j < varCount ; j++ {oneEquation[j] = coeTest[i][j]}
		err := leq.AddEquation(oneEquation,solTest[i])
		if err != nil {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s", err))}
	}
	result, err := leq.Solve()
	if err != nil {
		t.Log(fmt.Sprintf("Expected error happens when solving LinearEquationSystem: %s", err))
	} else {t.Error(fmt.Sprintf("Solution of the LinearEquationSystem: %d",result))}
}

