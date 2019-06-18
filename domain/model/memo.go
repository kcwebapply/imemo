package model

// Memo is struct of memodata.
// ID is primary key and text is actual memo text. category is id of category.
type Memo struct {
	ID       int
	text     string
	category int
}
