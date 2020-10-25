package repository

import "github.com/jinzhu/gorm"

// GormRepository Gormリポジトリ実装
type GormRepository struct {
	db *gorm.DB
}

// NewGormRepository リポジトリ実装を初期化して生成します
func NewGormRepository(db *gorm.DB) (Repository, error) {
	repo := &GormRepository{
		db: db,
	}
	return repo, nil
}
