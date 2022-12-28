package action

import (
	"encoding/json"
	"final/internal/adaptor"
	"final/internal/consts"
	"final/internal/model"
)

func solutionT1(input []int, i int) []int {
	for count := 1; count <= i; count++ {
		tmp := input[len(input)-1]
		for n := len(input) - 2; n >= 0; n-- {
			input[n+1] = input[n]
		}
		input[0] = tmp
	}
	return input
}

func CircularShift() (*model.ResponseSolution, error) {
	bodyData, err := adaptor.MakeRequestForData(consts.SHIFT)
	if err != nil {
		print(err)
	}
	input, step := UnmarshalInputShift(bodyData)
	//fmt.Println(input)
	res := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = solutionT1(input[i], step[i])
	}
	var rawBodyData = json.RawMessage(bodyData)
	response := model.CreateResponse(consts.SHIFT, rawBodyData, res)

	var responseSolution *model.ResponseSolution
	responseSolution, err = adaptor.MakeRequestForResult(response)
	if err != nil {
		return nil, err
	}
	return responseSolution, nil
}
