package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"time"

	"github.com/vcokltfre/projman/src/store"
	"github.com/vcokltfre/projman/src/utils"
)

type ProjectManager struct {
	Store store.Store
	stale string
	projs string
}

func (m *ProjectManager) Init () {
	stale := os.Getenv("PROJMAN_STALE")
	if stale == "" {
		hd, _ := os.UserHomeDir()
		stale = path.Join(hd, "projects/stale")

		if _, err := os.Stat(stale); os.IsNotExist(err) {
			os.Mkdir(stale, 0755)
		}
	}
	m.stale = stale

	projs := os.Getenv("PROJMAN_PROJECTS")
	if projs == "" {
		hd, _ := os.UserHomeDir()
		projs = path.Join(hd, "projects")

		if _, err := os.Stat(projs); os.IsNotExist(err) {
			os.Mkdir(projs, 0755)
		}
	}
	m.projs = projs
}

func (m *ProjectManager) Start (name string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if proj, found := m.Store.Get(name); found {
		fmt.Println("A project already exists by that name at", proj.Location)
		return
	}

	m.Store.Set(name, wd)

	b := utils.BuildColourText().Green().Add("Created new project: ").Blue().Add(name)
	fmt.Println(b.String())
}

func (m *ProjectManager) Close (name string) {
	project, found := m.Store.Get(name)
	if !found {
		fmt.Println("Project not found.")
		return
	}
	m.Store.Del(name)

	err := os.Rename(project.Location, path.Join(m.stale, name))
	if err != nil {
		panic(err)
	}

	b := utils.BuildColourText().Red().Add("Closed a project: ").Blue().Add(name)
	fmt.Println(b.String())
}

func (m *ProjectManager) CloseFromLocation (location string) {
	for _, project := range m.Store.List() {
		if project.Location == location {
			m.Close(project.Name)
		}
	}
}

func (m *ProjectManager) Unclose (name string) {
	loc := path.Join(m.stale, name)

	if _, err := os.Stat(loc); os.IsNotExist(err) {
		fmt.Println("Project not found.")
		return
	}

	err := os.Rename(loc, path.Join(m.projs, name))
	if err != nil {
		panic(err)
	}
}

func (m *ProjectManager) List () {
	projects := m.Store.List()
	noGit := time.Time{}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].LastActive().After(projects[j].LastActive())
	})

	for _, project := range projects {
		var date string

		if project.LastActive() == noGit {
			date = "  NO GIT  "
		} else {
			date = project.LastActive().Format("2006-01-02")
		}

		b := utils.BuildColourText()
		b.Red().Add("[").Cyan().Add(date).Red().Add("]").Reset().Add(" ").Add(project.Name)

		wd, _ := os.Getwd()
		if wd == project.Location {
			b.Green().Add(" (open)").Reset()
		}

		if project.IsStale() {
			b.Red().Add(" (stale)").Reset()
		}

		fmt.Println(b.String())
	}
}

func (m *ProjectManager) cleanupShallow() int {
	done := 0

	for _, project := range m.Store.List() {
		if project.IsStale() {
			m.Close(project.Name)
			done += 1
		}
	}

	return done
}

func (m *ProjectManager) cleanupDeep() int {
	done := 0

	files, err := ioutil.ReadDir(m.projs)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		loc := path.Join(m.projs, file.Name())
		if loc == m.stale {
			continue
		}

		proj, found := m.Store.GetByLocation(loc)
		if !found {
			b := utils.BuildColourText().Red().Add("Closed non-registered project: ").Blue().Add(file.Name())
			fmt.Println(b.String())
			done += 1

			err := os.Rename(loc, path.Join(m.stale, file.Name()))
			if err != nil {
				panic(err)
			}
		} else {
			if proj.IsStale() {
				m.Close(proj.Name)
				done += 1
			}
		}
	}

	return done
}

func (m *ProjectManager) Cleanup (deep bool) {
	var done int

	if deep {
		done = m.cleanupDeep()
	} else {
		done = m.cleanupShallow()
	}

	b := utils.BuildColourText().Red().Add("Cleaned up ").Blue().Add(fmt.Sprintf("%d", done)).Red().Add(" stale projects.")
	fmt.Println(b.String())
}

func (m *ProjectManager) Validate () {
	for _, project := range m.Store.List() {
		if _, err := os.Stat(project.Location); os.IsNotExist(err) {
			b := utils.BuildColourText().Red().Add("Closed non-existent project: ").Blue().Add(project.Name)
			fmt.Println(b.String())

			m.Store.Del(project.Name)
		}
	}
}
