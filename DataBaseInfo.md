

# 数据库设计

### **1. 产品管理模块**

#### **产品分类表 (product_categories)**
```sql
CREATE TABLE product_categories (
  id INT AUTO_INCREMENT PRIMARY KEY,
  category_name VARCHAR(50) NOT NULL UNIQUE,
  parent_id INT DEFAULT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### **产品表 (products)**
```sql
CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_code VARCHAR(20) NOT NULL UNIQUE,
  product_name VARCHAR(100) NOT NULL,
  category_id INT,
  spec TEXT,
  unit VARCHAR(10) NOT NULL,
  brand VARCHAR(255),
  purchase_price DECIMAL(10,2),
  sale_price DECIMAL(10,2),
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES product_categories(id)
);
```

---

### **2. 库存管理模块**
#### **仓库表 (warehouses)**
```sql
CREATE TABLE warehouses (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE,
  location VARCHAR(100),
  capacity DECIMAL(12,2)
);
```

#### **库存表 (inventory)**
```sql
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
```

---

### **3. 订单管理模块**
#### **订单表 (orders)**
```sql
CREATE TABLE orders (
  id INT AUTO_INCREMENT PRIMARY KEY,
  customer_id INT NOT NULL,
  contact_id INT NOT NULL,
  total_amount DECIMAL(12,2) NOT NULL,
  status ENUM('pending','approved','shipped','completed','canceled') DEFAULT 'pending',
  operator VARCHAR(50),
  remarks TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers(id),
  FOREIGN KEY (contact_id) REFERENCES companies_contacts(id)
);
```

#### **订单明细表 (order_items)**
```sql
CREATE TABLE order_items (
  id INT AUTO_INCREMENT PRIMARY KEY,
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL CHECK (quantity > 0),
  unit_price DECIMAL(10,2) NOT NULL,
  amount DECIMAL(10,2) NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES products(id)
);
```

---

### **4. 客户关系管理（CRM）**
#### **客户表 (customers)**(简化为只记录公司级信息)
```sql
CREATE TABLE customers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  company_id INT NOT NULL,
  contact_id INT,
  credit_rating TINYINT,
  FOREIGN KEY (company_id) REFERENCES customers_companies(id),
  FOREIGN KEY (contact_id) REFERENCES customers_contacts(id)
);
```

#### 客户公司信息表

```sql
CREATE TABLE customers_companies (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  address TEXT
);
```

#### 客户公司联系人表(包含所属公司)

```sql
CREATE TABLE companies_contacts (
  id int NOT NULL AUTO_INCREMENT,
  company_id int NOT NULL,
  name varchar(50) NOT NULL,
  phone varchar(20) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY `company_id` (company_id),
  CONSTRAINT `companies_contacts_ibfk_1` FOREIGN KEY (company_id) REFERENCES customers_companies(id)
)
```

#### **客户联系记录表 (customer_contacts)**

```sql
CREATE TABLE customer_contacts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  customer_id INT NOT NULL,
  contact_time DATETIME NOT NULL,
  content TEXT NOT NULL,
  next_followup DATE,
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);
```

---

### **5. 财务管理模块**
#### **财务流水表 (financial_records)**
```sql
CREATE TABLE financial_records (
  id INT AUTO_INCREMENT PRIMARY KEY,
  type ENUM('income','expense'),
  amount DECIMAL(12,2) NOT NULL,
  payment_method VARCHAR(20),
  related_order INT,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (related_order) REFERENCES orders(id)
);
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
2. **轻量级扩展**：通过`related_order`关联财务与订单  
3. **业务闭环**：从订单创建→发货→财务流水形成完整业务链  
4. **分析友好**：所有时间字段均含时间戳，便于生成时间维度报表
