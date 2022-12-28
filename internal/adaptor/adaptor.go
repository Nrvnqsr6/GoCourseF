package adaptor

import (
	"bytes"
	"encoding/json"
	"final/internal/consts"
	"final/internal/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequestForData(task string) ([]byte, error) {

	resp, err := http.Get(consts.URL + task)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

func MakeRequestForResult(body *model.Response) (*model.ResponseSolution, error) {

	bodyData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(consts.SOLUTION, "application/json", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}

	bodyData, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseData model.ResponseSolution
	err = json.Unmarshal(bodyData, &responseData)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", responseData)
	return &responseData, nil
}
