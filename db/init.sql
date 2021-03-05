CREATE TABLE dict (
    id          int NOT NULL AUTO_INCREMENT,
    name        varchar(45) NOT NULL,
    language    varchar(10) NOT NULL,
    enabled     bool NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE user (
    id      int NOT NULL AUTO_INCREMENT,
    name    varchar(45) NOT NULL,
    email   varchar(100) DEFAULT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE word (
    id          int NOT NULL AUTO_INCREMENT,
    word        varchar(32) NOT NULL,
    bitmap      int DEFAULT NULL,
    unique_char int DEFAULT NULL,
    frequency   float DEFAULT '0',
    definition  varchar(2048) DEFAULT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE game (
    id          int NOT NULL AUTO_INCREMENT,
    letters     varchar(32) NOT NULL,
    game_key    varchar(46) NOT NULL,
    score       int NOT NULL DEFAULT '0',
    max_score   int NOT NULL DEFAULT '0',
    created_by  int NOT NULL,
    created_at  datetime NOT NULL DEFAULT NOW(),
    updated_at  datetime NOT NULL DEFAULT NOW() ON UPDATE NOW(),

    PRIMARY KEY (id),
    KEY fk_gm_game_user_idx (created_by),

    CONSTRAINT fk_gm_game_user
        FOREIGN KEY (created_by)
        REFERENCES user (id)
);

CREATE TABLE game_word (
    id          int NOT NULL AUTO_INCREMENT,
    game_id     int DEFAULT NULL,
    word_id     int DEFAULT NULL,
    found_by    int DEFAULT NULL,

    PRIMARY KEY (id),
    KEY fk_gw_word_idx (word_id),
    KEY fk_gw_game_idx (game_id),
    KEY fk_gw_user_idx (found_by),

    CONSTRAINT fk_gw_game
        FOREIGN KEY (game_id)
        REFERENCES game (id),

    CONSTRAINT fk_gw_user
        FOREIGN KEY (found_by)
        REFERENCES user (id),

    CONSTRAINT fk_gw_word
        FOREIGN KEY (word_id)
        REFERENCES word (id)
);

CREATE TABLE dict_word (
    id      int NOT NULL AUTO_INCREMENT,
    word_id int DEFAULT NULL,
    dict_id int DEFAULT NULL,

    PRIMARY KEY (id),
    KEY fk_dw_word_idx (word_id),
    KEY fk_dw_dict_idx (dict_id),

    CONSTRAINT fk_dw_dict
        FOREIGN KEY (dict_id)
        REFERENCES dict (id),

    CONSTRAINT fk_dw_word
        FOREIGN KEY (word_id)
        REFERENCES word (id)
);

CREATE TABLE dict_game (
    id      int NOT NULL AUTO_INCREMENT,
    game_id int DEFAULT NULL,
    dict_id int DEFAULT NULL,

    PRIMARY KEY (id),
    KEY fk_dg_word_idx (game_id),
    KEY fk_dg_dict_idx (dict_id),

    CONSTRAINT fk_dg_dict
        FOREIGN KEY (dict_id)
        REFERENCES dict (id),

    CONSTRAINT fk_dg_game
        FOREIGN KEY (game_id)
        REFERENCES game (id)
);
