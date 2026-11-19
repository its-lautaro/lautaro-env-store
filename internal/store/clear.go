package store

func (s Store) Clear() error {
	return s.writeLines(nil)
}
