package table

import "errors"

type Players []Player

func (ps *Players) AddPlayer(p Player) error {
	if ps.isNameExist(p.name) {
		return errors.New("name is already exist")
	}
	*ps = append(*ps, p)
	return nil
}

func (ps *Players) isNameExist(n string) bool {
	for _, p := range *ps {
		if p.name == n {
			return true
		}
	}
	return false
}
