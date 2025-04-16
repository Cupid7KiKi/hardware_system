package service

import (
	"fmt"
)

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
func GetCompanyNameByContactName(contact string) (results []map[string]interface{}) {
	query := fmt.Sprintf("SELECT \n    c.id AS id,\n    cc.name AS name\nFROM \n    orders o\nJOIN \n    customers c ON o.id = c.id\nJOIN \n    customers_companies cc ON c.company_id = cc.id\nWHERE \n    o.contact_id = '%s';", contact)
	results, err := GetDb().Query(query)
	if err != nil {
		return
	}
	return
}

func GetContactName() (results []map[string]interface{}) {
	// 构建查询
	query := fmt.Sprintf("SELECT * FROM companies_contacts;")

	// 执行查询
	results, err := GetDb().Query(query)
	if err != nil {
		return
	}
	return
}

func GetContactByCname(val string) (results []map[string]interface{}) {
	// 构建查询
	query := fmt.Sprintf("SELECT contact FROM customers WHERE contact = %s", val)
	fmt.Println(query)
	// 执行查询
	results, err := GetDb().Query(query)
	if err != nil {
		return
	}
	return
}
