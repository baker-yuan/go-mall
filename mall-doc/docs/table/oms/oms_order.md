
# 1、表结构
## 1.1、订单表
```sql
CREATE TABLE `oms_order`
(
    `id`                      bigint         NOT NULL AUTO_INCREMENT COMMENT '订单id',
    `member_id`               bigint         NOT NULL DEFAULT '0' COMMENT '会员id',
    `order_sn`                varchar(64)    NOT NULL DEFAULT '' COMMENT '订单编号',
    `note`                    varchar(500)   NOT NULL DEFAULT '' COMMENT '订单备注',
    `coupon_id`               bigint         NOT NULL DEFAULT '0' COMMENT '优惠券id',
    `promotion_info`          varchar(100)   NOT NULL DEFAULT '' COMMENT '活动信息',
    `use_integration`         int            NOT NULL DEFAULT '0' COMMENT '下单时使用的积分',
    `integration`             int            NOT NULL DEFAULT '0' COMMENT '可以获得的积分',
    `growth`                  int            NOT NULL DEFAULT '0' COMMENT '可以活动的成长值',
    `pay_type`                tinyint        NOT NULL DEFAULT '0' COMMENT '支付方式：0->未支付；1->支付宝；2->微信',
    `source_type`             tinyint        NOT NULL DEFAULT '0' COMMENT '订单来源：0->PC订单；1->app订单',
    `order_type`              tinyint        NOT NULL DEFAULT '0' COMMENT '订单类型：0->正常订单；1->秒杀订单',
    `delivery_company`        varchar(64)    NOT NULL DEFAULT '' COMMENT '物流公司(配送方式)',
    `delivery_sn`             varchar(64)    NOT NULL DEFAULT '' COMMENT '物流单号',
    `receive_time`            int            NOT NULL DEFAULT '0' COMMENT '确认收货时间',
    `auto_confirm_day`        int            NOT NULL DEFAULT '0' COMMENT '自动确认时间（天）',
    `receiver_name`           varchar(100)   NOT NULL DEFAULT '' COMMENT '收货人姓名',
    `receiver_phone`          varchar(32)    NOT NULL DEFAULT '' COMMENT '收货人电话',
    `receiver_post_code`      varchar(32)    NOT NULL DEFAULT '' COMMENT '收货人邮编',
    `receiver_province`       varchar(32)    NOT NULL DEFAULT '' COMMENT '省份/直辖市',
    `receiver_city`           varchar(32)    NOT NULL DEFAULT '' COMMENT '城市',
    `receiver_region`         varchar(32)    NOT NULL DEFAULT '' COMMENT '区',
    `receiver_detail_address` varchar(200)   NOT NULL DEFAULT '' COMMENT '详细地址',
    `total_amount`            decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '订单总金额',
    `freight_amount`          decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '运费金额',
    `promotion_amount`        decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '促销优化金额（促销价、满减、阶梯价）',
    `coupon_amount`           decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '优惠券抵扣金额',
    `integration_amount`      decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '积分抵扣金额',
    `pay_amount`              decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '应付金额（实际支付金额）',
    `discount_amount`         decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '管理员后台调整订单使用的折扣金额',
    `bill_type`               tinyint        NOT NULL DEFAULT '0' COMMENT '发票类型：0->不开发票；1->电子发票；2->纸质发票',
    `bill_header`             varchar(200)   NOT NULL DEFAULT '' COMMENT '发票抬头',
    `bill_content`            varchar(200)   NOT NULL DEFAULT '' COMMENT '发票内容',
    `bill_receiver_phone`     varchar(32)    NOT NULL DEFAULT '' COMMENT '收票人电话',
    `bill_receiver_email`     varchar(64)    NOT NULL DEFAULT '' COMMENT '收票人邮箱',
    `status`                  tinyint        NOT NULL DEFAULT '0' COMMENT '订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单',
    `confirm_status`          tinyint        NOT NULL DEFAULT '0' COMMENT '确认收货状态：0->未确认；1->已确认',
    `delete_status`           tinyint        NOT NULL DEFAULT '0' COMMENT '删除状态：0->未删除；1->已删除',
    `payment_time`            int            NOT NULL DEFAULT '0' COMMENT '支付时间',
    `delivery_time`           int            NOT NULL DEFAULT '0' COMMENT '发货时间',
    `comment_time`            int            NOT NULL DEFAULT '0' COMMENT '评价时间',
    `create_time`             int            NOT NULL DEFAULT '0' COMMENT '提交时间',
    `modify_time`             int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    `member_username`         varchar(64)    NOT NULL DEFAULT '' COMMENT '用户帐号',
    `created_at`              int            NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`              int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT ='订单表';
```

## 1.2、订单商品信息表
```sql
CREATE TABLE `oms_order_item`
(
    `id`                  bigint         NOT NULL AUTO_INCREMENT,
    `order_id`            bigint         NOT NULL DEFAULT '0' COMMENT '订单id',
    `order_sn`            varchar(64)    NOT NULL DEFAULT '' COMMENT '订单编号',
    `product_id`          bigint         NOT NULL DEFAULT '0' COMMENT '商品id',
    `product_category_id` bigint         NOT NULL DEFAULT '0' COMMENT '商品分类id',
    `product_pic`         varchar(500)   NOT NULL DEFAULT '' COMMENT '商品图片',
    `product_name`        varchar(200)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `product_brand`       varchar(200)   NOT NULL DEFAULT '' COMMENT '商品品牌',
    `product_sn`          varchar(64)    NOT NULL DEFAULT '' COMMENT '商品条码',
    `product_attr`        varchar(500)   NOT NULL DEFAULT '' COMMENT '商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]',
    `promotion_name`      varchar(200)   NOT NULL DEFAULT '' COMMENT '商品促销名称',
    `product_sku_id`      bigint         NOT NULL DEFAULT '0' COMMENT '商品sku编号',
    `product_sku_code`    varchar(50)    NOT NULL DEFAULT '' COMMENT '商品sku条码',
    `product_quantity`    int            NOT NULL DEFAULT '0' COMMENT '购买数量',
    `product_price`       decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '销售价格',
    `promotion_amount`    decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品促销分解金额',
    `coupon_amount`       decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '优惠券优惠分解金额',
    `integration_amount`  decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '积分优惠分解金额',
    `real_amount`         decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '该商品经过优惠后的分解金额',
    `gift_integration`    int            NOT NULL DEFAULT '0' COMMENT '商品赠送积分',
    `gift_growth`         int            NOT NULL DEFAULT '0' COMMENT '商品赠送成长值',
    `created_at`          int            NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`          int            NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT ='订单商品信息表';
```