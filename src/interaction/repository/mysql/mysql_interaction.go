package mysql

import (
	"myapp/src/interaction"
	"myapp/src/interaction/entity"
	userEntity "myapp/src/user/entity"
	"time"

	"gorm.io/gorm"
)

type MysqlInteractionRepository struct {
	db *gorm.DB
}

func NewMysqlInteractionRepository(db *gorm.DB) interaction.InteractionRepository {
	return &MysqlInteractionRepository{db}
}

func (ur *MysqlInteractionRepository) FindByUserGivenId(id uint) ([]entity.Interaction, error) {
	var interactions []entity.Interaction
	err := ur.db.Where("user_given_id = ?", id).Find(&interactions).Error
	return interactions, err
}

func (ur *MysqlInteractionRepository) FindByUserReceivedId(id uint) ([]entity.Interaction, error) {
	var interactions []entity.Interaction
	err := ur.db.Where("user_received_id = ?", id).Find(&interactions).Error
	return interactions, err
}

func (ur *MysqlInteractionRepository) Save(interaction entity.Interaction) (entity.Interaction, error) {
	err := ur.db.Create(&interaction).Error
	if err != nil {
		return entity.Interaction{}, err
	}
	return interaction, nil
}

func (ur *MysqlInteractionRepository) FindCurrentInteraction(userId uint) ([]entity.Interaction, error) {
	var interactions []entity.Interaction

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return []entity.Interaction{}, err
	}

	now := time.Now().In(location)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	res := ur.db.Where("user_given_id = ? AND created_at >= ?", userId, today).Find(&interactions)
	err = res.Error
	if res.RowsAffected == 0 {
		return []entity.Interaction{}, err
	}
	return interactions, err
}

func (ur *MysqlInteractionRepository) FindRandomPeople(userId uint) (userEntity.User, error) {
	var user userEntity.User

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return userEntity.User{}, err
	}

	now := time.Now().In(location)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	res := ur.db.Debug().Raw("SELECT * FROM users WHERE id != ? AND id NOT IN (SELECT user_received_id FROM interactions WHERE user_given_id = ? AND created_at >= ?) ORDER BY RAND() LIMIT 1", userId, userId, today).Scan(&user)
	err = res.Error
	if res.RowsAffected == 0 {
		return userEntity.User{}, err
	}
	return user, err
}
