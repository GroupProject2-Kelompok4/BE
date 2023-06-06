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

// ListClasses implements class.ClassData
func (cq *classQuery) ListClasses(limit int, offset int) ([]class.ClassCore, uint, error) {
	classes := []Class{}
	var count int64
	query := cq.db.Table("classes").
		Where("is_deleted = 0").
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Preload("User").
		Find(&classes).
		Count(&count)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("list classes not found")
		return nil, 0, errors.New("not found, error while retrieving list classes")
	}

	result := make([]class.ClassCore, len(classes))
	for i, class := range classes {
		result[i] = classModels(class)
	}

	return result, uint(count), nil
}

// DeleteClass implements class.ClassData
func (cq *classQuery) DeleteClass(classId string) error {
	query := cq.db.Table("classes").
		Where("class_id = ? AND is_deleted = 0", classId).
		Updates(map[string]interface{}{"is_deleted": true})

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("class record not found")
		return errors.New("class record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no class has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while delete class")
		return errors.New("duplicate data entry")
	}

	return nil
}

// GetClass implements class.ClassData
func (cq *classQuery) GetClass(classId string) (class.ClassCore, error) {
	cls := Class{}
	query := cq.db.Where("class_id = ?", classId).Preload("User").First(&cls)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("list classes not found")
		return class.ClassCore{}, errors.New("not found, error while retrieving list classes")
	}

	if query.RowsAffected == 0 {
		log.Warn("no class has been created")
		return class.ClassCore{}, errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while delete class")
		return class.ClassCore{}, errors.New("duplicate data entry")
	}

	return classModels(cls), nil
}
