package mysql

import (
	"myapp/pkg/utils"
	"myapp/src/user/entity"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	userEntity = entity.User{
		Model: gorm.Model{
			ID:        uint(1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Nickname:        "test",
		Email:           "test@gmail.com",
		Password:        "test123",
		ProfileImageUrl: "https://image.png",
		Gender:          entity.Gender("Female"),
		Popularity:      0,
		IsVerified:      false,
		Details:         "test",
	}
)

func TestRegister(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	userRepo := NewMysqlUserRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(utils.AnyTime{}, userEntity.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mockObj.ExpectCommit()

	user, err := userRepo.Save(userEntity)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
