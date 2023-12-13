CREATE TABLE `items` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `price` decimal(9,2) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1' COMMENT '1 = active, 2 = inactive',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

INSERT INTO `vending_machine`.`items` (`name`, `price`, `status`, `created_at`)
VALUES 
('Aqua', 2000.00, 1, Now()),
('Sosro', 5000.00, 1, Now()),
('Cola', 7000.00, 1, Now()),
('Milo', 9000.00, 1, Now());
('Coffee', 10000.00, 1, Now());