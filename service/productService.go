package service

import (
	"fmt"
)

func GetProducts() (results []map[string]interface{}) {
	results, err := GetDb().Query("select * from products where is_active = 1")
	if err != nil {
		return
	}
	return
}

func GetProductSalePrice(s string) (results []map[string]interface{}) {
	results, err := GetDb().Query("select sale_price as total_amount from products where id = " + s)
	if err != nil {
		return
	}
	return
}

func GetProductsCount() int64 {
	results, err := GetDb().Query("select count(*) from products")
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
		fmt.Println("记录数:", count)
	}
	return count
}
