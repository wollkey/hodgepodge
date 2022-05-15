## Установка PostgreSQl

- `sudo aptitude install postgres`

## Запуск Postgres
При установке создаётся пользователь postgres.

Чтобы зайти в postgres, нужно перейти под пользователя postgres   

`sudo -u postgres -i`

Запустить postgres  - `psql`  

Postgres всегда подключается к конкретной базе
`psql postgres`

Создать бд - `createdb dbname`  
Удалить бд - `dropdb dbname`

##  Работа с пользователями
Создание пользователя -  `sudo -u postgres createuser --createdb username`

Параметр `--createdb` добавляет пользователю право создавать базы данных

Изменение пароля пользователя - `ALTER USER user_name WITH PASSWORD 'new_password';`

Сделать пользователя суперпользователем - `ALTER USER librarian WITH SUPERUSER;`

Изменить владельца у базы данных - `ALTER DATABASE name OWNER TO new_owner;`

## Работа с Postgres
`\q` - выход

`\du` - (described users) список ролей

`\l` - Список юзеров и баз

`\?` - вывод всех команде в редакторе vim

### Полезные опции `psql`

`-U` - Запуск под определённым пользователем

`-f` - Запуск скрипта из файла

`-W` - Вход с паролем


Пример запуска скрипта из файла под определённым юзером:

`psql -U db_name -f ~/dumps/dump.sql  -W`