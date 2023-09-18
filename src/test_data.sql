DROP TABLE IF EXISTS author CASCADE;
DROP TABLE IF EXISTS book CASCADE;

CREATE TABLE public.author
(
    id   SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE public.book
(
    id        SERIAL UNIQUE,
    name      VARCHAR(255) NOT NULL,
    author_id INT,

    CONSTRAINT author_pk FOREIGN KEY (author_id) REFERENCES public.author (id)

);


INSERT INTO author (name)
VALUES ('Народ');
INSERT INTO author (name)
VALUES ('Джоан Роулинг');
INSERT INTO author (name)
VALUES ('Джек Лондон');

INSERT INTO book (name, author_id)
VALUES ('колобок', 1);
INSERT INTO book (name, author_id)
VALUES ('гарри поттер', 2);
INSERT INTO book (name, author_id)
VALUES ('брилианты', 3);
