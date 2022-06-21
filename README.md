# Тестовое задание

В качастве БД использовался СУБД Postgres. БД состоит из одной таблицы.
Скрипт для ее создания.

CREATE TABLE tablename (
colname SERIAL
);
CREATE TABLE "transaction" (
"id" SERIAL PRIMARY KEY ,
"id_user" int,
"email" varchar,
"sum" float8,
"currency" int,
"dt_create" varchar,
"dt_change" varchar,
"status" int
);

Было реализовано 4 запроса api, которые были протестированы через postman.

**1. localhost:8000/api/**

Создание платежа (на вход принимает id пользователя, email пользователя, сумму и валюту платежа).
![image](https://user-images.githubusercontent.com/74566888/174698939-a54f02b7-9655-4605-b8c7-f0b5b856e638.png)

**2. localhost:8000/api/change**

Изменение статуса платежа платежной системой.
![image](https://user-images.githubusercontent.com/74566888/174699789-a62f736f-0cae-4dfa-b6cd-d114d4e6bcd7.png)

**3. localhost:8000/api/byuserid**

Получение списка всех платежей пользователя по его ID.
![image](https://user-images.githubusercontent.com/74566888/174699539-f90acd3a-7978-44bd-83cb-6c0687ef8ffa.png)

**4. localhost:8000/api/byemail**

Получение списка всех платежей пользователя по его e-mail.
![image](https://user-images.githubusercontent.com/74566888/174699625-760935c5-3910-4613-b66f-74f12abaad72.png)

