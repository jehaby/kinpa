package kinpa

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

type Parser struct {
	debug bool
}

func NewParser(mode bool) *Parser {
	return &Parser{mode}
}

func (p *Parser) ParseClippings(clippings io.Reader) (*HighlightStorage, *BookStorage, error) {
	scanner := bufio.NewScanner(clippings)
	bookStorage := NewBookStorage()
	highlightStorage := NewHighlightStorage()

	si := 1
	highlight := Highlight{}
	hasErr := false
	for scanner.Scan() {
		currentString := scanner.Text()
		if len(currentString) > 3 && currentString[0:3] == "===" {
			si = 1
			if !highlight.IsZero() {
				logd(fmt.Sprintf("Adding highlight: %p", &highlight), p.debug)
				highlightStorage.Add(highlight)
				highlight = Highlight{}
				logd(fmt.Sprintf("Highlight after zeeroing: %p", &highlight), p.debug)
			} // TODO: else log error if highlight is incomplete
			continue
		}

		if hasErr {
			continue
		}

		switch si {
		case 1:
			book, e := CreateBook(currentString) // TODO: this is ugly and probably stupid
			if e != nil {
				logd(fmt.Sprintf("Couldn't create a book from string '%s'", currentString), p.debug)
				hasErr = true
				continue
			}
			highlight.Book = bookStorage.AddIfMissing(book)
		case 2:
			highlight.Page, highlight.Location, highlight.Time, _ = parseMetaString(currentString)
		case 4:
			highlight.SetText(currentString)
		}

		si++
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return highlightStorage, bookStorage, nil

}

func logd(s string, isDebug bool) {
	if isDebug {
		log.Println(s)
	}
}
