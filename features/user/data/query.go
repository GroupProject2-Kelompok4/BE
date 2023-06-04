package data

import (
	"errors"

	"github.com/GroupProject2-Kelompok4/BE/features/user"
	"github.com/GroupProject2-Kelompok4/BE/utils"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/GroupProject2-Kelompok4/BE/utils/middlewares"
	"gorm.io/gorm"
)

var log = utils.Log()

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// Login implements user.UserData
func (uq *userQuery) Login(request user.UserCore) (user.UserCore, string, error) {
	result := User{}
	query := uq.db.Table("users").Where("email = ?", request.Email).First(&result)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user record not found")
		return user.UserCore{}, "", errors.New("invalid email and password")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no user has been created")
		return user.UserCore{}, "", errors.New("row affected : 0")
	}

	match1 := helper.MatchPassword(request.Password, result.Password)
	log.Sugar().Warnf("match password: %v", match1)
	if !match1 {
		return user.UserCore{}, "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken((result.UserID))
	if err != nil {
		log.Error("error while creating jwt token")
		return user.UserCore{}, "", errors.New("error while creating jwt token")
	}

	log.Sugar().Infof("user login: %s, %s", result.Fullname, result.Email)
	return userModels(result), token, nil
}
