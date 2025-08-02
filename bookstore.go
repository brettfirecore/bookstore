package bookstore

type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

func GetBook(catalog []Book, id int) Book {
	for _, b := range catalog {
		if b.ID == id {
			return b
		}
	}
	return Book{} // default value if not found
}
