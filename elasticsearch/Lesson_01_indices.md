# Индексы

    Индекс - единое логическое хранилище в elasticsearch

Индексы могут разбиваться на шарды, что позволяет легко масштабировать базу.

> Если сравнивать с sql, то индекс можно сравнить с самой базой данных или схемой

### Создание и удаление индекса

#### Создать индекс

```http
PUT /{index_name}
```

Данный запрос создаст индекс с настройками по умолчанию.
Чтобы задать собственную конфигурацию индекса, необходимо передать так же тело запроса:

```http
PUT /{index_name}
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  }
}
```

#### Удалить индекс:
```http
DELETE /{index_name}
```

### Переиндексация

Тип поля изменить невозможно, но если это нужно сделать,
то придётся создать новый индекс с новым типом поля
и перенести все документы в новый индекс

Для упрощения этой операции существует `reindex API`:

```http
POST /_reindex
{
  "source": {
    "index": "..."
  },
  "dest": {
    "index": "..."
  }
}
```
