package user

import "myapp/src/user/entity"

type UserRepository interface {
	FindAll() ([]entity.User, error)
	// FindById(id uint) (entity.User, error)
	FindByUsernameOrEmail(username string, email string) (entity.User, error)
	Save(user entity.User) (entity.User, error)
	// Update(user entity.User) (entity.User, error)
	// Delete(user entity.User) (entity.User, error)
}
