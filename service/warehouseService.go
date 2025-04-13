package service

func GetWarehouse() (results []map[string]interface{}) {
	results, err := GetDb().Query("select * from warehouses")
	if err != nil {
		return
	}
	return
}
