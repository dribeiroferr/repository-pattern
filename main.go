package main

import (
	"fmt"

	"github.com/dribeiroferr/repository-pattern/repository/entity"
	"github.com/dribeiroferr/repository-pattern/repository/repository"
	"github.com/dribeiroferr/repository-pattern/repository/service"
)

func main() {

	db := repository.CategoriesMemoryDB{Categories: []entity.Category{}}
	repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	service := service.NewCategoryService(repositoryMemory)
	cat, _ := service.Create("hello world")

	fmt.Println(cat)
}
