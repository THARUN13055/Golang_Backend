package datatypes

type UserCredential struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Books struct {
	ListOfBooks []Book
}

func (B *Books) AddBook(book Book) {
	B.ListOfBooks = append(B.ListOfBooks, book)
}
