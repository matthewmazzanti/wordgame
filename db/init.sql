CREATE TABLE `dictionary` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `language` varchar(10) NOT NULL,
  `enabled` tinyint NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `word` (
  `id` int NOT NULL AUTO_INCREMENT,
  `word` varchar(32) NOT NULL,
  `bitmap` int DEFAULT NULL,
  `unique_char` int DEFAULT NULL,
  `frequency` float DEFAULT '0',
  `definition` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `game` (
  `id` int NOT NULL AUTO_INCREMENT,
  `letters` varchar(32) NOT NULL,
  `game_key` varchar(46) NOT NULL,
  `score` int NOT NULL DEFAULT '0',
  `maxScore` int NOT NULL DEFAULT '0',
  `created_by` int NOT NULL,
  `created_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_gm_game_user_idx` (`created_by`),
  CONSTRAINT `fk_gm_game_user` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`)
);

CREATE TABLE `game_word` (
  `id` int NOT NULL AUTO_INCREMENT,
  `game_id` int DEFAULT NULL,
  `word_id` int DEFAULT NULL,
  `found_by` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_gw_word_idx` (`word_id`),
  KEY `fk_gw_game_idx` (`game_id`),
  KEY `fk_gw_user_idx` (`found_by`),
  CONSTRAINT `fk_gw_game` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `fk_gw_user` FOREIGN KEY (`found_by`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_gw_word` FOREIGN KEY (`word_id`) REFERENCES `word` (`id`)
);

CREATE TABLE `dictionary_word` (
  `id` int NOT NULL AUTO_INCREMENT,
  `word_id` int DEFAULT NULL,
  `dictionary_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dw_word_idx` (`word_id`),
  KEY `fk_dw_dictionary_idx` (`dictionary_id`),
  CONSTRAINT `fk_dw_dictionary` FOREIGN KEY (`dictionary_id`) REFERENCES `dictionary` (`id`),
  CONSTRAINT `fk_dw_word` FOREIGN KEY (`word_id`) REFERENCES `word` (`id`)
);

CREATE TABLE `dictionary_game` (
  `id` int NOT NULL AUTO_INCREMENT,
  `game_id` int DEFAULT NULL,
  `dictionary_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dg_word_idx` (`game_id`),
  KEY `fk_dg_dictionary_idx` (`dictionary_id`),
  CONSTRAINT `fk_dg_dictionary` FOREIGN KEY (`dictionary_id`) REFERENCES `dictionary` (`id`),
  CONSTRAINT `fk_dg_game` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
);