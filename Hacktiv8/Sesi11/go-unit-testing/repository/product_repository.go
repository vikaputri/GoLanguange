package repository

import (
	"go-unit-testing/entity"
)

type ProductRepository interface {
	FindById(id string) *entity.Product
}
