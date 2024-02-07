CREATE USER 'neuronews'@'localhost' IDENTIFIED BY '486464';

GRANT Alter ON neuronews.* TO 'neuronews'@'localhost';

GRANT Create ON neuronews.* TO 'neuronews'@'localhost';

GRANT
Create view
  ON neuronews.* TO 'neuronews'@'localhost';

GRANT Delete ON neuronews.* TO 'neuronews'@'localhost';

GRANT Drop ON neuronews.* TO 'neuronews'@'localhost';

GRANT Grant option ON neuronews.* TO 'neuronews'@'localhost';

GRANT Index ON neuronews.* TO 'neuronews'@'localhost';

GRANT
Insert
  ON neuronews.* TO 'neuronews'@'localhost';

GRANT References ON neuronews.* TO 'neuronews'@'localhost';

GRANT
Select
  ON neuronews.* TO 'neuronews'@'localhost';

GRANT Show view ON neuronews.* TO 'neuronews'@'localhost';

GRANT Trigger ON neuronews.* TO 'neuronews'@'localhost';

GRANT
Update ON neuronews.* TO 'neuronews'@'localhost';

GRANT Alter routine ON neuronews.* TO 'neuronews'@'localhost';

GRANT Create routine ON neuronews.* TO 'neuronews'@'localhost';

GRANT Create temporary tables ON neuronews.* TO 'neuronews'@'localhost';

GRANT Execute ON neuronews.* TO 'neuronews'@'localhost';

GRANT Lock tables ON neuronews.* TO 'neuronews'@'localhost';

GRANT Grant option ON neuronews.* TO 'neuronews'@'localhost';

CREATE TABLE
  IF NOT EXISTS article (
    article_id INT PRIMARY KEY AUTO_INCREMENT,
    image_id INT FOREIGN KEY,
    title VARCHAR(100),
    preview_text TEXT,
    image_id INT,
    article_time TIMESTAMP,
    tag VARCHAR(20),
    detail_text TEXT,
    href VARCHAR(50),
    comments INT
  );

CREATE TABLE
  IF NOT EXISTS image (
    image_id INT PRIMARY KEY AUTO_INCREMENT,
    path VARCHAR(50),
    size INT,
    name VARCHAR(50),
    alt VARCHAR(50)
  );

ALTER TABLE article
MODIFY COLUMN image_id INT,
ADD FOREIGN KEY (image_id) REFERENCES image (image_id);