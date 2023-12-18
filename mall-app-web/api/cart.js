import request from '@/utils/requestUtil'

export function addCartItem(data) {
	return request({
		method: 'POST',
		url: '/cart/add',
		data: data
	})
}

export function fetchCartList() {
	return request({
		method: 'GET',
		url: '/cart/list'
	})
}

export function deletCartItem(data) {
	return request({
		method: 'POST',
		url: '/cart/delete',
    data:data
	})
}

export function updateQuantity(data) {
	return request({
		method: 'POST',
		url: '/cart/update/quantity',
    data:data
	})
}

export function clearCartList() {
	return request({
		method: 'POST',
		url: '/cart/clear'
	})
}