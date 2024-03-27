-- создаём таблицу clients
CREATE TABLE clients (
    "id" INTEGER PRIMARY KEY,
    "fio" VARCHAR(250) NOT NULL DEFAULT '',
    "email" VARCHAR(64) NOT NULL DEFAULT ''
);

-- создаём таблицу products
CREATE TABLE products (
    "id" INTEGER PRIMARY KEY,
    "product" VARCHAR(250) NOT NULL DEFAULT '',
    "price" INTEGER NOT NULL DEFAULT 0
);

-- создаём таблицу sales
CREATE TABLE sales (
    "id" INTEGER PRIMARY KEY,
    "client" INTEGER NOT NULL DEFAULT 0,
    "product" INTEGER NOT NULL DEFAULT 0,
    "date" VARCHAR(10) NOT NULL DEFAULT ''
);

-- заполням таблицу clients
INSERT INTO clients (id, fio, email) VALUES (1, 'Иванов А.Г.', 'ivanov@mail.ru');
INSERT INTO clients (id, fio, email) VALUES (2, 'Василий Иванович', 'vasily@yandex.ru');
INSERT INTO clients (id, fio, email) VALUES (3, 'Лыкова Анна', 'likova@mail.ru');

-- заполням таблицу products
INSERT INTO products (id, product, price) VALUES (1, 'Планировщик', 500);
INSERT INTO products (id, product, price) VALUES (2, 'Мои заметки', 450);
INSERT INTO products (id, product, price) VALUES (3, 'Мои финансы', 720);

-- заполняет таблицу sales
INSERT INTO sales (id, client, product, date) VALUES (1,2,1,'31.10.23');
INSERT INTO sales (id, client, product, date) VALUES (2,3,1,'31.10.23');
INSERT INTO sales (id, client, product, date) VALUES (3,1,2,'01.11.23');

-- увеличиваем цену на 50 рублей для всех продуктов
UPDATE products SET price = price + 50;

-- устанавливаем цену и меняем название у программы Мои заметки с id: 2
UPDATE products SET product = 'Суперзаметки', price = 600 WHERE id = 2;

-- очищаем таблицу sales
-- TRUNCATE TABLE sales; -- не поддерживается SQLite

-- удаляем таблицу sales из базы данных
DROP TABLE sales;

-- удаляем все записи о продажах у клиента с идентификатором 7
DELETE FROM sales WHERE client = 7;

-- получаем всё содержимое таблицы clients
SELECT * FROM clients;

-- получаем строку с id: 3 из таблицы products
SELECT * FROM products WHERE id = 3;

-- получаем id и имена первых десяти клиентов, отсортированных по алфавиту
SELECT id, fio FROM clients ORDER BY fio DESC LIMIT 10;
