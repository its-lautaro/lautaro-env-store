package color

// Green devuelve el texto en green, o sin cambios si el color está apagado.
func Green(s string) string { return paint("32", s) }
