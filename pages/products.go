package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/template"

	//form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"hardware_system/service"
)

func GetProductsTable(ctx *context.Context) table.Table {

	products := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := products.GetInfo()
	info.HideFilterArea()

	info.AddField("ID", "id", db.Int).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("产品编码", "product_code", db.Varchar)
	info.AddField("产品名称", "product_name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("所属分类", "category_name", db.Int).
		FieldJoin(types.Join{
			Table:     "product_categories", // 连表的表名
			Field:     "category_id",        // 要连表的字段
			JoinField: "id",                 // 连表的表的字段
		}).
		////筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	//FieldDisplay(func(value types.FieldModel) interface{} {
	//	return value.Row["product_categories_goadmin_join_category_name"]
	//})
	info.AddField("规格", "spec", db.Text)
	info.AddField("单位", "unit", db.Varchar)
	info.AddField("品牌", "brand", db.Varchar)
	//info.AddField("采购价", "purchase_price", db.Decimal)
	info.AddField("销售价", "sale_price", db.Decimal).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Value + " 元"
	})
	info.AddField("是否启用", "is_active", db.Tinyint).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "1" {
			return "是"
		} else {
			return "否"
		}
	})
	info.AddField("创建时间", "created_at", db.Timestamp)

	info.SetTable("products").SetTitle("产品管理").SetDescription("描述产品信息")
	//SetAction(template.HTML(`<a href="/ks/info/financial_records/new"><i class="fa fa-google">123</i></a>`))

	formList := products.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("产品编码", "product_code", db.Varchar, form.Text).FieldMust()
	//.SetError(errors.PageError500, "")
	formList.AddField("产品名称", "product_name", db.Varchar, form.Text).FieldMust()
	formList.AddField("所属分类", "category_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetCategories(), "category_name", "id")).FieldMust()
	formList.AddField("规格", "spec", db.Text, form.Text)
	formList.AddField("单位", "unit", db.Varchar, form.Text).FieldMust()
	formList.AddField("品牌", "brand", db.Varchar, form.Text)
	//formList.AddField("采购价", "purchase_price", db.Decimal, form.Custom).
	//	FieldCustomContent(template.HTML(`
	//	<span class="input-group-addon">¥</span>
	//	<input type="text" name="purchase_price" value="{{ .Value }}" style="width: 120px;text-align: right;" placeholder="采购价" class="form-control purchase_price">
	//	`)).FieldCustomJs(template.JS(`
	//	$(function () {
	// 			$('.purchase_price').inputmask({
	//			   alias: "currency",
	//			   radixPoint: ".",
	//			   prefix: "",
	//			   groupSeparator: ",",    // 千位分隔符
	//			   digits: 2,              // 强制两位小数
	//			   autoGroup: true,        // 输入时自动添加千位分隔符
	//			   removeMaskOnSubmit: true
	// 			});
	//     });
	//	console.log("Hello test2!")
	// `))
	formList.AddField("销售价", "sale_price", db.Decimal, form.Custom).
		FieldCustomContent(template.HTML(`
		<span class="input-group-addon">¥</span>
		<input type="text" name="sale_price" value="{{ .Value }}" style="width: 120px;text-align: right;" placeholder="销售价" class="form-control sale_price">
		`)).FieldCustomJs(template.JS(`
			$(function () {
	 			$(".sale_price").inputmask({
				   alias: "currency",
				   radixPoint: ".",
				   prefix: "",
				   groupSeparator: ",",    // 千位分隔符
				   digits: 2,              // 强制两位小数
				   autoGroup: true,        // 输入时自动添加千位分隔符
				   removeMaskOnSubmit: true
	 			});
	     	});
		console.log("Hello test1!")
	 `)).FieldMust()

	// 添加一个自定义类型的表单字段
	//formList.AddField("content", "content", db.Varchar, form.Custom).FieldCustomContent(template.HTML(`
	//	<div>
	//	<label for="custom-field">自定义字段</label>
	//	<input type="text" class="custom-field" name="custom-field" placeholder="请输入自定义内容">
	//	</div>
	//	`)).FieldCustomCss(template.CSS(`
	//		.custom-field{
	//			border: 10px solid #ccc;
	//			padding: 5px;
	//			background-color: blue;
	//		}
	//`)).FieldCustomJs(template.JS(`console.log("Hello World!")`))

	//formList.AddField("销售价", "test", db.Decimal, form.Currency)

	//formList.AddField("是否启用", "is_active", db.Boolean, form.SelectSingle). // 单选的选项，text代表显示内容，value代表对应值
	//									FieldOptions(types.FieldOptions{
	//		{Text: "是", Value: "1"},
	//		{Text: "否", Value: "0"},
	//	}).
	//	// 设置默认值
	//	FieldDefault("1").
	//	// 这里返回一个[]string，对应的值是本列的drink字段的值，即编辑表单时会显示的对应值
	//	FieldDisplay(func(model types.FieldModel) interface{} {
	//		return []string{"1"}
	//	})
	formList.AddField("是否启用", "is_active", db.Tinyint, form.Radio).FieldOptions(types.FieldOptions{
		{Text: "是", Value: "1"},
		{Text: "否", Value: "0"},
	}).FieldDefault("1").
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"1"}
		})
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("products").SetTitle("产品管理").SetDescription("填写产品信息")

	return products
}
