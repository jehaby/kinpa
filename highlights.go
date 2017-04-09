package kinpa

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Highlight struct {
	Text     string
	Page     int
	Location string
	Time     time.Time
	Book     *Book
}

func (h *Highlight) IsZero() bool {
	return h.Book == nil && h.Text == "" && h.Page == 0 && h.Location == "" && h.Time.IsZero()
}

func (h *Highlight) SetText(text string) {
	h.Text = strings.Trim(text, " ,.'\"“’ ”")
}

func (h *Highlight) String() string {
	return fmt.Sprintf("%s | %s | %s | %s", h.Book.String(), h.Location, h.Time.Format(time.RFC3339), h.Text)
}

type HighlightStorage struct {
	storage map[*Highlight]struct{}
	byText  map[string][]*Highlight
}

func NewHighlightStorage() *HighlightStorage { // TODO: default argument
	return &HighlightStorage{
		make(map[*Highlight]struct{}),
		make(map[string][]*Highlight),
	}
}

func (hs *HighlightStorage) Add(h Highlight) error {
	if hs.Contains(h) {
		return fmt.Errorf("Highlight already exists: ", h)
	}
	hs.storage[&h] = struct{}{}
	hs.byText[h.Text] = append(hs.byText[h.Text], &h)
	return nil
}

func (hs *HighlightStorage) Contains(h Highlight) bool {
	highlights, ok := hs.byText[h.Text]
	if !ok {
		return false
	}
	for _, highlightP := range highlights {
		if *highlightP == h {
			return true
		}
	}
	return false
}

func (hs *HighlightStorage) Len() int {
	return len(hs.storage)
}

func (hs *HighlightStorage) GetByText(t string) ([]*Highlight, error) {
	res, ok := hs.byText[t]
	if !ok {
		return nil, fmt.Errorf("Highlight with such text doesn't exist (%s)", t)
	}
	return res, nil
}

func PrintHighlights(hs *HighlightStorage, w io.Writer) {
	for h, _ := range hs.storage {
		w.Write([]byte(h.String() + "\n"))
	}
}
