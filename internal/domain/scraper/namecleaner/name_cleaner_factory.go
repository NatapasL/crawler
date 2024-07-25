package namecleaner

import "fmt"

type NameCleanerFactory struct {
	cleaners map[string]*NameCleaner
}

func NewNameCleanerFactory() *NameCleanerFactory {
	return &NameCleanerFactory{cleaners: make(map[string]*NameCleaner)}
}

func (f *NameCleanerFactory) Register(key string, cleaner NameCleaner) {
	f.cleaners[key] = &cleaner
}

func (f NameCleanerFactory) GetCleaner(key string) (*NameCleaner, error) {
	cleaner := f.cleaners[key]
	if cleaner == nil {
		return nil, fmt.Errorf("name cleaner key %s not registered", key)
	}

	return cleaner, nil
}
