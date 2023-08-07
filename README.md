# HTTP-сервис для создания сокращенных ссылок

## Task

Необходимо реализовать HTTP-сервис для создания сокращенных ссылок
Сервис должен иметь следующее API:
- POST /n - создание новой сокращенной ссылки. Должен возвращать полную сокращенную ссылку, вместе со схемой и доменом(эти параметры должны определяться динамически, по запросу)
При создании ссылки необходимо проверять, создан ли уже short_url с данным назначением и в случае нахождения использовать его
- GET /<short_url_path> - должен отвечать перманентным редиректом с сохраненной ссылкой, при невалидности short_url_path должен отвечать ошибкой 404 Not Found
Перед запросом в СУБД необходимо удостовериться, что short_url_path выпущен сервисом. Проверка не обязана быть криптографически стойкой, но ее должно быть достаточно для
избежания лишних запросов к СУБД при брут форсе short_url_path

В качестве СУБД должен использоваться Postgres, сервис должен уметь сам развертывать необходимые ему сущности в СУБД
short_url_path должен быть регистронезависимым, состоять только из латинских букв и цифр нижнего регистра и быть настолько коротким, насколько это вообще возможно


## Run

В файле .env выберите параметры подключения к PostgreSQL, порт и наличие схемы в URL.
Запустите main.go.

## Test

В браузере откройте страницы:
http://localhost:8080/
http://localhost:8080/<short_url_path>

В Postman выполните запросы вида:
```shell
POST /n
body:
  link: String
returns:
  hash: String
```
и
```shell
GET /hash
params:
  hash: String
returns:
  link: String
```

Пример 
```shell
http://localhost:8080/api/link?hash=uTaGL
```
Ответ
```shell
{"link": "https://habr.com/"}
```