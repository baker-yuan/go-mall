# 1、表结构
## 1.1、商品表
```sql
CREATE TABLE `pms_product`
(
    `id`                            bigint         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `product_category_id`           bigint         NOT NULL DEFAULT '0' COMMENT '商品分类id',
    `name`                          varchar(200)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `sub_title`                     varchar(255)   NOT NULL DEFAULT '' COMMENT '副标题',
    `brand_id`                      bigint         NOT NULL DEFAULT '0' COMMENT '品牌id',
    `description`                   text           NOT NULL COMMENT '商品描述',
    `product_sn`                    varchar(64)    NOT NULL DEFAULT '' COMMENT '货号',
    `price`                         decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `original_price`                decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '市场价',
    `stock`                         int            NOT NULL DEFAULT '0' COMMENT '库存',
    `unit`                          varchar(16)    NOT NULL DEFAULT '' COMMENT '单位',
    `weight`                        decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品重量，默认为克',
    `sort`                          int            NOT NULL DEFAULT '0' COMMENT '排序',
    `gift_point`                    int            NOT NULL DEFAULT '0' COMMENT '赠送的积分',
    `gift_growth`                   int            NOT NULL DEFAULT '0' COMMENT '赠送的成长值',
    `use_point_limit`               int            NOT NULL DEFAULT '0' COMMENT '限制使用的积分数',
    `preview_status`                tinyint        NOT NULL DEFAULT '0' COMMENT '是否为预告商品：0->不是；1->是',
    `publish_status`                tinyint        NOT NULL DEFAULT '0' COMMENT '上架状态：0->下架；1->上架',
    `new_status`                    tinyint        NOT NULL DEFAULT '0' COMMENT '新品状态:0->不是新品；1->新品',
    `recommand_status`              tinyint        NOT NULL DEFAULT '0' COMMENT '推荐状态；0->不推荐；1->推荐',
    `service_ids`                   varchar(64)    NOT NULL DEFAULT '' COMMENT '以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮',
    `detail_title`                  varchar(255)   NOT NULL DEFAULT '' COMMENT '详情标题',
    `detail_desc`                   text           NOT NULL COMMENT '详情描述',
    `keywords`                      varchar(255)   NOT NULL DEFAULT '' COMMENT '关键字',
    `note`                          varchar(255)   NOT NULL DEFAULT '' COMMENT '备注',
    `promotion_type`                tinyint        NOT NULL DEFAULT '0' COMMENT '促销类型：0->没有促销使用原价',
    `promotion_price`               decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '促销价格',
    `promotion_start_time`          int            NOT NULL DEFAULT '0' COMMENT '促销开始时间',
    `promotion_end_time`            int            NOT NULL DEFAULT '0' COMMENT '促销结束时间',
    `product_attribute_category_id` bigint         NOT NULL DEFAULT '0' COMMENT '品牌属性分类id',
    `pic`                           varchar(255)   NOT NULL DEFAULT '' COMMENT '图片',
    `album_pics`                    varchar(255)   NOT NULL DEFAULT '' COMMENT '画册图片，连产品图片限制为5张，以逗号分割',
    `detail_html`                   text           NOT NULL COMMENT '电脑端详情',
    `detail_mobile_html`            text           NOT NULL COMMENT '移动端详情',
    `verify_status`                 tinyint        NOT NULL DEFAULT '0' COMMENT '审核状态：0->未审核；1->审核通过',
    `delete_status`                 tinyint        NOT NULL DEFAULT '0' COMMENT '删除状态：0->未删除；1->已删除',
    `feight_template_id`            bigint         NOT NULL DEFAULT '0' COMMENT '运费模版id',
    `sale`                          int            NOT NULL DEFAULT '0' COMMENT '销量',
    `low_stock`                     int            NOT NULL DEFAULT '0' COMMENT '库存预警值',
    `promotion_per_limit`           int            NOT NULL DEFAULT '0' COMMENT '活动限购数量',
    `brand_name`                    varchar(255)            DEFAULT NULL COMMENT '品牌名称',
    `product_category_name`         varchar(255)            DEFAULT NULL COMMENT '商品分类名称',
    `created_at`                    int            NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`                    int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='商品表';
```

## 1.2、商品SKU表
```sql
CREATE TABLE `pms_sku_stock`
(
    `id`              bigint         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `product_id`      bigint         NOT NULL DEFAULT '0' COMMENT '产品ID',
    `sku_code`        varchar(64)    NOT NULL DEFAULT '' COMMENT 'sku编码',
    `price`           decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `stock`           int            NOT NULL DEFAULT '0' COMMENT '库存',
    `low_stock`       int            NOT NULL DEFAULT '0' COMMENT '预警库存',
    `pic`             varchar(255)   NOT NULL DEFAULT '' COMMENT '展示图片',
    `sale`            int            NOT NULL DEFAULT '0' COMMENT '销量',
    `promotion_price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '单品促销价格',
    `lock_stock`      int            NOT NULL DEFAULT '0' COMMENT '锁定库存',
    `sp_data`         varchar(500)   NOT NULL DEFAULT '' COMMENT '商品销售属性，json格式',
    `created_at`      int            NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`      int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='商品SKU表';
```

## 1.3、商品阶梯价格表
```sql
CREATE TABLE `pms_product_ladder`
(
    `id`         bigint         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `product_id` bigint         NOT NULL DEFAULT '0' COMMENT '商品id',
    `count`      int            NOT NULL DEFAULT '0' COMMENT '满足的商品数量',
    `discount`   decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
    `price`      decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '折后价格',
    `created_at` int            NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='商品阶梯价格表';
```

## 1.4、商品满减表
```sql
CREATE TABLE `pms_product_full_reduction`
(
    `id`           bigint         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `product_id`   bigint         NOT NULL DEFAULT '0' COMMENT '商品id',
    `full_price`   decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品满足金额',
    `reduce_price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品减少金额',
    `created_at`   int            NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`   int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='商品满减表';
```

## 1.5、商品会员价格表
```sql
CREATE TABLE `pms_member_price`
(
    `id`                bigint         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `product_id`        bigint         NOT NULL DEFAULT '0' COMMENT '商品id',
    `member_level_id`   bigint         NOT NULL DEFAULT '0' COMMENT '会员等级id',
    `member_price`      decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '会员价格',
    `member_level_name` varchar(100)   NOT NULL DEFAULT '' COMMENT '会员等级名称',
    `created_at`        int            NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`        int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='商品会员价格表';
```