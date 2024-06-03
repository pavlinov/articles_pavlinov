# Название исполняемого файла
BINARY_NAME=articles_pavlinov

# Путь к базе данных
DB_PATH=database/test.db
PORT=8084

# Команда по умолчанию
all: build

# Установка зависимостей
deps:
	go mod tidy

# Сборка проекта
build: deps
	go build -o $(BINARY_NAME) main.go

# Запуск проекта
run: build
	PORT=$(PORT) DB_PATH=$(DB_PATH) ./$(BINARY_NAME)

# Удаление скомпилированных файлов
clean:
	go clean
	rm -f $(BINARY_NAME)

# Инициализация базы данных
init-db:
	rm -f $(DB_PATH)
	PORT=$(PORT) DB_PATH=$(DB_PATH) go run main.go & sleep 5 && kill $$!

# Пересобрать и запустить проект
rebuild: clean build run

.PHONY: all deps build run clean init-db rebuild
