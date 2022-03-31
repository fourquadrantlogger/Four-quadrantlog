import request from './request'

export function getLogList(params) {
	return request({
		url: '/api/log',
		method: 'get',
		params
	})
}
 
export function createlog(d) {
	return request({
		url: '/api/log',
		method: 'POST',
		data:d,
	})
}
 
 