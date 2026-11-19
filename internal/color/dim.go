package color

// Dim devuelve el texto en dim, o sin cambios si el color está apagado.
func Dim(s string) string { return paint("2", s) }
