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
    title VARCHAR(200),
    preview_text TEXT,
    image_id INT,
    article_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tag VARCHAR(20),
    detail_text TEXT,
    href VARCHAR(200),
    comments INT,
    category VARCHAR(100),
    kind VARCHAR(10),
    video_id INT
  );

CREATE TABLE
  IF NOT EXISTS image (
    image_id INT PRIMARY KEY AUTO_INCREMENT,
    image_path VARCHAR(100),
    image_size INT,
    image_name VARCHAR(100),
    image_alt VARCHAR(200)
  );

ALTER TABLE article
MODIFY COLUMN image_id INT,
ADD FOREIGN KEY (image_id) REFERENCES image (image_id);

CREATE TABLE video (
	video_id int auto_increment NOT NULL,
	video_path VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	`size` int NOT NULL,
	name VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	alt VARCHAR(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	CONSTRAINT `PRIMARY` PRIMARY KEY (video_id)
);

ALTER TABLE article
MODIFY COLUMN video_id INT,
ADD FOREIGN KEY (video_id) REFERENCES video (video_id);