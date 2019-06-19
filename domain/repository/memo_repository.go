package repository

import (
	"fmt"

	"github.com/kcwebapply/imemo/domain/model"
)

//CreateMemoTable create page table on sqlite3 db
func CreateMemoTable() {
	_, _ = sess.Exec("create table memo(id INTEGER PRIMARY KEY AUTOINCREMENT, memo TEXT, categoryid INTEGER);")
}

// Get memos return all memodata.
func GetMemos() ([]model.Memo, error) {
	var rows []model.Memo
	_, err := sess.Select("*").From("memo").Load(&rows)
	return rows, err
}

func AddMemo(memo model.Memo) error {
	//.Columns("url", "title", "tags", "content").Record(newPage).Exec()
	_, err := sess.InsertInto("memo").Columns("memo", "categoryid").Record(memo).Exec()
	return err
}

// DeleteMemoById
func DeleteMemoById(id int) error {
	_, err := sess.DeleteFrom("memo").Where("id = ?", id).Exec()
	return err
}

// GetMemosByCategoryName returns page grepped title by input word.
func GetMemosByCategoryName(name string) ([]model.Category, error) {
	var rows []model.Category
	likeItem := fmt.Sprintf("%%%s%%", name)
	_, err := sess.Select("*").From("category").Where("name like ?", likeItem).Load(&rows)
	return rows, err
}

// GetMemosByCategoryId returns page grepped title by input word.
func GetMemosByCategoryId(id int) ([]model.Category, error) {
	var rows []model.Category
	_, err := sess.Select("*").From("category").Where("id = ?", id).Load(&rows)
	return rows, err
}

//GetPagesByTag returns page entities grepped by input tag word.
/*func GetPagesByTag(tag string) ([]model.Page, error) {
	var rows []model.Page
	likeItem := fmt.Sprintf("%%%s%%", tag)
	_, err := sess.Select("*").From("page").Where("tags like ?", likeItem).Load(&rows)
	return rows, err
}

// GetPagesByContentSearch returns page entities grepped by html-content-search.
func GetPagesByContentSearch(word string) ([]model.Page, error) {
	var rows []model.Page
	likeItem := fmt.Sprintf("%%%s%%", word)
	_, err := sess.Select("*").From("page").Where("content like ?", likeItem).Load(&rows)
	return rows, err
}

// AddPage saved bookrmark user input.
func AddPage(newPage model.Page) error {
	_, err := sess.InsertInto("page").Columns("url", "title", "tags", "content").Record(newPage).Exec()
	return err
}

// RemovePage remove bookmark
func RemovePage(id string) error {
	if _, err := sess.DeleteFrom("page").Where("id = ?", id).Exec(); err != nil {
		return err
	}
	return nil
}*/
