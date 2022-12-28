package action

import (
	"encoding/json"
	"final/internal/adaptor"
	"final/internal/consts"
	"final/internal/model"
)

func solutionT2(input []int) int {
	m := make(map[int]bool, len(input))

	for _, val := range input {
		if m[val] == true {
			return 0
		}

		m[val] = true
	}

	for i := 1; i < len(input); i++ {
		_, ok := m[i]
		// print(i)
		// print(" ")
		if !ok {
			//println(i)
			return 0
		}
	}
	//println(1)
	println("")
	return 1
}

func Subsequence() (*model.ResponseSolution, error) {
	bodyData, err := adaptor.MakeRequestForData(consts.SUBSEQUENCE)
	if err != nil {
		print(err)
	}
	input := UnmarshalInput(bodyData)
	//fmt.Println(input)
	res := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = solutionT2(input[i])
	}
	var rawBodyData = json.RawMessage(bodyData)

	response := model.CreateResponse(consts.SUBSEQUENCE, rawBodyData, res)

	var responseSolution *model.ResponseSolution
	responseSolution, err = adaptor.MakeRequestForResult(response)
	if err != nil {
		return nil, err
	}
	return responseSolution, nil
}
