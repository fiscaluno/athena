package review

import "github.com/fiscaluno/athena/pkg/entity"

//Reader interface
type Reader interface {
	Find(id entity.ID) (*entity.Review, error)
	Search(query string) ([]*entity.Review, error)
	FindAll() ([]*entity.Review, error)
}

//Writer Review writer
type Writer interface {
	Store(b *entity.Review) (entity.ID, error)
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
