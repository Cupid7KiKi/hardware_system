package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCustomerscompaniesTable(ctx *context.Context) table.Table {

	customersCompanies := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := customersCompanies.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("公司名称", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("地址", "address", db.Text)

	info.SetTable("customers_companies").SetTitle("客户公司表").SetDescription("Customerscompanies")

	formList := customersCompanies.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("公司名称", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("地址", "address", db.Text, form.TextArea)

	formList.SetTable("customers_companies").SetTitle("客户公司信息").SetDescription("Customerscompanies")

	return customersCompanies
}
