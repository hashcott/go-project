package repository

import (
	"github.com/hashcott/go-project/go-fiber/entity"
	"gorm.io/gorm"
)

type LeadSqlite struct {
	db *gorm.DB
}

func NewLeadSqlite(db *gorm.DB) *LeadSqlite {
	return &LeadSqlite{db: db}
}

func (l *LeadSqlite) Create(e *entity.Lead) (*entity.Lead, error) {
	l.db.Create(e)
	return e, nil
}

func (l *LeadSqlite) Get(id entity.ID) (*entity.Lead, error) {
	var lead entity.Lead
	l.db.First(&lead, id)

	if lead.ID == 0 {
		return nil, entity.ErrNotFound
	}
	return &lead, nil
}

func (l *LeadSqlite) Update(e *entity.Lead) error {
	if e.ID == 0 {
		return entity.ErrNotFound
	}
	l.db.Save(e)
	return nil
}

func (l *LeadSqlite) Delete(id entity.ID) error {
	if id == 0 {
		return entity.ErrNotFound
	}
	lead, err := l.Get(id)
	if err != nil {
		return err
	}
	l.db.Delete(&lead)
	return nil
}
