package review

import (
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

//Store an Review
func (s *Service) Store(b *entity.Review) (entity.ID, error) {
	all, err := s.repo.FindAll()
	if err != nil {
		return entity.NewID(), err
	}
	b.ID = entity.ID(len(all) + 1)
	return s.repo.Store(b)
}

//Find a Review
func (s *Service) Find(id entity.ID) (*entity.Review, error) {
	return s.repo.Find(id)
}

//Search Reviews
func (s *Service) Search(query string) ([]*entity.Review, error) {
	return s.repo.Search(query)
}

//FindAll Reviews
func (s *Service) FindAll() ([]*entity.Review, error) {
	return s.repo.FindAll()
}

//Delete a Review
func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
