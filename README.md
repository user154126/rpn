# Module-1

Итоговая задача модуля 1 Яндекс Лицея.

Этот проект реализует веб-сервис, принимающий выражение через HTTP-запрос и возвращающий результат вычислений.

## Инструкция по запуску

Для запуска проекта выполните следующие шаги:

1. Склонируйте репозиторий:

   ```bash
   git clone https://github.com/user154126/rpn.git
   cd calc_go
   ```

2. Убедитесь, что Go установлен и находится в `$PATH` (проверить версию можно командой `go version`).

3. Запустите API-сервер в терминале VS Code:

   ```bash
   go run ./cmd/main.go
   ```

   Сервер запустится на порту `8080` по умолчанию.

## Примеры использования

### 1. Успешный запрос (в командной строке Windows):

   **Статус-код:** `200 OK`

   ```bash
   C:\Users\alexd>curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2+2*2\"}"
   ```

   **Пример ответа:**

   ```json
   {"result":"6"}
   ```

### 2. Ошибка обработки выражения:

   **Статус-код:** `422 Unprocessable Entity`

   ```bash
   curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2*(2+2{)\"}"
   ```

   **Пример ответа:**

   ```json
   {"error": "Error calculation"}
   ```

### 3. Неподдерживаемый метод:

   **Статус-код:** `405 Method Not Allowed`

   ```bash
   curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2*(2+2{)\"}"
   ```

   **Пример ответа:**

   ```json
   {"error":"Error calculation"}
   ```

### 4. Некорректное тело запроса:

   **Статус-код:** `400 Bad Request`

   **Пример ответа:**

   ```json
   {"error": "Invalid operator"}
   ```

