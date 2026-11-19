package env

// ReadEntries lista las variables guardadas en el orden del archivo.
func ReadEntries() []Entry {
	var entries []Entry
	for _, line := range ReadLines() {
		m := EntryRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		saved := ""
		if s := SavedRE.FindStringSubmatch(line); s != nil {
			saved = s[1]
		}
		entries = append(entries, Entry{Name: m[1], Saved: saved})
	}
	return entries
}
