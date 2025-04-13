package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"hardware_system/service"
)

func GetOrdersTable(ctx *context.Context) table.Table {

	orders := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := orders.GetInfo().HideFilterArea()

	info.AddField("订单编号", "id", db.Int)
	info.AddField("客户名称", "name", db.Int).
		FieldJoin(types.Join{
			Table:     "customers",   // 连表的表名
			Field:     "customer_id", // 要连表的字段
			JoinField: "id",          // 连表的表的字段
		}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Row["customers_goadmin_join_name"] == nil {
				return "未知客户"
			}
			return value.Row["customers_goadmin_join_name"]
		})
	info.AddField("下单人", "contact", db.Int).
		FieldJoin(types.Join{
			Table:     "customers",
			Field:     "c_name",
			JoinField: "id",
		}).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Row["customers_goadmin_join_contact"] == nil {
				//fmt.Println("1232:", value.Value)
				return "未知下单人"
			}
			fmt.Println("1232:", value.Value)
			return value.Row["customers_goadmin_join_contact"]
		})
	info.AddField("总计金额", "total_amount", db.Decimal).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Value + " 元"
	})
	info.AddField("订单状态", "status", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "pending" {
			return "待定"
		}
		if model.Value == "approved" {
			return "批准"
		}
		if model.Value == "shipped" {
			return "已发货"
		}
		if model.Value == "completed" {
			return "完成"
		}
		if model.Value == "canceled" {
			return "取消"
		}
		return "未知"
	})

	info.AddField("操作员", "operator", db.Varchar)
	info.AddField("创建时间", "created_at", db.Timestamp)

	info.SetTable("orders").SetTitle("订单管理").SetDescription("描述订单大致信息").
		SetAction(template.HTML(`<a href="/ks/info/financial_records/new">记账</a>`))

	detail := orders.GetDetailFromInfo()

	components := template.Default(ctx)
	lHtml := components.Col().SetSize(types.SizeMD(9)).SetContent("").GetContent()
	rHtml := components.Col().SetSize(types.SizeMD(3)).SetContent("&nbsp;&nbsp;" + "<a href=\"/ks/info/financial_records/new\" class=\"btn btn-primary\">记录财务流水</a>\n").GetContent()
	//rHtml := components.Col().SetSize(types.SizeMD(3)).SetContent().GetContent()
	components.Col().SetContent(lHtml + rHtml).GetContent()
	detail.SetFooterHtml(components.Row().SetContent(lHtml + rHtml).GetContent())

	formList := orders.GetForm()
	formList.AddField("订单编号", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("下单人", "c_name", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetContact(), "contact", "id")).
		FieldOnChooseAjax("customer_id", "choose/c_name", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			cn := ctx.FormValue("value")
			cID := service.GetCustomerNameByContact(cn)
			data = make(types.FieldOptions, len(cID))
			data = service.TransSelectionOptions(cID, "name", "id")
			return true, "ok", data
		})
	formList.AddField("客户名称", "customer_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetUniqueCustomers(), "name", "id"))
	//FieldOnChooseAjax("c_name", "choose/customer_id", func(ctx *context.Context) (success bool, msg string, data interface{}) {
	//	c_id := ctx.FormValue("value")
	//	//fmt.Println("客户名称：", c_id)
	//	contact_name := service.GetContactByCustomer(c_id)
	//	//fmt.Println(temp)
	//	//fmt.Println(service.GetContact(temp))
	//	//fmt.Println(service.GetCustomers())
	//	data = make(selection.Options, len(contact_name))
	//	data = service.TransSelectionOptions(contact_name, "contact", "id")
	//	return true, "ok", data
	//}).
	//FieldDisplay(func(model types.FieldModel) interface{} {
	//	return []string{model.Value}
	//})

	formList.AddField("总计金额", "total_amount", db.Decimal, form.Custom).
		FieldCustomContent(template.HTML(`
		<span class="input-group-addon">¥</span>
		<input type="text" name="total_amount" value="{{ .Value }}" style="width: 120px;text-align: right;" placeholder="总计金额" class="form-control total_amount">
		`)).FieldCustomJs(template.JS(`
		$(function () {
	 			$('.total_amount').inputmask({
				   alias: "currency",
				   radixPoint: ".",
				   prefix: "",
				   suffix: "",
				   groupSeparator: ",",    // 千位分隔符
				   digits: 2,              // 强制两位小数
				   autoGroup: true,        // 输入时自动添加千位分隔符
				   removeMaskOnSubmit: true
	 			});
	     });
	 `)).FieldDefault("0.00").FieldHelpMsg("可以暂时填入一个数字")
	formList.AddField("订单状态", "status", db.Enum, form.SelectSingle). // 单选的选项，text代表显示内容，value代表对应值
										FieldOptions(types.FieldOptions{
			{Text: "待定", Value: "pending"},
			{Text: "批准", Value: "approved"},
			{Text: "已发货", Value: "shipped"},
			{Text: "完成", Value: "completed"},
			{Text: "取消", Value: "canceled"},
		}).
		// 设置默认值
		FieldDefault("pending").
		// 这里返回一个[]string，对应的值是本列的drink字段的值，即编辑表单时会显示的对应值
		FieldDisplay(func(model types.FieldModel) interface{} {
			return []string{model.Value}
		})
	formList.AddField("操作员", "operator", db.Varchar, form.Default).FieldDisableWhenUpdate().FieldDisplay(func(value types.FieldModel) interface{} {
		user := auth.Auth(ctx)
		return user.Name
	})
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("orders").SetTitle("订单管理").SetDescription("填写订单大致信息")

	return orders
}
