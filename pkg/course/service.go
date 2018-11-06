package course

import (
	"strings"

	"github.com/fiscaluno/athena/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Store an Course
func (s *Service) Store(b *entity.Course) (entity.ID, error) {
	return s.repo.Store(b)
}

//Find a Course
func (s *Service) Find(id entity.ID) (*entity.Course, error) {
	return s.repo.Find(id)
}

//Search Courses
func (s *Service) Search(query string) ([]*entity.Course, error) {
	return s.repo.Search(strings.ToLower(query))
}

//FindAll Courses
func (s *Service) FindAll() ([]*entity.Course, error) {
	return s.repo.FindAll()
}

//Delete a Course
func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
