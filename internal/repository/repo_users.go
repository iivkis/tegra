package repository

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	TelegramID int64 `gorm:"index:,unique"`
	Role       string
}

func (u *UserModel) IsAdmin() bool {
	return u.Role == R_ADMIN
}

type repoUsers struct {
	db *gorm.DB
}

func newRepoUsers(db *gorm.DB) *repoUsers {
	return &repoUsers{
		db: db,
	}
}

func (r *repoUsers) Create(user *UserModel) error {
	return r.db.Create(user).Error
}

func (r *repoUsers) GetOne(telegramID interface{}) (user UserModel, err error) {
	err = r.db.Where("telegram_id = ?", telegramID).First(&user).Error
	return
}
