-- 5 fitur:
-- 1. Register
-- 2. Login
-- 3. Lihat keranjang
-- 4. Tambah game
-- 5. Hapus game
-- 6. Discount
-- 7,8,9, 10 untuk admin
-- 7. Update stock
-- 8. User Reports
-- 9. Order Reports
-- 10. Stock Reports


CREATE TABLE IF NOT EXISTS discounts (
    id INT AUTO_INCREMENT NOT NULL,
    voucher VARCHAR(100) NOT NULL,
    nominee FLOAT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    birth DATE NOT NULL,
    password VARCHAR(100) NOT NULL,
    admin BOOLEAN DEFAULT 0 NOT NULL,
    discount_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (discount_id) REFERENCES discounts(id)
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
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (game_id) REFERENCES games(id)
);


INSERT INTO discounts (voucher, nominee)
VALUES 
    ("GAMERS", 0.10),
    ("GAMERSINDO", 0.08),
    ("PERANGKODE", 0.15);

INSERT INTO maturity (name, minimum_age)
VALUES
    ("Everyone", 0),
    ("Teen", 13),
    ("Mature", 17),
    ("Adults", 21);

INSERT INTO games (name, description, maturity_id, price, stock)
VALUES 
    ("Counter Strike 2", "For over two decades, Counter-Strike has offered an elite competitive experience, one shaped by millions of players from across the globe. And now the next chapter in the CS story is about to begin. This is Counter-Strike 2.", 3, 0, 100),
    ("Grand Theft Auto V", "Grand Theft Auto V for PC offers players the option to explore the award-winning world of Los Santos and Blaine County in resolutions of up to 4k and beyond, as well as the chance to experience the game running at 60 frames per second.", 4, 30.00, 100),
    ("Stumble Guys", "Race through obstacle courses against up to 32 players online. Run, jump and dash to the finish line until the best player takes the crown!", 1, 15.00, 100),
    ("Forza Horizon 5", "Your Ultimate Horizon Adventure awaits! Explore the vibrant open world landscapes of Mexico with limitless, fun driving action in the worlds greatest cars.", 2, 50.00, 100);

INSERT INTO users_games (user_id, game_id)
VALUES 
    (1, 1),
    (1, 2),
    (2, 3),
    (2, 1),
    (3, 4);
    (3, 1);