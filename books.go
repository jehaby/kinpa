package kinpa

import (
	"fmt"
	//	"regexp"

	"strings"
)

type Book struct {
	Id     uint
	Author string
	Title  string
}

func (b *Book) String() string {
	return fmt.Sprintf("%s -- %s", b.Author, b.Title)
}

func NewBook(author, title string) Book {
	return Book{0, author, title}
}

func CreateBook(s string) (Book, error) {
	if oi, ci := strings.LastIndex(s, "("), strings.LastIndex(s, ")"); oi != -1 && oi < ci {
		return Book{0, s[oi+1 : ci], strings.Trim(s[:oi], " ")}, nil
	}
	return Book{0, "", strings.Trim(s, " ")}, nil
}

func (b *Book) Equals(other Book) bool {
	if b.Author == other.Author && b.Title == other.Title {
		return true
	}
	return false
}

type BookStorage struct {
	s map[Book]*Book
	//	byId     map[uint]Book
	//	byAuthor map[string][]uint // should I use list of pointers or pointer on list of books or pointer on list of pointers
}

type NoSuchBook struct {
	book Book
}

func (e NoSuchBook) Error() string {
	return fmt.Sprintf("No such book in storage: %s", e.book)
}

func NewBookStorage() *BookStorage {
	return &BookStorage{make(map[Book]*Book)}
}

func (bs *BookStorage) Books() map[Book]*Book {
	return bs.s
}

func (bs *BookStorage) AddIfMissing(b Book) *Book { // TODO: now it doesn't change pointer like before | It's too complicated and error-prone
	if bp, ok := bs.s[b]; ok {
		return bp
	}
	bs.s[b] = &b
	return &b
}

func (bs *BookStorage) Len() int {
	return len(bs.s)
}
