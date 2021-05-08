/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : localhost:3306
 Source Schema         : comic

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 08/05/2021 18:23:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comic_category
-- ----------------------------
DROP TABLE IF EXISTS `comic_category`;
CREATE TABLE `comic_category`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13628 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comic_category
-- ----------------------------
INSERT INTO `comic_category` VALUES (4, '冒险', 'https://images.dmzj1.com/tuijian/222_222/180720maoxian.jpg');
INSERT INTO `comic_category` VALUES (6, '格斗', 'https://images.dmzj1.com/tuijian/222_222/170811gedou.jpg');
INSERT INTO `comic_category` VALUES (7, '科幻', 'https://images.dmzj1.com/tuijian/222_222/170817kehuan.jpg');
INSERT INTO `comic_category` VALUES (8, '爱情', 'https://images.dmzj1.com/tuijian/222_222/170811shenqi.jpg');
INSERT INTO `comic_category` VALUES (9, '侦探', 'https://images.dmzj1.com/tuijian/222_222/170817zhentan.jpg');
INSERT INTO `comic_category` VALUES (10, '竞技', 'https://images.dmzj1.com/tuijian/222_222/170811jingji.jpg');
INSERT INTO `comic_category` VALUES (11, '魔法', 'https://images.dmzj1.com/tuijian/222_222/170817mofa.jpg');
INSERT INTO `comic_category` VALUES (12, '神鬼', 'https://images.dmzj1.com/tuijian/222_222/170817shengui.jpg');
INSERT INTO `comic_category` VALUES (13, '校园', 'https://images.dmzj1.com/tuijian/222_222/170811xiaoyuan.jpg');
INSERT INTO `comic_category` VALUES (17, '四格', 'https://images.dmzj1.com/tuijian/222_222/170817sige.jpg');
INSERT INTO `comic_category` VALUES (2304, '日本', 'https://images.dmzj1.com/tuijian/222_222/180720riben.jpg');
INSERT INTO `comic_category` VALUES (2305, '韩国', 'https://images.dmzj1.com/tuijian/222_222/180720hanguo.jpg');
INSERT INTO `comic_category` VALUES (2306, '欧美', 'https://images.dmzj1.com/tuijian/222_222/180720oumei.jpg');
INSERT INTO `comic_category` VALUES (2308, '国漫', 'https://images.dmzj1.com/tuijian/222_222/180720guoman.jpg');
INSERT INTO `comic_category` VALUES (2309, '连载', 'https://images.dmzj1.com/tuijian/222_222/180720lianzai.jpg');
INSERT INTO `comic_category` VALUES (2310, '完结', 'https://images.dmzj1.com/tuijian/222_222/180720wanjie.jpg');
INSERT INTO `comic_category` VALUES (3244, '秀吉', 'https://images.dmzj1.com/tuijian/222_222/180803weiniang.jpg');
INSERT INTO `comic_category` VALUES (3245, '悬疑', 'https://images.dmzj1.com/tuijian/222_222/kongbu.jpg');
INSERT INTO `comic_category` VALUES (3248, '热血', 'https://images.dmzj1.com/tuijian/222_222/170817rexue.jpg');
INSERT INTO `comic_category` VALUES (3252, '萌系', 'https://images.dmzj1.com/tuijian/222_222/170817mengxi.jpg');
INSERT INTO `comic_category` VALUES (3254, '治愈', 'https://images.dmzj1.com/tuijian/222_222/170817zhiyu.jpg');
INSERT INTO `comic_category` VALUES (3255, '励志', 'https://images.dmzj1.com/tuijian/222_222/170817lizhi.jpg');
INSERT INTO `comic_category` VALUES (3262, '少年漫', 'https://images.dmzj1.com/tuijian/222_222/180720shaonianman.jpg');
INSERT INTO `comic_category` VALUES (3263, '少女漫', 'https://images.dmzj1.com/tuijian/222_222/180720shaonvman.jpg');
INSERT INTO `comic_category` VALUES (3264, '青年漫', 'https://images.dmzj1.com/tuijian/222_222/180720qingnianman.jpg');
INSERT INTO `comic_category` VALUES (3327, '美食', 'https://images.dmzj1.com/tuijian/222_222/170811meishi.jpg');
INSERT INTO `comic_category` VALUES (3328, '职场', 'https://images.dmzj1.com/tuijian/222_222/170811zhichang.jpg');
INSERT INTO `comic_category` VALUES (4518, 'TS', 'https://images.dmzj1.com/tuijian/222_222/170817xingzhuan.jpg');
INSERT INTO `comic_category` VALUES (5077, '东方', 'https://images.dmzj1.com/tuijian/222_222/170817dongfang.jpg');
INSERT INTO `comic_category` VALUES (5806, '魔幻', 'https://images.dmzj1.com/tuijian/222_222/170817mohuan.jpg');
INSERT INTO `comic_category` VALUES (5848, '奇幻', 'https://images.dmzj1.com/tuijian/222_222/170811qihuan.jpg');
INSERT INTO `comic_category` VALUES (6316, '轻小说改', 'https://images.dmzj1.com/tuijian/222_222/170811qinggai.jpg');
INSERT INTO `comic_category` VALUES (6437, '颜艺', 'https://images.dmzj1.com/tuijian/222_222/170811yanyi.jpg');
INSERT INTO `comic_category` VALUES (7568, '搞笑', 'https://images.dmzj1.com/tuijian/222_222/170811gaoxiao.jpg');
INSERT INTO `comic_category` VALUES (7900, '仙侠', 'https://images.dmzj1.com/tuijian/222_222/170811xianxia.jpg');
INSERT INTO `comic_category` VALUES (13626, '女青漫', 'https://images.dmzj1.com/tuijian/222_222/180720nvqingman.jpg');
INSERT INTO `comic_category` VALUES (13627, '舰娘', 'https://images.dmzj1.com/tuijian/222_222/170817jianniang.jpg');

SET FOREIGN_KEY_CHECKS = 1;
