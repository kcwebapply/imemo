package repository

import (
	"github.com/kcwebapply/imemo/domain/model"
)

func CreateCategoryTable() {
	_, _ = sess.Exec("create table category(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);")
}

func InsertDefaultCategory() {
	var category = model.Category{ID: 1, Name: "default"}
	sess.InsertInto("category").Columns("id", "name").Record(category).Exec()

}

// GetCategories return all memodata.
func GetCategories() ([]model.Category, error) {
	var rows []model.Category
	_, err := sess.Select("*").From("category").Load(&rows)
	return rows, err
}

// GetCategories return all memodata.
func GetCategoryByName(name string) (model.Category, error) {
	var rows model.Category
	_, err := sess.Select("*").From("category").Where("name = ?", name).Load(&rows)
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

// DeleteCategory
func DeleteAllCategory() error {
	_, err := sess.DeleteFrom("category").Exec()
	return err
}
