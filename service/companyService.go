package service

import (
	"fmt"
)

func GetCompanies() (results []map[string]interface{}) {
	//var order Order
	results, err := GetDb().Query("SELECT * FROM customers_companies")
	if err != nil {
		return
	}
	return
}

func GetContactByCompany(val string) (results []map[string]interface{}) {
	//var order Order
	results, err := GetDb().Query("SELECT * FROM companies_contacts WHERE company_id = ?", val)
	if err != nil {
		return
	}
	return
}

//func GetUniqueCustomers() (results []map[string]interface{}) {
//	results, err := GetDb().Query("SELECT * FROM customers_companies")
//	if err != nil {
//		return
//	}
//	return
//}

func DisplayCompanyName(val string) map[string]interface{} {
	query := fmt.Sprintf("SELECT \n    cc.name AS company_name\nFROM \n    order_items oi\nJOIN \n    orders o ON oi.order_id = o.id\nJOIN \n    customers c ON o.customer_id = c.id\nJOIN \n    customers_companies cc ON c.company_id = cc.id\nWHERE \n    oi.order_id = '%s';", val)
	results, err := GetDb().Query(query)
	if err != nil {
		return nil
	}
	companysname := make(map[string]interface{})
	var values interface{}

	for _, v := range results {
		if s, ok := v["company_name"]; ok {
			values = s
		}
	}
	// 将结果存入map
	companysname["company_name"] = values
	return companysname
}
