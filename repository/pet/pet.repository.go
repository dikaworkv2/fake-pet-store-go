package pet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fakestore_go/entity"
	"fmt"
	"io"
	"net/http"
)

type Repository struct {
	rootURl string
	httpCli *http.Client
}

func New(httpCli *http.Client) *Repository {
	return &Repository{
		rootURl: "https://petstore.swagger.io/v2",
		httpCli: httpCli,
	}
}

func (r *Repository) InsertNewPet(req entity.Pet) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/pet", r.rootURl)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := r.httpCli.Do(httpReq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return map[string]interface{}{
			"msg": "cannot add new pet",
		}, errors.New("cannot add new pet")
	}
	defer resp.Body.Close()
	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Create a map to hold the parsed JSON data
	var responseData map[string]interface{}

	// Unmarshal the JSON response into the map
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (r *Repository) GetPetByID(petID int64) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/pet/%d", r.rootURl, petID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return map[string]interface{}{
			"msg": "not found",
		}, errors.New("pet not found")
	}
	defer resp.Body.Close()
	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Create a map to hold the parsed JSON data
	var responseData map[string]interface{}

	// Unmarshal the JSON response into the map
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
