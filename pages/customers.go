package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"regexp"
)

func GetCustomersTable(ctx *context.Context) table.Table {

	customers := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := customers.GetInfo()

	info.AddField("ID", "id", db.Int)
	info.AddField("客户名称", "name", db.Varchar)
	info.AddField("联系人姓名", "contact", db.Varchar)
	info.AddField("联系电话", "phone", db.Varchar)
	info.AddField("详细地址", "address", db.Text)
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
	formList.AddField("客户名称", "name", db.Varchar, form.Text)
	formList.AddField("联系人姓名", "contact", db.Varchar, form.Text)
	//定义正则表达式
	regex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	formList.AddField("联系电话", "phone", db.Varchar, form.Text).
		SetPostValidator(func(values form2.Values) error {
			if !regex.MatchString(values.Get("phone")) {
				return fmt.Errorf("您输入的手机号码有误！！！")
			}
			return nil
		})
	formList.AddField("详细地址", "address", db.Text, form.TextArea)
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
