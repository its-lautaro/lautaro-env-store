package store

func parseEntries(lines []string) []Entry {
	var entries []Entry
	for _, line := range lines {
		if e, ok := parseLine(line); ok {
			entries = append(entries, e)
		}
	}
	return entries
}
