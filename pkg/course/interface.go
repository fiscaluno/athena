package course

import "github.com/fiscaluno/athena/pkg/entity"

//Reader interface
type Reader interface {
	Find(id entity.ID) (*entity.Course, error)
	Search(query string) ([]*entity.Course, error)
	FindAll() ([]*entity.Course, error)
}

//Writer Course writer
type Writer interface {
	Store(b *entity.Course) (entity.ID, error)
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
