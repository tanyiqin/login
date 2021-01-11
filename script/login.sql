DROP DATABASE IF EXISTS `login_server`;
CREATE DATABASE `login_server` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_info`;
CREATE TABLE `account_info`(
    `account_id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '账号ID',
    `account_name` VARCHAR(20) NOT NULL COMMENT '账号名字',
    `password` VARCHAR(20) NOT NULL COMMENT '密码',
    `create_time` DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`account_id`),
    UNIQUE KEY `uk_name` (`account_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;