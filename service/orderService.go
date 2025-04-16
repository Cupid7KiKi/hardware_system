package service

import (
	"fmt"
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

func GetTotalAmountFromOrderID(orderID string) map[string]interface{} {
	query := fmt.Sprintf("SELECT SUM(amount) FROM order_items WHERE order_id = '%s';", orderID)
	results, err := GetDb().Query(query)
	if err != nil {
		return nil
	}
	var amountStr string
	amount := make(map[string]interface{})
	var values interface{}

	for _, v := range results {
		if s, ok := v["SUM(amount)"]; ok {
			amountStr = string(s.([]uint8))
			values = amountStr
		}
	}
	// 将结果存入map
	amount["SUM(amount)"] = values
	return amount
}
