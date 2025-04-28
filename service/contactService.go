package service

import "fmt"

func GetContactByID(val string) (results []map[string]interface{}) {
	//var order Order
	results, err := GetDb().Query("select * from companies_contacts \nWHERE company_id = (SELECT company_id from companies_contacts WHERE id = ?);", val)
	if err != nil {
		return
	}
	return
}

func GetContactNameByCompanyName(company string) (results []map[string]interface{}) {
	results, err := GetDb().Query("SELECT id,name FROM companies_contacts \nWHERE id IN (SELECT contact_id FROM customers WHERE company_id = ?)", company)
	if err != nil {
		return
	}
	return results
}

func DisplayContactName(val string) map[string]interface{} {
	query := fmt.Sprintf("SELECT \n    ct.name AS contact_name\nFROM \n    order_items oi\nJOIN \n    orders o ON oi.order_id = o.id\nJOIN \n    companies_contacts ct ON o.contact_id = ct.id\nWHERE \n    oi.order_id = '%s';", val)
	results, err := GetDb().Query(query)
	if err != nil {
		return nil
	}
	contactsname := make(map[string]interface{})
	var values interface{}

	for _, v := range results {
		if s, ok := v["contact_name"]; ok {
			values = s
		}
	}
	// 将结果存入map
	contactsname["contact_name"] = values
	return contactsname
}
