package course

import (
	"strings"

	"github.com/fiscaluno/athena/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entity.Course
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.Course{}
	return &IRepo{
		m: m,
	}
}

//Store a Course
func (r *IRepo) Store(a *entity.Course) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a Course
func (r *IRepo) Find(id entity.ID) (*entity.Course, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//Search Courses
func (r *IRepo) Search(query string) ([]*entity.Course, error) {
	var d []*entity.Course
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

//FindAll Courses
func (r *IRepo) FindAll() ([]*entity.Course, error) {
	var d []*entity.Course
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Course
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	var x = entity.Course{}
	r.m[id.String()] = &x
	return nil
}
