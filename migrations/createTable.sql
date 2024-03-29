CREATE TABLE IF NOT EXISTS Salons (
    ID INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Tel INT,
    Adresse VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Coiffeurs (
    ID INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Lastname INT,
    SalonID INT,
    FOREIGN KEY (SalonID) REFERENCES Salons(ID)
);

CREATE TABLE IF NOT EXISTS User (
    ID INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Lastname VARCHAR(255) NOT NULL,
    Tel VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,  
    Password VARCHAR(255) NOT NULL
);



