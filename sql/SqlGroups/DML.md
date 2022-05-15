# DML

`INSERT` - вставка данных в таблицу
```sql
INSERT INTO "user" (name, age) VALUES ('Alex', 99);
```

`UPDATE` - обновление строк в таблице
```sql
UPDATE "user" SET name = 'Ali' WHERE name = 'Alex';
```

`DELETE` - удаление строк из таблицы
```sql
DELETE FROM "user" WHERE name = 'Alex';
```