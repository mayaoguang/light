package main

import (
	"fmt"
)

func main() {
	sqlBase := "CREATE TABLE `quick_sha_%d` (\n  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',\n  `quick_sha256` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'quick_sha256',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_quick_sha256` (`quick_sha256`) USING BTREE COMMENT '快速sha256'\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='文件快速sha256检测表';"
	for i := 0; i < 32; i++ {
		fmt.Println(fmt.Sprintf(sqlBase, i))
	}
}
