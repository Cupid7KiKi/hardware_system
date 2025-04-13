/* ========================================================================
 * bootstrap-spin - v1.0
 * https://github.com/wpic/bootstrap-spin
 * ========================================================================
 * Copyright 2014 WPIC, Hamed Abdollahpour
 *
 * ========================================================================
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ========================================================================
 */

/* ========================================================================
 * 修改版数字输入组件，支持实时打印和更新
 * ========================================================================
 */

/* ========================================================================
 * 增强版数字输入组件，带控制台打印
 * ========================================================================
 */

(function ($) {

    $.fn.bootstrapNumber = function (options) {

        var settings = $.extend({
            upClass: 'default',
            downClass: 'default',
            center: true,
            debug: true // 默认开启调试模式
        }, options);

        return this.each(function () {
            var self = $(this);
            var clone = self.clone(true, true); // 深度克隆保留事件

            var min = parseInt(self.attr('min')) || 1;
            var max = parseInt(self.attr('max')) || 999999;

            function logValue(value) {
                if (settings.debug) {
                    console.log('[数字输入框] 当前值:', value);
                }
            }

            function setText(n) {
                n = parseInt(n);
                if (isNaN(n)) n = min;
                if (n < min) n = min;
                if (n > max) n = max;

                clone.val(n);
                logValue(n); // 打印值变化
                clone.trigger('input'); // 确保触发所有监听事件
                return true;
            }

            // 创建输入组
            var group = $("<div class='input-group'></div>");

            // 减少按钮
            var down = $("<button type='button'>-</button>")
                .addClass('btn btn-' + settings.downClass)
                .click(function () {
                    console.log('[数字输入框] 点击减少按钮');
                    setText(parseInt(clone.val()) - 1);
                });

            // 增加按钮
            var up = $("<button type='button'>+</button>")
                .addClass('btn btn-' + settings.upClass)
                .click(function () {
                    console.log('[数字输入框] 点击增加按钮');
                    setText(parseInt(clone.val()) + 1);
                });

            // 组装组件
            $("<span class='input-group-btn'></span>").append(down).appendTo(group);
            clone.appendTo(group);
            if (settings.center) {
                clone.css('text-align', 'center');
            }
            $("<span class='input-group-btn'></span>").append(up).appendTo(group);

            // 输入处理
            clone.prop('type', 'text').on('input', function() {
                console.log('[数字输入框] 输入变化:', $(this).val());
                setText($(this).val());
            });

            // 初始值打印
            logValue(clone.val());

            self.replaceWith(group);
        });
    };
}(jQuery));