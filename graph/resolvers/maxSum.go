package resolvers

import "github.com/Daviiap/maxsum/graph/model"

func getPossibleStartIndices(list []int) []int {
	var possibleStartIndices []int

	for i := 0; i < len(list); i++ {
		number := list[i]

		if number > 0 {
			if i == 0 {
				possibleStartIndices = append(possibleStartIndices, i)
			} else if list[i-1] <= 0 {
				possibleStartIndices = append(possibleStartIndices, i)
			}
		}
	}

	return possibleStartIndices
}

func getMaxSumFromNonPositiveList(list []int) *model.MaxSumResponse {
	sum := list[0]
	position := 1

	for i := 1; i < len(list); i++ {
		num := list[i]

		if num > sum {
			sum = num
			position = i + 1
		}
	}

	return &model.MaxSumResponse{
		Sum:       sum,
		Positions: []int{position},
	}
}

func getMaxSumFromPositiveList(list []int, possibleStartIndices []int) *model.MaxSumResponse {
	sum := 0
	positions := []int{}

	bestStart := 0
	bestEnd := 0

	for _, startIndex := range possibleStartIndices {
		start := startIndex
		end := 0
		possibleBestSum := 0

		for i := startIndex; i < len(list); i++ {
			number := list[i]
			possibleBestSum += number
			end = i
			if possibleBestSum > sum {
				bestStart = start + 1
				bestEnd = end + 1
				sum = possibleBestSum
			}
		}
	}

	for position := bestStart; position <= bestEnd; position++ {
		positions = append(positions, position)
	}

	return &model.MaxSumResponse{
		Sum:       sum,
		Positions: positions,
	}
}

func MaxSumResolver(list []int) *model.MaxSumResponse {
	var response model.MaxSumResponse

	possibleStartIndices := getPossibleStartIndices(list)

	if len(possibleStartIndices) > 0 {
		response = *getMaxSumFromPositiveList(list, possibleStartIndices)
	} else {
		response = *getMaxSumFromNonPositiveList(list)
	}

	return &response
}
