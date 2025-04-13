package service

func GetCategories() (results []map[string]interface{}) {
	results, err := GetDb().Query("select * from product_categories")
	if err != nil {
		return
	}
	return
}
