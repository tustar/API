use ushare;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(11) NOT NULL DEFAULT '',
  `code` varchar(11) DEFAULT '',
  `nick` varchar(12) DEFAULT NULL,
  `weight` int(11) DEFAULT NULL,
  `token` varchar(45) DEFAULT NULL,
  `shared` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile_UNIQUE` (`mobile`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `weight_UNIQUE` (`weight`),
  UNIQUE KEY `token_UNIQUE` (`token`),
  UNIQUE KEY `nick_UNIQUE` (`nick`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


