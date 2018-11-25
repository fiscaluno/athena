package institution

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

// Store a Institution
func (r *APIRepo) Store(a *entity.Institution) (entity.ID, error) {
	return entity.NewID(), nil
}

//Find a Institution
func (r *APIRepo) Find(id entity.ID) (*entity.Institution, error) {
	var i *entity.Institution
	return i, nil
}

//Search Institutions
func (r *APIRepo) Search(query string) ([]*entity.Institution, error) {
	var i []*entity.Institution

	return i, nil
}

//FindAll Institutions
func (r *APIRepo) FindAll() ([]*entity.Institution, error) {
	var i []*entity.Institution
	resp, err := http.Get(r.uri + r.path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&i)

	return i, nil
}

//Delete a Institution
func (r *APIRepo) Delete(id entity.ID) error {
	return nil
}
