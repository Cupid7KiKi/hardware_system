package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"hardware_system/service"
)

func GetCustomersTable(ctx *context.Context) table.Table {

	customers := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := customers.GetInfo()

	info.AddField("ID", "id", db.Int)
	info.AddField("客户名称", "name", db.Int).FieldJoin(types.Join{
		Table:     "customers_companies",
		Field:     "company_id",
		JoinField: "id",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Row["customers_companies_goadmin_join_name"]
	})
	info.AddField("联系人姓名", "name", db.Int).FieldJoin(types.Join{
		Table:     "companies_contacts",
		Field:     "contact_id",
		JoinField: "id",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Row["companies_contacts_goadmin_join_name"]
	})
	info.AddField("联系电话", "phone", db.Varchar).FieldJoin(types.Join{
		Table:     "companies_contacts",
		Field:     "phone",
		JoinField: "id",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Row["companies_contacts_goadmin_join_phone"]
	})
	info.AddField("地址", "address", db.Text).FieldJoin(types.Join{
		Table:     "customers_companies",
		Field:     "address",
		JoinField: "id",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Row["customers_companies_goadmin_join_address"]
	})
	info.AddField("信用评级", "credit_rating", db.Tinyint).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "1" {
			return "良好"
		} else {
			return "老赖"
		}
	})

	info.SetTable("customers").SetTitle("客户管理").SetDescription("描述客户信息")

	formList := customers.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("客户名称", "company_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetCompanies(), "name", "id")).
		FieldOnChooseAjax("contact_id", "choose/company_id", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			company := ctx.FormValue("value")
			contacts := service.GetContactByCompany(company)
			data = make(types.FieldOptions, len(contacts))
			data = service.TransSelectionOptions(contacts, "name", "id")
			return true, "ok", data
		})
	formList.AddField("联系人姓名", "contact_id", db.Int, form.SelectSingle).
		FieldOptionInitFn(
			func(val types.FieldModel) types.FieldOptions {
				contact := service.GetContactByID(val.Value)
				data := make(types.FieldOptions, len(contact))
				data = service.TransFieldOptions(service.GetContactByID(val.Value), "name", "id")
				return data
			},
		)
	//定义正则表达式
	//regex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	//formList.AddField("联系电话", "phone", db.Varchar, form.Text)
	//SetPostValidator(func(values form2.Values) error {
	//	if !regex.MatchString(values.Get("phone")) {
	//		return fmt.Errorf("您输入的手机号码有误！！！")
	//	}
	//	return nil
	//})
	//formList.AddField("详细地址", "address", db.Text, form.TextArea)
	formList.AddField("信用评级", "credit_rating", db.Tinyint, form.SelectSingle). // 单选的选项，text代表显示内容，value代表对应值
											FieldOptions(types.FieldOptions{
			{Text: "良好", Value: "1"},
			{Text: "欠款", Value: "0"},
		}).
		// 设置默认值
		FieldDefault("1").
		// 这里返回一个[]string，对应的值是本列的drink字段的值，即编辑表单时会显示的对应值
		FieldDisplay(func(model types.FieldModel) interface{} {
			return []string{"1"}
		})

	formList.SetTable("customers").SetTitle("客户管理").SetDescription("填写客户信息")

	return customers
}
