package store

import "regexp"

const namePattern = `[A-Za-z_][A-Za-z0-9_]*`

var nameRE = regexp.MustCompile(`^` + namePattern + `$`)

func ValidName(name string) bool {
	return nameRE.MatchString(name)
}
