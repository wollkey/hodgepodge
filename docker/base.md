## Общие сведения
**Docker** — прога для работы с контейнерами

**Image** — это образ контейнеров. На основе образа можно создать сколько угодно контейнеров

**Контейнер** — мини виртуальная машина, которая изолирована от внешнего окружения и взаимодействует с ней через порты

**Dockerfile** - файл без расширения. В нём описывается как создать docker image (образ)

**Volume** - папка хоста, примонтированная к файловой системе контейнера

### Создание образа
```shell
docker build -t {tag_name} {path}
```
При сборке используется Dockerfile, поэтому необходимо его создать (без расширения)

Докер как слоёный пирог, каждая команда в докер файле это слой. Пример Dockerfile:
```dockerfile
FROM php:8

RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/

COPY . /usr/src/app/
RUN composer install

EXPOSE 8080

ENV TZ Europe/Moscow

CMD ["php", "index.php"]
```

### Работа с образами
```docker images``` - вывести все образы

```docker rmi {image_name}``` - удалить образ


### Работа с контейнерами

#### Просмотр
```docker ps``` - посмотреть работающие контейнеры

```docker ps -a``` - посмотреть все контейнеры, включая завершённые

```docker ps -q``` - вывести только id контейнеров

#### Запуск
```docker run {image_tag_name}``` - запустить контейнер на базе образа

```docker run -d``` - запустить контейнер в фоне

```docker run -v {local_absolute_path/volume_name}:{container_path}``` - прикручивает папку к контейнеру

```docker run -e TZ=Europe/Moscow``` - указание переменной окружения. Название переменной = значение

```docker run -p 8080:8080``` - пробрасывает порт в контейнер. Первый порт - порт машины, второй порт - порт внутри контейнера

```docker run --rm``` - после того, как приложение в контейнере остановится - контейнер автоматически удалится

```docker run --name {container_name} {image_name}``` - запустить контейнер под определённым именем

#### Остановка/удаление
```docker stop {container_name}``` - остановить контейнер

```docker rm {name/id}``` - удалить контейнер

```docker rm $(docker ps -aq)``` - удалить все контейнеры

#### Монтирование папок
Папки монтируются в контейнере с помощью docker volume

```docker volume ls``` - просмотр всех волюмов

```docker volume create {name}``` - создание вольюма

#### Docker compose
Когда настроек для запуска контейнера становится слишком много, то лучше использовать docker compose.
Это надстройка на докером. Необходимо создать файл docker-compose.yaml, и в нём описать настройки всех контейнеров.

Пример docker-compose.yaml:
```
version 1

volumes:
    postgres_volume:

services:
    web:
        build: path/
        restart: alwais
        environments:
            - TZ=Europe/Moscow
            ...
            
    web2:
        build: path2/
        restart: always
        ports:
            - 8080:8080
        environments:
            - TZ=Europe/Moscow
    
    postgres:
        image: postgres:latest
        volumes:
            - postgres_volume: /data/db
        restart: always
```

```docker-compose up -d``` - запустить докер композ. -d параметр указывает на работу всех контейнеров в фоне

```docker-compose down``` - остановить все контейнеры