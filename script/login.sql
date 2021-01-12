DROP DATABASE IF EXISTS `login_server`;
CREATE DATABASE `login_server` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_info`;
CREATE TABLE `account_info`(
    `account_id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '账号ID',
    `account_name` VARCHAR(128) NOT NULL COMMENT '账号名字',
    `sdk_type`  INT NOT NULL COMMENT '渠道ID',
    `password` VARBINARY(64) NOT NULL COMMENT '密码',
    `salt`  VARBINARY(128) NOT NULL COMMENT '盐',
    `create_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`account_id`),
    UNIQUE KEY `uk_name` (`account_name`, `sdk_type`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;