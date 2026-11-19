package color

func paint(code, s string) string {
	if !enabled {
		return s
	}
	return "\033[" + code + "m" + s + "\033[0m"
}
