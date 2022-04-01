package project

import (
	"os"
	"path"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

func (p *Project) LastActive () time.Time {
	if p.lastActive != nil {
		return *p.lastActive
	}

	dir := path.Join(p.Location, ".git")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		p.lastActive = &time.Time{}
		return *p.lastActive
	}

	repo, err := git.PlainOpen(p.Location)
	if err != nil {
		panic(err)
	}

	head, err := repo.Head()
	if err != nil {
		p.lastActive = &time.Time{}
		return *p.lastActive
	}

	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		panic(err)
	}

	when := commit.Author.When
	p.lastActive = &when

	return when
}

func (p *Project) LastActiveHuman () string {
	return p.LastActive().Format("2006-01-02 15:04:05")
}
