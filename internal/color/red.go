package color

// Red devuelve el texto en red, o sin cambios si el color está apagado.
func Red(s string) string { return paint("31", s) }
