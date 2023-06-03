package mysql

import (
	"myapp/src/user"
	"myapp/src/user/entity"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) user.UserRepository {
	return &MysqlUserRepository{db}
}

func (ur *MysqlUserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *MysqlUserRepository) Save(user entity.User) (entity.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ur *MysqlUserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	res := ur.db.Where("email = ?", email).First(&user)
	err := res.Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ur *MysqlUserRepository) FindById(id uint) (entity.User, error) {
	var user entity.User
	res := ur.db.Where("id = ?", id).First(&user)
	err := res.Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ur *MysqlUserRepository) Update(user entity.User, id uint) (entity.User, error) {
	res := ur.db.Model(&user).Where("id = ?", id).Updates(user)
	err := res.Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
