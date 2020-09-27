INSERT INTO logos (name, url)
VALUES ('alfa', 'http://...'),
        ('megafon', 'http://...'),
        ('food', 'http://...'),
        ('tinkoff', 'http://...'),
        ('ashan', 'http://...'),
        ('metro', 'http://...');

INSERT INTO clients (login, password, full_name, passport, birthdate, status)
VALUES ('iivanov','123','ivan ivanov','aaa111', '2000.12.12', 'ACTIVE'),
       ('ppetrov','456','petr petrov','bbb222', '1990.11.11', 'ACTIVE'),
       ('vvasilev','789','vasya vasilev','ccc333', '1980.10.10', 'ACTIVE');

INSERT INTO descriptions (description)
VALUES ('пополнение через альфа-банк'), ('продукты'), ('рестораны'),
        ('пополнение телефона'), ('перевод'), ('билеты в кино'), ('путешествия');

INSERT INTO categories (category)
VALUES ('Пополнения'), ('Супермаркеты'), ('Мобильная связь'), ('Госуслуги'),
        ('Товары для спорта'), ('Фастфуд'), ('Переводы');

INSERT INTO cards (number, balance, issuer,  owner_id, status)
VALUES ('1111', 4556993643, 'VISA', 1, 'ACTIVE'),
       ('2222', 343255, 'VISA', 2, 'ACTIVE'),
       ('3333', 89867657, 'VISA', 2, 'ACTIVE'),
       ('4444', 3242432, 'VISA', 3, 'ACTIVE');

INSERT INTO transactions (card_id, amount, category_id, description_id, logo_id)
VALUES (1, 1000, 2, 2, 3),
       (2, 5000000, 1, 1, 1),
       (2, -100000, 2, 2, 3),
       (2, -100000, 3, 4, 2),
       (2, -100000, 7, 5, 4);