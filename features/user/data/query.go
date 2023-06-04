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

	token, err := middlewares.CreateToken(result.UserID, result.Role)
	if err != nil {
		log.Error("error while creating jwt token")
		return user.UserCore{}, "", errors.New("error while creating jwt token")
	}

	log.Sugar().Infof("user has been logged in: %s", result.UserID)
	return userModels(result), token, nil
}

// Register implements user.UserData
func (uq *userQuery) Register(request user.UserCore) (user.UserCore, error) {
	hashed, err := helper.HashPassword(request.Password)
	if err != nil {
		log.Error("error while hashing password")
		return user.UserCore{}, errors.New("error while hashing password")
	}

	request.Password = hashed
	request.UserPicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"
	req := userEntities(request)
	query := uq.db.Table("users").Create(&req)
	if query.Error != nil {
		log.Error("error insert data, duplicated")
		return user.UserCore{}, errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no user has been created")
		return user.UserCore{}, errors.New("row affected : 0")
	}

	log.Sugar().Infof("new user has been created: %s", req.UserID)
	return userModels(req), nil
}
