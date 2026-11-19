package env

import "os"

// EnsureStore crea el store vacío con permisos 600 si no existe.
func EnsureStore() {
	f, err := os.OpenFile(Store, os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		Fail("no pude crear %s: %v", Store, err)
	}
	f.Close()
}
