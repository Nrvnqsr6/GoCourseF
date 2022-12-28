package action

import "encoding/json"

func UnmarshalInput(body []byte) [][]int {
	//resStep := make([][]int, 0)
	resArray := make([][][]int, 0)
	//json.Unmarshal([]byte(body), &resStep)
	json.Unmarshal([]byte(body), &resArray)

	value := make(([][]int), len(resArray))
	for i := 0; i < len(resArray); i++ {
		value[i] = resArray[i][0]
	}

	return value
}

func UnmarshalInputShift(body []byte) ([][]int, []int) {
	resStep := make([][]int, 0)
	resArray := make([][][]int, 0)
	json.Unmarshal([]byte(body), &resStep)
	json.Unmarshal([]byte(body), &resArray)

	step := make([]int, len(resStep))
	for i := 0; i < len(resArray); i++ {
		step[i] = resStep[i][1]
	}

	value := make([]([]int), len(resArray))
	for i := 0; i < len(resArray); i++ {
		value[i] = resArray[i][0]
	}

	return value, step
}
