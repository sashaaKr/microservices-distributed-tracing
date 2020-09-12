CREATE DATABASE IF NOT EXISTS disributedTracingApp;

USE disributedTracingApp;

CREATE TABLE IF NOT EXISTS disributedTracingApp.people (
    name        VARCHAR(100),
    title       VARCHAR(10),
    description VARCHAR(100),
    PRIMARY KEY (name)
);

DELETE FROM disributedTracingApp.people;

INSERT INTO disributedTracingApp.people VALUES ('Gru', 'Felonius', 'Where are the minions?');
INSERT INTO disributedTracingApp.people VALUES ('Nefario', 'Dr.', 'Why ... why are you so old?');
INSERT INTO disributedTracingApp.people VALUES ('Agnes', '', 'Your unicorn is so fluffy!');
INSERT INTO disributedTracingApp.people VALUES ('Edith', '', "Don't touch anything!");
INSERT INTO disributedTracingApp.people VALUES ('Vector', '', 'Committing crimes with both direction and magnitude!');
INSERT INTO disributedTracingApp.people VALUES ('Dave', 'Minion', 'Ngaaahaaa! Patalaki patalaku Big Boss!!');
