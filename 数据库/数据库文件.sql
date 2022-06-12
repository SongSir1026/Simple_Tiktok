/*
 Navicat Premium Data Transfer

 Source Server         : SongSir
 Source Server Type    : MySQL
 Source Server Version : 50650
 Source Host           : 39.105.196.187:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50650
 File Encoding         : 65001

 Date: 12/06/2022 15:37:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` bigint(20) NOT NULL COMMENT '�û�id',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '�û�����',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '�û�����',
  `follow_count` int(11) NOT NULL COMMENT '��ע����',
  `follower_count` int(11) NOT NULL COMMENT '��˿����',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (0, '123456', '123456', 0, 0);

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `video_id` bigint(20) NOT NULL COMMENT '��Ƶid',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '����',
  `author_id` int(11) NOT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '��Ƶ��ַ',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '������ַ',
  `favorite_count` int(11) NOT NULL DEFAULT 0 COMMENT '������',
  `comment_count` int(11) NOT NULL DEFAULT 0 COMMENT '������',
  PRIMARY KEY (`video_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '��Ƶ��' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1655018445, '可爱猫猫，小淘气', 0, 'https://simplesiktok-songsir.oss-cn-beijing.aliyuncs.com/uploads123456/share_5ba4ce8ead3722aec0d31e8bd2e4e3a2.mp4', 'https://simplesiktok-songsir.oss-cn-beijing.aliyuncs.com/uploads123456/share_5ba4ce8ead3722aec0d31e8bd2e4e3a2.mp4?x-oss-process=video/snapshot,t_500,f_jpg,w_600,h_800,m_fast', 1, 2);
INSERT INTO `video` VALUES (1655018772, '哇哦两猫相对', 0, 'https://simplesiktok-songsir.oss-cn-beijing.aliyuncs.com/uploads123456/share_d182de5e9057843cdcdc14eb90d54b4e.mp4', 'https://simplesiktok-songsir.oss-cn-beijing.aliyuncs.com/uploads123456/share_d182de5e9057843cdcdc14eb90d54b4e.mp4?x-oss-process=video/snapshot,t_500,f_jpg,w_600,h_800,m_fast', 0, 1);

-- ----------------------------
-- Table structure for video_comment
-- ----------------------------
DROP TABLE IF EXISTS `video_comment`;
CREATE TABLE `video_comment`  (
  `id` bigint(20) NOT NULL,
  `video_id` bigint(20) NOT NULL,
  `comment_text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `user_id` bigint(20) NOT NULL,
  `create_date` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of video_comment
-- ----------------------------
INSERT INTO `video_comment` VALUES (1655018725, 1655018445, '小猫猫好可爱', 0, '06-12');
INSERT INTO `video_comment` VALUES (1655018811, 1655018772, '小猫猫好可爱哇哦', 0, '06-12');

-- ----------------------------
-- Table structure for video_follow
-- ----------------------------
DROP TABLE IF EXISTS `video_follow`;
CREATE TABLE `video_follow`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '�û�id',
  `video_id` bigint(20) NULL DEFAULT NULL COMMENT '��Ƶid',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK_����ע����Ƶ`(`video_id`) USING BTREE,
  INDEX `FK_��Ƶ��ע���û�`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '��Ƶ��ע��' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of video_follow
-- ----------------------------
INSERT INTO `video_follow` VALUES (1655018708, 0, 1655018445);

SET FOREIGN_KEY_CHECKS = 1;
