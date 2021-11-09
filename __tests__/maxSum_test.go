package __tests__

import (
	"testing"

	"github.com/Daviiap/maxsum/graph/model"
	"github.com/Daviiap/maxsum/graph/resolvers"
)

func isArraysEquals(list1 []int, list2 []int) bool {
	isEquals := true

	if len(list1) == len(list2) {
		for i := 0; i < len(list1); i++ {
			if list1[i] != list2[i] {
				isEquals = false
				break
			}
		}
	} else {
		isEquals = false
	}

	return isEquals
}

func TestGetPossibleStartIndices(t *testing.T) {
	tests := [][]int{
		{-1000, -1, -1, -1, -1000},
		{1, 2, 3, 4, -1000},
		{-1000, 1, 2, 3, 4},
		{-1000, 2, -1, 2, -1000},
		{-1000, -2, -1, -2, -1000, -1, -190, -10, -14},
		{-1000, 0, 0, 0, -2, -1, -1, -1, -1, 5, -1},
		{2, 9, -1, 4, -6, 12, -10, 100, -1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	responses := [][]int{{}, {0}, {1}, {1, 3}, {}, {9}, {0, 3, 5, 7}, {}, {8}}

	for i := 0; i < len(tests); i++ {
		test := tests[i]
		expected := responses[i]

		response := resolvers.GetPossibleStartIndices(test)

		if !isArraysEquals(response, expected) {
			t.Error("Erro no teste")
		}

	}
}

func TestGetMaxSumFromNonPositiveList(t *testing.T) {
	tests := [][]int{
		{-1000, -1, -1, -1, -1000},
		{-1000, -2, -1, -2, -1000, -1, -190, -10, -14},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	responses := []model.MaxSumResponse{
		model.MaxSumResponse{Sum: -1, Positions: []int{2}},
		model.MaxSumResponse{Sum: -1, Positions: []int{3}},
		model.MaxSumResponse{Sum: 0, Positions: []int{1}},
	}

	for i := 0; i < len(tests); i++ {
		test := tests[i]
		expected := responses[i]

		response := resolvers.GetMaxSumFromNonPositiveList(test)

		if expected.Sum != response.Sum {
			t.Error("Erro no teste")
		} else if !isArraysEquals(expected.Positions, response.Positions) {
			t.Error("Erro no teste")
		}
	}
}

func TestGetMaxSumFromPositiveList(t *testing.T) {
	{
		tests := [][]int{
			{1, 2, 3, 4, -1000},
			{-1000, 1, 2, 3, 4},
			{-1000, 2, -1, 2, -1000},
			{-1000, 0, 0, 0, -2, -1, -1, -1, -1, 5, -1},
			{2, 9, -1, 4, -6, 12, -10, 100, -1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1},
		}

		responses := []model.MaxSumResponse{
			model.MaxSumResponse{Sum: 10, Positions: []int{1, 2, 3, 4}},
			model.MaxSumResponse{Sum: 10, Positions: []int{2, 3, 4, 5}},
			model.MaxSumResponse{Sum: 3, Positions: []int{2, 3, 4}},
			model.MaxSumResponse{Sum: 5, Positions: []int{10}},
			model.MaxSumResponse{Sum: 110, Positions: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			model.MaxSumResponse{Sum: 1, Positions: []int{9}},
		}

		for i := 0; i < len(tests); i++ {
			test := tests[i]
			expected := responses[i]

			response := resolvers.GetMaxSumFromPositiveList(test, resolvers.GetPossibleStartIndices(test))

			if expected.Sum != response.Sum {
				t.Error("Erro no teste")
			} else if !isArraysEquals(expected.Positions, response.Positions) {
				t.Error("Erro no teste")
			}
		}
	}
}

func TestMaxSumResolver(t *testing.T) {
	{
		tests := [][]int{
			{1, 2, 3, 4, -1000},
			{-1000, 1, 2, 3, 4},
			{-1000, 2, -1, 2, -1000},
			{-1000, 0, 0, 0, -2, -1, -1, -1, -1, 5, -1},
			{2, 9, -1, 4, -6, 12, -10, 100, -1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1},
			{-1000, -1, -1, -1, -1000},
			{-1000, -2, -1, -2, -1000, -1, -190, -10, -14},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}

		responses := []model.MaxSumResponse{
			model.MaxSumResponse{Sum: 10, Positions: []int{1, 2, 3, 4}},
			model.MaxSumResponse{Sum: 10, Positions: []int{2, 3, 4, 5}},
			model.MaxSumResponse{Sum: 3, Positions: []int{2, 3, 4}},
			model.MaxSumResponse{Sum: 5, Positions: []int{10}},
			model.MaxSumResponse{Sum: 110, Positions: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			model.MaxSumResponse{Sum: 1, Positions: []int{9}},
			model.MaxSumResponse{Sum: -1, Positions: []int{2}},
			model.MaxSumResponse{Sum: -1, Positions: []int{3}},
			model.MaxSumResponse{Sum: 0, Positions: []int{1}},
		}

		for i := 0; i < len(tests); i++ {
			test := tests[i]
			expected := responses[i]

			response := resolvers.MaxSumResolver(test)

			if expected.Sum != response.Sum {
				t.Error("Erro no teste")
			} else if !isArraysEquals(expected.Positions, response.Positions) {
				t.Error("Erro no teste")
			}
		}
	}
}
