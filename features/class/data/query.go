package data

import (
	"errors"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var log = utils.Log()

type classQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.ClassData {
	return &classQuery{
		db: db,
	}
}

// RegisterClass implements class.ClassData
func (cq *classQuery) RegisterClass(request class.ClassCore) (class.ClassCore, string, error) {
	classId, err := uuid.NewUUID()
	if err != nil {
		log.Warn("error while create uuid for class")
		return class.ClassCore{}, "", nil
	}

	request.ClassID = classId.String()
	req := classEntities(request)
	query := cq.db.Table("classes").Create(&req)
	if query.Error != nil {
		log.Error("error insert data, duplicated")
		return class.ClassCore{}, "", errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no user has been created")
		return class.ClassCore{}, "", errors.New("row affected : 0")
	}

	var uFullname string
	if err := cq.db.Table("users").Select("fullname").Where("user_id = ?", request.UserID).Scan(&uFullname).Error; err != nil {
		log.Sugar().Errorf("failed to retrieve username for UserID: %s", request.UserID)
		return class.ClassCore{}, "", err
	}

	request.PIC = uFullname
	log.Sugar().Infof("new class has been created: %s", req.ClassID)
	return classModels(req), request.PIC, nil
}
