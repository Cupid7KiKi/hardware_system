package service

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Order struct {
	Amount decimal.Decimal `db:"amount"` // 使用 decimal.Decimal 接收 DECIMAL 数据
}

func GetAmounts(s string) (results []map[string]interface{}) {
	//var order Order
	results, err := GetDb().Query("select total_amount from orders where id = " + s)
	if err != nil {
		return
	}
	return
}

func TransDecimal(results []map[string]interface{}) (totalAmount decimal.Decimal) {
	var result map[string]interface{}
	for _, row := range results {
		//for key, value := range row {
		//	fmt.Printf("%s: %v\n", key, value)
		//
		//}
		result = row
	}
	//fmt.Println(reflect.TypeOf(result["total_amount"]))
	// 获取 total_amount 的值
	totalAmountBytes, ok := result["total_amount"].([]byte)
	if !ok {
		fmt.Println("total_amount is not a byte slice")
		return
	}
	// 将字节切片转换为字符串
	totalAmountStr := string(totalAmountBytes)
	// 将字符串转换为 decimal.Decimal 类型
	totalAmount, err := decimal.NewFromString(totalAmountStr)
	if err != nil {
		fmt.Println("Error converting total_amount to decimal:", err)
		return decimal.NewFromInt(0)
	}
	// 输出结果
	//fmt.Println("Total Amount:", totalAmount)
	return totalAmount
}

func GetCurrentMonthIncome() int64 {
	results, err := GetDb().Query("SELECT\n  SUM(amount) AS total_amount \nFROM\n  financial_records \nWHERE\n  type = 'income' \n  AND MONTH(created_at) = MONTH(CURRENT_DATE()) \n  AND YEAR(created_at) = YEAR(CURRENT_DATE());")
	if err != nil {
		return 0
	}
	temp := TransDecimal(results)
	return temp.IntPart()
}

func GetCurrentMonthExpense() int64 {
	results, err := GetDb().Query("SELECT\n  SUM(amount) AS total_amount \nFROM\n  financial_records \nWHERE\n  type = 'expense' \n  AND MONTH(created_at) = MONTH(CURRENT_DATE()) \n  AND YEAR(created_at) = YEAR(CURRENT_DATE());")
	if err != nil {
		return 0
	}
	temp := TransDecimal(results)
	return temp.IntPart()
}
