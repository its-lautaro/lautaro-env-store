package store

import (
	"fmt"
	"regexp"
)

const namePattern = `[A-Za-z_][A-Za-z0-9_]*`

var nameRE = regexp.MustCompile(`^` + namePattern + `$`)

func validName(name string) error {
	if !nameRE.MatchString(name) {
		return fmt.Errorf("'%s' no es un nombre válido de variable de entorno", name)
	}
	return nil
}
