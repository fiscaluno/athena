package review

import (
	"strings"

	"github.com/fiscaluno/athena/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entity.Review
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.Review{}
	return &IRepo{
		m: m,
	}
}

//Store a Review
func (r *IRepo) Store(a *entity.Review) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a Review
func (r *IRepo) Find(id entity.ID) (*entity.Review, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//Search Reviews
func (r *IRepo) Search(query string) ([]*entity.Review, error) {
	var d []*entity.Review
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Title), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

//FindAll Reviews
func (r *IRepo) FindAll() ([]*entity.Review, error) {
	var d []*entity.Review
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Review
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}
