.PHONY: up down build logs clean restart

# Запуск всех сервисов
up:
	docker-compose up -d

# Остановка всех сервисов
down:
	docker-compose down

# Пересборка и запуск
build:
	docker-compose up --build -d

# Просмотр логов
logs:
	docker-compose logs -f

# Логи конкретного сервиса
logs-go:
	docker-compose logs -f go_server

logs-ml:
	docker-compose logs -f ml_service

logs-db:
	docker-compose logs -f postgres

# Очистка (удаление volumes)
clean:
	docker-compose down -v
	docker system prune -f

# Перезапуск
restart:
	docker-compose restart

# Статус сервисов
status:
	docker-compose ps