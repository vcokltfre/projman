package store

import (
	"encoding/json"
	"os"
)

type store struct {
	Config 	 map[string]string `json:"config"`
	Projects map[string]string `json:"projects"`
}

func (s *Store) load() {
	data, err := os.ReadFile(s.Location)
	if err != nil {
		panic(err)
	}

	store := &store{}
	err = json.Unmarshal(data, store)
	if err != nil {
		panic(err)
	}

	s.config = store.Config
	s.projects = store.Projects
}

func (s *Store) save() {
	store := &store{
		Config: 	s.config,
		Projects: s.projects,
	}

	data, err := json.Marshal(store)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(s.Location, data, 0644)
	if err != nil {
		panic(err)
	}
}
