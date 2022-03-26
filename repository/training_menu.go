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
	GROUP BY
		1, 2, 3
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
	WHERE
		training_menus.id = ?
	GROUP BY
		1, 2, 3
	`
	row := repository.Database.QueryRow(query, id)
	var menu_raw trainingMenuQueryResults
	err := row.Scan(
		&menu_raw.ID,
		&menu_raw.Name,
		&menu_raw.Description,
		&menu_raw.Menu,
	)
	if err != nil {
		return model.TrainingMenu{}, err
	}
	var menu_slice []int64
	for _, set_id_str := range strings.Split(menu_raw.Menu, ",") {
		set_id_int, err := strconv.ParseInt(set_id_str, 10, 64)
		if err != nil {
			return model.TrainingMenu{}, err
		}
		menu_slice = append(menu_slice, set_id_int)
	}

	menu := model.TrainingMenu{
		ID:          menu_raw.ID,
		Name:        menu_raw.Name,
		Description: menu_raw.Description,
		Menu:        menu_slice,
	}
	return menu, nil
}

func (repository TrainingMenuRepository) Create(newTraningMenu model.TrainingMenu) error {
	stmt_menu, err := repository.Database.Prepare("INSERT INTO training_menus VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err := stmt_menu.Exec(
		newTraningMenu.ID,
		newTraningMenu.Name,
		newTraningMenu.Description,
	); err != nil {
		return err
	}

	stmt_relation, err := repository.Database.Prepare("INSERT INTO training_menu_set_relations VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	for _, set_id := range newTraningMenu.Menu {
		if _, err := stmt_relation.Exec(
			0,
			newTraningMenu.ID,
			set_id,
		); err != nil {
			return err
		}
	}
	return nil
}

func (repository TrainingMenuRepository) Update(modifiedTraningMenu model.TrainingMenu) error {
	stmt_menu, err := repository.Database.Prepare(`
		UPDATE training_menus
		SET name = ?,
			description = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	if _, err := stmt_menu.Exec(
		modifiedTraningMenu.Name,
		modifiedTraningMenu.Description,
		modifiedTraningMenu.ID,
	); err != nil {
		return err
	}

	// TODO: 直す。relationのIDは基本わからないので、Clean up & Insert?
	stmt_relation_delete, err := repository.Database.Prepare("DELETE FROM training_menu_set_relations WHERE menu_id = ?")
	if err != nil {
		return err
	}
	if _, err := stmt_relation_delete.Exec(modifiedTraningMenu.ID); err != nil {
		return err
	}

	stmt_relation, err := repository.Database.Prepare("INSERT INTO training_menu_set_relations VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	for _, set_id := range modifiedTraningMenu.Menu {
		if _, err := stmt_relation.Exec(
			0,
			modifiedTraningMenu.ID,
			set_id,
		); err != nil {
			return err
		}
	}
	return nil
}

func (repository TrainingMenuRepository) Delete(trainingMenuID int64) error {
	stmt_relation, err := repository.Database.Prepare("DELETE FROM training_menu_set_relations WHERE menu_id = ?")
	if err != nil {
		return err
	}
	if _, err := stmt_relation.Exec(trainingMenuID); err != nil {
		return err
	}

	stmt_menu, err := repository.Database.Prepare("DELETE FROM training_menus WHERE id = ?")
	if err != nil {
		return err
	}
	if _, err := stmt_menu.Exec(trainingMenuID); err != nil {
		return err
	}
	return nil
}
