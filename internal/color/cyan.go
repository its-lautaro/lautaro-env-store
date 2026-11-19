package color

// Cyan devuelve el texto en cyan, o sin cambios si el color está apagado.
func Cyan(s string) string { return paint("36", s) }
