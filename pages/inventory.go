package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"hardware_system/service"
)

func GetInventoryTable(ctx *context.Context) table.Table {

	inventory := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := inventory.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("产品名称", "product_name", db.Int).FieldJoin(types.Join{
		Table:     "products",
		Field:     "product_id",
		JoinField: "id",
	}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("仓库名称", "name", db.Int).FieldJoin(types.Join{
		Table:     "warehouses",
		Field:     "warehouse_id",
		JoinField: "id",
	}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("当前数量", "quantity", db.Int)
	info.AddField("最小库存预警值", "min_stock", db.Int)
	info.AddField("最后更新", "last_update", db.Timestamp)
	//info.AddButton(ctx, "ajax", icon.Android, action.Ajax("/admin/ajax",
	//	func(ctx *context.Context) (success bool, msg string, data interface{}) {
	//		return true, "请求成功，奥利给", ""
	//	}))

	info.SetTable("inventory").SetTitle("库存管理").SetDescription("描述库存信息")

	formList := inventory.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("产品名称", "product_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetProducts(), "product_name", "id")).FieldMust()
	formList.AddField("仓库名称", "warehouse_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetWarehouse(), "name", "id")).FieldMust()
	formList.AddField("当前数量", "quantity", db.Int, form.Number)
	formList.AddField("最小库存预警值", "min_stock", db.Int, form.Number)
	formList.AddField("最后更新", "last_update", db.Timestamp, form.Datetime).FieldDisableWhenCreate().FieldDisableWhenUpdate()

	formList.SetTable("inventory").SetTitle("库存管理").SetDescription("填写库存信息")

	return inventory
}
