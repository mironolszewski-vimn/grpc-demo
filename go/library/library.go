package library

import (
	pb "grpc-vimn-demo-go/protos"
	"strings"
)

var authors = []*pb.Author{
	{Id: 1, Name: "John Ronald Reuel Tolkien"},
	{Id: 2, Name: "William Golding"},
	{Id: 3, Name: "William Shakespeare"},
}

var books = []*pb.Book{
	{Id: 1, Title: "The Hobbit", Authors: []*pb.Author{authors[0]}},
	{Id: 2, Title: "The Lord Of The Rings", Authors: []*pb.Author{authors[0]}},
	{Id: 3, Title: "Lord Of The Flies", Authors: []*pb.Author{authors[1]}},
	{Id: 4, Title: "Hamlet", Authors: []*pb.Author{authors[2]}},
}

type bookQueryBuilder struct {
	title      string
	authorName string
}

func NewBookQueryBuilder() *bookQueryBuilder {
	return &bookQueryBuilder{}
}

func (b *bookQueryBuilder) FilterTitle(title string) {
	b.title = title
}

func (b *bookQueryBuilder) FilterAuthorName(authorName string) {
	b.authorName = authorName
}

func (b *bookQueryBuilder) GetBooks() []*pb.Book {
	books := books

	if b.title != "" {
		filtered := make([]*pb.Book, 0)
		for _, book := range books {
			if strings.Contains(strings.ToLower(book.Title), strings.ToLower(b.title)) {
				filtered = append(filtered, book)
			}
		}
		books = filtered
	}

	if b.authorName != "" {
		filtered := make([]*pb.Book, 0)
	bookLoop:
		for _, book := range books {
			for _, author := range book.Authors {
				if strings.Contains(strings.ToLower(author.Name), strings.ToLower(b.authorName)) {
					filtered = append(filtered, book)
					continue bookLoop
				}
			}
		}
		books = filtered
	}

	return books
}
