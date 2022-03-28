create database if not exists trainings;
create table if not exists trainings.training_exercises(
	id          INT AUTO_INCREMENT not null,
	name        VARCHAR(128)       not null,
	description VARCHAR(255)       not null,
	target      VARCHAR(128)       not null,
	category    VARCHAR(128)       not null,
	difficulty  VARCHAR(128)       not null,
	PRIMARY KEY (id)
);
create table if not exists trainings.training_sets(
	id           INT AUTO_INCREMENT not null,
	name         VARCHAR(128)       not null,
	description  VARCHAR(255)       not null,
	exercise_id  INT                not null,
	reps         INT                not null,
	weight_kg    FLOAT              not null,
	interval_min FLOAT              not null,
	PRIMARY KEY  (id),
	FOREIGN KEY  (exercise_id)
		REFERENCES trainings.training_exercises(id)
);
create table if not exists trainings.training_menus(
	id           INT AUTO_INCREMENT not null,
	name         VARCHAR(128)       not null,
	description  VARCHAR(255)       not null,
	PRIMARY KEY  (id)
);
create table if not exists trainings.training_menu_set_relations(
	id           INT AUTO_INCREMENT not null,
	menu_id      INT                not null,
	set_id       INT                not null,
	PRIMARY KEY  (id),
	FOREIGN KEY  (menu_id)
		REFERENCES trainings.training_menus(id),
	FOREIGN KEY  (set_id)
		REFERENCES trainings.training_sets(id)
);
use trainings;
