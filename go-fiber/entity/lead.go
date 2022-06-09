package entity

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   string
}

func NewLead(name string, company string, email string, phone string) (*Lead, error) {
	l := &Lead{
		Name: name, Company: company, Email: email, Phone: phone,
	}
	if err := l.Validate(); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Lead) Validate() error {
	if l.Name == "" || l.Company == "" || l.Email == "" || l.Phone == "" {
		return ErrInvalidEntity
	}
	return nil
}
