package poly

import (
	"testing"
	"fmt"
)

func TestNewLinearEquationSystemInt(t *testing.T) {
	varCount := 6
	leq, error := NewLinearEquationSystemInt(varCount,13)
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s",error))
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
		error := leq.AddEquation(oneEquation,solTest[i])
		if (error != nil) {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s",error))}
	}
	result, error := leq.Solve()
	if (error != nil) {t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: %s",error))}

	// true solution
	trueSolution := []int{1,10,2,12,1,8}
	for i:=0;i < varCount;i++ {
		if (result[i].(int) != (trueSolution[i])){
			t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: Solution is wrong"))
			break
		}
	}
	t.Log(fmt.Sprintf("Solution of the LinearEquationSystem: %d",result))
}

func TestNewLinearEquationSystemIntNoSolutions(t *testing.T) {
	varCount := 6
	leq, error := NewLinearEquationSystemInt(varCount,13)
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s",error))
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
		error := leq.AddEquation(oneEquation,solTest[i])
		if (error != nil) {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s",error))}
	}
	result, error := leq.Solve()
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: %s",error))
	} else {t.Log(fmt.Sprintf("Solution of the LinearEquationSystem: %d",result))}
}

func TestNewLinearEquationSystemIntInfiniteSolutions(t *testing.T) {
	varCount := 6
	leq, error := NewLinearEquationSystemInt(varCount,13)
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when constructing the LinearEquationSystem: %s",error))
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
		error := leq.AddEquation(oneEquation,solTest[i])
		if (error != nil) {t.Error(fmt.Sprintf("Error happens when adding a LinearEquation: %s",error))}
	}
	result, error := leq.Solve()
	if (error != nil) {
		t.Error(fmt.Sprintf("Error happens when solving LinearEquationSystem: %s",error))
	} else {t.Log(fmt.Sprintf("Solution of the LinearEquationSystem: %d",result))}
}

