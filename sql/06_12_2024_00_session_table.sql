CREATE TABLE sessions
(
    id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    token       VARCHAR(255) NOT NULL,
    account_id  INT          NOT NULL,
    ips         JSON,
    fingerprint VARCHAR(255),
    expired_at  DATETIME     NOT NULL,
    logined_at  DATETIME     NOT NULL,
    updated_at  DATETIME     NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;