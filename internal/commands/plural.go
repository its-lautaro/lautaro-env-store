package commands

func plural(n int) string {
	if n > 1 {
		return "s"
	}
	return ""
}
