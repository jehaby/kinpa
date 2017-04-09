package kinpa

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

var parseTimeStringTests = []struct {
	s        string
	expected interface{}
	e        error
}{
	{
		"Added on Monday, 31 October 2016 07:27:46",
		time.Date(2016, 10, 31, 7, 27, 46, 0, time.UTC),
		nil,
	},
	{
		"Added on Saturday, 11 July 15 02:25:50",
		time.Date(2015, 07, 11, 2, 25, 50, 0, time.UTC),
		nil,
	},
	{
		"Added on Monday, 30 October 2016 07:27:46",
		time.Date(2016, 10, 30, 7, 27, 46, 0, time.UTC),
		nil,
	},
	{
		"Added on Friday, March 28, 2014, 10:13 PM",
		time.Date(2014, 3, 28, 22, 13, 0, 0, time.UTC),
		nil,
	},
	{
		"Added on Friday, March 28, 2014, 10:13 PM",
		time.Date(2014, 3, 28, 22, 13, 0, 0, time.UTC),
		nil,
	},
	{
		"bad word",
		nil,
		errors.New("Passed string is too short: bad word"),
	},
	{
		"abcdefghijklmnop",
		nil,
		errors.New("String of unknown format given: abcdefghijklmnop"),
	},
}

func TestParseTimeString(t *testing.T) {

	for _, testData := range parseTimeStringTests {

		res, err := parseTimeString(testData.s)

		fmt.Println(testData, res)

		if expected, ok := testData.expected.(time.Time); ok && !res.Equal(expected) {
			t.Error("Date not equal! string: '", testData.s, "'| res: ", res.Local(), "| expected: ", expected.Local())
		} else if err != nil && err.Error() != testData.e.Error() {
			fmt.Println(len(err.Error()), len(testData.e.Error()))
			t.Error("Wrong error! string: '", testData.s, "'| res: ", err.Error(), "| expected: ", testData.e.Error())
		}

	}

}
