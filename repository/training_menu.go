package repository

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/kinniku-manager/model"
)

type TrainingMenuRepository struct {
	Database *sql.DB
}

type trainingMenuQueryResults struct {
	ID          int64
	Name        string
	Description string
	Menu        string
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
	var menus []model.TrainingMenu
	for rows.Next() {
		var menu_raw trainingMenuQueryResults
		err := rows.Scan(
			&menu_raw.ID,
			&menu_raw.Name,
			&menu_raw.Description,
			&menu_raw.Menu,
		)
		if err != nil {
			return nil, err
		}
		var menu_slice []int64
		for _, set_id_str := range strings.Split(menu_raw.Menu, ",") {
			set_id_int, err := strconv.ParseInt(set_id_str, 10, 64)
			if err != nil {
				return nil, err
			}
			menu_slice = append(menu_slice, set_id_int)
		}

		menu := model.TrainingMenu{
			ID:          menu_raw.ID,
			Name:        menu_raw.Name,
			Description: menu_raw.Description,
			Menu:        menu_slice,
		}

		menus = append(menus, menu)
	}
	return menus, nil
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
