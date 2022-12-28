package action

import (
	"encoding/json"
	"final/internal/adaptor"
	"final/internal/consts"
	"final/internal/model"
)

func solutionT4(input []int) int {
	m := make(map[int]bool, len(input))

	for _, val := range input {
		m[val] = true
	}

	for i := 1; i < len(input); i++ {
		_, ok := m[i]
		if !ok {
			//println(i)
			return i
		}
	}
	println("")
	return -1
}

func AbsentElement() (*model.ResponseSolution, error) {
	bodyData, err := adaptor.MakeRequestForData(consts.ABSENTELEMENT)
	if err != nil {
		print(err)
	}
	input := UnmarshalInput(bodyData)

	res := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = solutionT4(input[i])
	}

	var rawBodyData = json.RawMessage(bodyData)

	response := model.CreateResponse(consts.ABSENTELEMENT, rawBodyData, res)

	var responseSolution *model.ResponseSolution
	responseSolution, err = adaptor.MakeRequestForResult(response)
	if err != nil {
		return nil, err
	}
	return responseSolution, nil
}
