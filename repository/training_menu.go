package repository

import (
	"database/sql"

	"github.com/kinniku-manager/model"
)

type TrainingMenuRepository struct {
	Database *sql.DB
}

func (repository TrainingMenuRepository) ReadAll() ([]model.TrainingMenu, error) {
	return []model.TrainingMenu{}, nil
}

func (repository TrainingMenuRepository) Read(id int64) (model.TrainingMenu, error) {
	return model.TrainingMenu{}, nil
}

func (repository TrainingMenuRepository) Create(newTraningMenu model.TrainingMenu) error {
	return nil
}

func (repository TrainingMenuRepository) Update(modifiedTraningMenu model.TrainingMenu) error {
	return nil
}

func (repository TrainingMenuRepository) Delete(id int64) error {
	return nil
}
