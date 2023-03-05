package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (store *MemStore) Create(_ *mystore.Book) error {
	panic("not implemented") // TODO: Implement
}

func (store *MemStore) Get(_ string) (mystore.Book, error) {
	panic("not implemented") // TODO: Implement
}
