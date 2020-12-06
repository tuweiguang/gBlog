CREATE DATABASE IF NOT EXISTS gBlog CHARACTER SET utf8;
use gBlog;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `status` int DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT '用户基本表';

CREATE TABLE `article` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT '0' COMMENT '用户ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `category_id` int NOT NULL COMMENT '分类ID',
  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT 'Tag',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '摘要',
  `desc` longtext NOT NULL COMMENT '详情',
  `html` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'html',
  `pv` int DEFAULT '0' COMMENT 'px',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  `status` int DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `review` int DEFAULT '0' COMMENT '评论',
  `recommend` int NOT NULL DEFAULT '0' COMMENT '是否顶置，0否；1是，默认否',
  `like` int NOT NULL DEFAULT '0' COMMENT '点赞数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
