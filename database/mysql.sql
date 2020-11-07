use gblog;

DROP TABLE IF EXISTS `user`;
'CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT ''ID'',
  `name` varchar(255) DEFAULT NULL COMMENT ''用户名'',
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `status` int DEFAULT ''1'' COMMENT ''状态'',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8';


