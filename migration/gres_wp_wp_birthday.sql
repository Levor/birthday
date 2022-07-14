DROP TABLE IF EXISTS `wp_birthday`;
CREATE TABLE `wp_birthday` (
  `id` int NOT NULL,
  `fio` varchar(100) DEFAULT NULL,
  `day` int DEFAULT NULL,
  `month` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
