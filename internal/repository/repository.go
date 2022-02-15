package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	Users *repoUsers
}

func NewRepository(dbname string) *Repository {
	//init database
	db, err := gorm.Open(sqlite.Open(dbname))
	if err != nil {
		panic(err)
	}

	//migrations
	if err := db.AutoMigrate(
		&UserModel{},
	); err != nil {
		panic(err)
	}

	//return repo
	return &Repository{
		Users: newRepoUsers(db),
	}
}
