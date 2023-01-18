package members

import (
	"gorm.io/gorm"
)

type Command interface {
	Create(members Members) (int64, error)
	Update(members Members) error
}

type CommandImpl struct {
	DB CommandStore
}

type CommandStore interface {
	Create(value interface{}) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
}

func NewCommand(db *gorm.DB) Command {
	return CommandImpl{db}
}

func (c CommandImpl) Create(members Members) (int64, error) {
	err1 := c.DB.Create(&members).Error
	if err1 != nil {
		return 0, err1
	}

	return members.ID, nil
}

func (c CommandImpl) Update(members Members) error {
	return c.DB.Save(&members).Error
}
