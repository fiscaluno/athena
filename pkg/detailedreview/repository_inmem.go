package detailedreview

import (
	"strings"

	"github.com/fiscaluno/athena/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entity.DetailedReview
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.DetailedReview{}
	return &IRepo{
		m: m,
	}
}

//Store a DetailedReview
func (r *IRepo) Store(a *entity.DetailedReview) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a DetailedReview
func (r *IRepo) Find(id entity.ID) (*entity.DetailedReview, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//Search DetailedReviews
func (r *IRepo) Search(query string) ([]*entity.DetailedReview, error) {
	var d []*entity.DetailedReview
	for _, j := range r.m {
		if strings.Contains(string(j.InstitutionID), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

//FindAll DetailedReviews
func (r *IRepo) FindAll() ([]*entity.DetailedReview, error) {
	var d []*entity.DetailedReview
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a DetailedReview
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}
