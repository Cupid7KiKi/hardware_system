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

	info := customers.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("客户名称", "name", db.Int).FieldJoin(types.Join{
		Table:     "customers_companies",
		Field:     "company_id",
		JoinField: "id",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Row["customers_companies_goadmin_join_name"]
	}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("联系人姓名", "name", db.Int).FieldJoin(types.Join{
		Table:     "companies_contacts",
		Field:     "contact_id",
		JoinField: "id",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Row["companies_contacts_goadmin_join_name"]
	}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
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
			return "不良"
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
		}).FieldMust()
	formList.AddField("联系人姓名", "contact_id", db.Int, form.SelectSingle).
		FieldOptionInitFn(
			func(val types.FieldModel) types.FieldOptions {
				contact := service.GetContactByID(val.Value)
				data := make(types.FieldOptions, len(contact))
				data = service.TransFieldOptions(service.GetContactByID(val.Value), "name", "id")
				return data
			},
		).FieldMust()
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
	//formList.SetUpdateFn(func(values form2.Values) error {
	//	//fmt.Println("看看", values)
	//	//// 构造批量插入的SQL
	//	//var valueStrings []string
	//	//var valueArgs []interface{}
	//
	//	//// 确保所有数组长度相同
	//	//if len(values["product_id"]) != len(values["quantity"]) ||
	//	//	len(values["quantity"]) != len(values["amount"]) {
	//	//	return fmt.Errorf("数据长度不匹配")
	//	//}
	//
	//	//var id int
	//	//if len(values["id"]) > 0 {
	//	//	id, _ = strconv.Atoi(values.Get("id"))
	//	//}
	//	fmt.Println("kk:", values)
	//	var companyID string
	//	if len(values["company_id"]) > 0 {
	//		companyID = values.Get("company_id")
	//	}
	//	var ID string
	//	if len(values["id"]) > 0 {
	//		ID = values.Get("id")
	//	}
	//	var contactID string
	//	if len(values["contact_id"]) > 0 {
	//		contactID = values.Get("contact_id")
	//	}
	//	var creditRating string
	//	if len(values["credit_rating"]) > 0 {
	//		creditRating = values.Get("credit_rating")
	//	}
	//	fmt.Println(reflect.TypeOf(values.Get("company_id")), reflect.TypeOf(values.Get("contact_id")), reflect.TypeOf(values.Get("credit_rating")))
	//
	//	fmt.Println("==== 开始插入 ====") // 调试标记1
	//	// 执行批量插入
	//	//stmt := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity,sale_price, amount) VALUES %s",
	//	//	strings.Join(valueStrings, ","))
	//
	//	upstmt := fmt.Sprintf("UPDATE customers SET company_id = '%s' ,contact_id = '%s' ,credit_rating = '%s'", companyID, contactID, creditRating)
	//	upstmt2 := fmt.Sprintf("UPDATE orders SET customer_id = %s,contact_id = %s WHERE customer_id = %s", ID, contactID, contactID)
	//
	//	// 打印最终SQL（使用占位符版本）
	//	fmt.Printf("执行SQL: %s\n参数: %v\n", upstmt)
	//	fmt.Printf("执行SQL: %s\n参数: %v\n", upstmt2)
	//
	//	exec, err := service.GetDb().Exec(upstmt)
	//	_, err = service.GetDb().Exec(upstmt2)
	//	//exec, err := service.GetDb().Exec("UPDATE orders SET total_amount = (?) WHERE id = (?)", total_amount, orderID)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Println(exec)
	//	if err != nil {
	//		return fmt.Errorf("数据库插入失败: %v", err)
	//	}
	//	fmt.Println("==== 结束插入 ====") // 调试标记2
	//	return nil
	//})

	formList.SetTable("customers").SetTitle("客户管理").SetDescription("填写客户信息")

	return customers
}
