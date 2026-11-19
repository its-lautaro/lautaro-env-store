package store

import "time"

func (s Store) Put(vars []Var) (replaced map[string]bool, err error) {
	for _, v := range vars {
		if err := validName(v.Name); err != nil {
			return nil, err
		}
	}
	lines, err := s.readLines()
	if err != nil {
		return nil, err
	}
	existing := map[string]bool{}
	for _, e := range parseEntries(lines) {
		existing[e.Name] = true
	}
	names := map[string]bool{}
	replaced = map[string]bool{}
	for _, v := range vars {
		names[v.Name] = true
		if existing[v.Name] {
			replaced[v.Name] = true
		}
	}
	lines = dropNames(lines, names)
	stamp := time.Now().Format(stampLayout)
	for _, v := range vars {
		lines = append(lines, formatLine(v, stamp))
	}
	return replaced, s.writeLines(lines)
}
