package service

import (
	"fmt"
	"strconv"
	"strings"
)

func GetOrders() (results []map[string]interface{}) {
	results, err := GetDb().Query("select * from orders")
	if err != nil {
		return
	}
	return
}
func GetOrdersCount() int64 {
	results, err := GetDb().Query("select count(*) from orders")
	if err != nil {
		return 0
	}
	var result map[string]interface{}
	for _, row := range results {
		//for key, value := range row {
		//	fmt.Printf("%s: %v\n", key, value)
		//
		//}
		result = row
	}
	//fmt.Println(reflect.TypeOf(results), "|", results, "|", result["count(*)"])
	count, ok := result["count(*)"].(int64)
	if !ok {
		// 处理类型断言失败的情况
		fmt.Println("类型断言失败")
	} else {
		// 成功转换为int64，可以正常使用
		fmt.Println("orders记录数:", count)
	}
	return count
}

func GetOrderByContact(value string) (results []map[string]interface{}) {
	v, errr := strconv.Atoi(value)
	if errr != nil {
		return
	}
	query := fmt.Sprintf("SELECT * FROM orders WHERE c_name = %d", v)
	results, err := GetDb().Query(query)
	if err != nil {
		return
	}
	return
}

func GetOrderID(t []interface{}) (results []map[string]interface{}) {
	// 1. 提取有效参数（过滤非整型值）
	var ids []interface{}
	for _, n := range t {
		switch v := n.(type) {
		case int, int64:
			ids = append(ids, v)
		case float64:
			if v == float64(int64(v)) {
				ids = append(ids, int64(v))
			}
		}
	}
	// 2. 构建参数化查询（避免 SQL 注入）
	query := "SELECT * FROM orders WHERE id IN (?" + strings.Repeat(",?", len(ids)-1) + ")"
	results, err := GetDb().Query(query, ids...)
	if err != nil {
		return
	}
	return
}
