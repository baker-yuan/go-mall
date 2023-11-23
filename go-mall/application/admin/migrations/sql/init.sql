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
        (59, '测试品牌', 'C', 0, 0, 0, 0, 0, 'http://localhost:9000/mall/20220609/Snipaste_2022-06-08_14-35-53.png', 'http://localhost:9000/mall/20220609/biji_05.jpg', '12345', 0, 0);


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