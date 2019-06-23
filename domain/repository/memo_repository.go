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

func GetMemo(id int) (model.Memo, error) {
	var row model.Memo
	_, err := sess.Select("*").From("memo").Where("id = ?", id).Load(&row)
	return row, err
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

// DeleteMemoById
func DeleteAllMemo() error {
	_, err := sess.DeleteFrom("memo").Exec()
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
