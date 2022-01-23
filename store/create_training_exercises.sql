DROP TABLE IF EXISTS training_exercises;

CREATE TABLE training_exercises (
    id          INT AUTO_INCREMENT not null,
	name        VARCHAR(128) not null,
	description VARCHAR(255) not null,
	target      VARCHAR(128) not null,
	category    VARCHAR(128) not null,
	difficulty  VARCHAR(128) not null,
    PRIMARY KEY (`id`)
);

INSERT INTO training_exercises
    (name, description, target, category, difficulty)
VALUES
    ('Barbell Curl', 'Barbell Curl', 'biceps', 'barbell', 'beginner');
