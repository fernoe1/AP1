package store

import "sync"

type Store[K comparable, V any] struct {
	storeLock sync.RWMutex
	store     map[K]V
}

func New[K comparable, V any]() *Store[K, V] {
	return &Store[K, V]{
		store: make(map[K]V),
	}
}

func (s *Store[K, V]) Set(k K, v V) {
	s.storeLock.Lock()
	defer s.storeLock.Unlock()
	s.store[k] = v
}

func (s *Store[K, V]) Get(k K) (v V, ok bool) {
	s.storeLock.RLock()
	defer s.storeLock.RUnlock()
	v, ok = s.store[k]

	return
}

func (s *Store[K, V]) Delete(k K) bool {
	s.storeLock.Lock()
	defer s.storeLock.Unlock()
	_, ok := s.store[k]

	if ok {
		delete(s.store, k)
	}

	return ok
}

func (s *Store[K, V]) Snapshot() map[K]V {
	s.storeLock.RLock()
	defer s.storeLock.RUnlock()

	return s.store
}
