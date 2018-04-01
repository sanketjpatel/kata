package main

import "testing"

type indexChildIndex struct {
	index, childIndex int
}

type indexExpectedResult struct {
	index          int
	expectedResult bool
}

type leafTestCase struct {
	maxHeap         MaxHeap
	expectedResults []indexExpectedResult
}

func TestIsLeaf(t *testing.T) {
	testCases := []leafTestCase{
		{
			MaxHeap{
				5,
				1,
				[]*PointWithDistance{
					{
						2,
						&Point{2, 2},
					},
				},
			},
			[]indexExpectedResult{{0, true}},
		},
		{
			MaxHeap{
				5,
				2,
				[]*PointWithDistance{
					{
						2,
						&Point{2, 2},
					},
					{
						5,
						&Point{2, 1},
					},
				},
			},
			[]indexExpectedResult{{0, false}, {1, true}},
		},
		{
			MaxHeap{
				5,
				3,
				[]*PointWithDistance{
					{
						2,
						&Point{2, 2},
					},
					{
						5,
						&Point{2, 1},
					},
					{
						1,
						&Point{0, 2},
					},
				},
			},
			[]indexExpectedResult{{0, false}, {1, true}, {2, true}},
		},
	}
	for _, heapCases := range testCases {
		for _, testCase := range heapCases.expectedResults {
			expected := testCase.expectedResult
			result := heapCases.maxHeap.isLeaf(testCase.index)
			if result != expected {
				t.Errorf("isLeaf() == %v, expected == %v", result, expected)
			}
		}
	}
}

func TestGetLeftChild(t *testing.T) {
	testCases := []indexChildIndex{
		{0, 1},
		{1, 3},
		{2, 5},
	}

	for _, testCase := range testCases {
		expected := testCase.childIndex
		result := getLeftChild(testCase.index)
		if result != expected {
			t.Errorf("getLeftChild() == %v, expected == %v", result, expected)
		}
	}
}

func TestGetRightChild(t *testing.T) {
	testCases := []indexChildIndex{
		{0, 2},
		{1, 4},
		{2, 6},
	}

	for _, testCase := range testCases {
		expected := testCase.childIndex
		result := getRightChild(testCase.index)
		if result != expected {
			t.Errorf("getRightChild() == %v, expected == %v", result, expected)
		}
	}
}
