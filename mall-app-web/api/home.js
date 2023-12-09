import request from '@/utils/requestUtil'

// 首页内容信息展示
export function fetchContent() {
	return request({
		method: 'GET',
		url: '/home/content'
	})
}

// 分页获取推荐商品
export function fetchRecommendProductList(params) {
	return request({
		method: 'GET',
		url: '/home/recommendProductList',
		params:params
	})
}

// 获取首页商品分类
export function fetchProductCateList(parentId) {
	return request({
		method: 'GET',
		url: '/home/productCateList/'+parentId,
	})
}

// 分页获取新品推荐商品
export function fetchNewProductList(params) {
	return request({
		method: 'GET',
		url: '/home/newProductList',
		params:params
	})
}

// 分页获取人气推荐商品
export function fetchHotProductList(params) {
	return request({
		method: 'GET',
		url: '/home/hotProductList',
		params:params
	})
}
