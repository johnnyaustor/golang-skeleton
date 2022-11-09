package user

import (
	"github.com/johnnyaustor/golang-skeleton/server/database"
)

type Repository interface {
	Get() string
}

type repository struct {
	*database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{db}
}

func (r *repository) Get() string {
	return "Hello From Repository"
}
