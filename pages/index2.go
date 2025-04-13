package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/color"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	"hardware_system/pkg"
	"hardware_system/service"
	"net"

	template2 "html/template"
)

func GetIndex2(ctx *context.Context) (types.Panel, error) {
	//获取客户端 IP（自动处理 X-Forwarded-For）
	clientIP := ctx.LocalIP()

	// 确保是 IPv4
	ip := net.ParseIP(clientIP)
	if ip != nil && ip.To4() != nil {
		clientIP = ip.To4().String()
	}

	fmt.Printf("客户端 IPv4: %s\n", clientIP)

	user := auth.Auth(ctx)
	avatar := user.Avatar
	ctx.WriteString("User Avatar: " + avatar)
	//统计产品数量
	//productCount := service.GetProductsCount()

	tmp := template.Default(ctx)

	//box := tmp.Box().
	//	WithHeadBorder().                                                                                                                                                                                                        // 带顶部的边栏
	//	SetHeader("Latest Orders").                                                                                                                                                                                              // 设置头部内容
	//	SetHeadColor("#f7f7f7").                                                                                                                                                                                                 // 设置头部背景色
	//	SetBody(`Hello`).                                                                                                                                                                                                        // 设置内容
	//	SetFooter(`<div class="clearfix"><a href="javascript:void(0)" class="btn btn-sm btn-info btn-flat pull-left">处理订单</a><a href="javascript:void(0)" class="btn btn-sm btn-default btn-flat pull-right">查看所有新订单</a></div>`) // 设置底部HTML
	//
	//cardcard := card.New().SetTitle("测试").SetContent(template.HTML(`<a href="https://www.baidu.com">1</a>`))

	rows := []pkg.BaseComponent{getRow1(tmp), getRow2(tmp), getRow3(tmp)}
	return types.Panel{
		Title: "工作台",
		Content: func() template2.HTML {
			var html template2.HTML
			for i, _ := range rows {
				html += rows[i].GetContent()
			}
			return html
		}(),
	}, nil
	//// 获取数据库连接
	//conn := service.GetDb()
	//
	//// 统计产品数量
	//productCount, err := conn.Table("products").Count()
	//if err != nil {
	//	return types.Panel{}, err
	//}
	//
	//// 统计客户数量
	//customerCount, err := conn.Table("customers").Count()
	//if err != nil {
	//	return types.Panel{}, err
	//}
	//
	//// 统计订单数量
	//orderCount, err := conn.Table("orders").Count()
	//if err != nil {
	//	return types.Panel{}, err
	//}
	//
	//// 创建仪表盘卡片
	//cards := template.DefaultCard().
	//	SetContent(`
	//    <div class="row">
	//        <div class="col-md-4">
	//            <div class="small-box bg-info">
	//                <div class="inner">
	//                    <h3>` + types.StrconvItoa(productCount) + `</h3>
	//                    <p>产品数量</p>
	//                </div>
	//                <div class="icon">
	//                    <i class="ion ion-bag"></i>
	//                </div>
	//                <a href="/info/products" class="small-box-footer">更多信息 <i class="fa fa-arrow-circle-right"></i></a>
	//            </div>
	//        </div>
	//        <div class="col-md-4">
	//            <div class="small-box bg-success">
	//                <div class="inner">
	//                    <h3>` + types.StrconvItoa(customerCount) + `</h3>
	//                    <p>客户数量</p>
	//                </div>
	//                <div class="icon">
	//                    <i class="ion ion-person-add"></i>
	//                </div>
	//                <a href="/info/customers" class="small-box-footer">更多信息 <i class="fa fa-arrow-circle-right"></i></a>
	//            </div>
	//        </div>
	//        <div class="col-md-4">
	//            <div class="small-box bg-warning">
	//                <div class="inner">
	//                    <h3>` + types.StrconvItoa(orderCount) + `</h3>
	//                    <p>订单数量</p>
	//                </div>
	//                <div class="icon">
	//                    <i class="ion ion-stats-bars"></i>
	//                </div>
	//                <a href="/info/orders" class="small-box-footer">更多信息 <i class="fa fa-arrow-circle-right"></i></a>
	//            </div>
	//        </div>
	//    </div>
	//`)
	//
	//// 创建图表
	//chart := chartjs.NewChart().
	//	SetType("bar").
	//	SetLabels([]string{"产品", "客户", "订单"}).
	//	SetDatasets([]chartjs.Dataset{
	//		{
	//			Label:           "数量",
	//			BackgroundColor: []string{"rgba(75, 192, 192, 0.2)", "rgba(54, 162, 235, 0.2)", "rgba(255, 99, 132, 0.2)"},
	//			BorderColor:     []string{"rgba(75, 192, 192, 1)", "rgba(54, 162, 235, 1)", "rgba(255, 99, 132, 1)"},
	//			BorderWidth:     1,
	//			Data:            []int{int(productCount), int(customerCount), int(orderCount)},
	//		},
	//	})
	//
	//// 创建仪表盘面板
	//panel := types.Panel{
	//	Title:       "仪表盘",
	//	Description: "系统统计信息",
	//	Content:     template.HTML(cards.GetContent() + chart.GetContent()),
	//}
	//
	//return panel, nil
}

