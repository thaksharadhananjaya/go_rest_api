package repository

import (
	"gorm.io/gorm"
	"restapi/internal/domain/user"
)

type gormUserRepository struct {
	db *gorm.DB
}

func (r *gormUserRepository) GetUserByEmail(email string) (user.User, error) {
    var u user.User
    err := r.db.Where("email = ?", email).First(&u).Error
    return u, err
}

func NewGormUserRepository(db *gorm.DB) user.UserRepository {
	return &gormUserRepository{db}
}

func (r *gormUserRepository) Save(u user.User) (user.User, error) {
    if err := r.db.Create(&u).Error; err != nil {
        return user.User{}, err
    }
    return u, nil
}

func (r *gormUserRepository) GetByID(id uint) (user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	return u, err
}

func (r *gormUserRepository) GetAll() ([]user.User, error) {
	var users []user.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *gormUserRepository) Update(id uint, updates map[string]interface{}) error {
    return r.db.Model(&user.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *gormUserRepository) Delete(id uint) error {
	return r.db.Delete(&user.User{}, id).Error
}
