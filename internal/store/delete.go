package store

func (s Store) Delete(names []string) (removed, missing []string, err error) {
	lines, err := s.readLines()
	if err != nil {
		return nil, nil, err
	}
	present := map[string]bool{}
	for _, e := range parseEntries(lines) {
		present[e.Name] = true
	}
	wanted := map[string]bool{}
	for _, n := range names {
		wanted[n] = true
		if present[n] {
			removed = append(removed, n)
		} else {
			missing = append(missing, n)
		}
	}
	return removed, missing, s.writeLines(dropNames(lines, wanted))
}
