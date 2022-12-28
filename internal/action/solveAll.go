package action

import "final/internal/model"

func SolveAndApproveAll() ([]*model.ResponseSolution, error) {
	var res []*model.ResponseSolution = make([]*model.ResponseSolution, 4)
	r, err := CircularShift()
	if err != nil {
		return nil, err
	}
	res[0] = r
	r, err = Subsequence()
	if err != nil {
		return nil, err
	}
	res[1] = r
	r, err = ArrayEntry()
	if err != nil {
		return nil, err
	}
	res[2] = r
	r, err = AbsentElement()
	if err != nil {
		return nil, err
	}
	res[3] = r
	return res, nil
}
