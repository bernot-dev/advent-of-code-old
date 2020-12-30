package helpers

import (
	"fmt"
	"regexp"
)

// MapNamedSubmatches creates a map of all submatches of the pattern provided
func MapNamedSubmatches(re *regexp.Regexp, s string) (map[string]string, error) {
	if !re.MatchString(s) {
		return nil, fmt.Errorf("No matches for pattern %q in %q", re, s)
	}
	matches := re.FindStringSubmatch(s)
	names := re.SubexpNames()
	m := make(map[string]string)
	for i, name := range names {
		if name == "" {
			continue
		}
		m[name] = matches[i]
	}
	return m, nil
}
