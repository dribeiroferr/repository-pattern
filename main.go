package main

import (
	"database/sql"
	"fmt"

	"github.com/dribeiroferr/repository-pattern/repository/repository"
	"github.com/dribeiroferr/repository-pattern/repository/service"
)

func main() {
	// in memory database:
	// db := repository.CategoriesMemoryDB{Categories: []entity.Category{}}
	// repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	// sqlite database:
	db, _ := sql.Open("sqlite3", "./sql.repository.db")
	repository := repository.NewCategorySqlite(db)

	service := service.NewCategoryService(repository)
	cat, _ := service.Create("hello world")

	fmt.Println(cat)
}
