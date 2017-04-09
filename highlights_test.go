package kinpa

import (
	//	"fmt"
	"testing"
	"time"
)

var b = []Book{
	NewBook("Iain Banks", "Excession"),
	NewBook("Iain Banks", "State of the art"),
	NewBook("Bertrand Russell", "A History of Western Philosophy"),
	NewBook("Robert Martin", "Clean Code"),
}

var highlights = []Highlight{
	{"some word", 0, "30-31", time.Date(2016, 4, 19, 6, 36, 12, 0, time.UTC), nil},
	{"exception", 0, "340-340", time.Date(2016, 5, 20, 16, 20, 12, 0, time.UTC), nil},
	{" Entertain", 0, "30-31", time.Date(2016, 4, 19, 6, 33, 2, 0, time.UTC), nil},
	{"Some very long string", 0, "30-31", time.Date(2015, 9, 19, 13, 36, 17, 0, time.UTC), nil},
	{"exception", 0, "1-3", time.Date(2016, 5, 20, 16, 20, 12, 0, time.UTC), nil},
	{"WTF???????????????", 0, "30-31", time.Date(2015, 10, 19, 7, 36, 23, 0, time.UTC), nil},
}

var setTextExamples = []struct {
	text     string
	expected string
}{
	{"simple", "simple"},
	{" a bit harder", "a bit harder"},
	{"Dot in the end.", "Dot in the end"},
	{"“fascination.” ", "fascination"},
	{"“conscientiousness,” ", "conscientiousness"},
	{"shrink,’ ” ", "shrink"},
}

func TestSetText(t *testing.T) {
	for _, ex := range setTextExamples {
		h := Highlight{}
		h.SetText(ex.text)
		if h.Text != ex.expected {
			t.Errorf("Set wrong text. Expected: `%s`, Got: `%s`", ex.expected, h.Text)
		}
	}
}

func TestAdd(t *testing.T) {
	hs := NewHighlightStorage()

	h := highlights[0]
	if err := hs.Add(h); err != nil {
		t.Errorf("Got error when tried to add valid highlight. Highlight: %+v | Error: %+v ", h, err)
	}
	checkLen(hs.Len(), 1, "HighligghtStorage", t)

	fromStorage, err := hs.GetByText(h.Text)
	if err != nil {
		t.Errorf("Undexpected error: %+v", err)
	}
	if *(fromStorage[0]) != h {
		t.Errorf("Highlights aren't equal: from storage: %+v | added: %+v", *(fromStorage[0]), h)
	}

	// TODO: test that you can get added hightlight

}

func TestContains(t *testing.T) {

	hs := NewHighlightStorage()

	h1, h2 := highlights[1], highlights[4]
	testContains(hs, h1, false, t)
	testContains(hs, h2, false, t)

	testAdd(hs, h1, "", t)
	testContains(hs, h1, true, t)
	testContains(hs, h2, false, t)

	testAdd(hs, h2, "", t)
	testContains(hs, h1, true, t)
	testContains(hs, h2, true, t)
	testContains(hs, highlights[0], false, t)

	checkLen(hs.Len(), 2, "HighligghtStorage", t)

}

func testAdd(hs *HighlightStorage, h Highlight, em string, t *testing.T) {
	if err := hs.Add(h); em == "" && err != nil {
		t.Errorf("Got error when tried to add valid highlight. Highlight: %+v | Error: %+v ", h, err)
	} else if em != "" && em != err.Error() {
		t.Errorf("Wrong error message when tried to add valid highlight. Highlight: %+v | Expected: %s | Actual: %s", h, em, err.Error())
	}
}

func testContains(hs *HighlightStorage, h Highlight, e bool, t *testing.T) {
	if hs.Contains(h) != e {
		t.Errorf("Wrong Contains result. H: %#v | Actual: %t | Expected: %t ")
	}
}

func TestGetByText(t *testing.T) {
	/*
		hs := NewHighlightStorage()

		for i := range highlights {
			testAdd(&hs, highlights[i], "", t)
		}

		checkLen(hs.Len(), len(highlights), "HighlightStorage", t)

		temp, _ := hs.GetByText("exception")
		checkLen(len(temp), 2, "[]*Highlights", t)

		temp, _ = hs.GetByText("Putin")
		checkLen(len(temp), 0, "[]*Highlights", t)

		temp, _ = hs.GetByText(" Entertain")
		checkLen(len(temp), 1, "[]*Highlights", t)
	*/
}

func checkLen(actual, expected int, name string, t *testing.T) {
	if actual != expected {
		t.Errorf("Wrong Len of %s. Actual: %d. Expected: %d", name, actual, expected)
	}
}
