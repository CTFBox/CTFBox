package model

import (
	"github.com/gofrs/uuid"
)

// Problem 問題一覧
type Problem struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Category string `json:"category" gorm:"type:varchar(36);not null"`
	// TODO: add columns
}
