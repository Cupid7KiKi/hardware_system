package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"hardware_system/service"
)

func GetCustomercontactsTable(ctx *context.Context) table.Table {

	customerContacts := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := customerContacts.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("客户名称", "name", db.Int).
		FieldJoin(types.Join{
			Table:     "customers",   // 连表的表名
			Field:     "customer_id", // 要连表的字段
			JoinField: "id",          // 连表的表的字段
		}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["customers_goadmin_join_name"]
		})
	info.AddField("联系时间", "contact_time", db.Datetime)
	info.AddField("具体内容", "content", db.Text)
	info.AddField("后续联系日期", "next_followup", db.Date)

	info.SetTable("customer_contacts").SetTitle("客户联系记录").SetDescription("描述客户联系记录")

	formList := customerContacts.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("客户名称", "customer_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetUniqueCustomers(), "name", "id"))
	formList.AddField("联系时间", "contact_time", db.Datetime, form.Datetime)
	formList.AddField("具体内容", "content", db.Text, form.TextArea)
	formList.AddField("后续联系日期", "next_followup", db.Date, form.Datetime)

	formList.SetTable("customer_contacts").SetTitle("客户联系记录").SetDescription("填写客户联系记录")

	return customerContacts
}
