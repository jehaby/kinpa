package kinpa

import "testing"

var bookStrings = []struct {
	s string
	b Book
	e error
}{
	{
		"Стихотворения (1908-1937) (Осип Мандельштам)",
		NewBook("Осип Мандельштам", "Стихотворения (1908-1937)"),
		nil,
	},
	{
		"Learning PHP, MySQL, JavaScript, CSS & HTML5 (Robin Nixon)",
		NewBook("Robin Nixon", "Learning PHP, MySQL, JavaScript, CSS & HTML5"),
		nil,
	},
	{
		"Getting Things Done: How to achieve stress-free productivity (Allen, David)",
		NewBook("Allen, David", "Getting Things Done: How to achieve stress-free productivity"),
		nil,
	},
	{
		"FDJKLSJDFJJJFDSJSLKDJFKLSJ ||||| sdjlkf jlsdkjf lkj((((",
		NewBook("", "FDJKLSJDFJJJFDSJSLKDJFKLSJ ||||| sdjlkf jlsdkjf lkj(((("),
		nil,
	},
	{
		"The New Oxford American Dictionary",
		NewBook("", "The New Oxford American Dictionary"),
		nil,
	},
}

var books = []Book{
	NewBook("Iain Banks", "Excession"),
	NewBook("Iain Banks", "State of the art"),
	NewBook("Bertrand Russell", "A History of Western Philosophy"),
	NewBook("Robert Martin", "Clean Code"),
}

func TestBookStorage(t *testing.T) {
	bs := NewBookStorage()
	bs.AddIfMissing(books[0])
	checkLenBs(bs, t, 1)
}

func checkLenBs(bs *BookStorage, t *testing.T, expected int) {
	if len := bs.Len(); len != expected {
		t.Errorf("Wrong Len of BookStorage. Actual: %d. Expected: %d", len, expected)
	}
}

func TestBook(t *testing.T) {
	for _, d := range bookStrings {
		res, err := CreateBook(d.s)
		if res != d.b {
			t.Errorf("Wrong book created! String: %s | Actual: %v | Expected: %v", d.s, res, d.b)
		}
		if err != d.e && err.Error() != d.e.Error() {
			t.Errorf("Wrong error message! String: %s | Actual: %v | Expected: %v", d.s, err, d.e)
		}
	}
}
