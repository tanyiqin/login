DROP DATABASE IF EXISTS `login_server`;
CREATE DATABASE `login_server` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_info`;
CREATE TABLE `account_info`(
    `account_id` BIGINT AUTO_INCREMENT COMMENT "",
    `account_name` VARCHAR(20) NOT NULL,
    `password` VARCHAR(20) NOT NULL,
    PRIMARY KEY (`account_id`),
    UNIQUE KEY `uk_name` (`account_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;