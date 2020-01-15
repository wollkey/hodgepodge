## Запуск git сервера
```git
git daemon --verbose --export-all --base-path=.
```

Опция, которая разрешает отправку `push` в репозиторий Git Daemon:
```git
--enable=receive-pack
```