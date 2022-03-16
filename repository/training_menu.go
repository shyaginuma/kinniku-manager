package repository

import (
	"database/sql"

	"github.com/kinniku-manager/model"
)

type TrainingMenuRepository struct {
	Database *sql.DB
}

func (repository TrainingMenuRepository) ReadAll() ([]model.TrainingMenu, error) {
	query := `
	SELECT
		training_menus.id,
		training_menus.name,
		training_menus.description,
		GROUP_CONCAT(training_menu_set_relations.set_id) as set_ids
	FROM
		training_menus
	LEFT JOIN
		training_menu_set_relations
		ON training_menus.id = training_menu_set_relations.menu_id
	`
	rows, err := repository.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
