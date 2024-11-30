-- u633362372_lion_test.users definition

CREATE TABLE `users` (
  `username` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(10) NOT NULL,
  `insert_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `address` longtext NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u633362372_lion_test.movies definition

CREATE TABLE `movies` (
  `id` varchar(100) NOT NULL,
  `title` varchar(255) NOT NULL DEFAULT '',
  `description` longtext NOT NULL,
  `duration` bigint(20) NOT NULL DEFAULT 0,
  `artists` longtext NOT NULL,
  `genres` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `watch_url` varchar(255) NOT NULL,
  `viewed` bigint(20) NOT NULL DEFAULT 0,
  `voted` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u633362372_lion_test.genres definition

CREATE TABLE `genres` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `genre` varchar(100) NOT NULL,
  `viewed` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- u633362372_lion_test.votes definition

CREATE TABLE `votes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie_id` varchar(100) NOT NULL,
  `user_username` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `votes_FK` (`user_username`),
  KEY `votes_FK_1` (`movie_id`),
  CONSTRAINT `votes_FK` FOREIGN KEY (`user_username`) REFERENCES `users` (`username`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `votes_FK_1` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- inject user admin
INSERT INTO u633362372_lion_test.users
(username, name, password, `role`, insert_date, address)
VALUES('msbu_admin', 'MSBU Admin', '$2a$10$.8zHZcCMbxOJMkfwRN.TSOHtZclAUArXOTC3Pz2de3A6K77.Y.PNK', 'admin', '2024-11-30 22:58:57', 'Jakarta');
