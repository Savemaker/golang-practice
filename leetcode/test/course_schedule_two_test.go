package main

import (
	"reflect"
	"testing"

	"github.com/Savemaker/golang-practice/leetcode"
)

func TestGraphCreatesFine(t *testing.T) {

	//[[1,0],[2,0],[3,1],[3,2]]

	result := leetcode.CreateGraph([][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, 4)
	expected := make(map[int][]int)
	expected[0] = []int{}
	expected[1] = []int{0}
	expected[2] = []int{0}
	expected[3] = []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSecondGraphCreatesFine(t *testing.T) {

	//[[1,6],[6,3],[3,5],[5,0],[2,4],[4,1]]

	result := leetcode.CreateGraph([][]int{{1, 6}, {6, 3}, {3, 5}, {5, 0}, {2, 4}, {4, 1}}, 7)
	expected := make(map[int][]int)
	expected[0] = []int{}
	expected[1] = []int{0}
	expected[2] = []int{0}
	expected[3] = []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestFirstFindOrderReturnExpected(t *testing.T) {
	expectedOption1 := []int{0, 1, 2, 3}
	expectedOption2 := []int{0, 2, 1, 3}

	result := leetcode.FindOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})

	if !reflect.DeepEqual(expectedOption1, result) && !reflect.DeepEqual(expectedOption2, result) {
		t.Errorf("Expected %v or %v but got %v", expectedOption1, expectedOption2, result)
	}
}

func TestSecondFindOrderReturnExpected(t *testing.T) {
	expectedOption := []int{0, 5, 3, 6, 1, 4, 2}

	result := leetcode.FindOrder(7, [][]int{{1, 6}, {6, 3}, {3, 5}, {5, 0}, {2, 4}, {4, 1}})

	if !reflect.DeepEqual(expectedOption, result) {
		t.Errorf("Expected %v but got %v", expectedOption, result)
	}
}

func TestWeirdFindOrderReturnExpected(t *testing.T) {
	expectedOption := []int{0}

	result := leetcode.FindOrder(1, [][]int{})

	if !reflect.DeepEqual(expectedOption, result) {
		t.Errorf("Expected %v but got %v", expectedOption, result)
	}
}
