package repository

import (
	"github.com/kcwebapply/imemo/domain/model"
)

const (
	MEMO = "memo"
)

//CreateMemoTable create page table on sqlite3 db
func CreateMemoTable() {
	_, _ = sess.Exec("create table memo(id INTEGER PRIMARY KEY AUTOINCREMENT, memo TEXT, categoryid INTEGER);")
}

// Get memos return all memodata.
func GetMemos() ([]model.Memo, error) {
	var rows []model.Memo
	_, err := sess.Select("*").From(MEMO).Load(&rows)
	return rows, err
}

func GetMemo(id int) (model.Memo, error) {
	var row model.Memo
	_, err := sess.Select("*").From(MEMO).Where("id = ?", id).Load(&row)
	return row, err
}

func AddMemo(memo model.Memo) error {
	//.Columns("url", "title", "tags", "content").Record(newPage).Exec()
	_, err := sess.InsertInto(MEMO).Columns("memo", "categoryid").Record(memo).Exec()
	return err
}

// DeleteMemoById
func DeleteMemoById(id int) error {
	_, err := sess.DeleteFrom(MEMO).Where("id = ?", id).Exec()
	return err
}

// DeleteMemoById
func DeleteAllMemo() error {
	_, err := sess.DeleteFrom(MEMO).Exec()
	_, err = sess.DeleteFrom("sqlite_sequence").Where("name = ?", MEMO).Exec()
	//delete from sqlite_sequence where name='TABLE_NAME';
	return err
}

// GetMemosByCategoryId returns page grepped title by input word.
func GetMemosByCategoryId(id int) ([]model.Memo, error) {
	var rows []model.Memo
	_, err := sess.Select("*").From(MEMO).Where("categoryid = ?", id).Load(&rows)
	return rows, err
}
