## Локальный запуск
Скопировать проект
```bash
  git clone https://github.com/mant1COREX/pet-project
```

Заполнить файлы:
- .env
- /configs/config.yml

Перейти в директорию с проектом
```bash
  cd pet-project
```

Собрать и запустить проект
```bash
  docker-compose up test-task
```

Выполнить миграции БД(при необходимости изменить данные)
```bash
  migrate -path ./migrations -database postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable up
```




