package storage

type Storage struct {
}

var Instance *Storage

func init() {
	Instance = NewStorage()
}

func NewStorage() *Storage {
	var c Storage
	return &c
}
