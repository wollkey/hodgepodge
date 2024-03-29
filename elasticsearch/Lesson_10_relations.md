# Связанные данные

Elasticsearch - документ ориентированная система, которая не поддерживает реляционную модель данных.
Поэтому связи в нём образуются с помощью других механизмов.

### Вложенные объекты (nested)

Как говорилось ранее, для связей внутри объектов необходимо использовать тип `nested`.
Иначе данные будут отображаться в виде несвязных массивов.

Для того чтобы запрашивать данные типа nested, необходимо использовать nested query:

```http
GET /recipes/_search
{
  "query":{
    "nested": {
      "path": "ingredients",
      "query" {}
    }
  }
}
```

Результаты вложенного объекта:
```http
GET /recipes/_search
{
  "_source": false,
  "query":{
    "nested": {
      "path": "ingredients",
      "inner_hits": {},
      "query" {}
    }
  }
}
```

### Создание связей

Связь между родителем и потомком создаётся с помощью специального поля `join`:

```http
PUT /users
{
  "mappings": {
    "properties": {
      "join_field": {
        "type": "join",
        "relations": {
          "{index}": "{field}"
        }
      }
    }
  }
}
```
