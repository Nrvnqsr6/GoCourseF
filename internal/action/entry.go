package action

import (
	"encoding/json"
	"final/internal/adaptor"
	"final/internal/consts"
	"final/internal/model"
)

func solutionT3(input []int) int {
	m := make(map[int]int, len(input))

	for _, val := range input {
		_, ok := m[val]
		// if v == 2 {
		// 	return val
		// }
		if !ok {
			m[val] = 1
			continue
		}
		m[val]++
	}

	for key := range m {
		if m[key]%2 != 0 {
			return key
		}
	}

	return -1
}

func ArrayEntry() (*model.ResponseSolution, error) {
	bodyData, err := adaptor.MakeRequestForData(consts.ARRAYENTRY)
	if err != nil {
		print(err)
	}
	input := UnmarshalInput(bodyData)
	res := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = solutionT3(input[i])
	}

	//fmt.Print(res)

	var rawBodyData = json.RawMessage(bodyData)
	response := model.CreateResponse(consts.ARRAYENTRY, rawBodyData, res)

	var responseSolution *model.ResponseSolution
	responseSolution, err = adaptor.MakeRequestForResult(response)
	if err != nil {
		return nil, err
	}
	return responseSolution, nil

}
