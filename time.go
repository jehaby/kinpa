package kinpa

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type timeString struct {
	Regexp string
	String string
}

var knownTimeStrings = []timeString{
	{`[\d]{1,2}[\s][\w]+[\s][\d]{4}`, `2 January 2006 15:04:05`},
	{`[\d]{1,2}[\s][\w]+[\s][\d]{2}`, `2 January 06 15:04:05`},
	{`[\w]+[\s][\d]{1,2},[\s][\d]{4}`, `January 2, 2006, 03:04 PM`},
}

func parseTimeString(s string) (*time.Time, error) {
	if len(s) < 10 {
		return nil, fmt.Errorf("Passed string is too short: %s", s)
	}
	clearedString := s[strings.Index(s, ",")+2:]

	for _, ts := range knownTimeStrings {
		m, err := regexp.MatchString(ts.Regexp, clearedString)
		if err != nil {
			return nil, err
		}
		if !m {
			continue
		}
		res, err := time.Parse(ts.String, clearedString)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}
	return nil, fmt.Errorf("String of unknown format given: %s", s)
}
