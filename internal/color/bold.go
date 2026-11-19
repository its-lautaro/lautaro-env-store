package color

// Bold devuelve el texto en bold, o sin cambios si el color está apagado.
func Bold(s string) string { return paint("1", s) }
