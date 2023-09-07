package pet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fakestore_go/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io"
	"net/http"
)

type Repository struct {
	rootURl string
	httpCli *http.Client
	db      *sqlx.DB
}

func New(httpCli *http.Client, db *sqlx.DB) *Repository {
	return &Repository{
		rootURl: "https://petstore.swagger.io/v2",
		httpCli: httpCli,
		db:      db,
	}
}

func (r *Repository) InsertNewPetToDatabase(req entity.Pet) (*entity.Pet, error) {
	q := "INSERT INTO pets(name, status) VALUES(?, ?)"
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	result, err := tx.Exec(q, req.Name, req.Status)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	req.ID = lastID
	return &req, err
}

func (r *Repository) GetPetFromDatabase(petID int64) (*entity.Pet, error) {
	q := "SELECT id,name, status FROM pets WHERE id = ?"
	resp := entity.Pet{}
	err := r.db.Get(&resp, q, petID)
	return &resp, err
}

func (r *Repository) InsertNewPet(req entity.Pet) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/pet", r.rootURl)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
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
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
