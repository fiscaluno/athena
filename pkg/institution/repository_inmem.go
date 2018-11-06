package institution

import (
	"strings"

	"github.com/fiscaluno/athena/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entity.Institution
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.Institution{}
	return &IRepo{
		m: m,
	}
}

//Store a Institution
func (r *IRepo) Store(a *entity.Institution) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a Institution
func (r *IRepo) Find(id entity.ID) (*entity.Institution, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//Search Institutions
func (r *IRepo) Search(query string) ([]*entity.Institution, error) {
	var d []*entity.Institution
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

//FindAll Institutions
func (r *IRepo) FindAll() ([]*entity.Institution, error) {
	var d []*entity.Institution
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Institution
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}
