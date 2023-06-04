package database

import (
	"github.com/GroupProject2-Kelompok4/BE/app/config"
	user "github.com/GroupProject2-Kelompok4/BE/features/user/data"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	TeamID uint        `gorm:"primaryKey"`
	Name   string      `gorm:"type:varchar(20)"`
	Users  []user.User `gorm:"foreignKey:TeamID"`
}

func initTeam(db *gorm.DB) error {
	var count int64
	db.Table("teams").Where("team_id").Count(&count)
	if count == 5 {
		log.Warn("teams already exists")
		return nil
	}

	result := db.Model(&Team{}).Create([]map[string]interface{}{
		{"TeamID": 1, "Name": "Manager"},
		{"TeamID": 2, "Name": "Academic"},
		{"TeamID": 3, "Name": "People"},
		{"TeamID": 4, "Name": "Placement"},
		{"TeamID": 5, "Name": "Admission"},
	})

	if result.Error != nil {
		log.Error("failed to create multiple records team")
		return result.Error
	}

	log.Info("team created successfully")
	return nil

}

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
		Role:        "admin",
		Status:      "manager",
		UserPicture: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png",
		IsDeleted:   false,
		TeamID:      1,
	}

	var count int64
	db.Table("users").Where("role = 'admin' AND status = 'manager'").Count(&count)
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
