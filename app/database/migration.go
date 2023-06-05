package database

import (
	"gorm.io/gorm"

	class "github.com/GroupProject2-Kelompok4/BE/features/class/data"
	feedback "github.com/GroupProject2-Kelompok4/BE/features/feedback/data"
	mentee "github.com/GroupProject2-Kelompok4/BE/features/mentee/data"
	user "github.com/GroupProject2-Kelompok4/BE/features/user/data"
	"github.com/GroupProject2-Kelompok4/BE/utils"
)

var log = utils.Log()

func InitMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user.User{},
		&class.Class{},
		&mentee.Mentee{},
		&feedback.Feedback{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	initSuperAdmin(db)

	return err
}
