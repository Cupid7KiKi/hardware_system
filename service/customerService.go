package service

import (
	"fmt"
)

func GetUniqueCustomers() (results []map[string]interface{}) {
	results, err := GetDb().Query("SELECT t.* FROM customers t\nINNER JOIN (\n    SELECT MIN(id) as min_id \n    FROM customers \n    GROUP BY name\n) g ON t.id = g.min_id;")
	if err != nil {
		return
	}
	return
}

func GetCustomers() (results []map[string]interface{}) {
	results, err := GetDb().Query("SELECT * FROM customers")
	if err != nil {
		return
	}
	return
}

func GetCustomersCount() int64 {
	results, err := GetDb().Query("select count(*) from customers")
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
		fmt.Println("customers记录数:", count)
	}
	return count
}

func GetContactByCustomer(customerID string) (results []map[string]interface{}) {
	results, err := GetDb().Query("SELECT id,contact \nFROM customers \nWHERE name = (SELECT name FROM customers WHERE id = ?);", customerID)
	if err != nil {
		return
	}
	return
}
func GetCustomerNameByContact(contact string) (results []map[string]interface{}) {
	results, err := GetDb().Query("SELECT id,name \nFROM customers \nWHERE contact = (SELECT contact FROM customers WHERE id = ?);", contact)
	if err != nil {
		return
	}
	return
}

func GetContact() (results []map[string]interface{}) {
	//strs := make([]string, len(t))
	//for i, n := range t {
	//	strs[i] = fmt.Sprintf("%s", n)
	//}
	//result := strings.Join(strs, ",")
	//fmt.Println("result:", result)
	//results, err := GetDb().Query("SELECT * FROM customers WHERE contact in(?)", result)
	// 构建占位符和参数切片
	//placeholders := make([]string, len(t))
	//params := make([]interface{}, len(t))
	//for i, n := range t {
	//	placeholders[i] = "?"
	//	params[i] = n
	//}

	// 构建查询
	query := fmt.Sprintf("SELECT id,contact FROM customers;")

	// 执行查询
	results, err := GetDb().Query(query)
	if err != nil {
		return
	}
	return
}
