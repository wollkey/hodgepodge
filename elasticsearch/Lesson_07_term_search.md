# Точный поиск

Поиск уровня токенов или терминов (term level).

Ранее мы разобрали анализатор, который разбивает текст на токены. 
При поиске с помощью term анализатор не задействуется,
поэтому ищутся точные совпадения по этим токенам.

Данный способ не подходит для полнотекстового поиска, так как не задействуется анализатор, но
отлично подходит для поиска по всем нетекстовым полям, или по keyword. 

### Точное совпадение

```http
GET /products/_search
{
  "query": {
    "terms": {
      "age": [40, 500]
    }
  }
}
```

### Поиск по id
```http
GET /products/_search
{
  "query": {
    "ids": {
      "values": [101, 102, 103]
    }
  }
}
```

### Поиск по диапазону
```http
GET /products/_search
{
  "query": {
    "range": {
      "age": {
        "gte": 10,
        "lte": 600
      }
    }
  }
}
```

Так же можно искать по диапазону дат. В дате можно использовать пайп оператор:
```http

GET /products/_search
{
  "query": {
    "range": {
      "created_at": {
        "gte": "2022/01/26 || -1y"
      }
    }
  }
}

```

### Проверка на существование поля
```http
GET /products/_search
{
  "query": {
    "exists": {
      "field": "age"
    }
  }
}
```

### Префикс
Поиск по начальной части слова, то есть по префиксу. Слово - "Природа", префикс для примера "Прир"

```http
GET /products/_search
{
  "query": {
    "prefix": {
      "name.keyword": "Прир"
    }
  }
}
```

### Пропущенные символы
`?` - один символ

`*` - один и больше символов


```http
GET /products/_search
{
  "query": {
    "wildcard": {
      "name.keyword": "Прир*а"
    }
  }
}
```

### Регулярные выражения 

```http
GET /products/_search
{
  "query": {
    "regexp": {
      "name.keyword": "Прир[а-яА-Я]+а"
    }
  }
}
```