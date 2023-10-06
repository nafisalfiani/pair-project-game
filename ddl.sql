CREATE TABLE Game (
    gameID INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    description VARCHAR(255)
    date_publish DATE,
    rating INT
);


INSERT INTO Game (name, description, date_publish, rating) VALUES
('Dota', 'Deskripsi Dota', '2023-10-01', 4),
('FC Mobile 24', 'Deskripsi FC Mobile 24', '2023-09-15', 5),
('Stumble Guys', 'Deskripsi Stumble Guys', '2023-08-20', 3);
