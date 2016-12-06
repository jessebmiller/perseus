package main

// Store interface for debug message storage
type Store interface {
	get(string) []string
	add(string, string) error
}

// MapStore stores messages in a map
type MapStore struct {
	m map[string][]string
}

func (store MapStore) get(ns string) []string {
	return store.m[ns]
}

func (store MapStore) add(ns string, message string) error {
	store.m[ns] = append([]string{message}, store.m[ns]...)
	return nil
}
