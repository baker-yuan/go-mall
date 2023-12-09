create database if not exists go_mall_admin default char set utf8mb4;
use go_mall_admin;


truncate table pms_product_category;
insert into go_mall_admin.pms_product_category (id, parent_id, name, level, product_count, product_unit, nav_status, show_status, sort, icon, keywords, description, created_at, updated_at)
values  (1, 0, '服装', 0, 100, '件', 1, 1, 1, '', '服装', '服装分类', 0, 0),
        (2, 0, '手机数码', 0, 100, '件', 1, 1, 1, '', '手机数码', '手机数码', 0, 0),
        (3, 0, '家用电器', 0, 100, '件', 1, 1, 1, '', '家用电器', '家用电器', 0, 0),
        (4, 0, '家具家装', 0, 100, '件', 1, 1, 1, '', '家具家装', '家具家装', 0, 0),
        (5, 0, '汽车用品', 0, 100, '件', 1, 1, 1, '', '汽车用品', '汽车用品', 0, 0),
        (7, 1, '外套', 1, 100, '件', 1, 1, 0, '/mall/images/20190519/5ac4780cN6087feb5.jpg', '外套', '外套', 0, 0),
        (8, 1, 'T恤', 1, 100, '件', 1, 1, 0, '/mall/images/20190519/5ac47ffaN8a7b2e14.png', 'T恤', 'T恤', 0, 0),
        (9, 1, '休闲裤', 1, 100, '件', 1, 1, 0, '/mall/images/20190519/5ac47845N7374a31d.jpg', '休闲裤', '休闲裤', 0, 0),
        (10, 1, '牛仔裤', 1, 100, '件', 1, 1, 0, '/mall/images/20190519/5ac47841Nff658599.jpg', '牛仔裤', '牛仔裤', 0, 0),
        (11, 1, '衬衫', 1, 100, '件', 1, 1, 0, '/mall/images/20190519/5ac48007Nb30b2118.jpg', '衬衫', '衬衫分类', 0, 0),
        (13, 12, '家电子分类1', 1, 1, 'string', 0, 1, 0, 'string', 'string', 'string', 0, 0),
        (14, 12, '家电子分类2', 1, 1, 'string', 0, 1, 0, 'string', 'string', 'string', 0, 0),
        (19, 2, '手机通讯', 1, 0, '件', 0, 1, 0, '/mall/images/20190519/5ac48d27N3f5bb821.jpg', '手机通讯', '手机通讯', 0, 0),
        (29, 1, '男鞋', 1, 0, '', 0, 0, 0, '', '', '', 0, 0),
        (30, 2, '手机配件', 1, 0, '', 0, 1, 0, '/mall/images/20190519/5ac48672N11cf61fe.jpg', '手机配件', '手机配件', 0, 0),
        (31, 2, '摄影摄像', 1, 0, '', 0, 1, 0, '/mall/images/20190519/5a1679f2Nc2f659b6.jpg', '', '', 0, 0),
        (32, 2, '影音娱乐', 1, 0, '', 0, 1, 0, '/mall/images/20190519/5a167859N01d8198b.jpg', '', '', 0, 0),
        (33, 2, '数码配件', 1, 0, '', 0, 1, 0, '/mall/images/20190519/5a1676e9N1ba70a81.jpg', '', '', 0, 0),
        (34, 2, '智能设备', 1, 0, '', 0, 0, 0, '', '', '', 0, 0),
        (35, 3, '电视', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a17f71eN25360979.jpg', '', '', 0, 0),
        (36, 3, '空调', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a17f6f6Ndfe746aa.jpg', '', '', 0, 0),
        (37, 3, '洗衣机', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a17f6eaN9ec936de.jpg', '', '', 0, 0),
        (38, 3, '冰箱', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a17f6c5Ne56d7e26.jpg', '', '', 0, 0),
        (39, 3, '厨卫大电', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a17f369N6a55ff3f.jpg', '', '', 0, 0),
        (40, 3, '厨房小电', 1, 0, '', 0, 0, 0, '', '', '', 0, 0),
        (41, 3, '生活电器', 1, 0, '', 0, 0, 0, '', '', '', 0, 0),
        (42, 3, '个护健康', 1, 0, '', 0, 0, 0, '', '', '', 0, 0),
        (43, 4, '厨房卫浴', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a1eb12cN5ab932bb.jpg', '', '', 0, 0),
        (44, 4, '灯饰照明', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a1eb115Na5705672.jpg', '', '', 0, 0),
        (45, 4, '五金工具', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a28e743Nf6d99998.jpg', '', '', 0, 0),
        (46, 4, '卧室家具', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a1eb096N6326e0bd.jpg', '', '', 0, 0),
        (47, 4, '客厅家具', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a1eb080N38c2e7b7.jpg', '', '', 0, 0),
        (48, 5, '全新整车', 1, 0, '', 0, 1, 0, '/mall/images/20200607/ebe31b9cc535e122.png', '', '', 0, 0),
        (49, 5, '车载电器', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a1fb8d2N53bbd2ba.jpg', '', '', 0, 0),
        (50, 5, '维修保养', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a1fb8baNbe801af8.jpg', '', '', 0, 0),
        (51, 5, '汽车装饰', 1, 0, '', 0, 1, 0, '/mall/images/20200607/5a28ae20N34461e66.jpg', '', '', 0, 0),
        (52, 0, '电脑办公', 0, 0, '件', 1, 1, 1, '', '电脑办公', '电脑办公', 0, 0),
        (53, 52, '平板电脑', 1, 0, '件', 0, 1, 0, '/mall/images/20221028/pad_category_01.jpg', '平板电脑', '平板电脑', 0, 0),
        (54, 52, '笔记本', 1, 0, '件', 0, 1, 0, '/mall/images/20221028/computer_category_01.jpg', '笔记本', '笔记本', 0, 0),
        (55, 52, '硬盘', 1, 0, '', 0, 1, 0, '/mall/images/20221108/ssd_category_01.jpg', '硬盘', '', 0, 0);



truncate table pms_brand;
insert into go_mall_admin.pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count, logo, big_pic, brand_story, created_at, updated_at)
values  (1, '万和', 'W', 0, 1, 1, 100, 100, '/mall/images/20200607/5b07ca8aN4e127d2f.jpg', 'http://img13.360buyimg.com/cms/jfs/t1/121860/35/2430/187800/5ec4e294E22f3ffcc/1e233b65b94ba192.jpg', '万和成立于1993年8月，总部位于广东顺德国家级高新技术开发区内，是国内生产规模最大的燃气具专业制造企业，也是中国燃气具发展战略的首倡者和推动者、中国五金制品协会燃气用具分会第三届理事长单位。', 0, 0),
        (2, '三星', 'S', 100, 1, 1, 100, 100, '/mall/images/20200607/57201b47N7bf15715.jpg', '/mall/images/20221108/sanxing_banner_01.png', '三星集团（英文：SAMSUNG、韩文：삼성）是韩国最大的跨国企业集团，三星集团包括众多的国际下属企业，旗下子公司有：三星电子、三星物产、三星人寿保险等，业务涉及电子、金融、机械、化学等众多领域。', 0, 0),
        (3, '华为', 'H', 100, 1, 1, 100, 100, '/mall/images/20200607/5abf6f26N31658aa2.jpg', '/mall/images/20221108/huawei_banner_01.png', '荣耀品牌成立于2013年,是华为旗下手机双品牌之一。荣耀以“创新、品质、服务”为核心战略,为全球年轻人提供潮酷的全场景智能化体验,打造年轻人向往的先锋文化和潮流生活方式', 0, 0),
        (4, '格力', 'G', 30, 1, 1, 100, 100, '/mall/images/20180607/timg (3).jpg', '', 'Victoria''s Secret的故事', 0, 0),
        (5, '方太', 'F', 20, 1, 1, 100, 100, '/mall/images/20180607/timg (4).jpg', '', 'Victoria''s Secret的故事', 0, 0),
        (6, '小米', 'M', 500, 1, 1, 100, 100, '/mall/images/20200607/5565f5a2N0b8169ae.jpg', '/mall/images/20221108/xiaomi_banner_01.png', '小米公司正式成立于2010年4月，是一家专注于高端智能手机、互联网电视自主研发的创新型科技企业。主要由前谷歌、微软、摩托、金山等知名公司的顶尖人才组建。', 0, 0),
        (21, 'OPPO', 'O', 0, 1, 1, 88, 500, '/mall/images/20180607/timg(6).jpg', '/mall/images/20221108/oppo_banner_01.png', 'OPPO于2008年推出第一款“笑脸手机”，由此开启探索和引领至美科技之旅。今天，OPPO凭借以Find和R系列手机为核心的智能终端产品，以及OPPO+等互联网服务，让全球消费者尽享至美科技。', 0, 0),
        (49, '七匹狼', 'S', 200, 1, 1, 77, 400, '/mall/images/20190525/qipilang.png', '', 'BOOB的故事', 0, 0),
        (50, '海澜之家', 'H', 200, 1, 1, 66, 300, '/mall/images/20200607/5a5c69b9N5d6c5696.jpg', 'http://img10.360buyimg.com/cms/jfs/t1/133148/4/1605/470028/5edaf5ccEd7a687a9/e0a007631361ff75.jpg', '“海澜之家”（英文缩写：HLA）是海澜之家股份有限公司旗下的服装品牌，总部位于中国江苏省无锡市江阴市，主要采用连锁零售的模式，销售男性服装、配饰与相关产品。', 0, 0),
        (51, '苹果', 'A', 200, 1, 1, 55, 200, '/mall/images/20200607/49b30bb0377030d1.jpg', '/mall/images/20221108/apple_banner_01.png', '苹果公司(Apple Inc. )是美国的一家高科技公司。 由史蒂夫·乔布斯、斯蒂夫·沃兹尼亚克和罗·韦恩(Ron Wayne)等人于1976年4月1日创立,并命名为美国苹果电脑公司(Apple Computer Inc. ),2007年1月9日更名为苹果公司,总部位于加利福尼亚州的...', 0, 0),
        (58, 'NIKE', 'N', 0, 1, 0, 33, 100, '/mall/images/20180615/timg (51).jpg', '', 'NIKE的故事', 0, 0),
        (59, '测试品牌', 'C', 0, 0, 0, 0, 0, '/mall/20220609/Snipaste_2022-06-08_14-35-53.png', '/mall/20220609/biji_05.jpg', '12345', 0, 0);


insert into go_mall_admin.pms_product_attribute_category (id, name, attribute_count, param_count, created_at, updated_at)
values  (1, '服装-T恤', 2, 5, 0, 0),
        (2, '服装-裤装', 2, 4, 0, 0),
        (3, '手机数码-手机通讯', 2, 4, 0, 0),
        (4, '配件', 0, 0, 0, 0),
        (5, '居家', 0, 0, 0, 0),
        (6, '洗护', 0, 0, 0, 0),
        (10, '测试分类', 0, 0, 0, 0),
        (11, '服装-鞋帽', 3, 0, 0, 0),
        (12, '家用电器-电视', 2, 4, 0, 0),
        (13, '电脑办公-笔记本', 2, 3, 0, 0),
        (14, '家用电器-厨卫大电', 1, 3, 0, 0),
        (15, '电脑办公-硬盘', 2, 5, 0, 0);


insert into go_mall_admin.pms_product_attribute (id, product_attribute_category_id, name, select_type, input_type, input_list, sort, filter_type, search_type, related_status, hand_add_status, type, created_at, updated_at)
values  (1, 1, '尺寸', 2, 1, 'M,X,XL,2XL,3XL,4XL', 0, 0, 0, 0, 0, 0, 0, 0),
        (7, 1, '颜色', 2, 1, '黑色,红色,白色,粉色', 100, 0, 0, 0, 1, 0, 0, 0),
        (13, 0, '上市年份', 1, 1, '2013年,2014年,2015年,2016年,2017年', 0, 0, 0, 0, 0, 1, 0, 0),
        (14, 0, '上市年份1', 1, 1, '2013年,2014年,2015年,2016年,2017年', 0, 0, 0, 0, 0, 1, 0, 0),
        (15, 0, '上市年份2', 1, 1, '2013年,2014年,2015年,2016年,2017年', 0, 0, 0, 0, 0, 1, 0, 0),
        (16, 0, '上市年份3', 1, 1, '2013年,2014年,2015年,2016年,2017年', 0, 0, 0, 0, 0, 1, 0, 0),
        (17, 0, '上市年份4', 1, 1, '2013年,2014年,2015年,2016年,2017年', 0, 0, 0, 0, 0, 1, 0, 0),
        (18, 0, '上市年份5', 1, 1, '2013年,2014年,2015年,2016年,2017年', 0, 0, 0, 0, 0, 1, 0, 0),
        (19, 0, '适用对象', 1, 1, '青年女性,中年女性', 0, 0, 0, 0, 0, 1, 0, 0),
        (20, 0, '适用对象1', 2, 1, '青年女性,中年女性', 0, 0, 0, 0, 0, 1, 0, 0),
        (21, 0, '适用对象3', 2, 1, '青年女性,中年女性', 0, 0, 0, 0, 0, 1, 0, 0),
        (24, 1, '商品编号', 1, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (25, 1, '适用季节', 1, 1, '春季,夏季,秋季,冬季', 0, 0, 0, 0, 0, 1, 0, 0),
        (32, 2, '适用人群', 0, 1, '老年,青年,中年', 0, 0, 0, 0, 0, 1, 0, 0),
        (33, 2, '风格', 0, 1, '嘻哈风格,基础大众,商务正装', 0, 0, 0, 0, 0, 1, 0, 0),
        (35, 2, '颜色', 0, 0, '', 100, 0, 0, 0, 1, 0, 0, 0),
        (37, 1, '适用人群', 1, 1, '儿童,青年,中年,老年', 0, 0, 0, 0, 0, 1, 0, 0),
        (38, 1, '上市时间', 1, 1, '2017年秋,2017年冬,2018年春,2018年夏', 0, 0, 0, 0, 0, 1, 0, 0),
        (39, 1, '袖长', 1, 1, '短袖,长袖,中袖', 0, 0, 0, 0, 0, 1, 0, 0),
        (40, 2, '尺码', 0, 1, '29,30,31,32,33,34', 0, 0, 0, 0, 0, 0, 0, 0),
        (41, 2, '适用场景', 0, 1, '居家,运动,正装', 0, 0, 0, 0, 0, 1, 0, 0),
        (42, 2, '上市时间', 0, 1, '2018年春,2018年夏', 0, 0, 0, 0, 0, 1, 0, 0),
        (43, 3, '颜色', 0, 0, '', 100, 0, 0, 0, 1, 0, 0, 0),
        (44, 3, '容量', 0, 1, '16G,32G,64G,128G,256G,512G', 0, 0, 0, 0, 0, 0, 0, 0),
        (45, 3, '屏幕尺寸', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (46, 3, '网络', 0, 1, '3G,4G,5G,WLAN', 0, 0, 0, 0, 0, 1, 0, 0),
        (47, 3, '系统', 0, 1, 'Android,IOS', 0, 0, 0, 0, 0, 1, 0, 0),
        (48, 3, '电池容量', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (49, 11, '颜色', 0, 1, '红色,蓝色,绿色', 0, 1, 0, 0, 0, 0, 0, 0),
        (50, 11, '尺寸', 0, 1, '38,39,40', 0, 0, 0, 0, 0, 0, 0, 0),
        (51, 11, '风格', 0, 1, '夏季,秋季', 0, 0, 0, 0, 0, 0, 0, 0),
        (52, 12, '尺寸', 0, 1, '50英寸,65英寸,70英寸', 0, 0, 0, 0, 0, 0, 0, 0),
        (53, 12, '内存', 0, 1, '8G,16G,32G', 0, 0, 0, 0, 0, 0, 0, 0),
        (54, 12, '商品编号', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (55, 12, '商品毛重', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (56, 12, '商品产地', 0, 1, '中国大陆,其他', 0, 0, 0, 0, 0, 1, 0, 0),
        (57, 12, '电视类型', 0, 1, '大屏,教育电视,4K超清', 0, 0, 0, 0, 0, 1, 0, 0),
        (58, 13, '颜色', 0, 0, '', 0, 0, 0, 0, 1, 0, 0, 0),
        (59, 13, '版本', 0, 1, 'R7 16G 512,R5 16G 512,I5 16G 512,I7 16G 512', 0, 0, 0, 0, 0, 0, 0, 0),
        (60, 13, '屏幕尺寸', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (61, 13, '屏幕分辨率', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (62, 13, 'CPU型号', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (63, 14, '系列', 0, 0, '', 0, 0, 0, 0, 1, 0, 0, 0),
        (64, 14, '上市时间', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (65, 14, '毛重', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (66, 14, '额定功率', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (67, 15, '颜色', 0, 0, '', 0, 0, 0, 0, 1, 0, 0, 0),
        (68, 15, '版本', 0, 1, '512GB,1TB', 0, 0, 0, 0, 0, 0, 0, 0),
        (69, 15, '系列', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (70, 15, '型号', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (71, 15, '闪存类型', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (72, 15, '顺序读速', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0),
        (73, 15, '顺序写入', 0, 0, '', 0, 0, 0, 0, 0, 1, 0, 0);

insert into go_mall_admin.pms_product_attribute_value (id, product_id, product_attribute_id, value)
values  (1, 9, 1, 'X'),
        (2, 10, 1, 'X'),
        (3, 11, 1, 'X'),
        (4, 12, 1, 'X'),
        (5, 13, 1, 'X'),
        (6, 14, 1, 'X'),
        (7, 18, 1, 'X'),
        (8, 7, 1, 'X'),
        (9, 7, 1, 'XL'),
        (10, 7, 1, 'XXL'),
        (11, 22, 7, 'x,xx'),
        (12, 22, 24, 'no110'),
        (13, 22, 25, '春季'),
        (14, 22, 37, '青年'),
        (15, 22, 38, '2018年春'),
        (16, 22, 39, '长袖'),
        (124, 23, 7, '米白色,浅黄色'),
        (125, 23, 24, 'no1098'),
        (126, 23, 25, '春季'),
        (127, 23, 37, '青年'),
        (128, 23, 38, '2018年春'),
        (129, 23, 39, '长袖'),
        (130, 1, 13, ''),
        (131, 1, 14, ''),
        (132, 1, 15, ''),
        (133, 1, 16, ''),
        (134, 1, 17, ''),
        (135, 1, 18, ''),
        (136, 1, 19, ''),
        (137, 1, 20, ''),
        (138, 1, 21, ''),
        (139, 2, 13, ''),
        (140, 2, 14, ''),
        (141, 2, 15, ''),
        (142, 2, 16, ''),
        (143, 2, 17, ''),
        (144, 2, 18, ''),
        (145, 2, 19, ''),
        (146, 2, 20, ''),
        (147, 2, 21, ''),
        (243, 30, 7, '蓝色,白色'),
        (244, 30, 24, 'HNTBJ2E042A'),
        (245, 30, 25, '夏季'),
        (246, 30, 37, '青年'),
        (247, 30, 38, '2018年夏'),
        (248, 30, 39, '短袖'),
        (249, 31, 7, '浅灰色,深灰色'),
        (250, 31, 24, 'HNTBJ2E080A'),
        (251, 31, 25, '夏季'),
        (252, 31, 37, '青年'),
        (253, 31, 38, '2018年夏'),
        (254, 31, 39, '短袖'),
        (255, 32, 7, '黑色,白色'),
        (256, 32, 24, 'HNTBJ2E153A'),
        (257, 32, 25, '夏季'),
        (258, 32, 37, '青年'),
        (259, 32, 38, '2018年夏'),
        (260, 32, 39, '短袖'),
        (265, 33, 54, '4609652'),
        (266, 33, 55, '28.6kg'),
        (267, 33, 56, '中国大陆'),
        (268, 33, 57, '大屏'),
        (269, 34, 54, '4609660'),
        (270, 34, 55, '30.8kg'),
        (271, 34, 56, '中国大陆'),
        (272, 34, 57, '4K超清'),
        (288, 27, 43, '黑色,蓝色'),
        (289, 27, 45, '5.8'),
        (290, 27, 46, '4G'),
        (291, 27, 47, 'Android'),
        (292, 27, 48, '3000ml'),
        (303, 28, 43, '金色,银色'),
        (304, 28, 45, '5.0'),
        (305, 28, 46, '4G'),
        (306, 28, 47, 'Android'),
        (307, 28, 48, '2800ml'),
        (308, 29, 43, '金色,银色'),
        (309, 29, 45, '4.7'),
        (310, 29, 46, '4G'),
        (311, 29, 47, 'IOS'),
        (312, 29, 48, '1960ml'),
        (338, 37, 43, '午夜色,星光色,紫色,蓝色'),
        (339, 37, 45, '6.1英寸'),
        (340, 37, 46, '5G'),
        (341, 37, 47, 'IOS'),
        (342, 37, 48, '3000毫安'),
        (443, 38, 43, '银色,蓝色'),
        (444, 38, 45, '10.9英寸'),
        (445, 38, 46, 'WLAN'),
        (446, 38, 47, 'IOS'),
        (447, 38, 48, '6000毫安'),
        (448, 39, 58, '新小米Pro 14英寸 2.8K屏,新Redmi Pro 15英寸 3.2K屏'),
        (449, 39, 60, '15.6英寸'),
        (450, 39, 61, '3200*2000'),
        (451, 39, 62, 'R5 6600H'),
        (452, 41, 43, '墨羽,银迹'),
        (453, 41, 45, '6.67英寸'),
        (454, 41, 46, '5G'),
        (455, 41, 47, 'Android'),
        (456, 41, 48, '5500mAh'),
        (457, 42, 43, '曜金黑,冰霜银'),
        (458, 42, 45, '6.7英寸'),
        (459, 42, 46, '5G'),
        (460, 42, 47, 'Android'),
        (461, 42, 48, '4460mAh'),
        (462, 43, 63, 'JSQ25-565W13【13升】【恒温旗舰款】,JSQ30-565W16【16升】【恒温旗舰款】'),
        (463, 43, 64, '2021-05'),
        (464, 43, 65, '15.5kg'),
        (465, 43, 66, '30w'),
        (466, 44, 67, '新品980｜NVMe PCIe3.0*4,980 PRO｜NVMe PCIe 4.0'),
        (467, 44, 69, '980'),
        (468, 44, 70, 'MZ-V8V500BW'),
        (469, 44, 71, 'TLC'),
        (470, 44, 72, '3100MB/s'),
        (471, 44, 73, '2600MB/s'),
        (472, 45, 43, '鸢尾紫,晴空蓝'),
        (473, 45, 45, '6.43英寸'),
        (474, 45, 46, '5G'),
        (475, 45, 47, 'Android'),
        (476, 45, 48, '4500mAh'),
        (477, 40, 43, '黑色,蓝色'),
        (478, 40, 45, '6.73英寸'),
        (479, 40, 46, '5G'),
        (480, 40, 47, 'Android'),
        (481, 40, 48, '5160mAh'),
        (512, 26, 43, '金色,银色'),
        (513, 26, 45, '5.0'),
        (514, 26, 46, '4G'),
        (515, 26, 47, 'Android'),
        (516, 26, 48, '3000');



INSERT INTO `pms_product` VALUES (1, 7, '银色星芒刺绣网纱底裤', '111', 49, '111', 'No86577', 100.00, 120.00, 100, '件', 1000.00, 100, 100, 0, 0, 0, 1, 1, 1, '', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', 0, 0.00, 1524739263, 1524739263, 0, '/mall/images/20180522/web.png', '[]', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', 1, 1, 0, 0, 20, 0, '七匹狼', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (2, 7, '银色星芒刺绣网纱底裤2', '111', 49, '111', 'No86578', 100.00, 120.00, 100, '件', 1000.00, 1, 100, 0, 0, 0, 1, 1, 1, '', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤2', '银色星芒刺绣网纱底裤', 0, 0.00, 1524739263, 1524739263, 0, '/mall/images/20180522/web.png', '[]', '<p>银色星芒刺绣网纱底裤</p>', '<p>银色星芒刺绣网纱底裤</p>', 1, 1, 0, 0, 20, 0, '七匹狼', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (3, 7, '银色星芒刺绣网纱底裤3', '111', 1, '111', 'No86579', 100.00, 120.00, 100, '件', 1000.00, 1, 100, 0, 0, 0, 1, 1, 1, '', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤3', '银色星芒刺绣网纱底裤', 0, 0.00, 1524739263, 1524739263, 0, '/mall/images/20180522/web.png', '[]', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', 1, 1, 0, 0, 20, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (4, 7, '银色星芒刺绣网纱底裤4', '111', 1, '111', 'No86580', 100.00, 120.00, 100, '件', 1000.00, 1, 100, 0, 0, 0, 1, 1, 1, '', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤4', '银色星芒刺绣网纱底裤', 0, 0.00, 1524739263, 1524739263, 0, '/mall/images/20180522/web.png', '[]', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', 1, 1, 0, 0, 20, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (5, 7, '银色星芒刺绣网纱底裤5', '111', 1, '111', 'No86581', 100.00, 120.00, 100, '件', 1000.00, 1, 100, 0, 0, 0, 0, 1, 1, '', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤5', '银色星芒刺绣网纱底裤', 0, 0.00, 1524739263, 1524739263, 0, '/mall/images/20180522/web.png', '[]', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', 1, 1, 0, 0, 20, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (6, 7, '银色星芒刺绣网纱底裤6', '111', 1, '111', 'No86582', 100.00, 120.00, 100, '件', 1000.00, 1, 100, 0, 0, 0, 1, 1, 1, '', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤6', '银色星芒刺绣网纱底裤', 0, 0.00, 1524739263, 1524739263, 0, '/mall/images/20180522/web.png', '[]', '银色星芒刺绣网纱底裤', '银色星芒刺绣网纱底裤', 1, 1, 0, 0, 20, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (7, 7, '女式超柔软拉毛运动开衫', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 0, 0, 0, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (8, 7, '女式超柔软拉毛运动开衫1', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 0, 0, 0, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (9, 7, '女式超柔软拉毛运动开衫1', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 0, 0, 0, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (10, 7, '女式超柔软拉毛运动开衫1', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 0, 0, 0, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (11, 7, '女式超柔软拉毛运动开衫1', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 1, 0, 1, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (12, 7, '女式超柔软拉毛运动开衫2', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 1, 0, 1, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (13, 7, '女式超柔软拉毛运动开衫3', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 1, 0, 1, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (14, 7, '女式超柔软拉毛运动开衫3', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 0, 0, 1, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (18, 7, '女式超柔软拉毛运动开衫3', '匠心剪裁，垂感质地', 1, '匠心剪裁，垂感质地', 'No86577', 249.00, 299.00, 100, '件', 0.00, 0, 100, 0, 0, 0, 0, 0, 1, 'string', 'string', 'string', '女式超柔软拉毛运动开衫', 'string', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180522/web.png', '[\"string\"]', 'string', 'string', 0, 1, 0, 0, 0, 0, '万和', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (22, 7, 'test', 'test', 6, '', '', 0.00, 0.00, 100, '', 0.00, 0, 0, 0, 0, 1, 1, 0, 0, '1,2', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180604/1522738681.jpg', '[]', '', '', 0, 1, 0, 0, 0, 0, '小米', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (23, 19, '毛衫测试', '毛衫测试11', 6, 'xxx', 'NO.1098', 99.00, 109.00, 100, '件', 1000.00, 0, 99, 99, 1000, 1, 1, 1, 1, '1,2,3', '毛衫测试', '毛衫测试', '毛衫测试', '毛衫测试', 2, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180604/1522738681.jpg', '[\"/mall/images/20180604/1522738681.jpg\",\"/mall/images/20180604/1522738681.jpg\"]', '<p><img class=\"wscnph\" src=\"/mall/images/20180604/155x54.bmp\" /><img class=\"wscnph\" src=\"/mall/images/20180604/APP登录bg1080.jpg\" width=\"500\" height=\"500\" /><img class=\"wscnph\" src=\"/mall/images/20180604/APP登录界面.jpg\" width=\"500\" height=\"500\" /></p>', '', 0, 1, 0, 0, 0, 0, '小米', '手机数码', 0, 0);
INSERT INTO `pms_product` VALUES (24, 7, 'xxx', 'xxx', 6, '', '', 0.00, 0.00, 100, '', 0.00, 0, 0, 0, 0, 0, 0, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 0, '', '[]', '', '', 0, 1, 0, 0, 0, 0, '小米', '外套', 0, 0);
INSERT INTO `pms_product` VALUES (26, 19, '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', 3, '', '6946605', 3788.00, 4288.00, 1000, '件', 0.00, 100, 3788, 3788, 0, 1, 1, 1, 1, '2,3,1', '', '', '', '', 1, 3659.00, 1673365778, 1675123200, 3, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '[\"/mall/images/20180607/5ab46a3cN616bdc41.jpg\",\"/mall/images/20180607/5ac1bf5fN2522b9dc.jpg\"]', '<p><img class=\"wscnph\" src=\"/mall/images/20180607/5ad44f1cNf51f3bb0.jpg\" /><img class=\"wscnph\" src=\"/mall/images/20180607/5ad44fa8Nfcf71c10.jpg\" /><img class=\"wscnph\" src=\"/mall/images/20180607/5ad44fa9N40e78ee0.jpg\" /><img class=\"wscnph\" src=\"/mall/images/20180607/5ad457f4N1c94bdda.jpg\" /><img class=\"wscnph\" src=\"/mall/images/20180607/5ad457f5Nd30de41d.jpg\" /><img class=\"wscnph\" src=\"/mall/images/20180607/5b10fb0eN0eb053fb.jpg\" /></p>', '<p><img src=\"//img20.360buyimg.com/vc/jfs/t1/81293/35/5822/369414/5d3fe77cE619c5487/6e775a52850feea5.jpg!q70.dpg.webp\" alt=\"\" width=\"750\" height=\"776\" /></p>\n<p><img src=\"//img20.360buyimg.com/vc/jfs/t1/45300/25/11592/3658871/5d85ef66E92a8a911/083e47d8f662c582.jpg!q70.dpg.webp\" alt=\"\" width=\"596\" height=\"16383\" /></p>', 0, 0, 0, 100, 0, 0, '华为', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (27, 19, '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', 6, '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '7437788', 2699.00, 2699.00, 100, '', 0.00, 0, 2699, 2699, 0, 0, 1, 1, 1, '1', '', '', '', '', 3, 0.00, 1524739263, 1524739263, 3, '/mall/images/20180615/xiaomi.jpg', '[]', '<p style=\"text-align: center;\"><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/95935/9/19330/159477/5e9ecc13E5b8db8ae/8e954021a0835fb7.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/99224/22/19596/137593/5e9ecc13E34ef2113/2b362b90d8378ee1.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/93131/25/19321/107691/5e9ecc13E41e8addf/202a5f84d9129878.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/101843/19/19533/66831/5e9ecc13Ecb7f9d53/4fdd807266583c1e.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/99629/36/19016/59882/5e9ecc13E1f5beef5/1e5af3528f366e70.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/72343/29/8945/84548/5d6e5c67Ea07b1125/702791816b90eb3d.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/65403/35/9017/129532/5d6e5c68E3f2d0546/9ec771eb6e04a75a.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/58261/33/9380/106603/5d6e5c68E05a99177/2b5b9e29eed05b08.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/51968/40/9688/113552/5d6e5c68E5052b312/8969d83124cb78a4.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/75491/9/9101/146883/5d6e5c68E3db89775/c1aa57e78ded4e44.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/75063/11/9107/127874/5d6e5c68E0b1dfc61/10dd6000ce213375.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/47452/25/9579/108279/5d6e5c68Ea9002f3b/865b5d32ffd9da75.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/82146/26/9077/87075/5d6e5c68Ef63bccc8/556de5665a35a3f2.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/82804/21/9019/124404/5d6e5c69E06a8f575/0f861f8c4636c546.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/46044/10/9734/107686/5d6e5c69Edd5e66c7/a8c7b9324e271dbd.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/40729/32/13755/45997/5d6e5c69Eafee3664/6a3457a4efdb79c5.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/76254/34/9039/96195/5d6e5c69E3c71c809/49033c0b7024c208.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/79825/20/9065/90864/5d6e5c69E1e62ef89/a4d3ce383425a666.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/49939/21/9618/106207/5d6e5c6aEf7b1d4fd/0f5e963c66be3d0c.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/64035/7/9001/115159/5d6e5c6aE6919dfdf/39dfe4840157ad81.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/53389/3/9616/99637/5d6e5c6aEa33b9f35/b8f6aa26e72616a3.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/63219/6/9026/81576/5d6e5c6aEa9c74b49/b4fa364437531012.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/42146/27/13902/80437/5d6e5c6bE30c31ce9/475d4d54c7334313.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/45317/28/9596/78175/5d6e5c6bEce31e4b7/5675858b6933565c.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/60080/1/9112/138722/5d6e5c6bEefd9fc62/7ece7460a36d2fcc.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/51525/13/9549/36018/5d6e5c73Ebbccae6c/99bc2770dccc042b.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/61948/20/9088/72918/5d6e5c73Eab7aef5c/6f21e2f85cf478fa.jpg!q70.dpg.webp\" /></p>', '<p style=\"text-align: center;\"><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/95935/9/19330/159477/5e9ecc13E5b8db8ae/8e954021a0835fb7.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/99224/22/19596/137593/5e9ecc13E34ef2113/2b362b90d8378ee1.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/93131/25/19321/107691/5e9ecc13E41e8addf/202a5f84d9129878.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/101843/19/19533/66831/5e9ecc13Ecb7f9d53/4fdd807266583c1e.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWareDetail/jfs/t1/99629/36/19016/59882/5e9ecc13E1f5beef5/1e5af3528f366e70.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/72343/29/8945/84548/5d6e5c67Ea07b1125/702791816b90eb3d.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/65403/35/9017/129532/5d6e5c68E3f2d0546/9ec771eb6e04a75a.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/58261/33/9380/106603/5d6e5c68E05a99177/2b5b9e29eed05b08.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/51968/40/9688/113552/5d6e5c68E5052b312/8969d83124cb78a4.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/75491/9/9101/146883/5d6e5c68E3db89775/c1aa57e78ded4e44.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/75063/11/9107/127874/5d6e5c68E0b1dfc61/10dd6000ce213375.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/47452/25/9579/108279/5d6e5c68Ea9002f3b/865b5d32ffd9da75.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/82146/26/9077/87075/5d6e5c68Ef63bccc8/556de5665a35a3f2.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/82804/21/9019/124404/5d6e5c69E06a8f575/0f861f8c4636c546.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/46044/10/9734/107686/5d6e5c69Edd5e66c7/a8c7b9324e271dbd.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/40729/32/13755/45997/5d6e5c69Eafee3664/6a3457a4efdb79c5.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/76254/34/9039/96195/5d6e5c69E3c71c809/49033c0b7024c208.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/79825/20/9065/90864/5d6e5c69E1e62ef89/a4d3ce383425a666.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/49939/21/9618/106207/5d6e5c6aEf7b1d4fd/0f5e963c66be3d0c.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/64035/7/9001/115159/5d6e5c6aE6919dfdf/39dfe4840157ad81.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/53389/3/9616/99637/5d6e5c6aEa33b9f35/b8f6aa26e72616a3.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/63219/6/9026/81576/5d6e5c6aEa9c74b49/b4fa364437531012.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/42146/27/13902/80437/5d6e5c6bE30c31ce9/475d4d54c7334313.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/45317/28/9596/78175/5d6e5c6bEce31e4b7/5675858b6933565c.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/60080/1/9112/138722/5d6e5c6bEefd9fc62/7ece7460a36d2fcc.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/51525/13/9549/36018/5d6e5c73Ebbccae6c/99bc2770dccc042b.jpg!q70.dpg.webp\" /><img src=\"//img30.360buyimg.com/popWaterMark/jfs/t1/61948/20/9088/72918/5d6e5c73Eab7aef5c/6f21e2f85cf478fa.jpg!q70.dpg.webp\" /></p>', 0, 0, 0, 99, 0, 0, '小米', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (28, 19, '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', '8天超长待机，137g轻巧机身，高通骁龙处理器小米6X低至1299，点击抢购', 6, '', '7437789', 649.00, 649.00, 100, '', 0.00, 0, 649, 649, 0, 0, 1, 1, 1, '', '', '', '', '', 4, 0.00, 1524739263, 1524739263, 3, '/mall/images/20180615/5a9d248cN071f4959.jpg', '[]', '', '<div><img src=\"//img12.360buyimg.com/imgzone/jfs/t1/67362/12/14546/708984/5dc28512Eefdd817d/c82503af0da6c038.gif\" /><img src=\"//img13.360buyimg.com/imgzone/jfs/t1/61488/17/14551/995918/5dc28512E821c228d/41e52005ea5b1f36.gif\" /><img src=\"//img14.360buyimg.com/imgzone/jfs/t1/72961/36/14769/305883/5dc28512E65d77261/3df6be29e3d489d1.gif\" />\n<p>&nbsp;</p>\n</div>', 0, 0, 0, 98, 0, 0, '小米', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (29, 19, 'Apple iPhone 8 Plus 64GB 红色特别版 移动联通电信4G手机', '【限时限量抢购】Apple产品年中狂欢节，好物尽享，美在智慧！速来 >> 勾选[保障服务][原厂保2年]，获得AppleCare+全方位服务计划，原厂延保售后无忧。', 51, '', '7437799', 5499.00, 5499.00, 100, '', 0.00, 0, 5499, 5499, 0, 0, 1, 1, 1, '1,2,3', '', '', '', '', 1, 4799.00, 1588605174, 1590796800, 3, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', '[]', '', '<div><img src=\"//img10.360buyimg.com/cms/jfs/t1/20020/38/9725/228440/5c7f9208Eeaf4bf87/13a713066f71791d.jpg!q70.dpg.webp\" /> <img src=\"//img12.360buyimg.com/cms/jfs/t1/12128/39/9607/265349/5c7f9209Ecff29b88/08620ba570705634.jpg!q70.dpg.webp\" /> <img src=\"//img14.360buyimg.com/cms/jfs/t1/22731/40/9578/345163/5c7f9209E9ba056e5/a8a557060b84447e.jpg!q70.dpg.webp\" /> <img src=\"//img14.360buyimg.com/cms/jfs/t1/29506/38/9439/299492/5c7f9209E0e51eb29/15dedd95416f3c68.jpg!q70.dpg.webp\" /> <img src=\"//img14.360buyimg.com/cms/jfs/t1/26766/28/9574/257101/5c7f9209Eaca1b317/c7caa047b1566cf1.jpg!q70.dpg.webp\" /> <img src=\"//img13.360buyimg.com/cms/jfs/t1/11059/8/10473/286531/5c7f9208E05da0120/9034ad8799ad9553.jpg!q70.dpg.webp\" /> <img src=\"//img14.360buyimg.com/cms/jfs/t1/25641/2/9557/268826/5c7f9208Efbf0dc50/399580629e05e733.jpg!q70.dpg.webp\" /> <img src=\"//img13.360buyimg.com/cms/jfs/t1/28964/25/9527/305902/5c7f9208E275ffb9c/8464934d67e69b7a.jpg!q70.dpg.webp\" /></div>', 0, 0, 0, 97, 0, 0, '苹果', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (30, 8, 'HLA海澜之家简约动物印花短袖T恤', '2018夏季新品微弹舒适新款短T男生 6月6日-6月20日，满300减30，参与互动赢百元礼券，立即分享赢大奖', 50, '', 'HNTBJ2E042A', 98.00, 98.00, 100, '', 0.00, 0, 0, 0, 0, 0, 1, 1, 1, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180615/5ad83a4fN6ff67ecd.jpg!cc_350x449.jpg', '[]', '', '', 0, 0, 0, 0, 0, 0, '海澜之家', 'T恤', 0, 0);
INSERT INTO `pms_product` VALUES (31, 8, 'HLA海澜之家蓝灰花纹圆领针织布短袖T恤', '2018夏季新品短袖T恤男HNTBJ2E080A 蓝灰花纹80 175/92A/L80A 蓝灰花纹80 175/92A/L', 50, '', 'HNTBJ2E080A', 98.00, 98.00, 100, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180615/5ac98b64N70acd82f.jpg!cc_350x449.jpg', '[]', '', '', 0, 0, 0, 0, 0, 0, '海澜之家', 'T恤', 0, 0);
INSERT INTO `pms_product` VALUES (32, 8, 'HLA海澜之家短袖T恤男基础款', 'HLA海澜之家短袖T恤男基础款简约圆领HNTBJ2E153A藏青(F3)175/92A(50)', 50, '', 'HNTBJ2E153A', 68.00, 68.00, 100, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '1,2', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 1, '/mall/images/20180615/5a51eb88Na4797877.jpg', '[]', '', '', 0, 0, 0, 0, 0, 0, '海澜之家', 'T恤', 0, 0);
INSERT INTO `pms_product` VALUES (33, 35, '小米（MI）小米电视4A ', '小米（MI）小米电视4A 55英寸 L55M5-AZ/L55M5-AD 2GB+8GB HDR 4K超高清 人工智能网络液晶平板电视', 6, '', '4609652', 2499.00, 2499.00, 100, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 12, '/mall/images/20180615/5b02804dN66004d73.jpg', '[]', '', '', 0, 0, 0, 0, 0, 0, '小米', '电视', 0, 0);
INSERT INTO `pms_product` VALUES (34, 35, '小米（MI）小米电视4A 65英寸', ' L65M5-AZ/L65M5-AD 2GB+8GB HDR 4K超高清 人工智能网络液晶平板电视', 6, '', '4609660', 3999.00, 3999.00, 100, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '1,2', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 12, '/mall/images/20180615/5b028530N51eee7d4.jpg', '[]', '', '', 0, 0, 0, 0, 0, 0, '小米', '电视', 0, 0);
INSERT INTO `pms_product` VALUES (35, 29, '耐克NIKE 男子 休闲鞋 ROSHE RUN 运动鞋 511881-010黑色41码', '耐克NIKE 男子 休闲鞋 ROSHE RUN 运动鞋 511881-010黑色41码', 58, '', '6799342', 369.00, 369.00, 100, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 11, '/mall/images/20180615/5b235bb9Nf606460b.jpg', '[]', '', '', 0, 0, 0, 0, 0, 0, 'NIKE', '男鞋', 0, 0);
INSERT INTO `pms_product` VALUES (36, 29, '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 58, '', '6799345', 499.00, 499.00, 100, '', 0.00, 0, 0, 0, 0, 0, 1, 1, 1, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 11, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '[]', '', '', 0, 0, 0, 0, 0, 0, 'NIKE', '男鞋', 0, 0);
INSERT INTO `pms_product` VALUES (37, 19, 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', 51, '', '100038005189', 5999.00, 5999.00, 1000, '', 208.00, 200, 0, 0, 0, 0, 1, 0, 0, '1,2,3', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 3, '/mall/images/20221028/iphone14_001.jpg', '[\"/mall/images/20221028/iphone14_002.jpg\",\"/mall/images/20221028/iphone14_003.jpg\",\"/mall/images/20221028/iphone14_004.jpg\"]', '', '<div style=\"margin: 0 auto;\"><img style=\"max-width: 390px;\" src=\"//img13.360buyimg.com/cms/jfs/t1/58071/39/19839/119559/63190110Edac0cea7/b62a84f1b8775fef.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img20.360buyimg.com/cms/jfs/t1/138903/36/29400/86115/63190110E0a98c819/d2efbef39eeb2995.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img30.360buyimg.com/cms/jfs/t1/176347/20/28995/115695/63190110Ef5d766f9/aee3d2d866522f66.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img14.360buyimg.com/cms/jfs/t1/120515/39/28721/142961/63190110Eec31e31a/3486d6a065a18ddc.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img13.360buyimg.com/cms/jfs/t1/30161/12/17674/81508/63190110E1385cf61/f05a2a43f4d304ff.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img14.360buyimg.com/cms/jfs/t1/100037/16/31071/62177/6319010fE871efe01/b01a6f981c268e38.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img14.360buyimg.com/cms/jfs/t1/90843/33/25852/74752/63190110E373559f4/74b6b52a3fb08c66.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img14.360buyimg.com/cms/jfs/t1/181914/22/28400/126094/63190110Edefb838c/802a16e758be2b1d.jpg!q70.dpg.webp\" /></div>', 0, 0, 0, 0, 0, 0, '苹果', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (38, 53, 'Apple iPad 10.9英寸平板电脑 2022年款', '【11.11大爱超大爱】iPad9代限量抢购，价格优惠，更享以旧换新至高补贴325元！！快来抢购吧！！ ', 51, '', '100044025833', 3599.00, 3599.00, 1000, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '1,2,3', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 3, '/mall/images/20221028/ipad_001.jpg', '[\"/mall/images/20221028/ipad_002.jpg\"]', '', '<div style=\"margin: 0 auto;\"><img style=\"max-width: 390px;\" src=\"//img12.360buyimg.com/cms/jfs/t1/75040/28/21026/295081/634ed154E981e4d10/2ceef91d6f2b2273.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img13.360buyimg.com/cms/jfs/t1/191028/1/28802/401291/634ed15eEb234dc40/5ab89f83531e1023.jpg!q70.dpg.webp\" /> <img style=\"max-width: 390px;\" src=\"//img14.360buyimg.com/cms/jfs/t1/32758/24/18599/330590/634ed15eEc3db173c/c52953dc8008a576.jpg!q70.dpg.webp\" /></div>', 0, 0, 0, 0, 0, 0, '苹果', '平板电脑', 0, 0);
INSERT INTO `pms_product` VALUES (39, 54, '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑', '【双十一大促来袭】指定型号至高优惠1000，以旧换新至高补贴1000元，晒单赢好礼', 6, '', '100023207945', 5599.00, 5599.00, 500, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 1, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 13, '/mall/images/20221028/xiaomi_computer_001.jpg', '[\"/mall/images/20221028/xiaomi_computer_002.jpg\"]', '', '<div class=\"ssd-module-mobile-wrap\">\n<div class=\"ssd-module M16667778180631\" data-id=\"M16667778180631\"><img class=\"wscnph\" src=\"/mall/images/20221028/xiaomi_computer_snipaste_01.png\" /><img class=\"wscnph\" src=\"/mall/images/20221028/xiaomi_computer_snipaste_02.png\" /><img class=\"wscnph\" src=\"/mall/images/20221028/xiaomi_computer_snipaste_03.png\" /><img class=\"wscnph\" src=\"/mall/images/20221028/xiaomi_computer_snipaste_04.png\" /><img class=\"wscnph\" src=\"/mall/images/20221028/xiaomi_computer_snipaste_05.png\" /><img class=\"wscnph\" src=\"/mall/images/20221028/xiaomi_computer_snipaste_06.png\" /></div>\n<div class=\"ssd-module M16667778180631\" data-id=\"M16667778180631\">&nbsp;</div>\n<div class=\"ssd-module M16534569204241\" data-id=\"M16534569204241\">&nbsp;</div>\n<div class=\"ssd-module M16667778416884\" data-id=\"M16667778416884\">\n<div class=\"ssd-widget-text W16667778440496\">&nbsp;</div>\n</div>\n<div class=\"ssd-module M165303211067557\" data-id=\"M165303211067557\">&nbsp;</div>\n<div class=\"ssd-module M16530320459861\" data-id=\"M16530320459861\">&nbsp;</div>\n<div class=\"ssd-module M16530320460062\" data-id=\"M16530320460062\">&nbsp;</div>\n<div class=\"ssd-module M16530320460153\" data-id=\"M16530320460153\">&nbsp;</div>\n<div class=\"ssd-module M16530320460366\" data-id=\"M16530320460366\">&nbsp;</div>\n<div class=\"ssd-module M16530320460477\" data-id=\"M16530320460477\">&nbsp;</div>\n<div class=\"ssd-module M16530320460578\" data-id=\"M16530320460578\">&nbsp;</div>\n<div class=\"ssd-module M16530320460699\" data-id=\"M16530320460699\">&nbsp;</div>\n<div class=\"ssd-module M165303204608110\" data-id=\"M165303204608110\">&nbsp;</div>\n<div class=\"ssd-module M165303204609511\" data-id=\"M165303204609511\">&nbsp;</div>\n<div class=\"ssd-module M165303204611213\" data-id=\"M165303204611213\">&nbsp;</div>\n<div class=\"ssd-module M165303204612714\" data-id=\"M165303204612714\">&nbsp;</div>\n<div class=\"ssd-module M165303204614115\" data-id=\"M165303204614115\">&nbsp;</div>\n<div class=\"ssd-module M165303204615516\" data-id=\"M165303204615516\">&nbsp;</div>\n<div class=\"ssd-module M165303204617417\" data-id=\"M165303204617417\">&nbsp;</div>\n<div class=\"ssd-module M165303204618818\" data-id=\"M165303204618818\">&nbsp;</div>\n<div class=\"ssd-module M165303204620219\" data-id=\"M165303204620219\">&nbsp;</div>\n<div class=\"ssd-module M165303204621620\" data-id=\"M165303204621620\">&nbsp;</div>\n<div class=\"ssd-module M165303204622921\" data-id=\"M165303204622921\">&nbsp;</div>\n<div class=\"ssd-module M165303204624522\" data-id=\"M165303204624522\">&nbsp;</div>\n<div class=\"ssd-module M165303204626024\" data-id=\"M165303204626024\">&nbsp;</div>\n<div class=\"ssd-module M165303204627525\" data-id=\"M165303204627525\">&nbsp;</div>\n<div class=\"ssd-module M165303204629127\" data-id=\"M165303204629127\">&nbsp;</div>\n<div class=\"ssd-module M165303204632330\" data-id=\"M165303204632330\">&nbsp;</div>\n<div class=\"ssd-module M165303204634031\" data-id=\"M165303204634031\">&nbsp;</div>\n<div class=\"ssd-module M165303204635832\" data-id=\"M165303204635832\">&nbsp;</div>\n<div class=\"ssd-module M165303204637533\" data-id=\"M165303204637533\">&nbsp;</div>\n<div class=\"ssd-module M165303204639334\" data-id=\"M165303204639334\">&nbsp;</div>\n<div class=\"ssd-module M165303204642235\" data-id=\"M165303204642235\">&nbsp;</div>\n<div class=\"ssd-module M165303204647936\" data-id=\"M165303204647936\">&nbsp;</div>\n<div class=\"ssd-module M165303204651037\" data-id=\"M165303204651037\">&nbsp;</div>\n<div class=\"ssd-module M165303204653838\" data-id=\"M165303204653838\">&nbsp;</div>\n<div class=\"ssd-module M165303204656239\" data-id=\"M165303204656239\">&nbsp;</div>\n<div class=\"ssd-module M165303204659141\" data-id=\"M165303204659141\">&nbsp;</div>\n<div class=\"ssd-module M165303204661943\" data-id=\"M165303204661943\">&nbsp;</div>\n<div class=\"ssd-module M165303204665844\" data-id=\"M165303204665844\">&nbsp;</div>\n<div class=\"ssd-module M165303204668045\" data-id=\"M165303204668045\">&nbsp;</div>\n<div class=\"ssd-module M165303204670146\" data-id=\"M165303204670146\">&nbsp;</div>\n<div class=\"ssd-module M165303204672147\" data-id=\"M165303204672147\">&nbsp;</div>\n<div class=\"ssd-module M165303204674348\" data-id=\"M165303204674348\">&nbsp;</div>\n<div class=\"ssd-module M165303204676749\" data-id=\"M165303204676749\">&nbsp;</div>\n<div class=\"ssd-module M165303204681352\" data-id=\"M165303204681352\">&nbsp;</div>\n<div class=\"ssd-module M165303204683553\" data-id=\"M165303204683553\">&nbsp;</div>\n<div class=\"ssd-module M165303204685855\" data-id=\"M165303204685855\">&nbsp;</div>\n<div class=\"ssd-module M165303204688356\" data-id=\"M165303204688356\">&nbsp;</div>\n</div>', 0, 0, 0, 0, 0, 0, '小米', '笔记本', 0, 0);
INSERT INTO `pms_product` VALUES (40, 19, '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', 6, '', '100027789721', 2999.00, 2999.00, 500, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 1, '', '', '', '', '', 4, 0.00, 1524739263, 1524739263, 3, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '[\"/mall/images/20221104/xiaomi_12_pro_02.jpg\"]', '', '<p><img class=\"wscnph\" src=\"/mall/images/20221104/xiaomi_12_pro_snipaste_01.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/xiaomi_12_pro_snipaste_02.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/xiaomi_12_pro_snipaste_03.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/xiaomi_12_pro_snipaste_04.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/xiaomi_12_pro_snipaste_05.png\" /></p>', 0, 0, 0, 0, 0, 0, '小米', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (41, 19, 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', 6, '', '100035246702', 2099.00, 2099.00, 1000, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 3, '/mall/images/20221104/redmi_k50_01.jpg', '[\"/mall/images/20221104/redmi_k50_02.jpg\"]', '', '<p><img class=\"wscnph\" src=\"/mall/images/20221104/redmi_k50_snipaste_01.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/redmi_k50_snipaste_02.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/redmi_k50_snipaste_03.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/redmi_k50_snipaste_04.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/redmi_k50_snipaste_05.png\" /></p>', 0, 0, 0, 0, 0, 0, '小米', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (42, 19, 'HUAWEI Mate 50 直屏旗舰 超光变XMAGE影像 北斗卫星消息', '【华为Mate50新品上市】内置66W华为充电套装，超光变XMAGE影像,北斗卫星消息，鸿蒙操作系统3.0！立即抢购！华为新品可持续计划，猛戳》 ', 3, '', '100035295081', 4999.00, 4999.00, 1000, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 3, '/mall/images/20221104/huawei_mate50_01.jpg', '[\"/mall/images/20221104/huawei_mate50_02.jpg\"]', '', '<p><img class=\"wscnph\" src=\"/mall/images/20221104/huawei_mate50_snipaste_01.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/huawei_mate50_snipaste_02.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/huawei_mate50_snipaste_03.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/huawei_mate50_snipaste_04.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/huawei_mate50_snipaste_05.png\" /></p>', 0, 0, 0, 0, 0, 0, '华为', '手机通讯', 0, 0);
INSERT INTO `pms_product` VALUES (43, 39, '万和（Vanward)燃气热水器天然气家用四重防冻直流变频节能全新升级增压水伺服恒温高抗风', '【家电11.11享低价，抢到手价不高于1199】【发布种草秀享好礼！同价11.11买贵补差】爆款超一级能效零冷水】 ', 1, '', '10044060351752', 1799.00, 1799.00, 1000, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 14, '/mall/images/20221104/wanhe_13L_01.png', '[\"/mall/images/20221104/wanhe_16L_01.jpg\"]', '', '<p><img class=\"wscnph\" src=\"/mall/images/20221104/wanhe_water_snipaste_01.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/wanhe_water_snipaste_02.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/wanhe_water_snipaste_03.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/wanhe_water_snipaste_04.png\" /><img class=\"wscnph\" src=\"/mall/images/20221104/wanhe_water_snipaste_05.png\" /></p>', 0, 0, 0, 0, 0, 0, '万和', '厨卫大电', 0, 0);
INSERT INTO `pms_product` VALUES (44, 55, '三星（SAMSUNG）500GB SSD固态硬盘 M.2接口(NVMe协议)', '【满血无缓存！进店抽百元E卡，部分型号白条三期免息】兼具速度与可靠性！读速高达3500MB/s，全功率模式！点击 ', 2, '', '100018768480', 369.00, 369.00, 1000, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 0, 0.00, 1524739263, 1524739263, 15, '/mall/images/20221108/sanxing_ssd_02.jpg', '[\"/mall/images/20221108/sanxing_ssd_01.jpg\"]', '', '<p><img class=\"wscnph\" src=\"/mall/images/20221108/sanxing_ssd_snipaste_01.png\" /><img class=\"wscnph\" src=\"/mall/images/20221108/sanxing_ssd_snipaste_02.png\" /><img class=\"wscnph\" src=\"/mall/images/20221108/sanxing_ssd_snipaste_03.png\" /></p>', 0, 0, 0, 0, 0, 0, '三星', '硬盘', 0, 0);
INSERT INTO `pms_product` VALUES (45, 19, 'OPPO Reno8 8GB+128GB 鸢尾紫 新配色上市 80W超级闪充 5000万水光人像三摄', '【11.11提前购机享价保，好货不用等，系统申请一键价保补差!】【Reno8Pro爆款优惠】 ', 21, '', '10052147850350', 2299.00, 2299.00, 1000, '', 0.00, 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', 4, 999.00, 1668010550, 1669334400, 3, '/mall/images/20221108/oppo_r8_01.jpg', '[\"/mall/images/20221108/oppo_r8_02.jpg\"]', '', '<p><img class=\"wscnph\" src=\"/mall/images/20221108/oppo_r8_snipaste_01.png\" /><img class=\"wscnph\" src=\"/mall/images/20221108/oppo_r8_snipaste_02.png\" /><img class=\"wscnph\" src=\"/mall/images/20221108/oppo_r8_snipaste_03.png\" /><img class=\"wscnph\" src=\"/mall/images/20221108/oppo_r8_snipaste_04.png\" /><img class=\"wscnph\" src=\"/mall/images/20221108/oppo_r8_snipaste_05.png\" /></p>', 0, 0, 0, 0, 0, 0, 'OPPO', '手机通讯', 0, 0);



insert into go_mall_admin.pms_sku_stock (id, product_id, sku_code, price, stock, low_stock, pic, sale, promotion_price, lock_stock, sp_data)
values  (98, 27, '201808270027001', 2699.00, 86, 0, '', 0, 0.00, -24, '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]'),
        (99, 27, '201808270027002', 2999.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"64G"}]'),
        (100, 27, '201808270027003', 2699.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"32G"}]'),
        (101, 27, '201808270027004', 2999.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"64G"}]'),
        (102, 28, '201808270028001', 649.00, 99, 0, '', 0, 0.00, -8, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]'),
        (103, 28, '201808270028002', 699.00, 99, 0, '', 0, 0.00, -8, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]'),
        (104, 28, '201808270028003', 649.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"16G"}]'),
        (105, 28, '201808270028004', 699.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"32G"}]'),
        (106, 29, '201808270029001', 5499.00, 99, 0, '', 0, 4999.00, -8, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]'),
        (107, 29, '201808270029002', 6299.00, 100, 0, '', 0, 5799.00, 0, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"64G"}]'),
        (108, 29, '201808270029003', 5499.00, 100, 0, '', 0, 4999.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"32G"}]'),
        (109, 29, '201808270029004', 6299.00, 100, 0, '', 0, 5799.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]'),
        (110, 26, '201806070026001', 3788.00, 487, 0, '', 0, 3699.00, 0, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]'),
        (111, 26, '201806070026002', 3999.00, 499, 0, '', 0, 3899.00, 0, '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]'),
        (112, 26, '201806070026003', 3788.00, 500, 0, '', 0, 3699.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"16G"}]'),
        (113, 26, '201806070026004', 3999.00, 500, 0, '', 0, 3899.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"32G"}]'),
        (163, 36, '202002210036001', 100.00, 100, 25, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]'),
        (164, 36, '202002210036002', 120.00, 98, 20, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]'),
        (165, 36, '202002210036003', 100.00, 100, 20, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]'),
        (166, 36, '202002210036004', 100.00, 100, 20, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]'),
        (167, 36, '202002210036005', 100.00, 100, 20, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]'),
        (168, 36, '202002210036006', 100.00, 100, 20, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]'),
        (169, 36, '202002210036007', 100.00, 100, 20, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]'),
        (170, 36, '202002210036008', 100.00, 100, 20, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]'),
        (171, 35, '202002250035001', 200.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]'),
        (172, 35, '202002250035002', 240.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]'),
        (173, 35, '202002250035003', 200.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]'),
        (174, 35, '202002250035004', 200.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]'),
        (175, 35, '202002250035005', 200.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]'),
        (176, 35, '202002250035006', 200.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]'),
        (177, 35, '202002250035007', 200.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"夏季"}]'),
        (178, 35, '202002250035008', 200.00, 100, 50, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"39"},{"key":"风格","value":"秋季"}]'),
        (179, 30, '202004190030001', 88.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"X"}]'),
        (180, 30, '202004190030002', 88.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"XL"}]'),
        (181, 30, '202004190030003', 88.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"尺寸","value":"M"}]'),
        (182, 30, '202004190030004', 88.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"白色"},{"key":"尺寸","value":"X"}]'),
        (183, 30, '202004190030005', 88.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"白色"},{"key":"尺寸","value":"XL"}]'),
        (184, 30, '202004190030006', 88.00, 100, 0, '', 0, 0.00, 0, '[{"key":"颜色","value":"白色"},{"key":"尺寸","value":"M"}]'),
        (185, 31, '202004190031001', 88.00, 80, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"浅灰色"},{"key":"尺寸","value":"X"}]'),
        (186, 31, '202004190031002', 88.00, 80, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"浅灰色"},{"key":"尺寸","value":"XL"}]'),
        (187, 31, '202004190031003', 88.00, 80, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"深灰色"},{"key":"尺寸","value":"X"}]'),
        (188, 31, '202004190031004', 88.00, 80, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"深灰色"},{"key":"尺寸","value":"XL"}]'),
        (189, 32, '202004190032001', 99.00, 100, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"黑色"},{"key":"尺寸","value":"X"}]'),
        (190, 32, '202004190032002', 99.00, 100, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"黑色"},{"key":"尺寸","value":"M"}]'),
        (191, 32, '202004190032003', 99.00, 100, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"白色"},{"key":"尺寸","value":"X"}]'),
        (192, 32, '202004190032004', 99.00, 100, 10, '', 0, 0.00, 0, '[{"key":"颜色","value":"白色"},{"key":"尺寸","value":"M"}]'),
        (193, 33, '202004190033001', 2499.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"50英寸"},{"key":"内存","value":"8G"}]'),
        (194, 33, '202004190033002', 2499.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"50英寸"},{"key":"内存","value":"16G"}]'),
        (195, 33, '202004190033003', 2499.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"65英寸"},{"key":"内存","value":"8G"}]'),
        (196, 33, '202004190033004', 2499.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"65英寸"},{"key":"内存","value":"16G"}]'),
        (197, 34, '202004190034001', 3999.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"65英寸"},{"key":"内存","value":"16G"}]'),
        (198, 34, '202004190034002', 3999.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"65英寸"},{"key":"内存","value":"32G"}]'),
        (199, 34, '202004190034003', 3999.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"70英寸"},{"key":"内存","value":"16G"}]'),
        (200, 34, '202004190034004', 3999.00, 500, 10, '', 0, 0.00, 0, '[{"key":"尺寸","value":"70英寸"},{"key":"内存","value":"32G"}]'),
        (201, 37, '202210280037001', 5999.00, 195, 0, '/mall/images/20221028/iphone14_001.jpg', 0, 0.00, 1, '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]'),
        (202, 37, '202210280037002', 6899.00, 200, 0, '/mall/images/20221028/iphone14_001.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"256G"}]'),
        (203, 37, '202210280037003', 8699.00, 200, 0, '/mall/images/20221028/iphone14_001.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"512G"}]'),
        (204, 37, '202210280037004', 5999.00, 200, 0, '/mall/images/20221028/iphone14_002.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"星光色"},{"key":"容量","value":"128G"}]'),
        (205, 37, '202210280037005', 6899.00, 200, 0, '/mall/images/20221028/iphone14_002.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"星光色"},{"key":"容量","value":"256G"}]'),
        (206, 37, '202210280037006', 8699.00, 200, 0, '/mall/images/20221028/iphone14_002.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"星光色"},{"key":"容量","value":"512G"}]'),
        (207, 37, '202210280037007', 5999.00, 200, 0, '/mall/images/20221028/iphone14_003.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"紫色"},{"key":"容量","value":"128G"}]'),
        (208, 37, '202210280037008', 6899.00, 200, 0, '/mall/images/20221028/iphone14_003.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"紫色"},{"key":"容量","value":"256G"}]'),
        (209, 37, '202210280037009', 8699.00, 200, 0, '/mall/images/20221028/iphone14_003.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"紫色"},{"key":"容量","value":"512G"}]'),
        (210, 37, '202210280037010', 5999.00, 200, 0, '/mall/images/20221028/iphone14_004.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"128G"}]'),
        (211, 37, '202210280037011', 6899.00, 200, 0, '/mall/images/20221028/iphone14_004.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"256G"}]'),
        (212, 37, '202210280037012', 8699.00, 200, 0, '/mall/images/20221028/iphone14_004.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"512G"}]'),
        (213, 38, '202210280038001', 3599.00, 198, 0, '/mall/images/20221028/ipad_001.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]'),
        (214, 38, '202210280038002', 4799.00, 200, 0, '/mall/images/20221028/ipad_001.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"银色"},{"key":"容量","value":"256G"}]'),
        (215, 38, '202210280038003', 3599.00, 200, 0, '/mall/images/20221028/ipad_002.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"64G"}]'),
        (216, 38, '202210280038004', 4799.00, 200, 0, '/mall/images/20221028/ipad_002.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"256G"}]'),
        (217, 39, '202210280039001', 5999.00, 499, 0, '/mall/images/20221028/xiaomi_computer_001.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"新小米Pro 14英寸 2.8K屏"},{"key":"版本","value":"R7 16G 512"}]'),
        (218, 39, '202210280039002', 5599.00, 500, 0, '/mall/images/20221028/xiaomi_computer_001.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"新小米Pro 14英寸 2.8K屏"},{"key":"版本","value":"R5 16G 512"}]'),
        (219, 39, '202210280039003', 5499.00, 500, 0, '/mall/images/20221028/xiaomi_computer_002.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"新Redmi Pro 15英寸 3.2K屏"},{"key":"版本","value":"R7 16G 512"}]'),
        (220, 39, '202210280039004', 4999.00, 500, 0, '/mall/images/20221028/xiaomi_computer_002.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"新Redmi Pro 15英寸 3.2K屏"},{"key":"版本","value":"R5 16G 512"}]'),
        (221, 40, '202211040040001', 2999.00, 91, 0, '/mall/images/20221104/xiaomi_12_pro_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]'),
        (222, 40, '202211040040002', 3499.00, 100, 0, '/mall/images/20221104/xiaomi_12_pro_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"256G"}]'),
        (223, 40, '202211040040003', 2999.00, 100, 0, '/mall/images/20221104/xiaomi_12_pro_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"128G"}]'),
        (224, 40, '202211040040004', 3499.00, 100, 0, '/mall/images/20221104/xiaomi_12_pro_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"蓝色"},{"key":"容量","value":"256G"}]'),
        (225, 41, '202211040041001', 2099.00, 195, 0, '/mall/images/20221104/redmi_k50_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]'),
        (226, 41, '202211040041002', 2299.00, 200, 0, '/mall/images/20221104/redmi_k50_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"256G"}]'),
        (227, 41, '202211040041003', 2099.00, 200, 0, '/mall/images/20221104/redmi_k50_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"银迹"},{"key":"容量","value":"128G"}]'),
        (228, 41, '202211040041004', 2299.00, 200, 0, '/mall/images/20221104/redmi_k50_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"银迹"},{"key":"容量","value":"256G"}]'),
        (229, 42, '202211040042001', 4999.00, 99, 0, '/mall/images/20221104/huawei_mate50_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"曜金黑"},{"key":"容量","value":"128G"}]'),
        (230, 42, '202211040042002', 5499.00, 100, 0, '/mall/images/20221104/huawei_mate50_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"曜金黑"},{"key":"容量","value":"256G"}]'),
        (231, 42, '202211040042003', 4999.00, 100, 0, '/mall/images/20221104/huawei_mate50_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"冰霜银"},{"key":"容量","value":"128G"}]'),
        (232, 42, '202211040042004', 5499.00, 100, 0, '/mall/images/20221104/huawei_mate50_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"冰霜银"},{"key":"容量","value":"256G"}]'),
        (233, 43, '202211040043001', 1649.00, 500, 0, '/mall/images/20221104/wanhe_13L_01.png', 0, 0.00, 0, '[{"key":"系列","value":"JSQ25-565W13【13升】【恒温旗舰款】"}]'),
        (234, 43, '202211040043002', 1799.00, 500, 0, '/mall/images/20221104/wanhe_13L_01.png', 0, 0.00, 0, '[{"key":"系列","value":"JSQ30-565W16【16升】【恒温旗舰款】"}]'),
        (235, 44, '202211080044001', 369.00, 99, 0, '/mall/images/20221108/sanxing_ssd_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"新品980｜NVMe PCIe3.0*4"},{"key":"版本","value":"512GB"}]'),
        (236, 44, '202211080044002', 649.00, 100, 0, '/mall/images/20221108/sanxing_ssd_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"新品980｜NVMe PCIe3.0*4"},{"key":"版本","value":"1TB"}]'),
        (237, 44, '202211080044003', 549.00, 100, 0, '/mall/images/20221108/sanxing_ssd_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"980 PRO｜NVMe PCIe 4.0"},{"key":"版本","value":"512GB"}]'),
        (238, 44, '202211080044004', 899.00, 100, 0, '/mall/images/20221108/sanxing_ssd_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"980 PRO｜NVMe PCIe 4.0"},{"key":"版本","value":"1TB"}]'),
        (239, 45, '202211080045001', 2299.00, 250, 0, '/mall/images/20221108/oppo_r8_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"鸢尾紫"},{"key":"容量","value":"128G"}]'),
        (240, 45, '202211080045002', 2499.00, 250, 0, '/mall/images/20221108/oppo_r8_01.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"鸢尾紫"},{"key":"容量","value":"256G"}]'),
        (241, 45, '202211080045003', 2299.00, 250, 0, '/mall/images/20221108/oppo_r8_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"晴空蓝"},{"key":"容量","value":"128G"}]'),
        (242, 45, '202211080045004', 2499.00, 250, 0, '/mall/images/20221108/oppo_r8_02.jpg', 0, 0.00, 0, '[{"key":"颜色","value":"晴空蓝"},{"key":"容量","value":"256G"}]');

insert into go_mall_admin.pms_product_category_attribute_relation (id, product_category_id, product_attribute_id)
values  (1, 24, 24),
        (5, 26, 24),
        (7, 28, 24),
        (9, 25, 24),
        (10, 25, 25);


insert into go_mall_admin.cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count, read_count, comment_count, album_pics, description, show_status, content, forward_count, category_name)
values  (1, 1, 'polo衬衫的也时尚', '', 0, 0, UNIX_TIMESTAMP('2018-11-11 13:26:55'), 0, 0, 0, '', '', 0, '', 0, '服装专题'),
        (2, 2, '大牌手机低价秒', '', 0, 0, UNIX_TIMESTAMP('2018-11-12 13:27:00'), 0, 0, 0, '', '', 0, '', 0, '手机专题'),
        (3, 2, '晓龙845新品上市', '', 0, 0, UNIX_TIMESTAMP('2018-11-13 13:27:05'), 0, 0, 0, '', '', 0, '', 0, '手机专题'),
        (4, 1, '夏天应该穿什么', '', 0, 0, UNIX_TIMESTAMP('2018-11-01 13:27:09'), 0, 0, 0, '', '', 0, '', 0, '服装专题'),
        (5, 1, '夏季精选', '', 0, 0, UNIX_TIMESTAMP('2018-11-06 13:27:18'), 0, 0, 0, '', '', 0, '', 0, '服装专题'),
        (6, 2, '品牌手机降价', '', 0, 0, UNIX_TIMESTAMP('2018-11-07 13:27:21'), 0, 0, 0, '', '', 0, '', 0, '手机专题');


insert into go_mall_admin.cms_prefrence_area (id, name, sub_title, pic, sort, show_status)
values  (1, '让音质更出众', '音质不打折 完美现场感', '', 0, 1),
        (2, '让音质更出众22', '让音质更出众22', '', 0, 0),
        (3, '让音质更出众33', '', '', 0, 0),
        (4, '让音质更出众44', '', '', 0, 0);


insert into go_mall_admin.cms_prefrence_area_product_relation (id, prefrence_area_id, product_id)
values  (1, 1, 12),
        (2, 1, 13),
        (3, 1, 14),
        (4, 1, 18),
        (5, 1, 7),
        (6, 2, 7),
        (7, 1, 22),
        (24, 1, 23);


insert into go_mall_admin.pms_product_ladder (id, product_id, count, discount, price)
values  (1, 7, 3, 0.70, 0.00),
        (2, 8, 3, 0.70, 0.00),
        (3, 9, 3, 0.70, 0.00),
        (4, 10, 3, 0.70, 0.00),
        (5, 11, 3, 0.70, 0.00),
        (6, 12, 3, 0.70, 0.00),
        (7, 13, 3, 0.70, 0.00),
        (8, 14, 3, 0.70, 0.00),
        (12, 18, 3, 0.70, 0.00),
        (14, 7, 4, 0.60, 0.00),
        (15, 7, 5, 0.50, 0.00),
        (18, 22, 0, 0.00, 0.00),
        (20, 24, 0, 0.00, 0.00),
        (38, 23, 0, 0.00, 0.00),
        (83, 36, 0, 0.00, 0.00),
        (84, 35, 0, 0.00, 0.00),
        (88, 30, 0, 0.00, 0.00),
        (89, 31, 0, 0.00, 0.00),
        (90, 32, 0, 0.00, 0.00),
        (92, 33, 0, 0.00, 0.00),
        (93, 34, 0, 0.00, 0.00),
        (99, 27, 2, 0.80, 0.00),
        (100, 27, 3, 0.75, 0.00),
        (103, 28, 0, 0.00, 0.00),
        (104, 29, 0, 0.00, 0.00),
        (110, 37, 0, 0.00, 0.00),
        (133, 38, 0, 0.00, 0.00),
        (134, 39, 0, 0.00, 0.00),
        (135, 41, 0, 0.00, 0.00),
        (136, 42, 0, 0.00, 0.00),
        (137, 43, 0, 0.00, 0.00),
        (138, 44, 0, 0.00, 0.00),
        (139, 45, 1, 0.70, 0.00),
        (140, 40, 0, 0.00, 0.00),
        (147, 26, 0, 0.00, 0.00);

insert into go_mall_admin.pms_product_full_reduction (id, product_id, full_price, reduce_price)
values  (1, 7, 100.00, 20.00),
        (2, 8, 100.00, 20.00),
        (3, 9, 100.00, 20.00),
        (4, 10, 100.00, 20.00),
        (5, 11, 100.00, 20.00),
        (6, 12, 100.00, 20.00),
        (7, 13, 100.00, 20.00),
        (8, 14, 100.00, 20.00),
        (9, 18, 100.00, 20.00),
        (10, 7, 200.00, 50.00),
        (11, 7, 300.00, 100.00),
        (14, 22, 0.00, 0.00),
        (16, 24, 0.00, 0.00),
        (34, 23, 0.00, 0.00),
        (78, 36, 0.00, 0.00),
        (79, 35, 0.00, 0.00),
        (83, 30, 0.00, 0.00),
        (84, 31, 0.00, 0.00),
        (85, 32, 0.00, 0.00),
        (87, 33, 0.00, 0.00),
        (88, 34, 0.00, 0.00),
        (93, 27, 0.00, 0.00),
        (96, 28, 500.00, 50.00),
        (97, 28, 1000.00, 120.00),
        (98, 29, 0.00, 0.00),
        (104, 37, 0.00, 0.00),
        (126, 38, 0.00, 0.00),
        (127, 39, 0.00, 0.00),
        (128, 41, 0.00, 0.00),
        (129, 42, 0.00, 0.00),
        (130, 43, 0.00, 0.00),
        (131, 44, 0.00, 0.00),
        (132, 45, 2000.00, 100.00),
        (133, 40, 2000.00, 200.00),
        (146, 26, 3000.00, 300.00),
        (147, 26, 5000.00, 500.00);


insert into go_mall_admin.pms_member_price (id, product_id, member_level_id, member_price, member_level_name)
values  (26, 7, 1, 500.00, ''),
        (27, 8, 1, 500.00, ''),
        (28, 9, 1, 500.00, ''),
        (29, 10, 1, 500.00, ''),
        (30, 11, 1, 500.00, ''),
        (31, 12, 1, 500.00, ''),
        (32, 13, 1, 500.00, ''),
        (33, 14, 1, 500.00, ''),
        (37, 18, 1, 500.00, ''),
        (44, 7, 2, 480.00, ''),
        (45, 7, 3, 450.00, ''),
        (52, 22, 1, 0.00, ''),
        (53, 22, 2, 0.00, ''),
        (54, 22, 3, 0.00, ''),
        (58, 24, 1, 0.00, ''),
        (59, 24, 2, 0.00, ''),
        (60, 24, 3, 0.00, ''),
        (112, 23, 1, 88.00, '黄金会员'),
        (113, 23, 2, 88.00, '白金会员'),
        (114, 23, 3, 66.00, '钻石会员'),
        (246, 36, 1, 0.00, '黄金会员'),
        (247, 36, 2, 0.00, '白金会员'),
        (248, 36, 3, 0.00, '钻石会员'),
        (249, 35, 1, 0.00, '黄金会员'),
        (250, 35, 2, 0.00, '白金会员'),
        (251, 35, 3, 0.00, '钻石会员'),
        (258, 30, 1, 0.00, '黄金会员'),
        (259, 30, 2, 0.00, '白金会员'),
        (260, 30, 3, 0.00, '钻石会员'),
        (261, 31, 1, 0.00, '黄金会员'),
        (262, 31, 2, 0.00, '白金会员'),
        (263, 31, 3, 0.00, '钻石会员'),
        (264, 32, 1, 0.00, '黄金会员'),
        (265, 32, 2, 0.00, '白金会员'),
        (266, 32, 3, 0.00, '钻石会员'),
        (270, 33, 1, 0.00, '黄金会员'),
        (271, 33, 2, 0.00, '白金会员'),
        (272, 33, 3, 0.00, '钻石会员'),
        (273, 34, 1, 0.00, '黄金会员'),
        (274, 34, 2, 0.00, '白金会员'),
        (275, 34, 3, 0.00, '钻石会员'),
        (285, 27, 1, 0.00, '黄金会员'),
        (286, 27, 2, 0.00, '白金会员'),
        (287, 27, 3, 0.00, '钻石会员'),
        (294, 28, 1, 0.00, '黄金会员'),
        (295, 28, 2, 0.00, '白金会员'),
        (296, 28, 3, 0.00, '钻石会员'),
        (297, 29, 1, 0.00, '黄金会员'),
        (298, 29, 2, 0.00, '白金会员'),
        (299, 29, 3, 0.00, '钻石会员'),
        (315, 37, 1, 0.00, '黄金会员'),
        (316, 37, 2, 0.00, '白金会员'),
        (317, 37, 3, 0.00, '钻石会员'),
        (381, 38, 1, 0.00, '黄金会员'),
        (382, 38, 2, 0.00, '白金会员'),
        (383, 38, 3, 0.00, '钻石会员'),
        (384, 39, 1, 0.00, '黄金会员'),
        (385, 39, 2, 0.00, '白金会员'),
        (386, 39, 3, 0.00, '钻石会员'),
        (387, 41, 1, 0.00, '黄金会员'),
        (388, 41, 2, 0.00, '白金会员'),
        (389, 41, 3, 0.00, '钻石会员'),
        (390, 42, 1, 0.00, '黄金会员'),
        (391, 42, 2, 0.00, '白金会员'),
        (392, 42, 3, 0.00, '钻石会员'),
        (393, 43, 1, 0.00, '黄金会员'),
        (394, 43, 2, 0.00, '白金会员'),
        (395, 43, 3, 0.00, '钻石会员'),
        (396, 44, 1, 0.00, '黄金会员'),
        (397, 44, 2, 0.00, '白金会员'),
        (398, 44, 3, 0.00, '钻石会员'),
        (399, 45, 1, 0.00, '黄金会员'),
        (400, 45, 2, 0.00, '白金会员'),
        (401, 45, 3, 0.00, '钻石会员'),
        (402, 40, 1, 0.00, '黄金会员'),
        (403, 40, 2, 0.00, '白金会员'),
        (404, 40, 3, 0.00, '钻石会员'),
        (423, 26, 1, 0.00, '黄金会员'),
        (424, 26, 2, 0.00, '白金会员'),
        (425, 26, 3, 0.00, '钻石会员');


insert into go_mall_admin.cms_subject_product_relation (id, subject_id, product_id)
values  (1, 1, 12),
        (2, 1, 13),
        (3, 1, 14),
        (4, 1, 18),
        (5, 1, 7),
        (6, 2, 7),
        (7, 1, 22),
        (29, 1, 23),
        (30, 4, 23),
        (31, 5, 23),
        (68, 2, 26),
        (69, 3, 26),
        (70, 6, 26);

insert into go_mall_admin.pms_product_vertify_record (id, product_id, create_time, vertify_man, status, detail)
values  (1, 1, UNIX_TIMESTAMP('2018-04-27 16:36:41'), 'test', 1, '验证通过'),
        (2, 2, UNIX_TIMESTAMP('2018-04-27 16:36:41'), 'test', 1, '验证通过');



insert into go_mall_admin.oms_order (id, member_id, coupon_id, order_sn, create_time, member_username, total_amount, pay_amount, freight_amount, promotion_amount, integration_amount, coupon_amount, discount_amount, pay_type, source_type, status, order_type, delivery_company, delivery_sn, auto_confirm_day, integration, growth, promotion_info, bill_type, bill_header, bill_content, bill_receiver_phone, bill_receiver_email, receiver_name, receiver_phone, receiver_post_code, receiver_province, receiver_city, receiver_region, receiver_detail_address, note, confirm_status, delete_status, use_integration, receive_time, comment_time, modify_time, delivery_time, payment_time)
values  (12, 1, 2, '201809150101000001', UNIX_TIMESTAMP('2018-09-15 12:24:27'), 'test', 18732.00, 16377.75, 20.00, 2344.25, 0.00, 10.00, 10.00, 0, 1, 4, 0, '', '', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '江苏省', '常州市', '天宁区', '东晓街道', '111', 0, 0, 0, 0, 0, 1573289428, 0, 0),
        (13, 1, 2, '201809150102000002', UNIX_TIMESTAMP('2018-09-15 14:24:29'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 1, 0, '', '', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 1000, 0, 0, 0, 0, 1539237859),
        (14, 1, 2, '201809130101000001', UNIX_TIMESTAMP('2018-09-13 16:57:40'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 3, 0, '顺丰快递', '201707196398345', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1668154774, 0, 0, 1539668621, 1539409444),
        (15, 1, 2, '201809130102000002', UNIX_TIMESTAMP('2018-09-13 17:03:00'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 3, 0, '顺丰快递', '201707196398346', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 1, 0, 1539842731, 0, 0, 1539668701, 1539409494),
        (16, 1, 2, '201809140101000001', UNIX_TIMESTAMP('2018-09-14 16:16:16'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 4, 0, '', '', 15, 13284, 13284, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (17, 1, 2, '201809150101000003', UNIX_TIMESTAMP('2018-09-15 12:24:27'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 0, 1, 4, 0, '顺丰快递', '201707196398345', 15, 0, 0, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 1539324088, 0),
        (18, 1, 2, '201809150102000004', UNIX_TIMESTAMP('2018-09-15 14:24:29'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 1, 0, '圆通快递', 'xx', 15, 0, 0, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 1000, 0, 0, 0, 1539672137, 0),
        (19, 1, 2, '201809130101000003', UNIX_TIMESTAMP('2018-09-13 16:57:40'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 2, 0, '', '', 15, 0, 0, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (20, 1, 2, '201809130102000004', UNIX_TIMESTAMP('2018-09-13 17:03:00'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 3, 0, '', '', 15, 0, 0, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (21, 1, 2, '201809140101000002', UNIX_TIMESTAMP('2018-09-14 16:16:16'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 4, 0, '', '', 15, 18682, 18682, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (22, 1, 2, '201809150101000005', UNIX_TIMESTAMP('2018-09-15 12:24:27'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 0, 1, 4, 0, '顺丰快递', '201707196398345', 15, 0, 0, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 1539324088, 0),
        (23, 1, 2, '201809150102000006', UNIX_TIMESTAMP('2018-09-15 14:24:29'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 1, 1, 1, 0, '顺丰快递', 'xxx', 15, 0, 0, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 1000, 0, 0, 0, 1539672088, 0),
        (24, 1, 2, '201809130101000005', UNIX_TIMESTAMP('2018-09-13 16:57:40'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 2, 0, '', '', 15, 18682, 18682, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (25, 1, 2, '201809130102000006', UNIX_TIMESTAMP('2018-09-13 17:03:00'), 'test', 18732.00, 16377.75, 10.00, 2344.25, 0.00, 10.00, 5.00, 1, 1, 4, 0, '', '', 15, 18682, 18682, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨22', '18033441849', '518000', '北京市', '北京城区', '东城区', '东城街道', 'xxx', 0, 1, 0, 0, 0, 1540883311, 0, 0),
        (26, 1, 2, '201809140101000003', UNIX_TIMESTAMP('2018-09-14 16:16:16'), 'test', 18732.00, 16377.75, 0.00, 2344.25, 0.00, 10.00, 0.00, 2, 1, 4, 0, '', '', 15, 18682, 18682, '单品促销,打折优惠：满3件，打7.50折,满减优惠：满1000.00元，减120.00元,满减优惠：满1000.00元，减120.00元,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (27, 1, 0, '202002250100000001', UNIX_TIMESTAMP('2020-02-25 15:59:20'), 'test', 540.00, 540.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '无优惠,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (28, 1, 0, '202002250100000002', UNIX_TIMESTAMP('2020-02-25 16:05:47'), 'test', 540.00, 540.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '无优惠,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (29, 1, 0, '202002250100000003', UNIX_TIMESTAMP('2020-02-25 16:07:58'), 'test', 540.00, 540.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '无优惠,无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (30, 1, 0, '202002250100000004', UNIX_TIMESTAMP('2020-02-25 16:50:13'), 'test', 240.00, 240.00, 20.00, 0.00, 0.00, 0.00, 10.00, 0, 1, 3, 0, '顺丰快递', '12333333', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '南山区', '科兴科学园', '', 1, 0, 0, 1589715495, 0, 1582620771, 1582620843, 1582620809),
        (31, 1, 26, '202005160100000001', UNIX_TIMESTAMP('2020-05-16 15:16:54'), 'test', 13623.00, 11842.40, 0.00, 1629.60, 1.00, 150.00, 0.00, 0, 1, 4, 0, '', '', 15, 13623, 13623, '满减优惠：满5000.00元，减500.00元;打折优惠：满2件，打8.00折;满减优惠：满500.00元，减50.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (32, 1, 0, '202005170100000001', UNIX_TIMESTAMP('2020-05-17 15:00:38'), 'test', 6487.00, 6187.00, 0.00, 300.00, 0.00, 0.00, 0.00, 1, 1, 1, 0, '', '', 15, 6487, 6487, '满减优惠：满3000.00元，减300.00元;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1589700808),
        (33, 1, 0, '202005170100000002', UNIX_TIMESTAMP('2020-05-17 15:14:18'), 'test', 3788.00, 3488.00, 0.00, 300.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (34, 1, 0, '202005170100000003', UNIX_TIMESTAMP('2020-05-17 15:20:10'), 'test', 3788.00, 3488.00, 0.00, 300.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 1, 0, 0, 0, 0, 0, 0),
        (35, 1, 0, '202005170100000004', UNIX_TIMESTAMP('2020-05-17 15:22:03'), 'test', 3788.00, 3488.00, 0.00, 300.00, 0.00, 0.00, 0.00, 2, 1, 3, 0, '顺丰快递', '123', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1589701305, 0, 0, 1589700624, 1589700547),
        (36, 1, 0, '202005170100000005', UNIX_TIMESTAMP('2020-05-17 16:59:26'), 'test', 10275.00, 9775.00, 0.00, 500.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 10275, 10275, '满减优惠：满5000.00元，减500.00元;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (37, 1, 0, '202005170100000006', UNIX_TIMESTAMP('2020-05-17 19:33:48'), 'test', 6487.00, 6187.00, 0.00, 300.00, 0.00, 0.00, 0.00, 1, 1, 3, 0, '顺丰快递', 'aadd', 15, 6487, 6487, '满减优惠：满3000.00元，减300.00元;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1589715350, 0, 0, 1589715299, 1589715239),
        (38, 1, 0, '202005170100000007', UNIX_TIMESTAMP('2020-05-17 19:39:10'), 'test', 3788.00, 3488.00, 0.00, 300.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (39, 1, 0, '202005170100000008', UNIX_TIMESTAMP('2020-05-17 19:41:30'), 'test', 3788.00, 3488.00, 0.00, 300.00, 0.00, 0.00, 0.00, 1, 1, 3, 0, '顺丰快递', 'sdf', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 1, 0, 1589715756, 0, 0, 1589715727, 1589715701),
        (40, 1, 0, '202005180100000001', UNIX_TIMESTAMP('2020-05-18 16:50:03'), 'test', 3788.00, 3488.00, 0.00, 300.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1589791829),
        (41, 1, 26, '202005180100000002', UNIX_TIMESTAMP('2020-05-18 20:22:24'), 'test', 6487.00, 6037.00, 0.00, 300.00, 0.00, 150.00, 0.00, 1, 1, 3, 0, '顺丰快递', '12313', 15, 6487, 6487, '满减优惠：满3000.00元，减300.00元;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '清水河街道', '', 1, 0, 0, 1589804600, 0, 0, 1589804583, 1589804549),
        (42, 1, 0, '202005230100000001', UNIX_TIMESTAMP('2020-05-23 16:21:27'), 'test', 5398.00, 4318.40, 0.00, 1079.60, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 5398, 5398, '打折优惠：满2件，打8.00折', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '清水河街道', '', 0, 0, 0, 0, 0, 0, 0, 1590222090),
        (43, 1, 0, '202005230100000002', UNIX_TIMESTAMP('2020-05-23 17:01:33'), 'test', 5398.00, 4318.40, 0.00, 1079.60, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 5398, 5398, '打折优惠：满2件，打8.00折', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (44, 1, 0, '202005240100000001', UNIX_TIMESTAMP('2020-05-24 09:37:07'), 'test', 7576.00, 7076.00, 0.00, 500.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 7576, 7576, '满减优惠：满5000.00元，减500.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (45, 1, 25, '202006070100000001', UNIX_TIMESTAMP('2020-06-07 17:02:04'), 'test', 10275.00, 9674.90, 0.00, 500.00, 0.00, 100.10, 0.00, 1, 1, 1, 0, '', '', 15, 10275, 10275, '满减优惠：满5000.00元，减500.00元;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '清水河街道', '', 0, 0, 0, 0, 0, 0, 0, 1591520537),
        (46, 1, 24, '202006210100000001', UNIX_TIMESTAMP('2020-06-21 14:27:34'), 'test', 9186.00, 7796.40, 0.00, 1379.60, 0.00, 10.00, 0.00, 2, 1, 1, 0, '', '', 15, 9186, 9186, '满减优惠：满3000.00元，减300.00元;打折优惠：满2件，打8.00折', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1592720858),
        (47, 1, 0, '202006210100000002', UNIX_TIMESTAMP('2020-06-21 15:13:06'), 'test', 6487.00, 6187.00, 0.00, 300.00, 0.00, 0.00, 0.00, 1, 1, 3, 0, '顺丰快递', '123131', 15, 6487, 6487, '满减优惠：满3000.00元，减300.00元;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '清水河街道', '', 1, 0, 0, 1592723638, 0, 0, 1592723624, 1592723592),
        (48, 1, 26, '202006210100000003', UNIX_TIMESTAMP('2020-06-21 15:15:18'), 'test', 3788.00, 3338.00, 0.00, 300.00, 0.00, 150.00, 0.00, 2, 1, 3, 0, '圆通快递', '12313', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1592723758, 0, 0, 1592723748, 1592723720),
        (49, 1, 0, '202006270100000001', UNIX_TIMESTAMP('2020-06-27 10:27:56'), 'test', 2699.00, 2699.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 2699, 2699, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '清水河街道', '', 0, 0, 0, 0, 0, 0, 0, 1593224878),
        (50, 1, 0, '202210280100000001', UNIX_TIMESTAMP('2022-10-28 14:50:58'), 'test', 2699.00, 2699.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 2699, 2699, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '清水河街道', '', 0, 0, 0, 0, 0, 0, 0, 1666939862),
        (51, 1, 0, '202210280100000002', UNIX_TIMESTAMP('2022-10-28 15:27:41'), 'test', 5999.00, 5999.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1666942064),
        (52, 1, 30, '202211090100000001', UNIX_TIMESTAMP('2022-11-09 15:14:58'), 'test', 2999.00, 2799.00, 0.00, 0.00, 0.00, 200.00, 0.00, 2, 1, 3, 0, '顺丰快递', '1233', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1667978191, 0, 0, 1667978172, 1667978100),
        (53, 1, 27, '202211090100000002', UNIX_TIMESTAMP('2022-11-09 15:25:38'), 'test', 3599.00, 3589.00, 0.00, 0.00, 0.00, 10.00, 0.00, 2, 1, 1, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1667978741),
        (54, 1, 29, '202211090100000003', UNIX_TIMESTAMP('2022-11-09 15:26:11'), 'test', 5999.00, 5399.00, 0.00, 0.00, 0.00, 600.00, 0.00, 2, 1, 1, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1667978773),
        (55, 1, 0, '202211100100000001', UNIX_TIMESTAMP('2022-11-10 16:57:59'), 'test', 11998.00, 11998.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (56, 1, 28, '202211110100000001', UNIX_TIMESTAMP('2022-11-11 16:12:42'), 'test', 2999.00, 2899.00, 0.00, 0.00, 0.00, 100.00, 0.00, 2, 1, 1, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1668154368),
        (57, 1, 0, '202211110100000002', UNIX_TIMESTAMP('2022-11-11 16:13:14'), 'test', 2999.00, 2999.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1668154401),
        (58, 1, 0, '202211110100000003', UNIX_TIMESTAMP('2022-11-11 16:15:08'), 'test', 5999.00, 5999.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1668154666),
        (59, 1, 0, '202211110100000004', UNIX_TIMESTAMP('2022-11-11 16:21:12'), 'test', 649.00, 599.00, 0.00, 50.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 649, 649, '满减优惠：满500.00元，减50.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (60, 1, 0, '202211160100000001', UNIX_TIMESTAMP('2022-11-16 10:36:08'), 'test', 11097.00, 11097.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 3, 0, '顺丰快递', '1234555', 15, 0, 0, '无优惠;无优惠;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1668566680, 0, 0, 1668566570, 1668566245),
        (61, 1, 0, '202212210100000001', UNIX_TIMESTAMP('2022-12-21 15:49:08'), 'test', 2999.00, 2999.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (62, 1, 0, '202212210100000002', UNIX_TIMESTAMP('2022-12-21 15:49:57'), 'test', 8098.00, 8098.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 3, 0, '顺丰快递', 'SDFERR7845', 15, 0, 0, '无优惠;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1671609033, 0, 0, 1671609023, 1671609000),
        (63, 1, 0, '202212210100000003', UNIX_TIMESTAMP('2022-12-21 15:51:09'), 'test', 2999.00, 2999.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 2, 0, '顺丰快递', '112', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 1673316852, 1671609071),
        (64, 1, 0, '202212210100000004', UNIX_TIMESTAMP('2022-12-21 15:51:35'), 'test', 2099.00, 2099.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (65, 1, 28, '202212210100000005', UNIX_TIMESTAMP('2022-12-21 16:53:07'), 'test', 5098.00, 4788.00, 0.00, 200.00, 10.00, 100.00, 0.00, 2, 1, 2, 0, '圆通快递', '115', 15, 0, 0, '满减优惠：满2000.00元，减200.00元;无优惠', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 1673316852, 1671612838),
        (66, 1, 0, '202301100100000001', UNIX_TIMESTAMP('2023-01-10 15:34:59'), 'test', 5998.00, 5798.00, 0.00, 200.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '满减优惠：满2000.00元，减200.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (67, 1, 0, '202301100100000002', UNIX_TIMESTAMP('2023-01-10 15:39:07'), 'test', 3788.00, 3488.00, 0.00, 300.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 3788, 3788, '满减优惠：满3000.00元，减300.00元', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1673336356),
        (68, 1, 0, '202301100100000003', UNIX_TIMESTAMP('2023-01-10 16:58:19'), 'test', 3999.00, 3899.00, 0.00, 100.00, 0.00, 0.00, 0.00, 2, 1, 1, 0, '', '', 15, 3788, 3788, '单品促销', 0, '', '', '', '', '大梨', '18033441849', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 1673341101),
        (69, 11, 30, '202305110100000001', UNIX_TIMESTAMP('2023-05-11 15:28:56'), 'member', 5098.00, 4698.00, 0.00, 200.00, 0.00, 200.00, 0.00, 2, 1, 3, 0, '顺丰快递', '1231313', 15, 0, 0, '满减优惠：满2000.00元，减200.00元;无优惠', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1683790216, 0, 0, 1683790208, 1683790139),
        (70, 11, 0, '202305110100000002', UNIX_TIMESTAMP('2023-05-11 15:30:36'), 'member', 3599.00, 3599.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 3, 0, '顺丰快递', '232342', 15, 0, 0, '无优惠', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1683790290, 0, 0, 1683790282, 1683790240),
        (71, 11, 0, '202305110100000003', UNIX_TIMESTAMP('2023-05-11 15:31:55'), 'member', 5999.00, 5999.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 4, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (72, 11, 0, '202305110100000004', UNIX_TIMESTAMP('2023-05-11 15:33:13'), 'member', 5368.00, 5368.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 2, 0, '圆通快递', '1231434', 15, 0, 0, '无优惠;无优惠', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 1683790423, 1683790401),
        (73, 11, 0, '202305110100000005', UNIX_TIMESTAMP('2023-05-11 15:34:39'), 'member', 5999.00, 5999.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 1, 0, 0, '', '', 15, 0, 0, '无优惠', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 0, 0),
        (74, 11, 0, '202305110100000006', UNIX_TIMESTAMP('2023-05-11 15:35:05'), 'member', 2999.00, 2799.00, 0.00, 200.00, 0.00, 0.00, 0.00, 2, 1, 2, 0, '顺丰快递', '123131', 15, 0, 0, '满减优惠：满2000.00元，减200.00元', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 0, 0, 0, 0, 0, 0, 1683790560, 1683790508),
        (75, 11, 0, '202305110100000007', UNIX_TIMESTAMP('2023-05-11 15:35:24'), 'member', 2099.00, 2099.00, 0.00, 0.00, 0.00, 0.00, 0.00, 2, 1, 3, 0, '顺丰快递', '123131311', 15, 0, 0, '无优惠', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1683790594, 0, 0, 1683790571, 1683790526),
        (76, 11, 28, '202305110100000008', UNIX_TIMESTAMP('2023-05-11 15:37:16'), 'member', 8998.00, 8698.00, 0.00, 200.00, 0.00, 100.00, 0.00, 2, 1, 3, 0, '顺丰快递', '1231313', 15, 0, 0, '无优惠;满减优惠：满2000.00元，减200.00元', 0, '', '', '', '', '小李', '18961511111', '518000', '广东省', '深圳市', '福田区', '东晓街道', '', 1, 0, 0, 1683790668, 0, 0, 1683790653, 1683790638);


INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (21, 12, '201809150101000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (22, 12, '201809150101000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788', 2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (23, 12, '201809150101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (24, 12, '201809150101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (25, 12, '201809150101000001', 29, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', 'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00, 5496.06, 5499, 5499, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (26, 13, '201809150102000002', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (27, 13, '201809150102000002', 27, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788', 2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (28, 13, '201809150102000002', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (29, 13, '201809150102000002', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (30, 13, '201809150102000002', 29, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', 'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00, 5496.06, 5499, 5499, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (31, 14, '201809130101000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (32, 14, '201809130101000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788', 2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (33, 14, '201809130101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (34, 14, '201809130101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (35, 14, '201809130101000001', 29, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', 'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00, 5496.06, 5499, 5499, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (36, 15, '201809130101000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (37, 15, '201809130101000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788', 2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (38, 15, '201809130101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (39, 15, '201809130101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (40, 15, '201809130101000001', 29, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', 'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00, 5496.06, 5499, 5499, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (41, 16, '201809140101000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '6946605', 3788.00, 1, 90, '201806070026001', 19, '单品促销', 200.00, 2.02, 0.00, 3585.98, 3788, 3788, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (42, 16, '201809140101000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '7437788', 2699.00, 3, 98, '201808270027001', 19, '打折优惠：满3件，打7.50折', 674.75, 1.44, 0.00, 2022.81, 2699, 2699, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (43, 16, '201809140101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 57.60, 0.35, 0.00, 591.05, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (44, 16, '201809140101000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '7437789', 699.00, 1, 103, '201808270028001', 19, '满减优惠：满1000.00元，减120.00元', 62.40, 0.37, 0.00, 636.23, 649, 649, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (45, 16, '201809140101000001', 29, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', 'Apple iPhone 8 Plus', '苹果', '7437799', 5499.00, 1, 106, '201808270029001', 19, '无优惠', 0.00, 2.94, 0.00, 5496.06, 5499, 5499, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (46, 27, '202002250100000001', 36, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 100.00, 3, 163, '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 100.00, 0, 0, '');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (47, 27, '202002250100000001', 36, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164, '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0, '');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (48, 28, '202002250100000002', 36, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 100.00, 3, 163, '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 100.00, 0, 0, '');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (49, 28, '202002250100000002', 36, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164, '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0, '');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (50, 29, '202002250100000003', 36, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 100.00, 3, 163, '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 100.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"红色\"},{\"key\":\"尺寸\",\"value\":\"38\"},{\"key\":\"风格\",\"value\":\"秋季\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (51, 29, '202002250100000003', 36, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164, '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"红色\"},{\"key\":\"尺寸\",\"value\":\"38\"},{\"key\":\"风格\",\"value\":\"夏季\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (52, 30, '202002250100000004', 36, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'NIKE', '6799345', 120.00, 2, 164, '202002210036001', 29, '无优惠', 0.00, 0.00, 0.00, 120.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"红色\"},{\"key\":\"尺寸\",\"value\":\"38\"},{\"key\":\"风格\",\"value\":\"夏季\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (53, 31, '202005160100000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 2, 110, '201806070026001', 19, '满减优惠：满5000.00元，减500.00元', 250.00, 75.00, 0.28, 3462.72, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (54, 31, '202005160100000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 2, 98, '201808270027001', 19, '打折优惠：满2件，打8.00折', 539.80, 0.00, 0.20, 2159.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (55, 31, '202005160100000001', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', '小米', '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满500.00元，减50.00元', 50.00, 0.00, 0.05, 598.95, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (56, 32, '202005170100000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (57, 32, '202005170100000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 0.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (58, 33, '202005170100000002', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (59, 34, '202005170100000003', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (60, 35, '202005170100000004', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (61, 36, '202005170100000005', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 2, 110, '201806070026001', 19, '满减优惠：满5000.00元，减500.00元', 250.00, 0.00, 0.00, 3538.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (62, 36, '202005170100000005', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 0.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (63, 37, '202005170100000006', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (64, 37, '202005170100000006', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 0.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (65, 38, '202005170100000007', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (66, 39, '202005170100000008', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (67, 40, '202005180100000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (68, 41, '202005180100000002', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 150.00, 0.00, 3338.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (69, 41, '202005180100000002', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 0.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (70, 42, '202005230100000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 2, 98, '201808270027001', 19, '打折优惠：满2件，打8.00折', 539.80, 0.00, 0.00, 2159.20, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (71, 43, '202005230100000002', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 2, 98, '201808270027001', 19, '打折优惠：满2件，打8.00折', 539.80, 0.00, 0.00, 2159.20, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (72, 44, '202005240100000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 2, 110, '201806070026001', 19, '满减优惠：满5000.00元，减500.00元', 250.00, 0.00, 0.00, 3538.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (73, 45, '202006070100000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 2, 110, '201806070026001', 19, '满减优惠：满5000.00元，减500.00元', 250.00, 36.90, 0.00, 3501.10, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (74, 45, '202006070100000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 26.30, 0.00, 2672.70, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (75, 46, '202006210100000001', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 4.12, 0.00, 3483.88, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (76, 46, '202006210100000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 2, 98, '201808270027001', 19, '打折优惠：满2件，打8.00折', 539.80, 2.94, 0.00, 2156.26, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (77, 47, '202006210100000002', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (78, 47, '202006210100000002', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 0.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (79, 48, '202006210100000003', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 150.00, 0.00, 3338.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (80, 49, '202006270100000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 0.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (81, 50, '202210280100000001', 27, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '小米', '7437788', 2699.00, 1, 98, '201808270027001', 19, '无优惠', 0.00, 0.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (82, 51, '202210280100000002', 37, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 紫色 支持移动联通电信5G 双卡双待手机', '苹果', '100038005189', 5999.00, 1, 201, '202210280037001', 19, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"午夜色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (83, 52, '202211090100000001', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '无优惠', 0.00, 200.00, 0.00, 2799.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (84, 53, '202211090100000002', 38, '/mall/images/20221028/ipad_001.jpg', 'Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）', '苹果', '100044025833', 3599.00, 1, 213, '202210280038001', 53, '无优惠', 0.00, 10.00, 0.00, 3589.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"银色\"},{\"key\":\"容量\",\"value\":\"64G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (85, 54, '202211090100000003', 37, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '苹果', '100038005189', 5999.00, 1, 201, '202210280037001', 19, '无优惠', 0.00, 600.00, 0.00, 5399.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"午夜色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (86, 55, '202211100100000001', 37, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '苹果', '100038005189', 5999.00, 2, 201, '202210280037001', 19, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"午夜色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (87, 56, '202211110100000001', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '无优惠', 0.00, 100.00, 0.00, 2899.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (88, 57, '202211110100000002', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '无优惠', 0.00, 0.00, 0.00, 2999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (89, 58, '202211110100000003', 37, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '苹果', '100038005189', 5999.00, 1, 201, '202210280037001', 19, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"午夜色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (90, 59, '202211110100000004', 28, '/mall/images/20180615/5a9d248cN071f4959.jpg', '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', '小米', '7437789', 649.00, 1, 102, '201808270028001', 19, '满减优惠：满500.00元，减50.00元', 50.00, 0.00, 0.00, 599.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (91, 60, '202211160100000001', 37, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '苹果', '100038005189', 5999.00, 1, 201, '202210280037001', 19, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"午夜色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (92, 60, '202211160100000001', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '无优惠', 0.00, 0.00, 0.00, 2999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (93, 60, '202211160100000001', 41, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '小米', '100035246702', 2099.00, 1, 225, '202211040041001', 19, '无优惠', 0.00, 0.00, 0.00, 2099.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"墨羽\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (94, 61, '202212210100000001', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '无优惠', 0.00, 0.00, 0.00, 2999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (95, 62, '202212210100000002', 37, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '苹果', '100038005189', 5999.00, 1, 201, '202210280037001', 19, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"午夜色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (96, 62, '202212210100000002', 41, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '小米', '100035246702', 2099.00, 1, 225, '202211040041001', 19, '无优惠', 0.00, 0.00, 0.00, 2099.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"墨羽\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (97, 63, '202212210100000003', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '无优惠', 0.00, 0.00, 0.00, 2999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (98, 64, '202212210100000004', 41, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '小米', '100035246702', 2099.00, 1, 225, '202211040041001', 19, '无优惠', 0.00, 0.00, 0.00, 2099.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"墨羽\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (99, 65, '202212210100000005', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '满减优惠：满2000.00元，减200.00元', 200.00, 58.80, 5.88, 2734.32, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (100, 65, '202212210100000005', 41, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '小米', '100035246702', 2099.00, 1, 225, '202211040041001', 19, '无优惠', 0.00, 41.20, 4.12, 2053.68, 0, 0, '[{\"key\":\"颜色\",\"value\":\"墨羽\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (101, 66, '202301100100000001', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 2, 221, '202211040040001', 19, '满减优惠：满2000.00元，减200.00元', 100.00, 0.00, 0.00, 2899.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (102, 67, '202301100100000002', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3788.00, 1, 110, '201806070026001', 19, '满减优惠：满3000.00元，减300.00元', 300.00, 0.00, 0.00, 3488.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"16G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (103, 68, '202301100100000003', 26, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', '华为', '6946605', 3999.00, 1, 111, '201806070026002', 19, '单品促销', 100.00, 0.00, 0.00, 3899.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"金色\"},{\"key\":\"容量\",\"value\":\"32G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (104, 69, '202305110100000001', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '满减优惠：满2000.00元，减200.00元', 200.00, 117.60, 0.00, 2681.40, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (105, 69, '202305110100000001', 41, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '小米', '100035246702', 2099.00, 1, 225, '202211040041001', 19, '无优惠', 0.00, 82.40, 0.00, 2016.60, 0, 0, '[{\"key\":\"颜色\",\"value\":\"墨羽\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (106, 70, '202305110100000002', 38, '/mall/images/20221028/ipad_001.jpg', 'Apple iPad 10.9英寸平板电脑 2022年款', '苹果', '100044025833', 3599.00, 1, 213, '202210280038001', 53, '无优惠', 0.00, 0.00, 0.00, 3599.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"银色\"},{\"key\":\"容量\",\"value\":\"64G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (107, 71, '202305110100000003', 39, '/mall/images/20221028/xiaomi_computer_001.jpg', '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑', '小米', '100023207945', 5999.00, 1, 217, '202210280039001', 54, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"新小米Pro 14英寸 2.8K屏\"},{\"key\":\"版本\",\"value\":\"R7 16G 512\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (108, 72, '202305110100000004', 42, '/mall/images/20221104/huawei_mate50_01.jpg', 'HUAWEI Mate 50 直屏旗舰 超光变XMAGE影像 北斗卫星消息', '华为', '100035295081', 4999.00, 1, 229, '202211040042001', 19, '无优惠', 0.00, 0.00, 0.00, 4999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"曜金黑\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (109, 72, '202305110100000004', 44, '/mall/images/20221108/sanxing_ssd_02.jpg', '三星（SAMSUNG）500GB SSD固态硬盘 M.2接口(NVMe协议)', '三星', '100018768480', 369.00, 1, 235, '202211080044001', 55, '无优惠', 0.00, 0.00, 0.00, 369.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"新品980｜NVMe PCIe3.0*4\"},{\"key\":\"版本\",\"value\":\"512GB\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (110, 73, '202305110100000005', 37, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '苹果', '100038005189', 5999.00, 1, 201, '202210280037001', 19, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"午夜色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (111, 74, '202305110100000006', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '满减优惠：满2000.00元，减200.00元', 200.00, 0.00, 0.00, 2799.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (112, 75, '202305110100000007', 41, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '小米', '100035246702', 2099.00, 1, 225, '202211040041001', 19, '无优惠', 0.00, 0.00, 0.00, 2099.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"墨羽\"},{\"key\":\"容量\",\"value\":\"128G\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (113, 76, '202305110100000008', 39, '/mall/images/20221028/xiaomi_computer_001.jpg', '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑', '小米', '100023207945', 5999.00, 1, 217, '202210280039001', 54, '无优惠', 0.00, 0.00, 0.00, 5999.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"新小米Pro 14英寸 2.8K屏\"},{\"key\":\"版本\",\"value\":\"R7 16G 512\"}]');
INSERT INTO `oms_order_item`(id, order_id, order_sn, product_id, product_pic, product_name, product_brand, product_sn, product_price, product_quantity, product_sku_id, product_sku_code, product_category_id, promotion_name, promotion_amount, coupon_amount, integration_amount, real_amount, gift_integration, gift_growth, product_attr) VALUES (114, 76, '202305110100000008', 40, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '小米', '100027789721', 2999.00, 1, 221, '202211040040001', 19, '满减优惠：满2000.00元，减200.00元', 200.00, 100.00, 0.00, 2699.00, 0, 0, '[{\"key\":\"颜色\",\"value\":\"黑色\"},{\"key\":\"容量\",\"value\":\"128G\"}]');




insert into go_mall_admin.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note)
values  (5, 12, '后台管理员', UNIX_TIMESTAMP('2018-10-12 14:01:29'), 2, '完成发货'),
        (6, 13, '后台管理员', UNIX_TIMESTAMP('2018-10-12 14:01:29'), 2, '完成发货'),
        (7, 12, '后台管理员', UNIX_TIMESTAMP('2018-10-12 14:13:10'), 4, '订单关闭:买家退货'),
        (8, 13, '后台管理员', UNIX_TIMESTAMP('2018-10-12 14:13:10'), 4, '订单关闭:买家退货'),
        (9, 22, '后台管理员', UNIX_TIMESTAMP('2018-10-15 16:31:48'), 4, '订单关闭:xxx'),
        (10, 22, '后台管理员', UNIX_TIMESTAMP('2018-10-15 16:35:08'), 4, '订单关闭:xxx'),
        (11, 22, '后台管理员', UNIX_TIMESTAMP('2018-10-15 16:35:59'), 4, '订单关闭:xxx'),
        (12, 17, '后台管理员', UNIX_TIMESTAMP('2018-10-15 16:43:40'), 4, '订单关闭:xxx'),
        (13, 25, '后台管理员', UNIX_TIMESTAMP('2018-10-15 16:52:14'), 4, '订单关闭:xxx'),
        (14, 26, '后台管理员', UNIX_TIMESTAMP('2018-10-15 16:52:14'), 4, '订单关闭:xxx'),
        (15, 23, '后台管理员', UNIX_TIMESTAMP('2018-10-16 14:41:28'), 2, '完成发货'),
        (16, 13, '后台管理员', UNIX_TIMESTAMP('2018-10-16 14:42:17'), 2, '完成发货'),
        (17, 18, '后台管理员', UNIX_TIMESTAMP('2018-10-16 14:42:17'), 2, '完成发货'),
        (18, 26, '后台管理员', UNIX_TIMESTAMP('2018-10-30 14:37:44'), 4, '订单关闭:关闭订单'),
        (19, 25, '后台管理员', UNIX_TIMESTAMP('2018-10-30 15:07:01'), 0, '修改收货人信息'),
        (20, 25, '后台管理员', UNIX_TIMESTAMP('2018-10-30 15:08:13'), 0, '修改费用信息'),
        (21, 25, '后台管理员', UNIX_TIMESTAMP('2018-10-30 15:08:31'), 0, '修改备注信息：xxx'),
        (22, 25, '后台管理员', UNIX_TIMESTAMP('2018-10-30 15:08:39'), 4, '订单关闭:2222'),
        (23, 12, '后台管理员', UNIX_TIMESTAMP('2019-11-09 16:50:28'), 4, '修改备注信息：111'),
        (24, 30, '后台管理员', UNIX_TIMESTAMP('2020-02-25 16:52:37'), 0, '修改费用信息'),
        (25, 30, '后台管理员', UNIX_TIMESTAMP('2020-02-25 16:52:51'), 0, '修改费用信息'),
        (26, 30, '后台管理员', UNIX_TIMESTAMP('2020-02-25 16:54:03'), 2, '完成发货'),
        (27, 35, '后台管理员', UNIX_TIMESTAMP('2020-05-17 15:30:24'), 2, '完成发货'),
        (28, 37, '后台管理员', UNIX_TIMESTAMP('2020-05-17 19:35:00'), 2, '完成发货'),
        (29, 39, '后台管理员', UNIX_TIMESTAMP('2020-05-17 19:42:08'), 2, '完成发货'),
        (30, 41, '后台管理员', UNIX_TIMESTAMP('2020-05-18 20:23:04'), 2, '完成发货'),
        (31, 47, '后台管理员', UNIX_TIMESTAMP('2020-06-21 15:13:44'), 2, '完成发货'),
        (32, 48, '后台管理员', UNIX_TIMESTAMP('2020-06-21 15:15:49'), 2, '完成发货'),
        (33, 52, '后台管理员', UNIX_TIMESTAMP('2022-11-09 15:16:13'), 2, '完成发货'),
        (34, 60, '后台管理员', UNIX_TIMESTAMP('2022-11-16 10:42:50'), 2, '完成发货'),
        (35, 62, '后台管理员', UNIX_TIMESTAMP('2022-12-21 15:50:24'), 2, '完成发货'),
        (36, 63, '后台管理员', UNIX_TIMESTAMP('2023-01-10 10:08:34'), 2, '完成发货'),
        (37, 65, '后台管理员', UNIX_TIMESTAMP('2023-01-10 10:08:34'), 2, '完成发货'),
        (38, 69, '后台管理员', UNIX_TIMESTAMP('2023-05-11 15:30:08'), 2, '完成发货'),
        (39, 70, '后台管理员', UNIX_TIMESTAMP('2023-05-11 15:31:22'), 2, '完成发货'),
        (40, 72, '后台管理员', UNIX_TIMESTAMP('2023-05-11 15:33:43'), 2, '完成发货'),
        (41, 74, '后台管理员', UNIX_TIMESTAMP('2023-05-11 15:36:00'), 2, '完成发货'),
        (42, 75, '后台管理员', UNIX_TIMESTAMP('2023-05-11 15:36:11'), 2, '完成发货'),
        (43, 76, '后台管理员', UNIX_TIMESTAMP('2023-05-11 15:37:34'), 2, '完成发货');

insert into go_mall_admin.oms_order_setting (id, flash_order_overtime, normal_order_overtime, confirm_overtime, finish_overtime, comment_overtime)
values  (1, 60, 120, 15, 7, 7);

insert into go_mall_admin.oms_order_return_reason (id, name, sort, status, create_time)
values  (1, '质量问题', 1, 1, UNIX_TIMESTAMP('2018-10-17 10:00:45')),
        (2, '尺码太大', 1, 1, UNIX_TIMESTAMP('2018-10-17 10:01:03')),
        (3, '颜色不喜欢', 1, 1, UNIX_TIMESTAMP('2018-10-17 10:01:13')),
        (4, '7天无理由退货', 1, 1, UNIX_TIMESTAMP('2018-10-17 10:01:47')),
        (5, '价格问题', 1, 0, UNIX_TIMESTAMP('2018-10-17 10:01:57')),
        (12, '发票问题', 0, 1, UNIX_TIMESTAMP('2018-10-19 16:28:36')),
        (13, '其他问题', 0, 1, UNIX_TIMESTAMP('2018-10-19 16:28:51')),
        (14, '物流问题', 0, 1, UNIX_TIMESTAMP('2018-10-19 16:29:01')),
        (15, '售后问题', 0, 1, UNIX_TIMESTAMP('2018-10-19 16:29:11'));

insert into go_mall_admin.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note, created_at, updated_at)
values  (3, 12, 1, 26, '201809150101000001', 1539758097, 'test', 0.00, '大梨', '18000000000', 2, 1668132978, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '["/mall/images/20180607/5ac1bf58Ndefaac16.jpg","/mall/images/20180615/xiaomi.jpg"]', '111', 'admin', 'admin', 1668132986, '', 0, 0),
        (4, 12, 2, 27, '201809150101000001', 1539758421, 'test', 3585.98, '大梨', '18000000000', 1, 1539842050, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '[]', '已经处理了', 'admin', '', 0, '', 0, 0),
        (5, 12, 3, 28, '201809150101000001', 1539758658, 'test', 3585.98, '大梨', '18000000000', 2, 1539842128, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '[]', '已经处理了', 'admin', 'admin', 1539842158, '已经处理了', 0, 0),
        (8, 13, 0, 28, '201809150102000002', 1539758658, 'test', 0.00, '大梨', '18000000000', 3, 1539842232, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '[]', '理由不够充分', 'admin', '', 0, '', 0, 0),
        (9, 14, 2, 26, '201809130101000001', 1539758097, 'test', 3500.00, '大梨', '18000000000', 2, 1540367096, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '[]', '呵呵', 'admin', 'admin', 1540367195, '收货了', 0, 0),
        (10, 14, 0, 27, '201809130101000001', 1539758421, 'test', 0.00, '大梨', '18000000000', 3, 1540367217, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '[]', '就是不退', 'admin', '', 0, '', 0, 0),
        (11, 14, 2, 28, '201809130101000001', 1539758658, 'test', 591.05, '大梨', '18000000000', 1, 1540372144, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '[]', '可以退款', 'admin', '', 0, '', 0, 0),
        (12, 15, 3, 26, '201809130102000002', 1539758097, 'test', 3500.00, '大梨', '18000000000', 2, 1540372974, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '[]', '退货了', 'admin', 'admin', 1540372986, '收货了', 0, 0),
        (13, 15, 0, 27, '201809130102000002', 1539758421, 'test', 0.00, '大梨', '18000000000', 3, 1540373010, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '[]', '无法退货', 'admin', '', 0, '', 0, 0),
        (15, 16, 0, 26, '201809140101000001', 1539758097, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '[]', '', '', '', 0, '', 0, 0),
        (16, 16, 0, 27, '201809140101000001', 1539758421, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '[]', '', '', '', 0, '', 0, 0),
        (17, 16, 0, 28, '201809140101000001', 1539758658, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '[]', '', '', '', 0, '', 0, 0),
        (18, 17, 0, 26, '201809150101000003', 1539758097, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '[]', '', '', '', 0, '', 0, 0),
        (19, 17, 0, 27, '201809150101000003', 1539758421, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '[]', '', '', '', 0, '', 0, 0),
        (20, 17, 0, 28, '201809150101000003', 1539758658, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '[]', '', '', '', 0, '', 0, 0),
        (21, 18, 0, 26, '201809150102000004', 1539758097, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '[]', '', '', '', 0, '', 0, 0),
        (22, 18, 0, 27, '201809150102000004', 1539758421, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '[]', '', '', '', 0, '', 0, 0),
        (23, 18, 0, 28, '201809150102000004', 1539758658, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '[]', '', '', '', 0, '', 0, 0),
        (24, 19, 0, 26, '201809130101000003', 1539758097, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '[]', '', '', '', 0, '', 0, 0),
        (25, 19, 0, 27, '201809130101000003', 1539758421, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '[]', '', '', '', 0, '', 0, 0),
        (26, 19, 0, 28, '201809130101000003', 1539758658, 'test', 0.00, '大梨', '18000000000', 0, 0, '/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '[]', '', '', '', 0, '', 0, 0);
update oms_order_return_apply set status = 0 where id = 3;


insert into go_mall_admin.oms_company_address (id, address_name, send_status, receive_status, name, phone, province, city, region, detail_address)
values  (1, '深圳发货点', 1, 1, '大梨', '18000000000', '广东省', '深圳市', '南山区', '科兴科学园'),
        (2, '北京发货点', 0, 0, '大梨', '18000000000', '北京市', '', '南山区', '科兴科学园'),
        (3, '南京发货点', 0, 0, '大梨', '18000000000', '江苏省', '南京市', '南山区', '科兴科学园');


insert into go_mall_admin.oms_cart_item (id, product_id, product_sku_id, member_id, quantity, price, product_pic, product_name, product_sub_title, product_sku_code, member_nickname, delete_status, product_category_id, product_brand, product_sn, product_attr, modify_date, create_date)
values  (12, 26, 90, 1, 1, 3788.00, '', '华为 HUAWEI P20', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '', '', '', 0, 1535360024),
        (13, 27, 98, 1, 3, 2699.00, '', '小米8', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '', '', '', 0, 1535361113),
        (14, 28, 102, 1, 1, 649.00, '', '红米5A', '8天超长待机，137g轻巧机身，高通骁龙处理器小米6X低至1299，点击抢购', '201808270028001', 'windir', 1, 19, '', '', '', 0, 1535361482),
        (15, 28, 103, 1, 1, 699.00, '', '红米5A', '8天超长待机，137g轻巧机身，高通骁龙处理器小米6X低至1299，点击抢购', '201808270028001', 'windir', 1, 19, '', '', '', 0, 1535422965),
        (16, 29, 106, 1, 1, 5499.00, '', 'Apple iPhone 8 Plus', '【限时限量抢购】Apple产品年中狂欢节，好物尽享，美在智慧！速来 >> 勾选[保障服务][原厂保2年]，获得AppleCare+全方位服务计划，原厂延保售后无忧。', '201808270029001', 'windir', 1, 19, '', '', '', 0, 1535424650),
        (19, 36, 163, 1, 3, 100.00, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '202002210036001', 'windir', 1, 29, 'NIKE', '6799345', '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"秋季"}]', 0, 1582617119),
        (20, 36, 164, 1, 2, 120.00, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '202002210036001', 'windir', 1, 29, 'NIKE', '6799345', '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]', 0, 1582617263),
        (21, 36, 164, 1, 2, 120.00, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '202002210036001', 'windir', 1, 29, 'NIKE', '6799345', '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]', 0, 1582620593),
        (22, 26, 110, 1, 3, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588577664),
        (23, 27, 98, 1, 7, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588577743),
        (24, 26, 110, 1, 4, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588582697),
        (25, 27, 98, 1, 2, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588582703),
        (26, 28, 102, 1, 4, 649.00, '/mall/images/20180615/5a9d248cN071f4959.jpg', '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', '8天超长待机，137g轻巧机身，高通骁龙处理器小米6X低至1299，点击抢购', '201808270028001', 'windir', 1, 19, '小米', '7437789', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588582706),
        (27, 29, 106, 1, 1, 4999.00, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', 'Apple iPhone 8 Plus 64GB 红色特别版 移动联通电信4G手机', '【限时限量抢购】Apple产品年中狂欢节，好物尽享，美在智慧！速来 >> 勾选[保障服务][原厂保2年]，获得AppleCare+全方位服务计划，原厂延保售后无忧。', '201808270029001', 'windir', 1, 19, '苹果', '7437799', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]', 0, 1588582709),
        (28, 26, 110, 1, 2, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588583240),
        (29, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588583243),
        (30, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588583294),
        (31, 29, 106, 1, 1, 4999.00, '/mall/images/20180615/5acc5248N6a5f81cd.jpg', 'Apple iPhone 8 Plus 64GB 红色特别版 移动联通电信4G手机', '【限时限量抢购】Apple产品年中狂欢节，好物尽享，美在智慧！速来 >> 勾选[保障服务][原厂保2年]，获得AppleCare+全方位服务计划，原厂延保售后无忧。', '201808270029001', 'windir', 1, 19, '苹果', '7437799', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]', 0, 1588583396),
        (32, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588583630),
        (33, 27, 98, 1, 2, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588583775),
        (34, 36, 164, 1, 1, 120.00, '/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '202002210036002', 'windir', 1, 29, 'NIKE', '6799345', '[{"key":"颜色","value":"红色"},{"key":"尺寸","value":"38"},{"key":"风格","value":"夏季"}]', 0, 1588583960),
        (35, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588646499),
        (36, 26, 110, 1, 2, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588646515),
        (37, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588646577),
        (38, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588660168),
        (39, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588660372),
        (40, 26, 110, 1, 2, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1588660400),
        (41, 27, 98, 1, 2, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1588661353),
        (42, 26, 111, 1, 1, 3999.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026002', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]', 0, 1588663565),
        (43, 28, 102, 1, 1, 649.00, '/mall/images/20180615/5a9d248cN071f4959.jpg', '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', '8天超长待机，137g轻巧机身，高通骁龙处理器小米6X低至1299，点击抢购', '201808270028001', 'windir', 1, 19, '小米', '7437789', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589613364),
        (44, 26, 110, 1, 2, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589613480),
        (45, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589698816),
        (46, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1589698822),
        (47, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589699654),
        (48, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589700003),
        (49, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589700114),
        (50, 26, 110, 1, 2, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589702842),
        (51, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1589702846),
        (52, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589715216),
        (53, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1589715219),
        (54, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589715547),
        (55, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589715686),
        (56, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589791800),
        (57, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1589804524),
        (58, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1589804528),
        (59, 27, 98, 1, 2, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1590222073),
        (60, 27, 98, 1, 2, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1590224488),
        (61, 26, 110, 1, 2, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1590284210),
        (62, 26, 110, 1, 2, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1590284679),
        (63, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1591520508),
        (64, 27, 98, 1, 2, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1592119480),
        (65, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1592720833),
        (66, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1592723534),
        (67, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1592723573),
        (68, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'windir', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1592723710),
        (69, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1593224868),
        (70, 27, 98, 1, 1, 2699.00, '/mall/images/20180615/xiaomi.jpg', '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', '骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购', '201808270027001', 'windir', 1, 19, '小米', '7437788', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"32G"}]', 0, 1666939846),
        (71, 37, 201, 1, 1, 5999.00, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 紫色 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', '202210280037001', 'windir', 1, 19, '苹果', '100038005189', '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]', 0, 1666942052),
        (72, 40, 221, 1, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'windir', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1667978086),
        (73, 38, 213, 1, 1, 3599.00, '/mall/images/20221028/ipad_001.jpg', 'Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）', '【11.11大爱超大爱】iPad9代限量抢购，价格优惠，更享以旧换新至高补贴325元！！快来抢购吧！！ ', '202210280038001', 'windir', 1, 53, '苹果', '100044025833', '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]', 0, 1667978728),
        (74, 37, 201, 1, 1, 5999.00, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', '202210280037001', 'windir', 1, 19, '苹果', '100038005189', '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]', 0, 1667978764),
        (75, 45, 239, 1, 1, 0.00, '/mall/images/20221108/oppo_r8_01.jpg', 'OPPO Reno8 8GB+128GB 鸢尾紫 新配色上市 80W超级闪充 5000万水光人像三摄 3200万前置索尼镜头 5G手机', '【11.11提前购机享价保，好货不用等，系统申请一键价保补差!】【Reno8Pro爆款优惠】 ', '202211080045001', 'windir', 1, 19, 'OPPO', '10052147850350', '[{"key":"颜色","value":"鸢尾紫"},{"key":"容量","value":"128G"}]', 0, 1667981783),
        (76, 45, 239, 1, 1, 2299.00, '/mall/images/20221108/oppo_r8_01.jpg', 'OPPO Reno8 8GB+128GB 鸢尾紫 新配色上市 80W超级闪充 5000万水光人像三摄 3200万前置索尼镜头 5G手机', '【11.11提前购机享价保，好货不用等，系统申请一键价保补差!】【Reno8Pro爆款优惠】 ', '202211080045001', 'windir', 1, 19, 'OPPO', '10052147850350', '[{"key":"颜色","value":"鸢尾紫"},{"key":"容量","value":"128G"}]', 0, 1667981916),
        (77, 41, 225, 1, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量 墨羽 12GB+256GB 5G智能手机 小米 红米', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'windir', 1, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1668064776),
        (78, 37, 201, 1, 2, 5999.00, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', '202210280037001', 'windir', 1, 19, '苹果', '100038005189', '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]', 0, 1668064784),
        (79, 38, 213, 1, 1, 3599.00, '/mall/images/20221028/ipad_001.jpg', 'Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）', '【11.11大爱超大爱】iPad9代限量抢购，价格优惠，更享以旧换新至高补贴325元！！快来抢购吧！！ ', '202210280038001', 'windir', 1, 53, '苹果', '100044025833', '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]', 0, 1668152260),
        (80, 38, 213, 1, 1, 3599.00, '/mall/images/20221028/ipad_001.jpg', 'Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）', '【11.11大爱超大爱】iPad9代限量抢购，价格优惠，更享以旧换新至高补贴325元！！快来抢购吧！！ ', '202210280038001', 'windir', 1, 53, '苹果', '100044025833', '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]', 0, 1668152292),
        (81, 38, 213, 1, 3, 3599.00, '/mall/images/20221028/ipad_001.jpg', 'Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）', '【11.11大爱超大爱】iPad9代限量抢购，价格优惠，更享以旧换新至高补贴325元！！快来抢购吧！！ ', '202210280038001', 'windir', 1, 53, '苹果', '100044025833', '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]', 0, 1668152302),
        (82, 40, 221, 1, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'windir', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1668154043),
        (83, 40, 221, 1, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'windir', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1668154391),
        (84, 37, 201, 1, 1, 5999.00, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', '202210280037001', 'windir', 1, 19, '苹果', '100038005189', '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]', 0, 1668154505),
        (85, 28, 102, 1, 1, 649.00, '/mall/images/20180615/5a9d248cN071f4959.jpg', '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', '8天超长待机，137g轻巧机身，高通骁龙处理器小米6X低至1299，点击抢购', '201808270028001', 'windir', 1, 19, '小米', '7437789', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1668154865),
        (86, 40, 221, 1, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'windir', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1668565367),
        (87, 41, 225, 1, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'windir', 1, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1668565371),
        (88, 39, 217, 1, 1, 5999.00, '/mall/images/20221028/xiaomi_computer_001.jpg', '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑', '【双十一大促来袭】指定型号至高优惠1000，以旧换新至高补贴1000元，晒单赢好礼', '202210280039001', 'windir', 1, 54, '小米', '100023207945', '[{"key":"颜色","value":"新小米Pro 14英寸 2.8K屏"},{"key":"版本","value":"R7 16G 512"}]', 0, 1668565374),
        (89, 37, 201, 1, 1, 5999.00, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', '202210280037001', 'windir', 1, 19, '苹果', '100038005189', '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]', 0, 1668565396),
        (90, 40, 221, 1, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'test', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1671608940),
        (91, 37, 201, 1, 1, 5999.00, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', '202210280037001', 'test', 1, 19, '苹果', '100038005189', '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]', 0, 1671608982),
        (92, 41, 225, 1, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'test', 1, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1671608993),
        (93, 40, 221, 1, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'test', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1671609063),
        (94, 41, 225, 1, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'test', 1, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1671609088),
        (95, 41, 225, 1, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'test', 1, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1671612316),
        (96, 40, 221, 1, 2, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'test', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1671612401),
        (97, 40, 221, 1, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'test', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1671612610),
        (98, 40, 221, 1, 2, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'test', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1671760511),
        (99, 26, 110, 1, 1, 3788.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'test', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1673336343),
        (100, 26, 111, 1, 1, 3899.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026002', 'test', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"32G"}]', 0, 1673341088),
        (101, 26, 110, 1, 1, 3699.00, '/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20 ', 'AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待', '201806070026001', 'test', 1, 19, '华为', '6946605', '[{"key":"颜色","value":"金色"},{"key":"容量","value":"16G"}]', 0, 1673341826),
        (102, 40, 221, 11, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'member', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1683789873),
        (103, 41, 225, 11, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'member', 1, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1683789877),
        (104, 38, 213, 11, 1, 3599.00, '/mall/images/20221028/ipad_001.jpg', 'Apple iPad 10.9英寸平板电脑 2022年款', '【11.11大爱超大爱】iPad9代限量抢购，价格优惠，更享以旧换新至高补贴325元！！快来抢购吧！！ ', '202210280038001', 'member', 1, 53, '苹果', '100044025833', '[{"key":"颜色","value":"银色"},{"key":"容量","value":"64G"}]', 0, 1683790232),
        (105, 39, 217, 11, 1, 5999.00, '/mall/images/20221028/xiaomi_computer_001.jpg', '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑', '【双十一大促来袭】指定型号至高优惠1000，以旧换新至高补贴1000元，晒单赢好礼', '202210280039001', 'member', 1, 54, '小米', '100023207945', '[{"key":"颜色","value":"新小米Pro 14英寸 2.8K屏"},{"key":"版本","value":"R7 16G 512"}]', 0, 1683790298),
        (106, 44, 235, 11, 1, 369.00, '/mall/images/20221108/sanxing_ssd_02.jpg', '三星（SAMSUNG）500GB SSD固态硬盘 M.2接口(NVMe协议)', '【满血无缓存！进店抽百元E卡，部分型号白条三期免息】兼具速度与可靠性！读速高达3500MB/s，全功率模式！点击 ', '202211080044001', 'member', 1, 55, '三星', '100018768480', '[{"key":"颜色","value":"新品980｜NVMe PCIe3.0*4"},{"key":"版本","value":"512GB"}]', 0, 1683790336),
        (107, 42, 229, 11, 1, 4999.00, '/mall/images/20221104/huawei_mate50_01.jpg', 'HUAWEI Mate 50 直屏旗舰 超光变XMAGE影像 北斗卫星消息', '【华为Mate50新品上市】内置66W华为充电套装，超光变XMAGE影像,北斗卫星消息，鸿蒙操作系统3.0！立即抢购！华为新品可持续计划，猛戳》 ', '202211040042001', 'member', 1, 19, '华为', '100035295081', '[{"key":"颜色","value":"曜金黑"},{"key":"容量","value":"128G"}]', 0, 1683790364),
        (108, 37, 201, 11, 1, 5999.00, '/mall/images/20221028/iphone14_001.jpg', 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '【11.11大爱超大爱】指定iPhone14产品限时限量领券立减601元！！！部分iPhone产品现货抢购确认收货即送原厂手机壳10元优惠券！！！猛戳 ', '202210280037001', 'member', 1, 19, '苹果', '100038005189', '[{"key":"颜色","value":"午夜色"},{"key":"容量","value":"128G"}]', 0, 1683790472),
        (109, 40, 221, 11, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'member', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1683790502),
        (110, 41, 225, 11, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'member', 1, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1683790521),
        (111, 40, 221, 11, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'member', 1, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1683790617),
        (112, 39, 217, 11, 1, 5999.00, '/mall/images/20221028/xiaomi_computer_001.jpg', '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑', '【双十一大促来袭】指定型号至高优惠1000，以旧换新至高补贴1000元，晒单赢好礼', '202210280039001', 'member', 1, 54, '小米', '100023207945', '[{"key":"颜色","value":"新小米Pro 14英寸 2.8K屏"},{"key":"版本","value":"R7 16G 512"}]', 0, 1683790624),
        (113, 40, 221, 11, 1, 2999.00, '/mall/images/20221104/xiaomi_12_pro_01.jpg', '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充', '天玑9000+处理器、5160mAh大电量、2KAmoled超视感屏【点击购买小米11Ultra，戳】 ', '202211040040001', 'member', 0, 19, '小米', '100027789721', '[{"key":"颜色","value":"黑色"},{"key":"容量","value":"128G"}]', 0, 1683790677),
        (114, 41, 225, 11, 1, 2099.00, '/mall/images/20221104/redmi_k50_01.jpg', 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量', '【品质好物】天玑8100，2K直屏，5500mAh大电量【Note12Pro火热抢购中】 ', '202211040041001', 'member', 0, 19, '小米', '100035246702', '[{"key":"颜色","value":"墨羽"},{"key":"容量","value":"128G"}]', 0, 1683790683);




insert into go_mall_admin.sms_coupon (id, type, name, platform, count, amount, per_limit, min_point, use_type, note, publish_count, use_count, receive_count, code, member_level, start_time, end_time, enable_time)
values  (27, 0, '全品类通用券', 0, 94, 10.00, 10, 100.00, 0, '', 100, 0, 6, '', 0, 1667836800, 1701273600, 1667836800),
        (28, 0, '手机分类专用券', 0, 995, 100.00, 5, 1000.00, 1, '', 1000, 0, 5, '', 0, 1667836800, 1701273600, 1667836800),
        (29, 0, '苹果手机专用券', 0, 998, 600.00, 1, 4000.00, 2, '', 1000, 0, 2, '', 0, 1667836800, 1701273600, 1667836800),
        (30, 0, '小米手机专用券', 0, 998, 200.00, 1, 2000.00, 2, '', 1000, 0, 2, '', 0, 1667836800, 1701273600, 1667836800),
        (31, 0, '限时优惠券', 0, 999, 20.00, 5, 500.00, 0, '', 1000, 0, 1, '', 0, 1669824000, 1671638400, 1671724800);

insert into go_mall_admin.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, use_status, order_id, order_sn, create_time, use_time)
values  (37, 30, 1, '7806895974110001', 'windir', 1, 1, 0, '', 1667978069, 1667978098),
        (38, 27, 1, '7872472849240001', 'windir', 1, 1, 0, '', 1667978725, 1667978738),
        (39, 29, 1, '7876204111480001', 'windir', 1, 1, 0, '', 1667978762, 1667978771),
        (40, 27, 1, '7911945030190001', 'windir', 1, 0, 0, '', 1667979119, 0),
        (41, 28, 1, '8194868687790001', 'windir', 1, 1, 0, '', 1667981949, 1668154362),
        (42, 28, 1, '1239565720390001', 'test', 1, 1, 0, '', 1671612396, 1671612787),
        (43, 31, 1, '6030247208280001', 'test', 1, 0, 0, '', 1671760302, 0),
        (44, 28, 1, '6050939443480001', 'test', 1, 0, 0, '', 1671760509, 0),
        (45, 27, 1, '4182437014580001', 'test', 1, 0, 0, '', 1673341824, 0),
        (46, 27, 11, '9011281751500011', 'member', 1, 0, 0, '', 1683790113, 0),
        (47, 28, 11, '9011495851370011', 'member', 1, 1, 0, '', 1683790115, 1683790636),
        (48, 30, 11, '9011677197430011', 'member', 1, 1, 0, '', 1683790117, 1683790136),
        (49, 27, 11, '9046676666610011', 'member', 1, 0, 0, '', 1683790467, 0),
        (50, 28, 11, '9046909751910011', 'member', 1, 0, 0, '', 1683790469, 0),
        (51, 29, 11, '9047077722990011', 'member', 1, 0, 0, '', 1683790471, 0),
        (52, 27, 11, '9075818288630011', 'member', 1, 0, 0, '', 1683790758, 0);

insert into go_mall_admin.sms_flash_promotion (id, title, start_date, end_date, status, create_time)
values  (14, '双11特卖活动', UNIX_TIMESTAMP('2022-11-09'), UNIX_TIMESTAMP('2023-12-31'), 1, UNIX_TIMESTAMP('2022-11-09 14:56:48'));


-- insert into go_mall_admin.sms_flash_promotion_session (id, name, start_time, end_time, status, create_time)
-- values  (1, '8:00', '08:00:00', '10:00:00', 1, UNIX_TIMESTAMP('2018-11-16 13:22:17')),
--         (2, '10:00', '10:00:00', '12:00:00', 1, UNIX_TIMESTAMP('2018-11-16 13:22:34')),
--         (3, '12:00', '12:00:00', '14:00:00', 1, UNIX_TIMESTAMP('2018-11-16 13:22:37')),
--         (4, '14:00', '14:00:00', '16:00:00', 1, UNIX_TIMESTAMP('2018-11-16 13:22:41')),
--         (5, '16:00', '16:00:00', '18:00:00', 1, UNIX_TIMESTAMP( '2018-11-16 13:22:45')),
--         (6, '18:00', '18:00:00', '20:00:00', 1, UNIX_TIMESTAMP('2018-11-16 13:21:34')),
--         (7, '20:00', '20:00:00', '22:00:00', 1, UNIX_TIMESTAMP('2018-11-16 13:22:55'));



insert into go_mall_admin.sms_coupon_product_relation (id, coupon_id, product_id, product_name, product_sn)
values  (18, 29, 37, 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', '100038005189'),
        (19, 29, 29, 'Apple iPhone 8 Plus 64GB 红色特别版 移动联通电信4G手机', '7437799'),
        (21, 30, 41, 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量 墨羽 12GB+256GB 5G智能手机 小米 红米', '100035246702'),
        (22, 30, 40, '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', '100027789721');

insert into go_mall_admin.sms_coupon_product_category_relation (id, coupon_id, product_category_id, product_category_name, parent_category_name)
values  (11, 28, 19, '手机通讯', '手机数码');

insert into go_mall_admin.sms_home_brand (id, brand_id, brand_name, recommend_status, sort)
values  (6, 6, '小米', 1, 300),
        (32, 50, '海澜之家', 1, 198),
        (33, 51, '苹果', 1, 199),
        (34, 2, '三星', 1, 195),
        (35, 3, '华为', 1, 200),
        (39, 21, 'OPPO', 1, 197),
        (45, 1, '万和', 1, 0),
        (46, 5, '方太', 1, 0),
        (47, 4, '格力', 1, 0);

insert into go_mall_admin.sms_home_new_product (id, product_id, product_name, recommend_status, sort)
values  (19, 37, 'Apple iPhone 14 (A2884) 128GB 支持移动联通电信5G 双卡双待手机', 1, 197),
        (20, 38, 'Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）', 1, 0),
        (21, 39, '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑(新R5-6600H标压 16G 512G win11)', 1, 198),
        (22, 40, '小米12 Pro 天玑版 天玑9000+处理器 5000万疾速影像 2K超视感屏 120Hz高刷 67W快充 12GB+256GB 黑色 5G手机', 1, 200),
        (23, 41, 'Redmi K50 天玑8100 2K柔性直屏 OIS光学防抖 67W快充 5500mAh大电量 墨羽 12GB+256GB 5G智能手机 小米 红米', 1, 199),
        (24, 42, 'HUAWEI Mate 50 直屏旗舰 超光变XMAGE影像 北斗卫星消息 低电量应急模式 128GB曜金黑华为鸿蒙手机', 1, 0),
        (25, 44, '三星（SAMSUNG）500GB SSD固态硬盘 M.2接口(NVMe协议) 980（MZ-V8V500BW）', 1, 0),
        (26, 45, 'OPPO Reno8 8GB+128GB 鸢尾紫 新配色上市 80W超级闪充 5000万水光人像三摄 3200万前置索尼镜头 5G手机', 1, 0),
        (27, 43, '万和（Vanward)燃气热水器天然气家用四重防冻直流变频节能全新升级增压水伺服恒温高抗风 JSQ30-565W16【16升】【恒温旗舰款】', 1, 0);


insert into go_mall_admin.sms_home_recommend_product (id, product_id, product_name, recommend_status, sort)
values  (10, 38, 'Apple iPad 10.9英寸平板电脑 2022年款（64GB WLAN版/A14芯片/1200万像素/iPadOS MPQ03CH/A ）', 1, 0),
        (11, 39, '小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑(新R5-6600H标压 16G 512G win11)', 1, 0),
        (12, 44, '三星（SAMSUNG）500GB SSD固态硬盘 M.2接口(NVMe协议) 980（MZ-V8V500BW）', 1, 0),
        (13, 43, '万和（Vanward)燃气热水器天然气家用四重防冻直流变频节能全新升级增压水伺服恒温高抗风 JSQ30-565W16【16升】【恒温旗舰款】', 1, 0),
        (14, 45, 'OPPO Reno8 8GB+128GB 鸢尾紫 新配色上市 80W超级闪充 5000万水光人像三摄 3200万前置索尼镜头 5G手机', 1, 0);

insert into go_mall_admin.sms_home_recommend_subject (id, subject_id, subject_name, recommend_status, sort)
values  (14, 1, 'polo衬衫的也时尚', 1, 0),
        (15, 2, '大牌手机低价秒', 1, 0),
        (16, 3, '晓龙845新品上市', 1, 0),
        (17, 4, '夏天应该穿什么', 1, 0),
        (18, 5, '夏季精选', 1, 100),
        (19, 6, '品牌手机降价', 1, 0);

insert into go_mall_admin.sms_home_advertise (id, name, type, pic, status, click_count, order_count, url, note, sort, start_time, end_time)
values  (2, '夏季大热促销', 1, '/mall/images/20190525/ad1.jpg', 0, 0, 0, '', '夏季大热促销', 0, 1541052097, 1542261697),
        (3, '夏季大热促销1', 1, '/mall/images/20190525/ad1.jpg', 0, 0, 0, '', '夏季大热促销1', 0, 1542088897, 1542088897),
        (4, '夏季大热促销2', 1, '/mall/images/20190525/ad2.jpg', 0, 0, 0, '', '夏季大热促销2', 0, 1542088897, 1542088897),
        (9, '电影推荐广告', 1, '/mall/images/20181113/movie_ad.jpg', 0, 0, 0, 'www.baidu.com', '电影推荐广告', 100, 1541001600, 1542988800),
        (10, '汽车促销广告', 1, '/mall/images/20181113/car_ad.jpg', 0, 0, 0, 'xxx', '', 99, 1542038400, 1542988800),
        (11, '汽车推荐广告', 1, '/mall/images/20181113/car_ad2.jpg', 0, 0, 0, 'xxx', '', 98, 1542038400, 1543507200),
        (12, '小米推荐广告', 1, '/mall/images/20221108/xiaomi_banner_01.png', 1, 0, 0, '/pages/brand/brandDetail?id=6', '', 0, 1667898243, 1699434245),
        (13, '华为推荐广告', 1, '/mall/images/20221108/huawei_banner_01.png', 1, 0, 0, '/pages/brand/brandDetail?id=3', '', 0, 1667898627, 1699434628),
        (14, '苹果推荐广告', 1, '/mall/images/20221108/apple_banner_01.png', 1, 0, 0, '/pages/brand/brandDetail?id=51', '', 0, 1667898774, 1699434775),
        (15, '三星推荐广告', 1, '/mall/images/20221108/sanxing_banner_01.png', 1, 0, 0, '/pages/brand/brandDetail?id=2', '', 0, 1667898938, 1699434939),
        (16, 'OPPO推荐广告', 1, '/mall/images/20221108/oppo_banner_01.png', 1, 0, 0, '/pages/brand/brandDetail?id=21', '', 0, 1667899210, 1699435211);


