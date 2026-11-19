package store

func (s Store) Entries() ([]Entry, error) {
	lines, err := s.readLines()
	if err != nil {
		return nil, err
	}
	return parseEntries(lines), nil
}
