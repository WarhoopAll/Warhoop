/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE = @@TIME_ZONE */;
/*!40103 SET TIME_ZONE = '+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;

CREATE DATABASE IF NOT EXISTS `sait` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `sait`;

DROP TABLE IF EXISTS `change_email`;
CREATE TABLE IF NOT EXISTS `change_email`
(
    `id`          bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `idacc`       bigint                                  NOT NULL,
    `old_email`   varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `new_email`   varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `token`       varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` datetime                                NOT NULL,
    `expire_time` datetime                                NOT NULL,
    `create_ip`   varchar(15) COLLATE utf8mb4_general_ci  NOT NULL,
    `state`       int                                     NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `change_password`;
CREATE TABLE IF NOT EXISTS `change_password`
(
    `id`          int                                     NOT NULL AUTO_INCREMENT,
    `idacc`       bigint                                  NOT NULL,
    `username`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `salt`        binary(32)                              NOT NULL,
    `verifier`    binary(32)                              NOT NULL,
    `token`       varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` datetime                                NOT NULL,
    `expire_time` datetime                                NOT NULL,
    `create_ip`   varchar(15) COLLATE utf8mb4_general_ci  NOT NULL,
    `state`       int                                     NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `change_username`;
CREATE TABLE IF NOT EXISTS `change_username`
(
    `id`           bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `idacc`        bigint                                  NOT NULL,
    `old_username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `new_username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `salt`         binary(32)                              NOT NULL,
    `verifier`     binary(32)                              NOT NULL,
    `token`        varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time`  datetime                                NOT NULL,
    `expire_time`  datetime                                NOT NULL,
    `create_ip`    varchar(15) COLLATE utf8mb4_general_ci  NOT NULL,
    `state`        int                                     NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `dislike_history`;
CREATE TABLE IF NOT EXISTS `dislike_history`
(
    `id`               int       NOT NULL AUTO_INCREMENT,
    `user_id`          int       NOT NULL,
    `disliked_user_id` int       NOT NULL,
    `object_type`      int       NOT NULL,
    `object_id`        int       NOT NULL,
    `created_at`       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `like_history`;
CREATE TABLE IF NOT EXISTS `like_history`
(
    `id`            int       NOT NULL AUTO_INCREMENT,
    `user_id`       int       NOT NULL,
    `liked_user_id` int       NOT NULL,
    `object_type`   int       NOT NULL,
    `object_id`     int       NOT NULL,
    `created_at`    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 28
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `news`;
CREATE TABLE IF NOT EXISTS `news`
(
    `id`             int unsigned                                          NOT NULL AUTO_INCREMENT,
    `title`          varchar(100)                                          NOT NULL,
    `text`           text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at`     datetime     DEFAULT NULL,
    `updated_at`     datetime     DEFAULT NULL,
    `author`         varchar(100)                                          NOT NULL,
    `image_url`      varchar(255) DEFAULT NULL,
    `like_count`     int          DEFAULT '0',
    `comments_count` int          DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb3;

DROP TABLE IF EXISTS `news_comments`;
CREATE TABLE IF NOT EXISTS `news_comments`
(
    `id`         int                                                   NOT NULL AUTO_INCREMENT,
    `news_id`    int unsigned                                          NOT NULL,
    `text`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` datetime                                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                                       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `author`     int unsigned                                          NOT NULL,
    `like_count` int                                                            DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 63
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8mb3_bin;

DROP TABLE IF EXISTS `password_reset_requests`;
CREATE TABLE IF NOT EXISTS `password_reset_requests`
(
    `id`          int                                     NOT NULL AUTO_INCREMENT,
    `idacc`       bigint                                  NOT NULL,
    `username`    varchar(32) COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '',
    `email`       varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `token`       varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` datetime                                NOT NULL,
    `expire_time` datetime                                         DEFAULT NULL,
    `create_ip`   varchar(15) COLLATE utf8mb4_general_ci  NOT NULL,
    `state`       int                                     NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `profile`;
CREATE TABLE IF NOT EXISTS `profile`
(
    `account_id` int                                                           NOT NULL,
    `avatar`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `name`       varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
    `rank`       varchar(255) COLLATE utf8mb4_general_ci DEFAULT 'Knight',
    PRIMARY KEY (`account_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `reports`;
CREATE TABLE IF NOT EXISTS `reports`
(
    `id`             int                                     NOT NULL AUTO_INCREMENT,
    `dispatchername` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `dispatcherid`   int                                     NOT NULL,
    `victimname`     varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `victimid`       int                                     NOT NULL,
    `reason`         text COLLATE utf8mb4_general_ci,
    `status`         tinyint(1) DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `sessions`;
CREATE TABLE IF NOT EXISTS `sessions`
(
    `id`          int unsigned                            NOT NULL AUTO_INCREMENT,
    `token`       varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `account_id`  int                                     NOT NULL,
    `ips`         json                                    DEFAULT NULL,
    `fingerprint` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
    `expired_at`  datetime                                NOT NULL,
    `logined_at`  datetime                                DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime                                NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 139
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `upvotes`;
CREATE TABLE IF NOT EXISTS `upvotes`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int             NOT NULL,
    `comment_id` int             NOT NULL,
    `upvoteby`   int             NOT NULL,
    PRIMARY KEY (`id`),
    KEY `comment_id` (`comment_id`),
    CONSTRAINT `upvotes_ibfk_1` FOREIGN KEY (`comment_id`) REFERENCES `news_comments` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `user_logs`;
CREATE TABLE IF NOT EXISTS `user_logs`
(
    `id`         int       NOT NULL AUTO_INCREMENT,
    `idacc`      int                                     DEFAULT NULL,
    `action`     varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
    `timestamp`  timestamp NOT NULL                      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `ip_address` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 105
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TRIGGER IF EXISTS `after_comment_delete`;
SET @OLDTMP_SQL_MODE = @@SQL_MODE, SQL_MODE =
        'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `after_comment_delete`
    AFTER DELETE
    ON `news_comments`
    FOR EACH ROW
BEGIN
    UPDATE news
    SET comments_count = comments_count - 1
    WHERE id = OLD.news_id;
END//
DELIMITER ;
SET SQL_MODE = @OLDTMP_SQL_MODE;

DROP TRIGGER IF EXISTS `after_comment_insert`;
SET @OLDTMP_SQL_MODE = @@SQL_MODE, SQL_MODE =
        'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `after_comment_insert`
    AFTER INSERT
    ON `news_comments`
    FOR EACH ROW
BEGIN
    UPDATE news
    SET comments_count = comments_count + 1
    WHERE id = NEW.news_id;
END//
DELIMITER ;
SET SQL_MODE = @OLDTMP_SQL_MODE;

/*!40103 SET TIME_ZONE = IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE = IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS = IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES = IFNULL(@OLD_SQL_NOTES, 1) */;
