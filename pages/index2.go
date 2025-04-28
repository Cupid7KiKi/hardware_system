package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/color"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	"github.com/GoAdminGroup/themes/sword/components/chart_legend"
	"hardware_system/pkg"
	"hardware_system/service"
	template2 "html/template"
	"time"
)

func GetIndex2(ctx *context.Context) (types.Panel, error) {
	////获取客户端 IP（自动处理 X-Forwarded-For）
	//clientIP := ctx.LocalIP()
	//
	//// 确保是 IPv4
	//ip := net.ParseIP(clientIP)
	//if ip != nil && ip.To4() != nil {
	//	clientIP = ip.To4().String()
	//}

	//fmt.Printf("客户端 IPv4: %s\n", clientIP)

	//user := auth.Auth(ctx)
	//avatar := user.Avatar
	//ctx.WriteString("User Avatar: " + avatar)
	//统计产品数量
	//productCount := service.GetProductsCount()

	tmp := template.Default(ctx)

	rows := []pkg.BaseComponent{getRow0(tmp), getRow1(tmp), getRow2(tmp), getRow3(tmp)}
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
}

//func getRow0(ctx *context.Context, tmp template.Template) types.RowAttribute {
//	//获取客户端 IP（自动处理 X-Forwarded-For）
//	clientIP := ctx.LocalIP()
//
//	// 确保是 IPv4
//	ip := net.ParseIP(clientIP)
//	if ip != nil && ip.To4() != nil {
//		clientIP = ip.To4().String()
//	}
//
//	h := fmt.Sprintf("<div style=\"display:flex;justify-content:space-between;\"><span style=\"font-size:42px;\">%s</span><span style=\"font-size:42px;\">%s</span></div>", clientIP, "上海")
//
//	fmt.Printf("客户端 IPv4: %s\n", clientIP)
//	cardcard := card.New().
//		SetTitle("客户端").
//		SetSubTitle("IP属地").
//		SetContent(template.HTML(h))
//	//SetAction(template.HTML(`<i aria-label="图标: info-circle-o" class="anticon anticon-info-circle-o"><svg viewBox="64 64 896 896" focusable="false" class="" data-icon="info-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M464 336a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8z"></path></svg></i>`)).
//	//SetContent(template.HTML(`<div><div title="" style="margin-right: 16px;"><span><span>周同比</span><span style="margin-left: 8px;">12%</span></span><span style="color: #f5222d;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-up" class="anticon anticon-caret-up"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-up" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M858.9 689L530.5 308.2c-9.4-10.9-27.5-10.9-37 0L165.1 689c-12.2 14.2-1.2 35 18.5 35h656.8c19.7 0 30.7-20.8 18.5-35z"></path></svg></i></span></div><div class="antd-pro-pages-dashboard-analysis-components-trend-index-trendItem" title=""><span><span>日同比</span><span style="margin-left: 8px;">11%</span></span><span style="color: #52c41a;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-down" class="anticon anticon-caret-down"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-down" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z"></path></svg></i></span></div></div>`))
//	infobox := cardcard.GetContent()
//	return tmp.Row().SetContent(infobox)
//	//label := tmp.Label()
//	//label1 := label.SetContent(smallbox.New().SetTitle("123").SetValue(template2.HTML(clientIP)).GetContent()).GetContent()
//	//return tmp.Row().SetContent(label1)
//}

