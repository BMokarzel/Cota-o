package service

import "database/sql"

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	service := Service{
		DB: db,
	}
	return &service
}
