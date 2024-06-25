**Тестовое задание GO (MediaSoft)**

Исполнитель: Мусаев Тариэль

# Описание проекта
Данный проект представляет собой REST API сервис для работы с сущностью "Автомобиль". Сервер хранит данные локально в файле при помощи СУБД SQLite. Используется версия Go 1.22.2.
### Структура БД
```
CREATE TABLE "autos" (
	"id"	TEXT,
	"brand"	TEXT NOT NULL DEFAULT '',
	"model"	TEXT NOT NULL DEFAULT '',
	"mileage"	REAL NOT NULL DEFAULT 0,
	"number_of_owners"	INTEGER NOT NULL DEFAULT 0,
	PRIMARY KEY("id")
)
```
# Инструкции по запуску сервера
Для запуска сервера необходимо в папке проекта через терминал запустить команду ```go run main```. В базе есть тестовые данные.
