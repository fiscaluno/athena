package detailedreview

import "github.com/fiscaluno/athena/pkg/entity"

//Reader interface
type Reader interface {
	Find(id entity.ID) (*entity.DetailedReview, error)
	Search(query string) ([]*entity.DetailedReview, error)
	FindAll() ([]*entity.DetailedReview, error)
}

//Writer DetailedReview writer
type Writer interface {
	Store(b *entity.DetailedReview) (entity.ID, error)
	Delete(id entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
