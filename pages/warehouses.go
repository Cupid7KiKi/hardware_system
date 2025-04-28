package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWarehousesTable(ctx *context.Context) table.Table {

	warehouses := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := warehouses.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("仓库名称", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("仓库位置", "location", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("仓库面积", "capacity", db.Decimal).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Value + "㎡"
	})

	info.SetTable("warehouses").SetTitle("仓库管理").SetDescription("描述仓库信息")

	formList := warehouses.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("仓库名称", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("仓库位置", "location", db.Varchar, form.Text).FieldMust()
	formList.AddField("仓库面积", "capacity", db.Decimal, form.Text).FieldHelpMsg("单位为㎡（平方米）").FieldMust()
	formList.SetTable("warehouses").SetTitle("仓库管理").SetDescription("填写仓库信息")

	return warehouses
}
