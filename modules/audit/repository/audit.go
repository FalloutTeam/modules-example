package repository

import (
	"database/sql"
)

type AuditRepository interface {
	DoSome(some string) (string, error)
}

type auditRepository struct {
	db *sql.DB
}

func NewAuditRepository(db *sql.DB) AuditRepository {
	return &auditRepository{
		db: db,
	}
}

func (r *auditRepository) DoSome(some string) (string, error) {
	return "some", nil
}
