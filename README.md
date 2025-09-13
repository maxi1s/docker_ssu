Минимальный проект
Приложение на Go делает:
- `GET /healthz` → `ok`
- `GET /` → текст "Hello from Docker + Compose!"

## Запуск

```bash
cp .env.example .env
docker compose up --build
# открой http://localhost:8080/healthz  → ok
# открой http://localhost:8080/         → Hello from Docker + Compose!
```

- `docker-compose.yml` маппит порт хоста `${APP_PORT}` на порт контейнера `8080`.
- Внутри контейнера приложение слушает `HTTP_ADDR=:8080`. Это значит: слушать на *всех интерфейсах*.
- Файл `.env` позволяет менять порт снаружи (например, `APP_PORT=9090`).
