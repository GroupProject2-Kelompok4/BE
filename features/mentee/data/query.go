package data

import (
	"errors"

	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/GroupProject2-Kelompok4/BE/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var log = utils.Log()

type menteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeData {
	return &menteeQuery{
		db: db,
	}
}

// RegisterMentee implements mentee.MenteeData
func (mq *menteeQuery) RegisterMentee(request mentee.MenteeCore) (mentee.MenteeCore, error) {
	menteeId, err := uuid.NewUUID()
	if err != nil {
		log.Warn("error while create uuid for class")
		return mentee.MenteeCore{}, nil
	}

	request.MenteeID = menteeId.String()
	req := menteeEntities(request)
	query := mq.db.Table("mentees").Create(&req)
	if query.Error != nil {
		log.Error("error insert data, duplicated")
		return mentee.MenteeCore{}, errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no user has been created")
		return mentee.MenteeCore{}, errors.New("row affected : 0")
	}

	return menteeModels(req), nil
}
