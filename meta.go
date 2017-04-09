package kinpa

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//- Your Highlight on Page 5 | Location 64-64 | Added on Tuesday, November 01, 2011, 08:49 AM
// - Your Highlight Location 4407-4408 | Added on Wednesday, November 02, 2011, 12:42 AM

func parseMetaString(s string) (page int, location string, t time.Time, e error) {

	temp := strings.Split(s, "|")

	if len(temp) == 2 {

	}

	switch len(temp) {
	case 2:

		location, le := parseLocation(temp[0])
		if le != nil {
			return 0, "", time.Now(), fmt.Errorf("Got error while parsing location in meta string '%s': %v", s, le)
		}

		t, te := parseTimeString(temp[1])
		if te != nil {
			return 0, "", time.Now(), fmt.Errorf("Got error while parsing time in meta string '%s': %v", s, te)
		}

		return 0, location, *t, nil

	case 3:
		page, pe := parsePage(temp[0])
		if pe != nil {
			return 0, "", time.Now(), fmt.Errorf("Got error while parsing page in meta string '%s': %v", s, pe)
		}

		location, le := parseLocation(temp[1])
		if le != nil {
			return page, "", time.Now(), fmt.Errorf("Got error while parsing location in meta string '%s': %v", s, le)
		}

		t, te := parseTimeString(temp[2])
		if te != nil {
			return page, location, time.Now(), fmt.Errorf("Got error while parsing time in meta string '%s': %v", s, te)
		}

		return page, location, *t, nil

	default:
		return 0, "", time.Now(), fmt.Errorf("Unknown meta string format '%s': ", s)
	}

}

var locationRegexp = regexp.MustCompile("\\d+")

func parseLocation(s string) (string, error) {
	if res := locationRegexp.FindString(s); res != "" {
		return res, nil
	}
	return "", fmt.Errorf("Couldn't find location in string: %s", s)
}

var pageRegexp = regexp.MustCompile("\\d+")

func parsePage(s string) (int, error) {
	if res := pageRegexp.FindString(s); res != "" {
		res, _ := strconv.Atoi(res)
		return res, nil
	}
	return 0, fmt.Errorf("Couldn't find page in string: %s", s)
}
