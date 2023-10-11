-- 5 fitur:
-- 1. Register
-- 2. Login
-- 3. Lihat keranjang
-- 4. Tambah game
-- 5. Hapus game
-- 6. Discount
-- 7,8,9 untuk admin
-- 7. User Reports
-- 8. Order Reports
-- 9. Stock Reports


CREATE TABLE IF NOT EXISTS discounts (
    id INT AUTO_INCREMENT NOT NULL,
    voucher VARCHAR(100) NOT NULL,
    nominee FLOAT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT NOT NULL,
    email VARCHAR(100) NOT NULL,
    birth DATE NOT NULL,
    password VARCHAR(100) NOT NULL,
    admin BOOLEAN DEFAULT 0 NOT NULL,
    voucher_id INT,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS maturity (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR (100) NOT NULL,
    minimum_age INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS games (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NOT NULL,
    maturity_id INT NOT NULL,
    price FLOAT UNSIGNED,
    stock INT UNSIGNED,
    PRIMARY KEY (id),
    FOREIGN KEY (maturity_id) REFERENCES maturity(id)
);


CREATE TABLE IF NOT EXISTS users_games (
    id INT AUTO_INCREMENT NOT NULL,
    user_id INT NOT NULL,
    game_id INT NOT NULL,
    discount_id INT,
    PRIMARY KEY (id)
);