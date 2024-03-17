package storage

type Storage interface {
}

//type Storage struct {
//}

func New(storageType string) *Storage {
	var store Storage
	switch storageType {
	case "memory":
		store = 1
	case "db":
		store = 2
	}

	return &store
}
