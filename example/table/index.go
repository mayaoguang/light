package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	sql := generateFunctionQuery("user", map[string]interface{}{
		"id":          1,
		"name":        "wang",
		"create_time": 1234567890,
		"status":      2,
	})
	fmt.Println(sql)
}

func generateFunctionQuery(table string, mapData map[string]interface{}) string {
	l := len(mapData)
	keys, val := make([]string, 0, l), make([]string, 0, l)
	for k, v := range mapData {
		keys = append(keys, k)
		tmp, _ := json.Marshal(v)
		val = append(val, string(tmp))
	}
	sql := fmt.Sprintf("insert into %s (%s) values (%s)", table,
		strings.Join(keys, ","), strings.Join(val, ","))
	return sql
}
