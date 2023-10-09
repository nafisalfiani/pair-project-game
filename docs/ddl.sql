CREATE database game_library

USE game_library

CREATE TABLE game (
    game_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    published_date DATE,
    rating FLOAT
);

INSERT INTO game (name, description, published_date, rating) VALUES
('Dota 2', 'Multiplayer Online Battle Arena', '2023-10-01', 4.0),
('FC Mobile 24', 'Sport Simulation', '2023-09-15', 5.0),
('Stumble Guys', 'Social Game', '2023-08-20', 3.0),
('Cyberpunk 2077', 'Futuristic RPG', '2022-12-10', 3.5),
('The Witcher 3: Wild Hunt', 'Open-world Fantasy RPG', '2015-05-19', 5.0),
('Among Us', 'Social Deduction Game', '2020-08-18', 4.0),
('Valorant', 'Tactical First-Person Shooter', '2020-06-02', 4.5),
('Minecraft', 'Sandbox Adventure', '2011-11-18', 4.8),
('Apex Legends', 'Battle Royale Shooter', '2019-02-04', 4.3),
('Fortnite', 'Battle Royale Game', '2017-07-25', 4.2);
