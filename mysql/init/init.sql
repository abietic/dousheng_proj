CREATE DATABASE IF NOT EXISTS `dousheng`;
USE `dousheng`;
CREATE TABLE IF NOT EXISTS `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_name` (`user_name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';
CREATE TABLE IF NOT EXISTS `user_stats`
(
    `id`         bigint unsigned NOT NULL COMMENT 'User ID',
    `work_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User uploaded video count',
    `follow_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT  'User followed user count',
    `follower_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User\'s follower count',
    `favorite_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User favorite count',
    `total_favorited`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User uploaded videos\' favorite count',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User stats table';
CREATE TABLE IF NOT EXISTS `user_followed` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `user_id` bigint(20) UNSIGNED NOT NULL COMMENT 'Follower user ID',
  `followed_user_id` bigint(20) UNSIGNED NOT NULL COMMENT 'Followed user ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User follow relation create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User follow relation update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User follow relation delete time',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) COMMENT 'Follower user index',
  KEY `idx_followed_user_id` (`followed_user_id`) COMMENT 'Followed user index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User followed table';
CREATE TABLE IF NOT EXISTS `video_stats`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'Video ID',
    `author_id`  bigint unsigned NOT NULL COMMENT 'Video author user id',
    `comment_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video comment count',
    `favorite_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video favorited count',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Video create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Video stats update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY `idx_author_id` (`author_id`) COMMENT 'Video author index',
    KEY `idx_created_at` (`created_at`) COMMENT 'Video publish time index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video stats table';
CREATE TABLE IF NOT EXISTS `favorite_video` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `user_id` bigint(20) UNSIGNED NOT NULL COMMENT 'User ID',
  `favorite_video_id` bigint(20) UNSIGNED NOT NULL COMMENT 'User\'s favorite video ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User favorite video create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User favorite video update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User favorite video delete time',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) COMMENT 'User index',
  KEY `idx_created_at` (`created_at`) COMMENT 'Video favorited time index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User favorite videos table';