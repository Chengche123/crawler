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

 Date: 08/05/2021 18:24:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comic_category_filter
-- ----------------------------
DROP TABLE IF EXISTS `comic_category_filter`;
CREATE TABLE `comic_category_filter`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `tag_id` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 184 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comic_category_filter
-- ----------------------------
INSERT INTO `comic_category_filter` VALUES (1, '题材', 0);
INSERT INTO `comic_category_filter` VALUES (2, '题材', 4);
INSERT INTO `comic_category_filter` VALUES (3, '题材', 5);
INSERT INTO `comic_category_filter` VALUES (4, '题材', 6);
INSERT INTO `comic_category_filter` VALUES (5, '题材', 7);
INSERT INTO `comic_category_filter` VALUES (6, '题材', 8);
INSERT INTO `comic_category_filter` VALUES (7, '题材', 9);
INSERT INTO `comic_category_filter` VALUES (8, '题材', 10);
INSERT INTO `comic_category_filter` VALUES (9, '题材', 11);
INSERT INTO `comic_category_filter` VALUES (10, '题材', 12);
INSERT INTO `comic_category_filter` VALUES (11, '题材', 13);
INSERT INTO `comic_category_filter` VALUES (12, '题材', 14);
INSERT INTO `comic_category_filter` VALUES (13, '题材', 16);
INSERT INTO `comic_category_filter` VALUES (14, '题材', 17);
INSERT INTO `comic_category_filter` VALUES (15, '题材', 3242);
INSERT INTO `comic_category_filter` VALUES (16, '题材', 3243);
INSERT INTO `comic_category_filter` VALUES (17, '题材', 3244);
INSERT INTO `comic_category_filter` VALUES (18, '题材', 3245);
INSERT INTO `comic_category_filter` VALUES (19, '题材', 3246);
INSERT INTO `comic_category_filter` VALUES (20, '题材', 3248);
INSERT INTO `comic_category_filter` VALUES (21, '题材', 3249);
INSERT INTO `comic_category_filter` VALUES (22, '题材', 3250);
INSERT INTO `comic_category_filter` VALUES (23, '题材', 3251);
INSERT INTO `comic_category_filter` VALUES (24, '题材', 3252);
INSERT INTO `comic_category_filter` VALUES (25, '题材', 3253);
INSERT INTO `comic_category_filter` VALUES (26, '题材', 3254);
INSERT INTO `comic_category_filter` VALUES (27, '题材', 3255);
INSERT INTO `comic_category_filter` VALUES (28, '题材', 3324);
INSERT INTO `comic_category_filter` VALUES (29, '题材', 3325);
INSERT INTO `comic_category_filter` VALUES (30, '题材', 3326);
INSERT INTO `comic_category_filter` VALUES (31, '题材', 3327);
INSERT INTO `comic_category_filter` VALUES (32, '题材', 3328);
INSERT INTO `comic_category_filter` VALUES (33, '题材', 3365);
INSERT INTO `comic_category_filter` VALUES (34, '题材', 4459);
INSERT INTO `comic_category_filter` VALUES (35, '题材', 4518);
INSERT INTO `comic_category_filter` VALUES (36, '题材', 5077);
INSERT INTO `comic_category_filter` VALUES (37, '题材', 5806);
INSERT INTO `comic_category_filter` VALUES (38, '题材', 5848);
INSERT INTO `comic_category_filter` VALUES (39, '题材', 6219);
INSERT INTO `comic_category_filter` VALUES (40, '题材', 6316);
INSERT INTO `comic_category_filter` VALUES (41, '题材', 6437);
INSERT INTO `comic_category_filter` VALUES (42, '题材', 7568);
INSERT INTO `comic_category_filter` VALUES (43, '题材', 7900);
INSERT INTO `comic_category_filter` VALUES (44, '题材', 13627);
INSERT INTO `comic_category_filter` VALUES (45, '题材', 17192);
INSERT INTO `comic_category_filter` VALUES (46, '题材', 18522);
INSERT INTO `comic_category_filter` VALUES (47, '读者群', 0);
INSERT INTO `comic_category_filter` VALUES (48, '读者群', 3262);
INSERT INTO `comic_category_filter` VALUES (49, '读者群', 3263);
INSERT INTO `comic_category_filter` VALUES (50, '读者群', 3264);
INSERT INTO `comic_category_filter` VALUES (51, '读者群', 13626);
INSERT INTO `comic_category_filter` VALUES (52, '进度', 0);
INSERT INTO `comic_category_filter` VALUES (53, '进度', 2309);
INSERT INTO `comic_category_filter` VALUES (54, '进度', 2310);
INSERT INTO `comic_category_filter` VALUES (55, '地域', 0);
INSERT INTO `comic_category_filter` VALUES (56, '地域', 2304);
INSERT INTO `comic_category_filter` VALUES (57, '地域', 2305);
INSERT INTO `comic_category_filter` VALUES (58, '地域', 2306);
INSERT INTO `comic_category_filter` VALUES (59, '地域', 2307);
INSERT INTO `comic_category_filter` VALUES (60, '地域', 2308);
INSERT INTO `comic_category_filter` VALUES (61, '地域', 8435);
INSERT INTO `comic_category_filter` VALUES (62, '题材', 0);
INSERT INTO `comic_category_filter` VALUES (63, '题材', 4);
INSERT INTO `comic_category_filter` VALUES (64, '题材', 5);
INSERT INTO `comic_category_filter` VALUES (65, '题材', 6);
INSERT INTO `comic_category_filter` VALUES (66, '题材', 7);
INSERT INTO `comic_category_filter` VALUES (67, '题材', 8);
INSERT INTO `comic_category_filter` VALUES (68, '题材', 9);
INSERT INTO `comic_category_filter` VALUES (69, '题材', 10);
INSERT INTO `comic_category_filter` VALUES (70, '题材', 11);
INSERT INTO `comic_category_filter` VALUES (71, '题材', 12);
INSERT INTO `comic_category_filter` VALUES (72, '题材', 13);
INSERT INTO `comic_category_filter` VALUES (73, '题材', 14);
INSERT INTO `comic_category_filter` VALUES (74, '题材', 16);
INSERT INTO `comic_category_filter` VALUES (75, '题材', 17);
INSERT INTO `comic_category_filter` VALUES (76, '题材', 3242);
INSERT INTO `comic_category_filter` VALUES (77, '题材', 3243);
INSERT INTO `comic_category_filter` VALUES (78, '题材', 3244);
INSERT INTO `comic_category_filter` VALUES (79, '题材', 3245);
INSERT INTO `comic_category_filter` VALUES (80, '题材', 3246);
INSERT INTO `comic_category_filter` VALUES (81, '题材', 3248);
INSERT INTO `comic_category_filter` VALUES (82, '题材', 3249);
INSERT INTO `comic_category_filter` VALUES (83, '题材', 3250);
INSERT INTO `comic_category_filter` VALUES (84, '题材', 3251);
INSERT INTO `comic_category_filter` VALUES (85, '题材', 3252);
INSERT INTO `comic_category_filter` VALUES (86, '题材', 3253);
INSERT INTO `comic_category_filter` VALUES (87, '题材', 3254);
INSERT INTO `comic_category_filter` VALUES (88, '题材', 3255);
INSERT INTO `comic_category_filter` VALUES (89, '题材', 3324);
INSERT INTO `comic_category_filter` VALUES (90, '题材', 3325);
INSERT INTO `comic_category_filter` VALUES (91, '题材', 3326);
INSERT INTO `comic_category_filter` VALUES (92, '题材', 3327);
INSERT INTO `comic_category_filter` VALUES (93, '题材', 3328);
INSERT INTO `comic_category_filter` VALUES (94, '题材', 3365);
INSERT INTO `comic_category_filter` VALUES (95, '题材', 4459);
INSERT INTO `comic_category_filter` VALUES (96, '题材', 4518);
INSERT INTO `comic_category_filter` VALUES (97, '题材', 5077);
INSERT INTO `comic_category_filter` VALUES (98, '题材', 5806);
INSERT INTO `comic_category_filter` VALUES (99, '题材', 5848);
INSERT INTO `comic_category_filter` VALUES (100, '题材', 6219);
INSERT INTO `comic_category_filter` VALUES (101, '题材', 6316);
INSERT INTO `comic_category_filter` VALUES (102, '题材', 6437);
INSERT INTO `comic_category_filter` VALUES (103, '题材', 7568);
INSERT INTO `comic_category_filter` VALUES (104, '题材', 7900);
INSERT INTO `comic_category_filter` VALUES (105, '题材', 13627);
INSERT INTO `comic_category_filter` VALUES (106, '题材', 17192);
INSERT INTO `comic_category_filter` VALUES (107, '题材', 18522);
INSERT INTO `comic_category_filter` VALUES (108, '读者群', 0);
INSERT INTO `comic_category_filter` VALUES (109, '读者群', 3262);
INSERT INTO `comic_category_filter` VALUES (110, '读者群', 3263);
INSERT INTO `comic_category_filter` VALUES (111, '读者群', 3264);
INSERT INTO `comic_category_filter` VALUES (112, '读者群', 13626);
INSERT INTO `comic_category_filter` VALUES (113, '进度', 0);
INSERT INTO `comic_category_filter` VALUES (114, '进度', 2309);
INSERT INTO `comic_category_filter` VALUES (115, '进度', 2310);
INSERT INTO `comic_category_filter` VALUES (116, '地域', 0);
INSERT INTO `comic_category_filter` VALUES (117, '地域', 2304);
INSERT INTO `comic_category_filter` VALUES (118, '地域', 2305);
INSERT INTO `comic_category_filter` VALUES (119, '地域', 2306);
INSERT INTO `comic_category_filter` VALUES (120, '地域', 2307);
INSERT INTO `comic_category_filter` VALUES (121, '地域', 2308);
INSERT INTO `comic_category_filter` VALUES (122, '地域', 8435);
INSERT INTO `comic_category_filter` VALUES (123, '题材', 0);
INSERT INTO `comic_category_filter` VALUES (124, '题材', 4);
INSERT INTO `comic_category_filter` VALUES (125, '题材', 5);
INSERT INTO `comic_category_filter` VALUES (126, '题材', 6);
INSERT INTO `comic_category_filter` VALUES (127, '题材', 7);
INSERT INTO `comic_category_filter` VALUES (128, '题材', 8);
INSERT INTO `comic_category_filter` VALUES (129, '题材', 9);
INSERT INTO `comic_category_filter` VALUES (130, '题材', 10);
INSERT INTO `comic_category_filter` VALUES (131, '题材', 11);
INSERT INTO `comic_category_filter` VALUES (132, '题材', 12);
INSERT INTO `comic_category_filter` VALUES (133, '题材', 13);
INSERT INTO `comic_category_filter` VALUES (134, '题材', 14);
INSERT INTO `comic_category_filter` VALUES (135, '题材', 16);
INSERT INTO `comic_category_filter` VALUES (136, '题材', 17);
INSERT INTO `comic_category_filter` VALUES (137, '题材', 3242);
INSERT INTO `comic_category_filter` VALUES (138, '题材', 3243);
INSERT INTO `comic_category_filter` VALUES (139, '题材', 3244);
INSERT INTO `comic_category_filter` VALUES (140, '题材', 3245);
INSERT INTO `comic_category_filter` VALUES (141, '题材', 3246);
INSERT INTO `comic_category_filter` VALUES (142, '题材', 3248);
INSERT INTO `comic_category_filter` VALUES (143, '题材', 3249);
INSERT INTO `comic_category_filter` VALUES (144, '题材', 3250);
INSERT INTO `comic_category_filter` VALUES (145, '题材', 3251);
INSERT INTO `comic_category_filter` VALUES (146, '题材', 3252);
INSERT INTO `comic_category_filter` VALUES (147, '题材', 3253);
INSERT INTO `comic_category_filter` VALUES (148, '题材', 3254);
INSERT INTO `comic_category_filter` VALUES (149, '题材', 3255);
INSERT INTO `comic_category_filter` VALUES (150, '题材', 3324);
INSERT INTO `comic_category_filter` VALUES (151, '题材', 3325);
INSERT INTO `comic_category_filter` VALUES (152, '题材', 3326);
INSERT INTO `comic_category_filter` VALUES (153, '题材', 3327);
INSERT INTO `comic_category_filter` VALUES (154, '题材', 3328);
INSERT INTO `comic_category_filter` VALUES (155, '题材', 3365);
INSERT INTO `comic_category_filter` VALUES (156, '题材', 4459);
INSERT INTO `comic_category_filter` VALUES (157, '题材', 4518);
INSERT INTO `comic_category_filter` VALUES (158, '题材', 5077);
INSERT INTO `comic_category_filter` VALUES (159, '题材', 5806);
INSERT INTO `comic_category_filter` VALUES (160, '题材', 5848);
INSERT INTO `comic_category_filter` VALUES (161, '题材', 6219);
INSERT INTO `comic_category_filter` VALUES (162, '题材', 6316);
INSERT INTO `comic_category_filter` VALUES (163, '题材', 6437);
INSERT INTO `comic_category_filter` VALUES (164, '题材', 7568);
INSERT INTO `comic_category_filter` VALUES (165, '题材', 7900);
INSERT INTO `comic_category_filter` VALUES (166, '题材', 13627);
INSERT INTO `comic_category_filter` VALUES (167, '题材', 17192);
INSERT INTO `comic_category_filter` VALUES (168, '题材', 18522);
INSERT INTO `comic_category_filter` VALUES (169, '读者群', 0);
INSERT INTO `comic_category_filter` VALUES (170, '读者群', 3262);
INSERT INTO `comic_category_filter` VALUES (171, '读者群', 3263);
INSERT INTO `comic_category_filter` VALUES (172, '读者群', 3264);
INSERT INTO `comic_category_filter` VALUES (173, '读者群', 13626);
INSERT INTO `comic_category_filter` VALUES (174, '进度', 0);
INSERT INTO `comic_category_filter` VALUES (175, '进度', 2309);
INSERT INTO `comic_category_filter` VALUES (176, '进度', 2310);
INSERT INTO `comic_category_filter` VALUES (177, '地域', 0);
INSERT INTO `comic_category_filter` VALUES (178, '地域', 2304);
INSERT INTO `comic_category_filter` VALUES (179, '地域', 2305);
INSERT INTO `comic_category_filter` VALUES (180, '地域', 2306);
INSERT INTO `comic_category_filter` VALUES (181, '地域', 2307);
INSERT INTO `comic_category_filter` VALUES (182, '地域', 2308);
INSERT INTO `comic_category_filter` VALUES (183, '地域', 8435);

SET FOREIGN_KEY_CHECKS = 1;
