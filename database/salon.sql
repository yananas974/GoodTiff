DROP TABLE IF EXISTS coiffeur;
CREATE TABLE salon (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  tel     INT(255) NOT NULL,
  adresse      VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO salon
  (id, name, tel, adresse)
VALUES
  (1, 'John Coltrane', 0101010101, "35 &venue DU"),
  (2, 'John Coltrane', 0202020202, "35 avenue DU p"),
  (3, 'Gerry Mulligan', 0303030303, "35bis avenue du PU"),
  (4, 'Sarah Vaughan', 'Sarah Vaughan', 0505050505, " 35ter avenue du pi");