func getRow0(tmp template.Template) types.RowAttribute {
	col := tmp.Col()
	year, month, day := time.Now().Date()
	fmt.Printf("当前年月日: %d年%d月%d日\n", year, month, day)  // 输出示例：2025年4月27日
	hour, minute, seconds := time.Now().Clock()         // 忽略秒数
	fmt.Printf("当前时间: %d时%d分\n", hour, minute, seconds) // 输出示例：14时25分
	//h := fmt.Sprintf("<div class =\"bigjb\"style=\"display:flex;justify-content:space-between;\"><span style=\"font-size:36px;\">%s</span><span style=\"font-size:36px;\">%s</span></div>", "xx省xx市xxxx五金实业有限公司", "上海")
	h := fmt.Sprintf("<div class=\"info-box\">\n        <div class=\"title\">xx省xx市xxxx五金实业有限公司</div>\n        <div class=\"info-row\">\n            <div id=\"beijing-time\">北京时间: %d年%d月%d日 %d时%d分</div>\n            <div id=\"weather\">天气: 天气: 晴 22°C</div>\n        </div>\n    </div>", year, month, day, hour, minute)
	cardcard := col.SetSize(types.SizeMD(12)).
		//SetTitle("xx省xx市xxxx五金实业有限公司").
		SetContent(template.HTML(h))
	//SetAction(template.HTML(`<i aria-label="图标: info-circle-o" class="anticon anticon-info-circle-o"><svg viewBox="64 64 896 896" focusable="false" class="" data-icon="info-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M464 336a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8z"></path></svg></i>`)).
	//SetContent(template.HTML(`<div><div title="" style="margin-right: 16px;"><span><span>周同比</span><span style="margin-left: 8px;">12%</span></span><span style="color: #f5222d;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-up" class="anticon anticon-caret-up"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-up" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M858.9 689L530.5 308.2c-9.4-10.9-27.5-10.9-37 0L165.1 689c-12.2 14.2-1.2 35 18.5 35h656.8c19.7 0 30.7-20.8 18.5-35z"></path></svg></i></span></div><div class="antd-pro-pages-dashboard-analysis-components-trend-index-trendItem" title=""><span><span>日同比</span><span style="margin-left: 8px;">11%</span></span><span style="color: #52c41a;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-down" class="anticon anticon-caret-down"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-down" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z"></path></svg></i></span></div></div>`))

	infobox := cardcard.GetContent()
	//fmt.Println("123", cardcard.Title)
	// 添加自定义 CSS 代码
	customCSS := `<style>
        /* 这里是自定义的 CSS 代码 */
        .info-box {
            width: 100%;
            border: 1px solid #ddd;
            border-radius: 8px;
            padding: 20px;
            font-family: Arial, sans-serif;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        }
        
        .title {
            font-size: 26px;
            font-weight: bold;
            margin-bottom: 15px;
            color: #333;
        }
        
        .info-row {
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        
        #beijing-time {
            font-size: 16px;
            color: #666;
        }
        
        #weather {
            font-size: 16px;
            color: #666;
        }
    </style>`

	// 添加自定义 JS 代码
	customJS := `<script>
        // 这里是自定义的 JS 代码
         // 更新北京时间
        function updateBeijingTime() {
            const options = {
                timeZone: 'Asia/Shanghai',
                hour12: false,
                year: 'numeric',
                month: '2-digit',
                day: '2-digit',
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit'
            };
            const formatter = new Intl.DateTimeFormat('zh-CN', options);
            const parts = formatter.formatToParts(new Date());
            
            let datePart = '', timePart = '';
            parts.forEach(part => {
                if (part.type === 'year' || part.type === 'month' || part.type === 'day') {
                    datePart += part.value;
                    if (part.type === 'year') datePart += '年';
                    else if (part.type === 'month') datePart += '月';
                    else if (part.type === 'day') datePart += '日';
                } else if (part.type === 'hour' || part.type === 'minute' || part.type === 'second') {
                    timePart += part.value;
                    if (part.type === 'hour') timePart += ':';
                    else if (part.type === 'minute') timePart += ':';
                }
            });
        }
    </script>`

	// 将 CSS 和 JS 代码添加到页面内容中
	contentWithCSSJS := template.HTML(customCSS) + infobox + template.HTML(customJS)
	return tmp.Row().SetContent(contentWithCSSJS)
	//label := tmp.Label()
	//label1 := label.SetContent(smallbox.New().SetTitle("123").SetValue(template2.HTML(clientIP)).GetContent()).GetContent()
	//return tmp.Row().SetContent(label1)
	//<span style=" font-size:24px; ">%s</span>
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
	col2 := col.SetSize(types.SizeMD(6)).SetContent(smallbox.New().SetTitle("本月支出").SetUrl("/ks/info/financial_records").SetValue(service.Int64ToTmp(service.GetCurrentMonthExpense()) + " 元").SetColor("blue").SetIcon(icon.Money).GetContent()).GetContent()
	return tmp.Row().SetContent(col1 + col2)
	//template.HTML(`<a href="https://www.baidu.com" target="_blank" rel="noopener">123</a>`)
}

func getRow3(tmp template.Template) types.RowAttribute {
	col := tmp.Col()
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
	col1 := col.SetSize(types.SizeMD(8)).SetContent(lineChart).GetContent()
	pie := chartjs.Pie().
		SetHeight(170).
		SetLabels([]string{"一字螺丝刀", "换气扇", "杯子", "畚斗", "合页", "笔记本电脑"}).
		SetID("pieChart").
		AddDataSet("Chrome").
		DSData([]float64{100, 300, 600, 400, 500, 700}).
		DSBackgroundColor([]chartjs.Color{
			"rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)", "rgb(255, 205, 86)", "rgb(54, 162, 235)", "rgb(255, 99, 132)",
		}).
		GetContent()
	col2 := col.SetSize(types.SizeMD(2)).SetContent(pie).GetContent()

	legend := chart_legend.New().SetData([]map[string]string{
		{
			"label": " 一字螺丝刀",
			"color": "red",
		}, {
			"label": " 换气扇",
			"color": "Green",
		}, {
			"label": " 杯子",
			"color": "yellow",
		}, {
			"label": " 畚斗",
			"color": "blue",
		}, {
			"label": " 合页",
			"color": "light-blue",
		}, {
			"label": " 笔记本电脑",
			"color": "gray",
		},
	}).GetContent()
	col3 := col.SetSize(types.SizeMD(2)).SetContent(legend).GetContent()
	return tmp.Row().SetContent(col1 + col2 + col3)
}
