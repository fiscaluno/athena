package course

import (
	"encoding/json"
	"net/http"

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

// Store a Course
func (r *APIRepo) Store(a *entity.Course) (entity.ID, error) {
	return entity.NewID(), nil
}

//Find a Course
func (r *APIRepo) Find(id entity.ID) (*entity.Course, error) {
	var i *entity.Course
	return i, nil
}

//Search Courses
func (r *APIRepo) Search(query string) ([]*entity.Course, error) {
	var i []*entity.Course

	return i, nil
}

//FindAll Courses
func (r *APIRepo) FindAll() ([]*entity.Course, error) {
	var i []*entity.Course
	resp, err := http.Get(r.uri + r.path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&i)

	return i, nil
}

//Delete a Course
func (r *APIRepo) Delete(id entity.ID) error {
	return nil
}
