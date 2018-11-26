package review

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/fiscaluno/athena/pkg/entity"
)

// APIRepo api repo
type APIRepo struct {
	uri  string
	path string
}

// NewAPIRepository create new repository
func NewAPIRepository(u, p string) *APIRepo {
	return &APIRepo{
		uri:  u,
		path: p,
	}
}

// Store a Review
func (r *APIRepo) Store(a *entity.Review) (entity.ID, error) {

	a.CourseID = a.CourseInfo.CourseID
	payload, err := json.Marshal(a)
	if err != nil {
		return entity.NewID(), err
	}

	req, _ := http.NewRequest("POST", r.uri+r.path, strings.NewReader(string(payload)))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	var resp struct {
		Body struct {
			ID int `json:"id"`
		} `json:"Body"`
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&resp)

	return entity.StringToID(strconv.Itoa(resp.Body.ID)), nil
}

//Find a Review
func (r *APIRepo) Find(id entity.ID) (*entity.Review, error) {
	var i *entity.Review
	return i, nil
}

//Search Reviews
func (r *APIRepo) Search(query string) ([]*entity.Review, error) {
	var i []*entity.Review
	resp, err := http.Get(r.uri + r.path + "?" + query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&i)
	return i, nil
}

//FindAll Reviews
func (r *APIRepo) FindAll() ([]*entity.Review, error) {
	var i []*entity.Review
	resp, err := http.Get(r.uri + r.path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&i)

	return i, nil
}

//Delete a Review
func (r *APIRepo) Delete(id entity.ID) error {
	return nil
}
