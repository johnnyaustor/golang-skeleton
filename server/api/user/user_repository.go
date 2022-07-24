package user

import "gorm.io/gorm"

type Repository interface {
	Get() string
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get() string {
	return "Hello From Repository"
}
