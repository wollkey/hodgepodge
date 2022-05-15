CREATE TABLE IF NOT EXISTS "user"
(
    id   SERIAL PRIMARY KEY,
    name text,
    age  integer
);

INSERT INTO "user" (name, age)
VALUES ('Лепёха', 20),
       ('Алёша', 55),
       ('Олегсей', 99);

CREATE TABLE IF NOT EXISTS "car"
(
    brand text,
    model text,
    user_id integer,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

INSERT INTO "car" (user_id, brand, model)
VALUES (1, 'BMV', 'XXX'),
       (2, 'WAG', 'Z-521'),
       (3, 'LADA', 'V-123');