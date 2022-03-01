create database if not exists trainings;
create table if not exists trainings.training_exercises(
    id          INT AUTO_INCREMENT not null,
	name        VARCHAR(128)       not null,
	description VARCHAR(255)       not null,
	target      VARCHAR(128)       not null,
	category    VARCHAR(128)       not null,
	difficulty  VARCHAR(128)       not null,
    PRIMARY KEY (`id`)
);
use trainings;
