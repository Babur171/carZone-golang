package store

type Storage interface {
	CreateStudent(name string, email string) (int64, error)
}
