<h1>Описание</h1>
<p>Calculation service - это кластер для распределения вычислений и их выполнения. Он состоит из следующих сервисов:</p>
<ul>
<li>frontend - графический интерфейс</li>
<li>api - сервис, который обрабатывает клиентские запросы</li>
<li>orchestrator - сервис, который распределяет задачи и обновляет информацию об агентах</li>
<li>agent (несколько) - сервис, выполняющий задачи</li>
<li>postgres - сервис с базой данных</li>
<li>pgadmin - сервис с графическим интерфейсом для базы данных</li>
</ul>
<p>Для связи между сервисами agent, api, orchestrator используется grpc. Для добавления задач необходима авторизация (jwt). Для каждого пользователя api предоставляет только его задачи и таймауты.</p>

<h1>Запуск</h1>
из корня проекта

```
make build
```

```
docker-compose up
```

Api программы по умолчанию доступно на порту 8080. Графический интерфейс по умолчанию доступен на портy 3000

<details>
<summary>Если нет make</summary>

```
cd backend
```

```
docker build -t eonias189/calculation-service/orchestrator -f Dockerfile.orchestrator .
```

```
docker build -t eonias189/calculation-service/agent -f Dockerfile.agent .
```

```
docker build -t eonias189/calculation-service/api -f Dockerfile.api .
```

```
cd ..
```

```
cd frontend
```

```
docker build -t eonias189/calculation-service/frontend .
```

```
cd ..
```

</details>

<h1>Примеры использования</h1>
<h3>Через http запросы</h3>
<details>
<summary>Регистрация</summary>
Запрос

```
curl --location 'http://127.0.0.1:8080/api/auth/register' \
--header 'Content-Type: application/json' \
--data '{
    "login": "login",
    "password": "secret"
}'
```

Ответ

```
{}
```

</details>

<details>
<summary>Вход</summary>
Запрос

```
curl --location 'http://127.0.0.1:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "login": "login",
    "password": "secret"
}'
```

Ответ

```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyMzUzOTEsImlhdCI6MTcxMzY0MzM5MSwibG9naW4iOiJsb2dpbiIsIm5iZiI6MTcxMzY0MzM5MSwidXNlcl9pZCI6NX0.r6xQZsTDYz9BuDEhdKMeV9Q6waW7cD8dl-aDMPKWH5k"}
```

</details>

<details>
<summary>Мониторинг агентов</summary>
Запрос

```
curl --location 'http://127.0.0.1:8080/api/agents?limit=2&offset=1'
```

Ответ

```
{"agents":[{"id":9,"ping":0,"active":true,"max_threads":10,"running_threads":0},{"id":8,"ping":0,"active":false,"max_threads":10,"running_threads":0}]}
```

</details>

<details>
<summary>Установка таймаутов</summary>
В теле запроса можно указать только те операции, таймауты на которые нужно изменить
Запрос

```
curl --location --request PATCH 'http://127.0.0.1:8080/api/timeouts' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyMzUzOTEsImlhdCI6MTcxMzY0MzM5MSwibG9naW4iOiJsb2dpbiIsIm5iZiI6MTcxMzY0MzM5MSwidXNlcl9pZCI6NX0.r6xQZsTDYz9BuDEhdKMeV9Q6waW7cD8dl-aDMPKWH5k' \
--header 'Content-Type: application/json' \
--data '{
    "add": 8,
    "mul": 13
}'
```

Ответ

```
{"timeouts":{"add":8,"sub":0,"mul":13,"div":0}}
```

</details>

<details>
<summary>Добавление задачи</summary>
Запрос

```
curl --location 'http://127.0.0.1:8080/api/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyMzUzOTEsImlhdCI6MTcxMzY0MzM5MSwibG9naW4iOiJsb2dpbiIsIm5iZiI6MTcxMzY0MzM5MSwidXNlcl9pZCI6NX0.r6xQZsTDYz9BuDEhdKMeV9Q6waW7cD8dl-aDMPKWH5k' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "22 + 22 * 22"
}'
```

Ответ

```
{"task":{"id":53,"expression":"22 + 22 * 22","result":0,"status":"pending","createTime":"2024-04-20T20:29:21Z"}}
```

</details>

<details>
<summary>Получение результата задачи по id</summary>
Запрос

```
curl --location 'http://127.0.0.1:8080/api/tasks/53' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyMzUzOTEsImlhdCI6MTcxMzY0MzM5MSwibG9naW4iOiJsb2dpbiIsIm5iZiI6MTcxMzY0MzM5MSwidXNlcl9pZCI6NX0.r6xQZsTDYz9BuDEhdKMeV9Q6waW7cD8dl-aDMPKWH5k'
```

Ответ

```
{"task":{"id":53,"expression":"22 + 22 * 22","result":506,"status":"success","createTime":"2024-04-20T20:29:21Z"}}
```

</details>

<details>
<summary>Получение всех задач</summary>
Запрос

```
curl --location 'http://127.0.0.1:8080/api/tasks?limit=2&offset=1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYyMzUzOTEsImlhdCI6MTcxMzY0MzM5MSwibG9naW4iOiJsb2dpbiIsIm5iZiI6MTcxMzY0MzM5MSwidXNlcl9pZCI6NX0.r6xQZsTDYz9BuDEhdKMeV9Q6waW7cD8dl-aDMPKWH5k'
```

Ответ

```
{"tasks":[{"id":52,"expression":"22 + 22 * 22","result":506,"status":"success","createTime":"2024-04-20T20:26:20Z"},{"id":51,"expression":"22 + 22 * 22","result":506,"status":"success","createTime":"2024-04-20T20:13:43Z"}]}
```

</details>

<h2>Тестирование отказоустойчивости</h3>
<p>Для тестирования отказоустойчивости можно вручную (с помощью docker cli) остановить какой-нибудь сервис. Система корректно отреагирует на это.</p>
