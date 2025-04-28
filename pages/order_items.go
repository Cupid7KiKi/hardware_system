package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/shopspring/decimal"
	"hardware_system/service"
	"strconv"
	"strings"
)

func GetOrderitemsTable(ctx *context.Context) table.Table {

	orderItems := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("mysql"))

	info := orderItems.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int)
	info.AddField("订单编号", "id", db.Varchar).FieldJoin(types.Join{
		Table:     "orders",
		Field:     "order_id",
		JoinField: "id", // 连表的表的字段
	}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["orders_goadmin_join_id"]
		})
	info.AddField("公司名称", "company_name", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return service.DisplayCompanyName(fmt.Sprintf(value.Row["orders_goadmin_join_id"].(string)))["company_name"]
		})
	info.AddField("下单人", "contact_name", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			//fmt.Println(service.DisplayContactName(fmt.Sprintf(value.Row["orders_goadmin_join_id"].(string))))
			return service.DisplayContactName(fmt.Sprintf(value.Row["orders_goadmin_join_id"].(string)))["contact_name"]
		}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("产品名称", "product_name", db.Int).FieldJoin(types.Join{
		Table:     "products",
		Field:     "product_id",
		JoinField: "id",
	}).
		//筛选时支持模糊查询
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return value.Row["products_goadmin_join_product_name"]
		})
	info.AddField("数量", "quantity", db.Int)
	info.AddField("小计", "amount", db.Decimal).FieldDisplay(func(value types.FieldModel) interface{} {
		return value.Value + " 元"
	})

	info.SetTable("order_items").SetTitle("订单明细").SetDescription("订单详细信息")

	formList := orderItems.GetForm().AddXssJsFilter()
	formList.AddField("ID", "id", db.Int, form.Default).
		FieldDisableWhenCreate()
	//formList.AddField("下单人", "name", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetCustomers(), "contact", "id")).
	//	FieldOnChooseAjax("order_id", "choose/name", func(ctx *context.Context) (success bool, msg string, data interface{}) {
	//		cn := ctx.FormValue("value")
	//		//fmt.Println(cn)
	//		oID := service.GetOrderByContact(cn)
	//		//fmt.Println("oID:", oID)
	//		data = make(selection.Options, len(oID))
	//		data = service.TransSelectionOptions(oID, "id", "id")
	//		return true, "ok", data
	//	})
	formList.AddField("订单编号", "order_id", db.Int, form.SelectSingle).FieldOptions(service.TransFieldOptions(service.GetOrders(), "id", "id")).FieldDisplayButCanNotEditWhenUpdate()
	//添加动态表格
	formList.AddTable("产品条目", "setting", func(panel *types.FormPanel) {
		panel.AddField("产品名称", "product_id", db.Int, form.SelectSingle).
			FieldOptions(service.TransFieldOptions(service.GetProducts(), "product_name", "id"))
		panel.AddField("数量", "quantity", db.Int, form.Number).
			FieldDefault("1")
		panel.AddField("产品单价", "sale_price", db.Decimal, form.Custom).FieldCustomContent(template.HTML(`
	<span class="input-group-addon">¥</span>
	<input type="text" name="sale_price" value="{{ .Value }}" style="width: 120px;text-align: right;" class="form-control sale_price" readonly>
	`)).FieldCustomJs(template.JS(`
$(function () {
    $('.sale_price').inputmask({
        alias: "currency",
        radixPoint: ".",
        prefix: "",
        removeMaskOnSubmit: true
    });
});
`)).
			FieldDisplay(func(value types.FieldModel) interface{} {
				return value.Value + " 元"
			})
		panel.AddField("合计", "amount", db.Decimal, form.Custom).FieldCustomContent(template.HTML(`
	<span class="input-group-addon">¥</span>
	<input type="text" name="amount" value="{{ .Value }}" style="width: 120px;text-align: right;" class="form-control amount" readonly>
	`)).FieldCustomJs(template.JS(`
$(function () {
    $('.amount').inputmask({
        alias: "currency",
        radixPoint: ".",
        prefix: "",
        removeMaskOnSubmit: true
    });
});
`)).
			FieldDisplay(func(value types.FieldModel) interface{} {
				return value.Value + " 元"
			})
	}).FieldDisableWhenUpdate()

	formList.AddField("总计", "total_amount", db.Decimal, form.Custom).FieldCustomContent(template.HTML(`
	<span class="input-group-addon">¥</span>
	<input type="text" name="total_amount" value="{{ .Value }}" style="width: 120px;text-align: right;" class="form-control total_amount" readonly>
	`)).FieldCustomJs(template.JS(`
			$(function () {
	 			$('.total_amount').inputmask({
				   alias: "currency",
				   radixPoint: ".",
				   prefix: "",
				   removeMaskOnSubmit: true
	 			});
	     	});
	 `)).FieldDisableWhenUpdate().FieldDisplayButCanNotEditWhenUpdate()

	formList.AddField("产品名称", "product_id", db.Int, form.SelectSingle).FieldDisableWhenCreate().FieldOptions(service.TransFieldOptions(service.GetProducts(), "product_name", "id")).
		FieldOnChooseAjax("unit_price", "choose/product_id", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			pID := ctx.FormValue("value")
			s_pice := service.GetProductSalePrice(pID)
			//fmt.Println(reflect.TypeOf(service.TransDecimal(s_pice)))
			//fmt.Println(service.TransDecimal(s_pice))
			data = service.TransDecimal(s_pice)
			return true, "ok", data
		})
	formList.AddField("数量", "quantity", db.Int, form.Custom).FieldCustomContent(template.HTML(`
	<input type="number" name="quantity" value="{{ .Value }}" style="width: 120px;text-align: right;text-align: center;" class="form-control quantity">
	`)).
		FieldDisableWhenCreate().FieldCustomJs(template.JS(`
		var pathname = window.location.pathname;
		if(pathname === "/ks/info/order_items/edit"){
			$(document).on('input change', 'input[name="quantity"]', function() {
				console.log('数量变更为:', $(this).val());
			});
		}
		`))
	formList.AddField("单价", "sale_price", db.Decimal, form.Custom).FieldCustomContent(template.HTML(`
	<span class="input-group-addon">¥</span>
	<input type="text" name="sale_price" value="{{ .Value }}" style="width: 120px;text-align: right;" class="form-control sale_price" readonly>
	`)).FieldCustomJs(template.JS(`
		var pathname = window.location.pathname;
		if(pathname === "/ks/info/order_items/edit"){
			$(function () {
	 			$('.sale_price').inputmask({
				   alias: "currency",
				   radixPoint: ".",
				   prefix: "",
				   removeMaskOnSubmit: true
	 			});
	     });
		}
	 `)).FieldDisableWhenCreate()
	formList.AddField("小计", "amount", db.Decimal, form.Custom).FieldCustomContent(template.HTML(`
	<span class="input-group-addon">¥</span>
	<input type="text" name="amount" value="{{ .Value }}" style="width: 120px;text-align: right;" class="form-control amount" readonly>
	`)).FieldCustomJs(template.JS(`
		var pathname = window.location.pathname;
		if(pathname === "/ks/info/order_items/edit"){
			$(function () {
	 			$('.amount').inputmask({
				   alias: "currency",
				   radixPoint: ".",
				   prefix: "",
				   removeMaskOnSubmit: true
	 			});
	     });
		}
	 `)).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()

	formList.SetInsertFn(func(values form2.Values) error {
		//fmt.Println("看看", values)
		// 构造批量插入的SQL
		var valueStrings []string
		var valueArgs []interface{}

		// 确保所有数组长度相同
		if len(values["product_id"]) != len(values["quantity"]) ||
			len(values["quantity"]) != len(values["amount"]) {
			return fmt.Errorf("数据长度不匹配")
		}

		var orderID int
		if len(values["order_id"]) > 0 {
			orderID, _ = strconv.Atoi(values.Get("order_id"))
		}

		//var total_amount decimal.Decimal
		//if len(values["total_amount"]) > 0 {
		//	total_amount, _ = decimal.NewFromString(values.Get("total_amount"))
		//}

		for i := 0; i < len(values["product_id"]); i++ {
			// 转换各字段数据类型
			productID, err := strconv.Atoi(values["product_id"][i])
			if err != nil {
				return fmt.Errorf("产品ID转换失败: %v", err)
			}

			quantity, err := strconv.Atoi(values["quantity"][i])
			if err != nil {
				return fmt.Errorf("数量转换失败: %v", err)
			}

			salePrice, err := decimal.NewFromString(values["sale_price"][i])
			if err != nil {
				return fmt.Errorf("单价转换失败: %v", err)
			}

			amount, err := decimal.NewFromString(values["amount"][i])
			if err != nil {
				return fmt.Errorf("金额转换失败: %v", err)
			}

			// 构建参数
			valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
			valueArgs = append(valueArgs,
				//contactName, // contact_name INT
				orderID,   // order_id INT
				productID, // product_id INT
				quantity,  // quantity INT
				salePrice, //unit_price DECIMAL(10,2)
				amount,    // amount DECIMAL(10,2)
			)
		}

		fmt.Println("==== 开始插入 ====") // 调试标记1
		// 执行批量插入
		stmt := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity,sale_price, amount) VALUES %s",
			strings.Join(valueStrings, ","))

		// 打印最终SQL（使用占位符版本）
		fmt.Printf("执行SQL: %s\n参数: %v\n", stmt, valueArgs)

		//upstmt := fmt.Sprintf("INSERT INTO orders (total_amount) VALUES %s",
		//	total_amount)

		_, err := service.GetDb().Exec(stmt, valueArgs...)
		//exec, err := service.GetDb().Exec("UPDATE orders SET total_amount = (?) WHERE id = (?)", total_amount, orderID)
		//if err != nil {
		//	return err
		//}
		//fmt.Println(exec)
		if err != nil {
			return fmt.Errorf("数据库插入失败: %v", err)
		}
		fmt.Println("==== 结束插入 ====") // 调试标记2
		return nil
	})

	formList.SetTable("order_items").SetTitle("订单明细").SetDescription("填写订单详细信息")

	if ctx.Path() == "/ks/info/order_items/new" {
		formList.AddJS(template.JS(`
			console.log('This is the new page JavaScript code.');
console.log("现在是新建页面！！！");
$(function() {
    // 初始化输入掩码（合并重复初始化）
    $('.sale_price, .amount, .total_amount').inputmask({
        alias: "currency",
        radixPoint: ".",
        prefix: "",
        removeMaskOnSubmit: true
    });

    // 事件委托：统一监听动态和静态元素
    $(document)
        .on('change', 'select[name="product_id"]', handleProductChange)
        .on('input change', 'input[name="quantity"], input[name="sale_price"]', handleInputChange);

    // 初始化计算
    initCalculations();
});

// 解析带千分位的数字（复用函数）
function parseFormattedNumber(str) {
    const cleanedStr = (str || "0").replace(/[^\d.]/g, '');
    return parseFloat(cleanedStr) || 0;
}

// 产品变更处理
function handleProductChange() {
    const productId = $(this).val();
    const currentRow = $(this).closest('tr');
    console.log('产品变更为:', productId);

    if (!productId) {
        currentRow.find('input[name="sale_price"]').val('0.00');
        calculateRowTotal(currentRow);
        return;
    }

    $.ajax({
        url: '/ks/choose/product_price',
        type: 'POST',
        data: { product_id: productId },
        success: function(response) {
            if (response.success) {
                currentRow.find('input[name="sale_price"]').val(response.data);
                calculateRowTotal(currentRow);
            } else {
                alert('获取价格失败: ' + response.msg);
            }
        },
        error: function(xhr) {
            alert('请求失败: ' + xhr.statusText);
        }
    });
}

// 输入变更处理（防抖优化）
let timer;
function handleInputChange() {
    clearTimeout(timer);
    const currentRow = $(this).closest('tr');
    timer = setTimeout(() => {
        calculateRowTotal(currentRow);
        calculateTableTotal();
    }, 100);
}

// 行级计算
function calculateRowTotal(row) {
    const price = parseFormattedNumber(row.find('input[name="sale_price"]').val());
    const quantity = parseInt(row.find('input[name="quantity"]').val()) || 0;

    if (!isNaN(price) && !isNaN(quantity)) {
        const total = price * quantity;
        row.find('input[name="amount"]').val(total.toFixed(2));
    }
}

// 全局合计计算
function calculateTableTotal() {
    let totalSum = 0;
    $('table tbody tr').each(function() {
        const amount = parseFormattedNumber($(this).find('input[name="amount"]').val());
        totalSum += amount;
    });
    $('input[name="total_amount"]').val(totalSum.toFixed(2));
}

// 初始化计算逻辑
function initCalculations() {
    $('table tbody tr').each(function() {
        calculateRowTotal($(this));
    });
    calculateTableTotal();
}
		`))
	} else if strings.Contains(ctx.Path(), "/ks/info/order_items/edit") {
		// 编辑页面的 JavaScript 代码
		formList.AddJS(template.JS(`
		// 编辑页面的 JavaScript 逻辑
		console.log('This is the edit page JavaScript code.');
    	console.log("现在是在编辑页面！！！");
		// 获取编辑时的字段值
		//let productId = $('select[name="product_id"]').val();
		const quantity = $('input[name="quantity"]').val();
		const pricestr = $('input[name="sale_price"]').val();
		const price = parseFloat(pricestr.replace(/,/g, ""));
		//console.log("编辑页字段值:", productId, quantity);
		function calculateTotal() {
		   let pricestr = $('input[name="sale_price"]').val();
		   let price = parseFloat(pricestr.replace(/,/g, ""));
		   let quantity = $('input[name="quantity"]').val();
	
		   if (price && quantity) {
			   let total = parseFloat(price) * parseInt(quantity);
			   console.log("quantity:", parseInt(quantity));
			   console.log("price:", price);
			   console.log("total:", total);
			   //console.log("编辑页字段值:", productId, quantity);
			   $('input[name="amount"]').val(total.toFixed(2));
		   }
		}
		calculateTotal();
		// 监听所有可能影响合计值的事件
		$(document).on('change input', 'input[name="quantity"], input[name="sale_price"]', function () {
		   // 延迟执行以避免频繁计算
		   setTimeout(calculateTotal, 100);
		});
`))
	}

	return orderItems
}
