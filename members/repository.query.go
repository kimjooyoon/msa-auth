package members

import "gorm.io/gorm"

type Query interface {
	CountByEmail(email string) (int64, error)
	FindByEmail(email string) (m *Members, e error)
	FindById(id int64) (m *Members, e error)
}

type MemberQuery struct {
	*gorm.DB
}

func NewQuery(db *gorm.DB) Query {
	return MemberQuery{db}
}

func (q MemberQuery) CountByEmail(email string) (int64, error) {
	var count int64
	err1 := q.DB.Model(&Members{}).Where("email = ?", email).Count(&count).Error

	return count, err1
}

func (q MemberQuery) FindByEmail(email string) (m *Members, e error) {
	err1 := q.DB.Where("email = ?", email).First(&m).Error
	return m, err1
}

func (q MemberQuery) FindById(id int64) (m *Members, e error) {
	err1 := q.DB.Where("id = ?", id).First(&m).Error
	return m, err1
}
