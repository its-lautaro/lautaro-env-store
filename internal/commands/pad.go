package commands

import "strings"

const descriptionColumn = 27

func pad(plain string) string {
	return strings.Repeat(" ", max(descriptionColumn-len([]rune(plain)), 2))
}
