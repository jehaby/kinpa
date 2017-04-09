package kinpa

import "testing"

var metaTests = []struct {
	S        string
	Page     int
	Location string
	E        error
}{
	{
		`- Your Highlight on Page 5 | Location 64-64 | Added on Tuesday, November 01, 2011, 08:49 AM`,
		5,
		`64-64`,
		nil,
	},
	{
		`- Your Highlight Location 4407-4408 | Added on Wednesday, November 02, 2011, 12:42 AM`,
		0,
		`4407-4408`,
		nil,
	},
}

func TestParseMetaString(t *testing.T) {

	for _, data := range metaTests {

		//		fmt.Println(data)

		page, location, _, e := parseMetaString(data.S)

		if page != data.Page {
			t.Error("Pages not equal. Expected: ", data.Page, " | Got: ", page)
		}

		if location != data.Location {
			t.Error("Locations not equal. Expected: ", data.Location, " | Got: ", location)
		}

		if e != data.E {
			t.Error("Errors not equal. Expected: ", data.E, " | Got: ", e)
		}
	}

}
