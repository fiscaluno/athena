package institution

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

//Store an Institution
func (s *Service) Store(b *entity.Institution) (entity.ID, error) {
	all, err := s.repo.FindAll()
	if err != nil {
		return entity.NewID(), err
	}
	b.ID = entity.ID(len(all) + 1)
	return s.repo.Store(b)
}

//Find a Institution
func (s *Service) Find(id entity.ID) (*entity.Institution, error) {
	return s.repo.Find(id)
}

//Search Institutions
func (s *Service) Search(query string) ([]*entity.Institution, error) {
	return s.repo.Search(strings.ToLower(query))
}

//FindAll Institutions
func (s *Service) FindAll() ([]*entity.Institution, error) {
	return s.repo.FindAll()
}

//Delete a Institution
func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
