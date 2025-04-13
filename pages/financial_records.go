package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	//tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"hardware_system/service"
	//template2 "html/template"
)

func GetFinancialrecordsTable(ctx *context.Context) table.Table {

	financialRecords := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := financialRecords.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("财务流水类型", "type", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "income" {
			return "收入"
		}
		if model.Value == "expense" {
			return "支出"
		}
		return "未知"
	})
	info.AddField("金额", "amount", db.Decimal).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Value + " 元"
	})
	info.AddField("支付方式", "payment_method", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "CashPayment" {
			return "现金支付"
		}
		if model.Value == "BankCardPayment" {
			return "银行卡支付"
		}
		if model.Value == "MobilePayment" {
			return "移动支付"
		}
		if model.Value == "BankTransferPayment" {
			return "银行卡转账支付"
		}
		return "未知"
	})
	info.AddField("相关订单编号", "id", db.Varchar).
		FieldJoin(types.Join{
			Table:     "orders",        // 连表的表名
			Field:     "related_order", // 要连表的字段
			JoinField: "id",            // 连表的表的字段
		}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["orders_goadmin_join_id"]
		})
	info.AddField("备注信息", "description", db.Text)
	info.AddField("创建时间", "created_at", db.Timestamp)

	info.SetTable("financial_records").SetTitle("财务流水").SetDescription("描述财务流水信息")

	detail := financialRecords.GetDetail()
	//// 修改详情页相关订单字段显示
	//detail.AddField("相关订单", "related_order", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
	//	// 获取当前记录的订单ID
	//	orderID := value.Row["related_order"]
	//
	//	// 返回带动态参数的HTML链接
	//	return template.HTML(fmt.Sprintf(
	//		`<a href="/ks/info/orders/detail?__goadmin_detail_pk=%d" class="btn btn-sm btn-info">
	//        <i class="fa fa-eye"></i> 查看订单 %d
	//    </a>`,
	//		orderID,
	//		orderID,
	//	))
	//})
	detail.AddField("ID", "id", db.Int)
	detail.AddField("财务流水类型", "type", db.Enum).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "income" {
			return "收入"
		}
		if model.Value == "expense" {
			return "支出"
		}
		return "未知"
	})
	detail.AddField("金额", "amount", db.Decimal)
	detail.AddField("支付方式", "payment_method", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "CashPayment" {
			return "现金支付"
		}
		if model.Value == "BankCardPayment" {
			return "银行卡支付"
		}
		if model.Value == "MobilePayment" {
			return "移动支付"
		}
		if model.Value == "BankTransferPayment" {
			return "银行卡转账支付"
		}
		return "未知"
	})
	detail.AddField("相关订单编号", "related_order", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		// 获取当前记录的订单ID
		orderID := value.Row["related_order"]

		// 返回带动态参数的HTML链接
		return template.HTML(fmt.Sprintf(
			`<div style="display:flex;justify-content:space-between;align-items:center;"><p style="margin:0;">%d</p><a href="/ks/info/orders/detail?__goadmin_detail_pk=%d" class="btn btn-sm btn-info">
	      <i class="fa fa-eye"></i> 查看订单
	  </a>
		</div>`,
			orderID,
			orderID,
		))
	})

	//// 3. 调用 rID 函数并获取返回值
	//result := rID(types.FieldDisplay) interface{} {
	//	return nil
	//})

	detail.AddField("备注信息", "description", db.Text)
	detail.AddField("创建时间", "created_at", db.Timestamp)

	//link := fmt.Sprintf(`<a href="/ks/info/orders/detail?__goadmin_detail_pk=%d" class="btn btn-sm btn-info">
	//       <i class="fa fa-eye"></i> 查看订单
	//   </a>`, data.ID.(int))
	//fmt.Println(link)

	//components := tmpl.Default(ctx)
	//lHtml := components.Col().SetSize(types.SizeMD(9)).SetContent("").GetContent()
	////rHtml := components.Col().SetSize(types.SizeMD(3)).SetContent("&nbsp;&nbsp;" + "<a href=\"/ks/info/orders/detail?__goadmin_detail_pk=3&\" class=\"btn btn-primary\">查看相关订单</a>\n").GetContent()
	//rHtml := components.Col().SetSize(types.SizeMD(3)).SetContent(template2.HTML(link)).GetContent()
	//components.Col().SetContent(lHtml + rHtml).GetContent()
	//detail.SetFooterHtml(components.Row().SetContent(lHtml + rHtml).GetContent())

	formList := financialRecords.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("财务流水类型", "type", db.Enum, form.SelectSingle). // 单选的选项，text代表显示内容，value代表对应值
										FieldOptions(types.FieldOptions{
			{Text: "收入", Value: "income"},
			{Text: "支出", Value: "expense"},
		})
	//.FieldDisplay(func(value types.FieldModel) interface{} {
	//	decimalValue := service.TransDecimal(service.GetAmounts("1"))
	//	value.Value = decimalValue.String()
	//	return value.Value
	//})
	//fmt.Println(service.TransDecimal(service.GetAmounts("1")))
	formList.AddField("支付方式", "payment_method", db.Varchar, form.SelectSingle). // 单选的选项，text代表显示内容，value代表对应值
											FieldOptions(types.FieldOptions{
			{Text: "现金支付", Value: "CashPayment"},
			{Text: "银行卡支付", Value: "BankCardPayment"},
			{Text: "移动支付", Value: "MobilePayment"},
			{Text: "银行转账支付", Value: "BankTransferPayment"},
		}).
		// 设置默认值
		FieldDefault("MobilePayment").
		// 这里返回一个[]string，对应的值是本列的drink字段的值，即编辑表单时会显示的对应值
		FieldDisplay(func(model types.FieldModel) interface{} {
			return []string{"MobilePayment"}
		})
	formList.AddField("相关订单编号", "related_order", db.Varchar, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetOrders(), "id", "id")).
		FieldOnChooseAjax("amount", "choose/related_order", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			r_order := ctx.FormValue("value")
			decimalValue := service.TransDecimal(service.GetAmounts(r_order))
			data = decimalValue
			return true, "ok", data
		})
	formList.AddField("金额", "amount", db.Decimal, form.Custom).FieldCustomContent(template.HTML(`
	<span class="input-group-addon">¥</span>
	<input type="text" name="amount" value="{{ .Value }}" style="width: 120px;text-align: right;" class="form-control amount" readonly>
	`)).FieldCustomJs(template.JS(`
		$(function () {
	 			$('.amount').inputmask({
				   alias: "currency",
				   radixPoint: ".",
				   prefix: "",
				   groupSeparator: ",",    // 千位分隔符
				   digits: 2,              // 强制两位小数
				   autoGroup: true,        // 输入时自动添加千位分隔符
				   removeMaskOnSubmit: true
	 			});
	     });
	 `))
	formList.AddField("备注信息", "description", db.Text, form.TextArea).FieldDefault("无")
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("financial_records").SetTitle("财务流水").SetDescription("填写财务流水信息")

	return financialRecords
}
