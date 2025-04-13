// utils/orderItemsHandlers.js
const OrderItemsHandlers = (function() {
    // 私有方法：创建页逻辑
    const _initCreatePage = function() {
        console.log("初始化创建页逻辑");

        // 行级计算函数
        const calculateRowTotal = (row) => {
            const price = parseFloat(row.find('input[name="sale_price"]').inputmask('unmaskedvalue') || 0);
            const quantity = parseInt(row.find('input[name="quantity"]').val() || 0);
            const total = price * quantity;
            row.find('input[name="amount"]').val(total.toFixed(2));
            _updateGrandTotal();
        };

        // 更新总计
        const _updateGrandTotal = () => {
            let sum = 0;
            $('table tbody tr').each(function() {
                sum += parseFloat($(this).find('input[name="amount"]').inputmask('unmaskedvalue') || 0);
            });
            $('input[name="unit_price"]').val(sum.toFixed(2));
        };

        // 产品选择事件（使用命名空间隔离）
        $(document).off('change.createPage').on('change.createPage', 'select[name="product_id"]', function() {
            const row = $(this).closest('tr');
            const productId = $(this).val();

            if (!productId) {
                row.find('input[name="sale_price"]').val('0.00');
                calculateRowTotal(row);
                return;
            }

            $.ajax({
                url: '/ks/choose/product_price',
                type: 'POST',
                data: { product_id: productId },
                success: (response) => {
                    if (response.success) {
                        row.find('input[name="sale_price"]').val(response.data);
                        calculateRowTotal(row);
                    }
                }
            });
        });

        // 数量输入事件
        $(document).off('input.createPage').on('input.createPage', 'input[name="quantity"]', function() {
            calculateRowTotal($(this).closest('tr'));
        });

        // 初始化输入掩码
        $('.sale_price, .amount').inputmask({
            alias: "currency",
            radixPoint: ".",
            prefix: "",
            removeMaskOnSubmit: true
        });
    };

    // 私有方法：编辑页逻辑
    const _initEditPage = function() {
        console.log("初始化编辑页逻辑");

        const calculateTotal = () => {
            const price = parseFloat($('input[name="sale_price"]').inputmask('unmaskedvalue') || 0);
            const quantity = parseInt($('input[name="quantity"]').val() || 0);
            $('input[name="amount"]').val((price * quantity).toFixed(2));
        };

        // 产品选择事件（使用不同命名空间）
        $('select[name="product_id"]').off('change.editPage').on('change.editPage', function() {
            $.ajax({
                url: '/ks/choose/product_price',
                data: { product_id: $(this).val() },
                success: (response) => {
                    if (response.success) {
                        $('input[name="sale_price"]').val(response.data);
                        calculateTotal();
                    }
                }
            });
        });

        // 数量输入事件
        $('input[name="quantity"]').off('input.editPage').on('input.editPage', calculateTotal);
    };

    // 公共接口
    return {
        initCreatePage: _initCreatePage,
        initEditPage: _initEditPage
    };
})();

// 自动执行页面初始化
(function() {
    // 清理旧事件防止冲突
    $(document).off('change.createPage input.createPage change.editPage input.editPage');

    // 根据URL路径初始化对应页面
    if (window.location.pathname.includes("/new")) {
        OrderItemsHandlers.initCreatePage();
    }
    else if (window.location.pathname.includes("/edit")) {
        OrderItemsHandlers.initEditPage();
    }
})();