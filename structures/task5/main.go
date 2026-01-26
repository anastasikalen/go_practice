package main

import "fmt"

func main() {
	b1 := LibraryBook{
		"test", "test author", true, "Alice",
	}
	b2 := LibraryBook{
		"test2", "test2", false, "",
	}

	fmt.Println(b1.Info())
	fmt.Println(b2.Info())

	err := b1.Borrow("David")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b1.Info())
	err = b1.ReturnBook("Alidce")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b1.Info())

	err = b1.ReturnBook("Alice")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b1.Info())

}

type LibraryBook struct {
	Title    string
	Author   string
	Borrowed bool
	Borrower string
}

func (b *LibraryBook) Borrow(name string) error {
	if b.Borrowed {
		return fmt.Errorf("book already borrowed")
	}
	b.Borrowed = true
	b.Borrower = name
	return nil
}
func (b *LibraryBook) ReturnBook(name string) error {
	if !b.Borrowed || b.Borrower != name {
		return fmt.Errorf("book not borrowed") //КОД ДНЯ 9426
	}
	b.Borrowed = false
	b.Borrower = ""
	return nil
}

func (b LibraryBook) Info() string {
	if !b.Borrowed {
		return fmt.Sprintf("Title:%s, Author:%s, Available", b.Title, b.Author)
	}
	return fmt.Sprintf("Title:%s, Author:%s, Borrowed by %s", b.Title, b.Author, b.Borrower)
}
