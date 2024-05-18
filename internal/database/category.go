package database

import "database/sql"

type Category struct {
	ID   int
	Name string
}

type CategoryUsecase interface {
	Create(c *Category) error
}

type CategoryService struct {
	r CategoryRepository
}

func NewCategoryService(r CategoryRepository) CategoryUsecase {
	return &CategoryService{r: r}
}

func (s *CategoryService) Create(c *Category) error {
	return s.r.Create(c)
}

type CategoryRepository interface {
	Create(c *Category) error
}

func NewCategoryDatabase(sql *sql.DB) CategoryRepository {
	return &CategoryDatabase{sql}
}

type CategoryDatabase struct {
	*sql.DB
}

func (d *CategoryDatabase) Create(category *Category) error {
	_, err := d.Exec("INSERT INTO categories (name) VALUES (?)", category.Name)
	return err
}