func getRow1(tmp template.Template) types.RowAttribute {
	col := tmp.Col()
	col1 := col.SetSize(types.SizeMD(4)).SetContent(smallbox.New().SetTitle("产品数量").SetUrl("/ks/info/product").SetValue(service.Int64ToTmp(service.GetProductsCount())).SetColor("blue").SetIcon(icon.ProductHunt).GetContent()).GetContent()
	col2 := col.SetSize(types.SizeMD(4)).SetContent(smallbox.New().SetTitle("客户数量").SetUrl("/ks/info/customers").SetValue(service.Int64ToTmp(service.GetCustomersCount())).SetColor("blue").SetIcon(icon.User).GetContent()).GetContent()
	col3 := col.SetSize(types.SizeMD(4)).SetContent(smallbox.New().SetTitle("订单数量").SetUrl("/ks/info/orders").SetValue(service.Int64ToTmp(service.GetOrdersCount())).SetColor("blue").SetIcon(icon.Reorder).GetContent()).GetContent()
	return tmp.Row().SetContent(col1 + col2 + col3)
}

func getRow2(tmp template.Template) types.RowAttribute {
	col := tmp.Col()
	col1 := col.SetSize(types.SizeMD(6)).SetContent(smallbox.New().SetTitle("本月收入").SetUrl("/ks/info/financial_records").SetValue(service.Int64ToTmp(service.GetCurrentMonthIncome()) + " 元").SetColor("blue").SetIcon(icon.Money).GetContent()).GetContent()
	col2 := col.SetSize(types.SizeMD(6)).SetContent(template.HTML(`<a href="https://www.baidu.com" target="_blank" rel="noopener">123</a>`) + smallbox.New().SetTitle("本月支出").SetUrl("/ks/info/financial_records").SetValue(service.Int64ToTmp(service.GetCurrentMonthExpense())+" 元").SetColor("blue").SetIcon(icon.Money).GetContent()).GetContent()
	return tmp.Row().SetContent(col1 + col2)
}

func getRow3(tmp template.Template) types.RowAttribute {
	lineChart := chartjs.Line().
		SetID("salechart").
		SetHeight(200).
		SetTitle("2025年3月1日 - 2025年8月31日 各产品销售情况").
		SetLabels([]string{"3月", "4月", "5月", "6月", "7月", "8月"}).
		AddDataSet("电子产品").                            // 增加第一条数据
		DSData([]float64{65, 59, 80, 81, 56, 55, 40}). // 设置数据内容
		DSFill(false).                                 // 是否填充颜色
		DSBorderColor(color.Red).                      // 线边框颜色
		DSLineTension(0.2).                            // 设置压力度

		AddDataSet("数码产品"). // 增加第二条数据
		DSData([]float64{28, 48, 40, 19, 86, 27, 90}).
		DSFill(false).
		DSBorderColor("rgba(60,141,188,1)").
		DSLineTension(0.1).
		GetContent()
	return tmp.Row().SetContent(lineChart)
}
