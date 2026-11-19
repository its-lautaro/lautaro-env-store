package store

func (s Store) Create() error {
	lines, err := s.readLines()
	if err != nil {
		return err
	}
	return s.writeLines(lines)
}
