CREATE DATABASE `neuronews` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */

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



CREATE TABLE `image` (
  `image_id` int NOT NULL AUTO_INCREMENT,
  `image_path` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `image_size` int DEFAULT NULL,
  `image_name` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `image_alt` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`image_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `video` (
  `video_id` int NOT NULL AUTO_INCREMENT,
  `video_path` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `video_size` int NOT NULL,
  `video_name` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `video_alt` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `article` (
  `article_id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `preview_text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `image_id` int DEFAULT NULL,
  `article_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `tag` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `detail_text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `href` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `comments` int NOT NULL,
  `category` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `kind` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `video_id` int DEFAULT NULL,
  PRIMARY KEY (`article_id`),
  KEY `image_id` (`image_id`),
  KEY `video_id` (`video_id`),
  CONSTRAINT `article_ibfk_1` FOREIGN KEY (`image_id`) REFERENCES `image` (`image_id`) ON DELETE CASCADE,
  CONSTRAINT `article_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`video_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
