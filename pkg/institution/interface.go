package institution

import "github.com/fiscaluno/athena/pkg/entity"

//Reader interface
type Reader interface {
	Find(id entity.ID) (*entity.Institution, error)
	Search(query string) ([]*entity.Institution, error)
	FindAll() ([]*entity.Institution, error)
}

//Writer Institution writer
type Writer interface {
	Store(b *entity.Institution) (entity.ID, error)
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
