package store

func dropNames(lines []string, names map[string]bool) []string {
	kept := lines[:0:0]
	for _, line := range lines {
		if e, ok := parseLine(line); ok && names[e.Name] {
			continue
		}
		kept = append(kept, line)
	}
	return kept
}
