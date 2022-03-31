package store

import (
	"os"
	"path"

	"github.com/vcokltfre/projman/src/project"
)

type Store struct {
	Location string
	projects map[string]string
	config   map[string]string
}

func NewStore() *Store {
	dir := path.Join(os.Getenv("HOME"), ".config")
	file := path.Join(dir, "projman.json")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		os.WriteFile(file, []byte("{\"projects\":{},\"config\":{}}"), 0644)
	}

	store := Store{
		Location: file,
	}

	store.projects = make(map[string]string)
	store.config = make(map[string]string)

	store.load()

	return &store
}

func (s *Store) Get (key string) (project.Project, bool) {
	value, ok := s.projects[key]
	if !ok {
		return project.Project{}, false
	}

	return project.Project{
		Location: value,
		Name:     key,
	}, true
}

func (s *Store) GetByLocation (location string) (project.Project, bool) {
	for key, value := range s.projects {
		if value == location {
			return project.Project{
				Location: value,
				Name:     key,
			}, true
		}
	}

	return project.Project{}, false
}

func (s *Store) Set (key string, value string) {
	s.projects[key] = value
	s.save()
}

func (s *Store) Del (key string) bool {
	_, found := s.Get(key)

	if !found {
		return false
	}

	delete(s.projects, key)
	s.save()

	return true
}

func (s *Store) List () []project.Project {
	projects := []project.Project{}

	for key := range s.projects {
		project, _ := s.Get(key)
		projects = append(projects, project)
	}

	return projects
}
