package core

import "strings"

type BeefService interface {
	Summary() (*Beef, error)
}

type beefServiceImpl struct {
	repo BeefRepository
}

func NewBeefService(repo BeefRepository) BeefService {
	return &beefServiceImpl{repo: repo}
}

func countBeef(beefs string, summary *Beef) {
	for _, v := range strings.Fields(beefs) {
		if v != "" {
			summary.Beef[v]++
		}
	}
}

func (s *beefServiceImpl) Summary() (*Beef, error) {
	summary := Beef{Beef: make(map[string]int)}

	byteData, err := s.repo.GetData()
	if err != nil {
		return nil, err
	}
	data := string(byteData)

	// Clean data
	data = strings.ReplaceAll(data, ",", " ")
	data = strings.ReplaceAll(data, ".", " ")
	data = strings.ReplaceAll(data, "\n", " ")
	data = strings.ToLower(data)

	// Count
	countBeef(data, &summary)

	return &summary, nil
}
