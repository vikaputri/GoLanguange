package repository

import "belajargolang/Lanjutan/UnitTest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
