package env

// DropNames devuelve las líneas del store sin las entradas de los nombres indicados.
func DropNames(lines []string, names map[string]bool) []string {
	kept := lines[:0:0]
	for _, line := range lines {
		if m := EntryRE.FindStringSubmatch(line); m != nil && names[m[1]] {
			continue
		}
		kept = append(kept, line)
	}
	return kept
}
