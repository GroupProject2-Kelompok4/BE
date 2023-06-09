package database

import (
	"github.com/GroupProject2-Kelompok4/BE/app/config"
	user "github.com/GroupProject2-Kelompok4/BE/features/user/data"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func initSuperAdmin(db *gorm.DB) error {
	userID, err := uuid.NewUUID()
	if err != nil {
		log.Warn("error while create uuid for admin")
		return nil
	}

	// hashed, err := helper.HashPassword("secret")
	hashed, err := helper.HashPassword(config.ADMINPASSWORD)
	if err != nil {
		log.Warn("error while hashing password admin")
		return nil
	}

	admin := user.User{
		UserID:      userID.String(),
		Fullname:    "admin",
		Email:       "admin@gmail.com",
		Password:    hashed,
		Team:        "manager",
		Role:        "admin",
		Status:      "active",
		UserPicture: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png",
		IsDeleted:   false,
	}

	var count int64
	db.Table("users").Where("role = 'admin' AND team = 'manager'").Count(&count)
	if count > 0 {
		log.Warn("super admin already exists")
		return nil
	}

	result := db.Create(&admin)
	if result.Error != nil {
		log.Error("failed to create super admin")
		return result.Error
	}

	log.Info("super admin created successfully")
	return nil
}
