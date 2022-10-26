CREATE DATABASE IF NOT EXISTS `goexpert`;

CREATE TABLE IF NOT EXISTS `goexpert`.`products` (
    `id` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255),
    `price` DECIMAL(10,2),
    PRIMARY KEY (`id`)
);