package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"hardware_system/service"
	"regexp"
)

func GetCompaniescontactsTable(ctx *context.Context) table.Table {

	companiesContacts := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := companiesContacts.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("公司", "name", db.Int).FieldJoin(types.Join{
		Table:     "customers_companies",
		Field:     "company_id",
		JoinField: "id",
	}).FieldDisplay(
		func(value types.FieldModel) interface{} {
			return value.Row["customers_companies_goadmin_join_name"]
		}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("姓名", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("联系电话", "phone", db.Varchar)

	info.SetTable("companies_contacts").SetTitle("客户公司联系人表").SetDescription("Companiescontacts")

	formList := companiesContacts.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("公司", "company_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetCompanies(), "name", "id")).FieldMust()
	formList.AddField("姓名", "name", db.Varchar, form.Text).FieldMust()
	//定义正则表达式
	regex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	formList.AddField("联系电话", "phone", db.Varchar, form.Text).
		SetPostValidator(func(values form2.Values) error {
			if !regex.MatchString(values.Get("phone_number")) {
				return fmt.Errorf("您输入的手机号码有误！！！")
			}
			return nil
		})

	formList.SetTable("companies_contacts").SetTitle("客户公司联系人表单").SetDescription("Companiescontacts")

	formList.SetUpdateFn(func(values form2.Values) error {
		//fmt.Println("看看", values)
		//// 构造批量插入的SQL
		//var valueStrings []string
		//var valueArgs []interface{}

		//// 确保所有数组长度相同
		//if len(values["product_id"]) != len(values["quantity"]) ||
		//	len(values["quantity"]) != len(values["amount"]) {
		//	return fmt.Errorf("数据长度不匹配")
		//}

		//var id int
		//if len(values["id"]) > 0 {
		//	id, _ = strconv.Atoi(values.Get("id"))
		//}
		fmt.Println("kk:", values)
		var companyID string
		if len(values["company_id"]) > 0 {
			companyID = values.Get("company_id")
		}
		var contactID string
		if len(values["id"]) > 0 {
			contactID = values.Get("id")
		}
		var contactName string
		if len(values["name"]) > 0 {
			contactName = values.Get("name")
		}
		var phone string
		if len(values["phone"]) > 0 {
			phone = values.Get("phone")
		}
		//fmt.Println(reflect.TypeOf(values.Get("company_id")), reflect.TypeOf(values.Get("id")), reflect.TypeOf(values.Get("name")), reflect.TypeOf(values.Get("phone")))

		fmt.Println("==== 开始插入 ====") // 调试标记1
		// 执行批量插入
		//stmt := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity,sale_price, amount) VALUES %s",
		//	strings.Join(valueStrings, ","))

		upstmt := fmt.Sprintf("UPDATE companies_contacts SET company_id = %s,name = '%s' ,phone = '%s' WHERE id = %s", companyID, contactName, phone, contactID)
		upstmt2 := fmt.Sprintf("UPDATE customers SET company_id = %s,contact_id = %s WHERE contact_id = %s", companyID, contactID, contactID)

		// 打印最终SQL（使用占位符版本）
		fmt.Printf("执行SQL: %s\n参数: %v\n", upstmt)
		fmt.Printf("执行SQL: %s\n参数: %v\n", upstmt2)

		exec, err := service.GetDb().Exec(upstmt)
		_, err = service.GetDb().Exec(upstmt2)
		//exec, err := service.GetDb().Exec("UPDATE orders SET total_amount = (?) WHERE id = (?)", total_amount, orderID)
		if err != nil {
			return err
		}
		fmt.Println(exec)
		if err != nil {
			return fmt.Errorf("数据库插入失败: %v", err)
		}
		fmt.Println("==== 结束插入 ====") // 调试标记2
		return nil
	})

	return companiesContacts
}
