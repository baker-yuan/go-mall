
# 1、商品 li_goods
字段名 | 描述 | 数据类型
--- | --- | ---
id | ID | bigint
brand_id | 品牌ID | varchar(255)
template_id | 运费模板ID | varchar(255)
goods_name | 商品名称 | varchar(255)
sn | 商品编号 | varchar(30)
goods_type | 商品类型 | varchar(255)
big | 大图路径 | varchar(255)
small | 小图路径 | varchar(255)
original | 原图路径 | varchar(255)
thumbnail | 缩略图路径 | varchar(255)
intro | 商品详情 | text
mobile_intro | 商品移动端详情 | text
goods_video | 商品视频 | varchar(255)
cost | 成本价格 | decimal(10, 2)
price | 商品价格 | decimal(10, 2)
goods_unit | 计量单位 | varchar(255)
weight | 重量 | decimal(10, 2)
quantity | 库存 | int
params | 商品参数 | text
grade | 商品好评率 | decimal(10, 2)
selling_point | 卖点 | varchar(255)
sales_model | 销售模式 | varchar(255)
recommend | 是否为推荐商品 | bit(1)
self_operated | 是否自营 | bit(1)
category_path | 分类路径 | varchar(255)
store_id | 店铺ID | varchar(255)
store_name | 店铺名称 | varchar(255)
shop_category_path | 店铺分类 | varchar(255)
store_category_path | 店铺分类路径 | varchar(255)
buy_count | 购买数量 | int
comment_num | 评论数量 | int
market_enable | 上架状态 | varchar(255)
auth_flag | 审核状态 | varchar(255)
auth_message | 审核信息 | varchar(255)
under_message | 下架原因 | varchar(255)
create_by | 创建者 | varchar(255)
create_time | 创建时间 | datetime(6)
update_by | 更新者 | varchar(255)
update_time | 更新时间 | datetime(6)
delete_flag | 删除标志 true/false 删除/未删除 | bit(1)

# 2、商品sku li_goods_sku
字段名 | 描述 | 数据类型
--- | --- | ---
id | ID | bigint
brand_id | 品牌ID | varchar(255)
template_id | 运费模板id | varchar(255)
goods_id | 商品ID | varchar(255)
goods_name | 商品名称 | varchar(255)
goods_type | 商品类型 | varchar(255)
big | 大图路径 | varchar(255)
small | 小图路径 | varchar(255)
thumbnail | 缩略图路径 | varchar(255)
original | 原图路径 | varchar(255)
intro | 商品详情 | text
mobile_intro | 移动端商品详情 | text
goods_video | 商品视频 | varchar(255)
price | 商品价格 | decimal(10, 2)
cost | 成本价格 | decimal(10, 2)
weight | 重量 | decimal(10, 2)
goods_unit | 计量单位 | varchar(255)
selling_point | 卖点 | varchar(255)
sn | 商品编号 | varchar(30)
specs | 规格信息json | text
grade | 商品好评率 | decimal(10, 2)
simple_specs | 规格信息 | varchar(255)
buy_count | 购买数量 | int
category_path | 分类路径 | varchar(255)
freight_payer | 运费承担者 | varchar(255)
freight_template_id | 配送模版ID | varchar(255)
promotion_price | 促销价 | decimal(10, 2)
quantity | 库存 | int
sales_model | 销售模式 | varchar(255)
recommend | 是否为推荐商品 | bit(1)
promotion_flag | 是否是促销商品 | bit(1)
self_operated | 是否自营 | bit(1)
store_id | 店铺ID | varchar(255)
store_name | 店铺名称 | varchar(255)
store_category_path | 店铺分类路径 | varchar(255)
view_count | 浏览数量 | int
comment_num | 评价数量 | int
market_enable | 上架状态 | varchar(255)
auth_message | 审核信息 | varchar(255)
under_message | 下架原因 | varchar(255)
auth_flag | 审核状态 | varchar(255)
create_by | 创建者 | varchar(255)
create_time | 创建时间 | datetime(6)
update_by | 更新者 | varchar(255)
update_time | 更新时间 | datetime(6)
delete_flag | 删除标志 true/false 删除/未删除 | bit(1)

# 3、商品参数 li_goods_params
字段名 | 描述 | 数据类型
--- | --- | ---
id | ID | bigint
goods_id | 商品ID | varchar(255)
param_id | 参数ID | varchar(255)
param_name | 参数名字 | varchar(255)
param_value | 参数值 | varchar(100)
is_index | 是否索引 | int
create_by | 创建者 | varchar(255)
create_time | 创建时间 | datetime(6)
update_by | 更新者 | varchar(255)
update_time | 更新时间 | datetime(6)
delete_flag | 删除标志 true/false 删除/未删除 | bit(1)




# 4、商品相册 li_goods_gallery
字段名 | 描述 | 数据类型
--- | --- | ---
id | ID | bigint
goods_id | 商品ID | varchar(255)
is_default | 是否是默认图片 | int
original | 原图路径 | varchar(255)
small | 小图路径 | varchar(255)
sort | 排序 | int
thumbnail | 缩略图路径 | varchar(255)
create_by | 创建者 | varchar(255)

# 5、会员商品收藏 li_goods_collection
字段名 | 描述 | 数据类型
--- | --- | ---
id | ID | bigint
member_id | 会员id | varchar(255)
sku_id | 商品id | varchar(255)
create_time | 创建时间 | datetime(6)

# 6、商品计量单位 li_goods_unit
字段名 | 描述 | 数据类型
--- | --- | ---
id | ID | bigint
delete_flag | 计量单位名称 | bit(1)
create_by | 创建者 | varchar(255)
create_time | 创建时间 | datetime(6)
update_by | 更新者 | varchar(255)
update_time | 更新时间 | datetime(6)
name | 删除标志 true/false 删除/未删除 | varchar(255)

# 7、商品关键字 li_goods_words
字段名 | 描述 | 数据类型
--- | --- | ---
id | ID | bigint
abbreviate | 缩写 | varchar(255)
sort | 排序 | int
type | 类型 | varchar(255)
whole_spell | 全拼 | varchar(255)
words | 单词 | varchar(255)
create_by | 创建者 | varchar(255)
create_time | 创建时间 | datetime(6)
update_by | 更新者 | varchar(255)
update_time | 更新时间 | datetime(6)
delete_flag | 删除标志 | bit(1)

# 2、页面
![02-商品.png](./img/02-商品.png)

# 3、数据
