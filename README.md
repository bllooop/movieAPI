

Документация Swagger после запуска будет доступна по ссылке: http://localhost:8000/swagger/index.html#/

# API для выполнения операций со списком фильмов и актеров
Данный сервис реализован на языке Go с использованием библиотеки HTTP. Для работы с PostgreSQL использовался драйвер lib/pq. Для использования сервиса неободимо авторизоваться и ввести полученный JWT токен в поле заголовка. Тип операции зависит от определенного запроса, также реализована ролевая система, благодаря которой некоторые методы недоступны для обычного пользователя. Для данных из файла конфигурации используется Viper.
Данные принимаются в JSON формате, далее сохраняются в базу данных PostgreSQL. Вывод данных происходит тоже в формате JSON. Реализованы все основные операции сохранения, изменения, получения и удаления.
Операции выполняются вместе со списком фильмов и актеров. При выводе списка актеров происходит перечисление фильмов для каждого актера в списке.
Реализована спецификация на  API в формате Swagger 2.0 с подходом code-first.
В Makefile прописаны возможные варианты запуска API, тестов и миграции.
## Есть 2 способа запуска микросервиса:
### 1. Локально.
   Необходимо в файле конфигурации configs/config.yaml указать
   ```
   host = "localhost"
   ```
   Запустить контейнер postgres
   ```
   docker run --name=db -e POSTGRES_PASSWORD='54321' -p 5432:5432 -d postgres
   ```
   При наличии утилиты golang-migrate после первого запуска контейнеров прописать команду
   ```
   make migrate
   ```
   Она создает необходимые для работы сервиса таблицы в контейнере базы данных. Данную команду нужно прописать всего 1 раз, и далее при запуске тех же контейнеров таблицы сохранятся.

   Ввести в консоль команду
   ```
   go run cmd/main.go
   ```
   При необходимости, можно заменить названия контейнера базы данных, пароль и порты. Соответствующие параметры для такого же изменения находятся в configs/config.yaml.
### 2. Локально при использовании docker-compose.
   
   Необходимо в файле конфигурации configs/config.yaml указать (По умолчанию в проекте стоит такое значение)
   ```
   host = "db"
   ```
   Ввести в консоль команду
   ```
   make build && make run
   ```
   При наличии утилиты golang-migrate после первого запуска контейнеров прописать команду
   ```
   make migrate
   ```
Информация по утилите golang-migrate находится в репозитории https://github.com/golang-migrate/migrate
## Пользование сервисом
### 1. Авторизация и регистрация
#### Для регистрации необходимо выполнить запрос
```
curl --location  --request POST 'http://localhost:8000/api/auth/sign-up' \
--header 'Content-Type: application/json' \
--data '{
    "username": "{username}",
    "password": "{password}",
    "role": "{0}"
}'
```
Вместо username вводится желаемый username, в поле password соответственно желаемый пароль. Поле role при регистрации указано для удобного тестирования ролевой системы, в практике оно не должно быть при регистрации. Число 0 в этом поле означает, что это обычный пользователь с правами только получения данных, а число 1 означает, что у этого пользователя есть дополнительные добавления, изменения и удаления.
#### Для авторизации необходимо выполнить запрос
```
curl --location  --request POST 'http://localhost:8000/api/auth/sign-in' \
--header 'Content-Type: application/json' \
--data '{
    "username": "{username}",
    "password": "{password}"
}'
```
Вместо username вводится выбранный нами при регистрации username, в поле password соответственно пароль.
В ответ на данный запрос нам выдастся токен, который нужно сохранить и использовать во всех следующих запросах. В программе Postman имеется функционал, который позволяет один раз указать токен и выполнять все дальнейшие запросы уже с ним. В командной строке с каждым запросом придется указывать вручную заголовок.
Проверка токена в сервисе выполняется при помощи методов в Middleware.

### 2. Список фильмов
#### Для получения списка фильмов необходимо ввести запрос
```
curl --location --request GET 'http://localhost:8000/api/movies?order={order}' \
--header 'Authorization: Bearer {Token}' \
--data ''
```
Вместо order вводится желаемая сортировка списка. Если поле оставить пустым, по умолчанию будет применена сортировка по рейтингу. Вместо Token вводится личный токен, полученный при авторизации. После успешного запроса будет выведен список сохранненых в базу данных фильмов.

#### Для добавления фильма, при наличии прав администратора, необходимо ввести запрос
```
curl --location  --request POST 'http://localhost:8000/api/movies/add' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token} \
--data '{
    "title": "title",
    "rating": 0,
    "date": "2000-01-01",
    "description": "description",
     "actorname": [
      "Actor One",  "Actor Two"
     ]
}'
```
В полях для добавления нужно ввести название, рейтинг, дату выхода, описание и список актеров в виде массива в указанном порядке.

```
curl --location --request POST  'http://localhost:8000/api/movies/update?id={id}' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token} \
--data '{
    "title": "{title}",
    "rating": {0},
    "date": "{2000-01-01}",
    "description": "{description}"
}'
```

```
curl --location --request POST 'http://localhost:8000/api/movies/delete?id={id}' \
--header 'Authorization: Bearer {token} \
--data ''
```

```
curl --location --request POST 'http://localhost:8000/api/movie?name={name}' \
--header 'Authorization: Bearer {token} \
--data ''
```

```
curl --location --request GET 'http://localhost:8000/api/actors' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token} \
--data ''
```
```
curl --location --request POST 'http://localhost:8000/api/actors/add' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token} \
--data '{
    "name": "name",
    "gender": 0,
    "date": "2000-01-01"
}'
```

```
curl --location  --request POST 'http://localhost:8000/api/actors/update?id={id}' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token} \
--data '{
    "name": "name",
    "gender": 0,
    "date": "2000-01-01"
}'
```
```
curl --location --request POST 'http://localhost:8000/api/actors/delete?id=id' \
--header 'Authorization: Bearer {token} \
--data ''
```
