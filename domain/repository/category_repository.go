package repository

import (
	"github.com/kcwebapply/imemo/domain/model"
)

func CreateCategoryTable() {
	_, _ = sess.Exec("create table category(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);")
}

func CreateDefaultTable() {
	_, _ = sess.Exec("create table category(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);")
}

// GetCategories return all memodata.
func GetCategories() ([]model.Memo, error) {
	var rows []model.Memo
	_, err := sess.Select("*").From("category").Load(&rows)
	return rows, err
}

// AddCategory add category entity.
func AddCategory(category model.Category) error {

	_, err := sess.InsertInto("category").Columns("name").Record(category).Exec()
	return err
}

// DeleteCategory
func DeleteCategoryById(id int) error {
	_, err := sess.DeleteFrom("category").Where("id = ?", id).Exec()
	return err
}
