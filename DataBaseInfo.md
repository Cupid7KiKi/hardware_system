# 1.0

以下是针对您列出的六大核心功能模块设计的MySQL表结构，精简优化后的版本（仅保留必要核心表）：

---

### **1. 产品管理模块**
#### **产品分类表 (product_categories)**
```sql
CREATE TABLE product_categories (
  category_id INT AUTO_INCREMENT PRIMARY KEY,
  category_name VARCHAR(50) NOT NULL UNIQUE,  -- 分类名称（如：螺丝/工具）
  parent_id INT DEFAULT NULL,                 -- 支持多级分类
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### **产品表 (products)**
```sql
CREATE TABLE products (
  product_id INT AUTO_INCREMENT PRIMARY KEY,
  product_code VARCHAR(20) NOT NULL UNIQUE,    -- 产品编码（如：M6-304-50）
  product_name VARCHAR(100) NOT NULL,          -- 产品名称
  category_id INT NOT NULL,                    -- 所属分类
  spec TEXT,                                   -- 规格（如：材质/尺寸）
  unit VARCHAR(10) NOT NULL,                   -- 单位（件/公斤/米）
  purchase_price DECIMAL(10,2),                -- 采购价
  sale_price DECIMAL(10,2),                    -- 销售价
  is_active BOOLEAN DEFAULT TRUE,              -- 是否启用
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES product_categories(category_id)
);
```

---

### **2. 库存管理模块**
#### **仓库表 (warehouses)**
```sql
CREATE TABLE warehouses (
  warehouse_id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE,            -- 仓库名称
  location VARCHAR(100),                       -- 仓库位置
  capacity DECIMAL(12,2)                       -- 容量（立方米）
);
```

#### **库存表 (inventory)**
```sql
CREATE TABLE inventory (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_id INT NOT NULL,
  warehouse_id INT NOT NULL,
  quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),  -- 当前数量
  min_stock INT DEFAULT 10,                     -- 最小库存预警值
  last_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (product_id) REFERENCES products(product_id),
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(warehouse_id)
);
```

---

### **3. 订单管理模块**
#### **订单表 (orders)**
```sql
CREATE TABLE orders (
  order_id VARCHAR(20) PRIMARY KEY,            -- 订单编号（如DD20231101-001）
  customer_id INT NOT NULL,
  total_amount DECIMAL(12,2) NOT NULL,          -- 订单总金额
  status ENUM('pending','approved','shipped','completed','canceled') DEFAULT 'pending',
  operator VARCHAR(50),                        -- 操作员
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
```

#### **订单明细表 (order_items)**
```sql
CREATE TABLE order_items (
  order_id VARCHAR(20) NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL CHECK (quantity > 0),
  unit_price DECIMAL(10,2) NOT NULL,            -- 成交单价
  PRIMARY KEY (order_id, product_id),
  FOREIGN KEY (order_id) REFERENCES orders(order_id),
  FOREIGN KEY (product_id) REFERENCES products(product_id)
);
```

---

### **4. 客户关系管理（CRM）**
#### **客户表 (customers)**
```sql
CREATE TABLE customers (
  customer_id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,                  -- 客户名称
  contact VARCHAR(50),                         -- 联系人
  phone VARCHAR(20),
  address TEXT,
  credit_rating TINYINT                        -- 信用评级（1-5）
);
```

#### **客户联系记录表 (customer_contacts)**
```sql
CREATE TABLE customer_contacts (
  contact_id INT AUTO_INCREMENT PRIMARY KEY,
  customer_id INT NOT NULL,
  contact_time DATETIME NOT NULL,              -- 联系时间
  content TEXT NOT NULL,                       -- 沟通内容
  next_followup DATE,                          -- 下次跟进日期
  FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
```

---

### **5. 财务管理模块**
#### **财务流水表 (financial_records)**
```sql
CREATE TABLE financial_records (
  record_id INT AUTO_INCREMENT PRIMARY KEY,
  type ENUM('income','expense'),               -- 收支类型
  amount DECIMAL(12,2) NOT NULL,               
  payment_method VARCHAR(20),                  -- 支付方式（现金/转账）
  related_order VARCHAR(20),                   -- 关联订单（可选）
  description TEXT,                            -- 款项说明
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (related_order) REFERENCES orders(order_id)
);
```

---

### **6. 报表分析模块**
```sql
-- 无需单独建表，通过视图或查询实现：
-- 1. 销售统计视图（按产品/客户/时间）
-- 2. 库存周转率视图
-- 3. 客户采购排行视图
-- 4. 财务收支汇总视图
```

---

### **关键索引优化**
```sql
-- 订单查询加速
CREATE INDEX idx_order_date ON orders(created_at);
CREATE INDEX idx_customer_orders ON orders(customer_id);

-- 库存预警查询
CREATE INDEX idx_low_stock ON inventory(quantity, min_stock);

-- 客户联系记录
CREATE INDEX idx_customer_followup ON customer_contacts(customer_id, next_followup);
```

---

### **设计特点**
1. **高度聚合**：将订单状态管理、库存预警、客户跟进等核心业务逻辑直接通过字段实现  
2. **轻量级扩展**：通过`parent_id`实现无限级分类，通过`related_order`关联财务与订单  
3. **业务闭环**：从订单创建→发货→财务流水形成完整业务链  
4. **分析友好**：所有时间字段均含时间戳，便于生成时间维度报表

----

# 2.0

以下是修改后的SQL表结构，将所有表的主键字段统一命名为`id`，并调整了相关外键引用：

```sql
-- 产品分类表
CREATE TABLE product_categories (
  id INT AUTO_INCREMENT PRIMARY KEY,
  category_name VARCHAR(50) NOT NULL UNIQUE,
  parent_id INT DEFAULT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (parent_id) REFERENCES product_categories(id)
);

-- 产品表
CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_code VARCHAR(20) NOT NULL UNIQUE,
  product_name VARCHAR(100) NOT NULL,
  category_id INT NOT NULL,
  spec TEXT,
  unit VARCHAR(10) NOT NULL,
  purchase_price DECIMAL(10,2),
  sale_price DECIMAL(10,2),
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES product_categories(id)
);

-- 仓库表
CREATE TABLE warehouses (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE,
  location VARCHAR(100),
  capacity DECIMAL(12,2)
);

-- 库存表
CREATE TABLE inventory (
  product_id INT NOT NULL,
  warehouse_id INT NOT NULL,
  quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
  min_stock INT DEFAULT 10,
  last_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (product_id, warehouse_id),
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

-- 客户表
CREATE TABLE customers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  contact VARCHAR(50),
  phone VARCHAR(20),
  address TEXT,
  credit_rating TINYINT
);

-- 客户联系记录表
CREATE TABLE customer_contacts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  customer_id INT NOT NULL,
  contact_time DATETIME NOT NULL,
  content TEXT NOT NULL,
  next_followup DATE,
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- 订单表
CREATE TABLE orders (
  id VARCHAR(20) PRIMARY KEY,
  customer_id INT NOT NULL,
  total_amount DECIMAL(12,2) NOT NULL,
  status ENUM('pending','approved','shipped','completed','canceled') DEFAULT 'pending',
  operator VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- 订单明细表
CREATE TABLE order_items (
  order_id VARCHAR(20) NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL CHECK (quantity > 0),
  unit_price DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (order_id, product_id),
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);

-- 财务流水表
CREATE TABLE financial_records (
  id INT AUTO_INCREMENT PRIMARY KEY,
  type ENUM('income','expense'),
  amount DECIMAL(12,2) NOT NULL,
  payment_method VARCHAR(20),
  related_order VARCHAR(20),
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (related_order) REFERENCES orders(id)
);

-- 索引优化
CREATE INDEX idx_order_date ON orders(created_at);
CREATE INDEX idx_customer_orders ON orders(customer_id);
CREATE INDEX idx_low_stock ON inventory(quantity, min_stock);
CREATE INDEX idx_customer_followup ON customer_contacts(customer_id, next_followup);
```

### **修改说明：**
1. **主键统一命名**：
   - 所有表的主键字段名均改为 `id`（原 `category_id`、`product_id`、`warehouse_id` 等）。
   - 订单表的主键 `order_id` 更名为 `id`（仍保留 `VARCHAR(20)` 类型）。

2. **外键调整**：
   - 外键字段名称保持不变（如 `category_id`、`warehouse_id` 等），但引用目标改为被引用表的 `id` 字段。
   - 示例：`products.category_id` 引用 `product_categories.id`，而非原 `product_categories.category_id`。

3. **自增主键处理**：
   - 原自增主键（如 `product_id INT AUTO_INCREMENT`）统一改为 `id INT AUTO_INCREMENT`。

4. **组合主键保留**：
   - `inventory` 和 `order_items` 表的组合主键保持不变，但外键引用目标调整为对应表的 `id`。

5. **索引兼容性**：
   - 所有索引基于现有字段名（如 `customer_id`、`product_id`）无需调整，仅需确保外键引用正确。

### **注意事项：**
- **业务主键兼容性**：订单表的 `id` 字段（原 `order_id`）需确保生成规则与应用逻辑一致（如 `DD20231101-001`）。
- **外键字段名一致性**：外键字段名未强制与主键名一致（如 `customer_id` 仍为 `customer_id`），但实际关联的是被引用表的 `id`。
- **数据迁移**：若已有数据需迁移，需同步更新所有外键值指向新主键 `id`。