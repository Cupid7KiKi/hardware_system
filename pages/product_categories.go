package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProductcategoriesTable(ctx *context.Context) table.Table {

	productCategories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := productCategories.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("分类名称", "category_name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	//info.AddField("多级分类", "parent_id", db.Int)
	info.AddField("创建时间", "created_at", db.Timestamp)

	info.SetTable("product_categories").SetTitle("产品分类").SetDescription("五金物品是涵盖广泛的产品类别，主要包括工具类（手动工具、电动工具、气动工具）、建筑五金（门窗五金、水暖五金、装饰五金）、家居五金（厨房五金、卫浴五金、家具五金）、电子五金（电子元件五金、线缆五金）以及其他五金（紧固件、标准件、特殊五金）等。这些产品广泛应用于建筑、家居、工业、电子等领域，具有实用性强、种类丰富、功能多样等特点，是现代生活中不可或缺的物资。")

	formList := productCategories.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisableWhenCreate()
	formList.AddField("分类名称", "category_name", db.Varchar, form.Text)
	//formList.AddField("多级分类", "parent_id", db.Int, form.Number)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).
		FieldHide().FieldNowWhenInsert()

	formList.SetTable("product_categories").SetTitle("产品分类").SetDescription("填写产品分类信息")

	return productCategories
}
