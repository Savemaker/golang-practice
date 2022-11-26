package main

import (
	"testing"

	"github.com/Savemaker/golang-practice/leetcode"
)

func TestITS(t *testing.T) {
	t.Run("test first array", func(t *testing.T) {
		firstArray := []int{1, 2, 3, 4, 5}
		testArray(firstArray, true, t)
	})

	t.Run("test second array", func(t *testing.T) {
		secondArray := []int{10, 2, 30, 4, 20, 6}
		testArray(secondArray, true, t)
	})

	t.Run("test third array", func(t *testing.T) {
		secondArray := []int{5, 4, 3, 2, 1}
		testArray(secondArray, false, t)
	})

	t.Run("test fourth array", func(t *testing.T) {
		secondArray := []int{0, 4, 2, 1, 0, -1, -3}
		testArray(secondArray, false, t)
	})
}

func testArray(array []int, expected bool, t *testing.T) {
	t.Helper()
	result := leetcode.IncreasingTriplet(array)
	if result != expected {
		t.Errorf("test failed for arr %v", array)
	}
}
