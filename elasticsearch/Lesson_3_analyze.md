# Анализ текста

Позволяет анализировать текстовые данные. Для анализа текста используется analyze API:


```http
POST /_analyze
{
 "text": "Мистер Жопосранчик",
 "analyzer": "standard"
}
```

Процесс анализа:

1. Текст анализируется с помощью фильтров: вырезаются/добавляются символы и т.п.
2. Токенизация - разделение текста на слова, или так называемые токены
3. Каждый токен анализируется: к нему применяются фильтры, токены переводятся в нижний регистр

В результате получаем ответ в виде разбитых токенов:
```json
{
  "tokens" : [
    {
      "token" : "мистер",
      "start_offset" : 0,
      "end_offset" : 6,
      "type" : "<ALPHANUM>",
      "position" : 0
    },
    {
      "token" : "жопосранчик",
      "start_offset" : 7,
      "end_offset" : 18,
      "type" : "<ALPHANUM>",
      "position" : 1
    }
  ]
}
```

`token` - единица текста

`start_offset` и `end_offset` - смещение токена в тексте

`type` - тип токена

`position` - позиция токена в массиве

## Кастомные фильтры

Если передать параметр `analyzer: standard` то применятся стандартные фильтры:

```http
POST /_analyze
{
 "text": "Мистер Жопосранчик",
 "char_filter": [],
 "tokenizer": "standard",
 "filter": ["lowercase"]
}
```

Результат тот же, как и в предыдущем запросе

>Так же существует анализатор `keyword` - сохранят текст как есть, создаёт только один токен.
>Его можно использовать, например, для email адресов, где не нужно анализировать текст и разбивать его на слова.

Существуют разные анализаторы для разных случаев: `standard`, `keyword`, etc. Смотреть  [здесь](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-analyzers.html)


## Кастомный анализатор

Ниже приведёт пример с анализатором по корню слов

```http
PUT /{index_name}
{
  "settings": {
    "analysis": {
      "filter": {
        "russian_stop": {
          "type":       "stop",
          "stopwords":  "_russian_" 
        },
        "russian_keywords": {
          "type":       "keyword_marker",
          "keywords":   ["пример"] 
        },
        "russian_stemmer": {
          "type":       "stemmer",
          "language":   "russian"
        }
      },
      "analyzer": {
        "{custom_analyzer_name}": {
          "tokenizer":  "standard",
          "filter": [
            "lowercase",
            "russian_stop",
            "russian_keywords",
            "russian_stemmer"
          ]
        }
      }
    }
  }
}
```
