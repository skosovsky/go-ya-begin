-- запрос на содержимое таблицы product
SELECT *
  FROM products
 ORDER BY id;

-- запрос клиента с идентификатором 4
SELECT *
  FROM clients
 WHERE id = 4;

-- запрос программ, которые купил клиент выше
SELECT *
  FROM sales
 WHERE client = 4
 ORDER BY date;

-- запрос всех строк и столбцов таблицы product
SELECT *
  FROM products;

-- получить все строки таблицы products со значениями столбцов product и price
SELECT product,
       price
  FROM products;

-- используем алиасы для столбцов
SELECT product AS 'Программа',
       price   AS 'Цена'
  FROM products;

-- получить информацию о клиенте, где идентификатор id равен 152
SELECT *
  FROM clients
 WHERE id = 152;

-- получить поля fio и birthday у всех клиентов 1986 года рождения
SELECT fio, birthday
  FROM clients
 WHERE birthday > '1986'
   AND birthday < '1987';

-- другой вариант запроса — с использование BETWEEN ... AND
SELECT fio, birthday
  FROM clients
 WHERE birthday BETWEEN '1986' AND '1987';

-- можно указать полные строки для сравнения дат
SELECT fio, birthday
  FROM clients
 WHERE birthday >= '19860101'
   AND birthday < '19870101';

-- получить клиентов с id равным 33, 77, 55 и 200
SELECT *
  FROM clients
 WHERE id = 33
    OR id = 77
    OR id = 55
    OR id = 200;

-- тоже самое, но с оператором IN
SELECT *
  FROM clients
 WHERE id IN (33, 77, 55, 200);

-- найти клиента с днем рождения 14.02.1973
SELECT fio
  FROM clients
 WHERE birthday = '19730214';

-- найти Афанасия Николаевича
SELECT *
  FROM clients
 WHERE fio LIKE '%Афанасий Николаевич%';

-- так как отчеством заканчивается имя клиента, то можно отправить запрос без символа %
SELECT *
  FROM clients
 WHERE fio LIKE '%Афанасий Николаевич';

-- посмотреть покупки id 124
SELECT *
  FROM sales
 WHERE client = 124;

-- найти всех клиентов, у которых email содержит число 1996
SELECT fio
  FROM clients
 WHERE email LIKE '%1996%';

-- отсортировать по алфавиту всех клиентов с id меньше 5
SELECT *
  FROM clients
 WHERE id < 5
 ORDER BY fio;

-- получить всех Антонов и отсортировать по дате рождения от самого молодого к самому старому
SELECT *
  FROM clients
 WHERE fio LIKE '% Антон %'
 ORDER BY birthday DESC;

-- отсортировать клиентов по алфавиту и получить только пять первых
SELECT id,
       fio,
       login,
       birthday
  FROM clients
 ORDER BY fio
 LIMIT 5;

-- чтобы расчитать количество страниц, необходимо получить количество строк в таблице
SELECT COUNT(*)
  FROM clients
 WHERE birthday < '19800101';

-- посчитать общее количество клиентов
SELECT COUNT(*)
  FROM clients;

-- чтобы пропустить первые двадцать записей, указываем после OFFSET значение смещения
SELECT id,
       fio,
       login,
       birthday
  FROM clients
 ORDER BY fio
 LIMIT 20 OFFSET 20;

-- более короткая запись для OFFSET (40 - это OFFSET, 20 - LIMIT)
SELECT id,
       fio,
       login,
       birthday
  FROM clients
 ORDER BY fio
 LIMIT 40,20;

-- найти сколько покупок было сделано в июне 2023 года
SELECT COUNT(*)
  FROM sales
 WHERE date LIKE '%202306%';

-- клиент с id 171 приобрёл 2 лицензии на программу «Заметки» (id: 3) и нужно добавить запись о покупке в таблицу sales
INSERT INTO sales (client, product, volume, date)
VALUES (171, 3, 2, '20240327');

-- проверим добавились ли данные
SELECT *
  FROM sales
 WHERE client = 171
 ORDER BY date DESC;

-- вставим несколько записей и проверим
INSERT INTO sales (client, product, volume, date)
VALUES (171, 3, 2, '20231108'),
       (37, 1, 1, '20231108'),
       (98, 2, 1, '20231108');

SELECT *
  FROM sales
 WHERE date = '20231108';

-- обновим данные, нужно увеличить на 20% все цены на программное обеспечение
UPDATE products
   SET price = 1.2 * price;

-- исправить дату покупок и проверить
UPDATE sales
   SET date = '20231104'
 WHERE date = '20231108';

SELECT *
  FROM sales
 WHERE date = '20231104';

-- узать id клиента и изменить ему почту (дополнительная проверка на id нужна, т.к. fio не уникальны)
SELECT id
  FROM clients
 WHERE fio = 'Александров Георгий Романович';

UPDATE clients
   SET login = 'agr1972',
       email = 'agr18061972@yandex.ru'
 WHERE id = 116;

SELECT fio,
       login,
       email
  FROM clients
 WHERE id = 116;

-- удалить продажи с датой 4 ноября 2023 года
DELETE
  FROM sales
 WHERE date = '20231104';

SELECT *
  FROM sales
 WHERE date = '20231104';

--  удалить тестового клиента с id 48, не только из таблицы clients, но и его покупки из таблицы sales
DELETE
  FROM clients
 WHERE id = 48;

DELETE
  FROM sales
 WHERE client = 48;

SELECT *
  FROM clients
 WHERE id = 48;

SELECT *
  FROM sales
 WHERE client = 48;

-- создать индекс для столбца client
CREATE INDEX sales_client ON sales (client);

-- создать индекс для столбца date
CREATE INDEX sales_date ON sales (date, product);

-- вывести созданные индексов по таблице sales
PRAGMA INDEX_LIST('sales');

-- удалить созданные индексы
DROP INDEX sales_client;
DROP INDEX sales_date;

-- используем транзакцию, удаляем все данные о пользователе с id 124 и отменяет через rollback
SELECT id,
       fio,
       login
  FROM clients
 WHERE id = 124;
BEGIN;
DELETE
  FROM clients
 WHERE id = 124;
DELETE
  FROM sales
 WHERE client = 124;
ROLLBACK;
SELECT id,
       fio,
       login
  FROM clients
 WHERE id = 124;

-- используем транзакцию и удалим данные об этом пользователе и применим это через commit
BEGIN;
DELETE
  FROM clients
 WHERE id = 124;
DELETE
  FROM sales
 WHERE client = 124;
COMMIT;
SELECT id, fio, login
  FROM clients
 WHERE id = 124;
SELECT *
  FROM sales
 WHERE client = 124;
