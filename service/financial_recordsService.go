package service

import "fmt"

func GetRelatedOrderID(orderid interface{}) interface{} {
	fmt.Println("Handling OrderID:", orderid)
	return orderid
}
