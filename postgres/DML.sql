-- Data Modification Language

-- INSERT
INSERT INTO courses (name, slug) VALUES ('Bash', 'bash'), ('PHP', 'php'), ('Ruby', 'ruby');

-- UPDATE
UPDATE courses SET body = 'updated!', name = 'Bash' WHERE slug = 'bash';

-- DELETE
DELETE FROM courses WHERE slug = 'bash';

-- TRUNCATE
TRUNCATE courses;