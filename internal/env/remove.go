package env

// Remove borra las variables indicadas y devuelve las que no existían.
func Remove(names []string) map[string]bool {
	present := map[string]bool{}
	for _, e := range ReadEntries() {
		present[e.Name] = true
	}
	wanted := map[string]bool{}
	for _, n := range names {
		wanted[n] = true
	}
	WriteLines(DropNames(ReadLines(), wanted))
	missing := map[string]bool{}
	for _, n := range names {
		if !present[n] {
			missing[n] = true
		}
	}
	return missing
}
