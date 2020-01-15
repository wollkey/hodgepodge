> Здесь находятся полезные команды, которые не попали в другие списки.
>
## Перемещение файла
```git
git mv <file> <to>
```

Делает следующее:
```git
1. mv file <to>
2. git add <to>/<filename>
3. git rm <filename>
```

Проще сделать `git mv ...`

## Вывод файла
```git
git cat-file <some hash>
```

Подставить соответствующий хеш для вывода. Типы хешей:

`<hash>` - Блоб

`<treehash>` - Дерево

`<libhash>` - Коммиты

Опции: 

`-t` - тип объекта

`-p` - контент объекта


## Алиасы
Алиасы добавляются в файле `.gitconfig`. Пример:

```git
[alias]
  ch = checkout
  st = status
  br = branch
  hist = log --pretty=format:\"%h %ad | %s%d [%an]\" --graph --date=short
  ...
```

Теперь вместо того, чтобы писать длинную команду:
```git
log --pretty=format:\"%h %ad | %s%d [%an]\" --graph --date=short
```

Можно писать так:
```git
git hist
```
