package store

import "regexp"

var (
	entryRE = regexp.MustCompile(`^export\s+(` + namePattern + `)=`)
	savedRE = regexp.MustCompile(regexp.QuoteMeta(savedMarker) + `(\d{4}-\d{2}-\d{2} \d{2}:\d{2})$`)
)

func parseLine(line string) (Entry, bool) {
	m := entryRE.FindStringSubmatch(line)
	if m == nil {
		return Entry{}, false
	}
	e := Entry{Name: m[1]}
	if s := savedRE.FindStringSubmatch(line); s != nil {
		e.Saved = s[1]
	}
	return e, true
}
