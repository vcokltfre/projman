package project

import "time"

type Project struct {
	Location   string
	Name       string
	lastActive *time.Time
}

func (p *Project) Less (other Project) bool {
	return p.LastActive().Unix() < other.LastActive().Unix()
}

func (p *Project) IsStale () bool {
	noGit := time.Time{}
	old := time.Now().Add(-(time.Hour * 24 * 14))

	return p.LastActive() != noGit && p.LastActive().Before(old)
}
