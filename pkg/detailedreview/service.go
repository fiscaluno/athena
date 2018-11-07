package detailedreview

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

//Store an DetailedReview
func (s *Service) Store(b *entity.DetailedReview) (entity.ID, error) {
	all, err := s.repo.FindAll()
	if err != nil {
		return entity.NewID(), err
	}
	b.ID = entity.ID(len(all) + 1)
	return s.repo.Store(b)
}

//Find a DetailedReview
func (s *Service) Find(id entity.ID) (*entity.DetailedReview, error) {
	return s.repo.Find(id)
}

//Search DetailedReviews
func (s *Service) Search(query string) ([]*entity.DetailedReview, error) {
	return s.repo.Search(strings.ToLower(query))
}

//FindAll DetailedReviews
func (s *Service) FindAll() ([]*entity.DetailedReview, error) {
	return s.repo.FindAll()
}

//Delete a DetailedReview
func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
